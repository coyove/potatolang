package potatolang

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"unsafe"

	"github.com/coyove/potatolang/parser"
)

func makeop(op byte, a, b uint16) uint32 {
	// 6 + 13 + 13
	return uint32(op)<<26 + uint32(a&0x1fff)<<13 + uint32(b&0x1fff)
}

func op(x uint32) (op byte, a, b uint16) {
	op = byte(x >> 26)
	a = uint16(x>>13) & 0x1fff
	b = uint16(x) & 0x1fff
	return
}

func u32Bytes(p []uint32) []byte {
	r := reflect.SliceHeader{}
	r.Cap = cap(p) * 4
	r.Len = len(p) * 4
	r.Data = (*reflect.SliceHeader)(unsafe.Pointer(&p)).Data
	return *(*[]byte)(unsafe.Pointer(uintptr(unsafe.Pointer(&r))))
}

func u32FromBytes(p []byte) []uint32 {
	if m := len(p) % 4; m != 0 {
		p = append(p, 0, 0, 0, 0)
		p = p[:len(p)+1-m]
	}
	r := reflect.SliceHeader{}
	r.Cap = cap(p) / 4
	r.Len = len(p) / 4
	r.Data = (*reflect.SliceHeader)(unsafe.Pointer(&p)).Data
	return *(*[]uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(&r))))
}

type posVByte []byte

func (p *posVByte) appendABC(a, b uint32, c uint16) {
	trunc := func(v uint32) (w byte, buf [4]byte) {
		if v < 256 {
			w = 0
		} else if v < 256*256 {
			w = 1
		} else if v < 256*256*256 {
			w = 2
		} else {
			w = 3
		}
		binary.LittleEndian.PutUint32(buf[:], v)
		return
	}
	aw, ab := trunc(a)
	bw, bb := trunc(b)
	x := aw<<6 + bw<<4
	cw := 0
	c--
	if c < 14 {
		x |= byte(c)
	} else if c < 256 {
		cw = 1
		x |= 0xe
	} else {
		cw = 2
		x |= 0xf
	}

	*p = append(*p, x)
	*p = append(*p, ab[:aw+1]...)
	*p = append(*p, bb[:bw+1]...)
	if cw == 1 {
		*p = append(*p, byte(c))
	} else if cw == 2 {
		*p = append(*p, byte(c), byte(c>>8))
	}
}

func (p posVByte) readABC(i int) (next int, a, b uint32, c uint16) {
	x := p[i]

	i++
	buf := [4]byte{}
	copy(buf[:], p[i:i+1+int(x>>6)])
	a = binary.LittleEndian.Uint32(buf[:])

	i = i + 1 + int(x>>6)
	buf = [4]byte{}
	copy(buf[:], p[i:i+1+int(x<<2>>6)])
	b = binary.LittleEndian.Uint32(buf[:])

	i = i + 1 + int(x<<2>>6)
	x &= 0xf
	if x < 14 {
		c = uint16(x + 1)
		next = i
	} else {
		buf = [4]byte{}
		copy(buf[:], p[i:i+int(x-13)])
		c = binary.LittleEndian.Uint16(buf[:2]) + 1
		next = i + int(x-13)
	}
	return
}

type packet struct {
	data   []uint32
	pos    posVByte
	source string
}

func newpacket() packet {
	return packet{data: make([]uint32, 0, 1), pos: make(posVByte, 0, 1)}
}

func (b *packet) Clear() {
	b.data = b.data[:0]
}

func (p *packet) Write(buf packet) {
	datalen := len(p.data)
	p.data = append(p.data, buf.data...)
	i := 0
	for i < len(buf.pos) {
		var a, b uint32
		var c uint16
		i, a, b, c = buf.pos.readABC(i)
		p.pos.appendABC(a+uint32(datalen), b, c)
	}
	p.source = buf.source
}

func (b *packet) WriteRaw(buf []uint32) { b.data = append(b.data, buf...) }

func (b *packet) Write64(v uint64) { b.data = append(b.data, uint32(v>>32), uint32(v)) }

func (b *packet) Write32(v uint32) { b.data = append(b.data, v) }

func (b *packet) WriteOP(op byte, opa, opb uint16) { b.data = append(b.data, makeop(op, opa, opb)) }

func (b *packet) WritePos(p parser.Meta) {
	b.pos.appendABC(uint32(len(b.data)), p.Line, p.Column)
	if p.Source != "" {
		b.source = p.Source
	}
}

func (b *packet) WriteDouble(v float64) {
	d := *(*uint64)(unsafe.Pointer(&v))
	b.Write64(d)
}

func (b *packet) WriteString(v string) {
	b.Write64(uint64(len(v)))
	b.WriteRaw(u32FromBytes([]byte(v)))
}

func (b *packet) TruncateLast(n int) {
	if len(b.data) > n {
		b.data = b.data[:len(b.data)-n]
	}
}

func (b *packet) WriteConsts(consts []interface{}) {
	// const table struct:
	// all values are placed sequentially
	// for numbers other than MaxUint64, they will be written directly
	// for MaxUint64, it will be written twice
	// for strings, a MaxUint64 will be written first, then the string
	for _, k := range consts {
		switch k := k.(type) {
		case float64:
			n := k
			if math.Float64bits(n) == math.MaxUint64 {
				b.Write64(math.MaxUint64)
				b.Write64(math.MaxUint64)
			} else {
				b.WriteDouble(n)
			}
		case string:
			b.Write64(math.MaxUint64)
			b.WriteString(k)
		}
	}
}

func (b *packet) Len() int {
	return len(b.data)
}

func crRead(data []uint32, cursor *uint32, len int) []uint32 {
	*cursor += uint32(len)
	return data[*cursor-uint32(len) : *cursor]
}

func crRead32(data []uint32, cursor *uint32) uint32 {
	*cursor++
	return data[*cursor-1]
}

func crRead64(data []uint32, cursor *uint32) uint64 {
	*cursor += 2
	return uint64(data[*cursor-2])<<32 + uint64(data[*cursor-1])
}

func crReadDouble(data []uint32, cursor *uint32) float64 {
	d := crRead64(data, cursor)
	return *(*float64)(unsafe.Pointer(&d))
}

func crReadString(data []uint32, cursor *uint32) string {
	x := crRead32(data, cursor)
	return crReadStringLen(data, int(x), cursor)
}

func crReadStringLen(data []uint32, length int, cursor *uint32) string {
	return string(crReadBytesLen(data, length, cursor))
}

func crReadBytesLen(data []uint32, length int, cursor *uint32) []byte {
	buf := crRead(data, cursor, int((length+3)/4))
	return u32Bytes(buf)[:length]
}

var singleOp = map[byte]string{
	OP_ASSERT:   "assert",
	OP_ADD:      "add",
	OP_SUB:      "sub",
	OP_MUL:      "mul",
	OP_DIV:      "div",
	OP_MOD:      "mod",
	OP_EQ:       "eq",
	OP_NEQ:      "neq",
	OP_LESS:     "less",
	OP_LESS_EQ:  "less-eq",
	OP_LEN:      "len",
	OP_COPY:     "copy",
	OP_LOAD:     "load",
	OP_STORE:    "store",
	OP_NOT:      "not",
	OP_BIT_NOT:  "bit-not",
	OP_BIT_AND:  "bit-and",
	OP_BIT_OR:   "bit-or",
	OP_BIT_XOR:  "bit-xor",
	OP_BIT_LSH:  "bit-lsh",
	OP_BIT_RSH:  "bit-rsh",
	OP_BIT_URSH: "bit-ursh",
	OP_TYPEOF:   "typeof",
	OP_SLICE:    "slice",
	OP_POP:      "pop",
}

func (c *Closure) crPrettify(tab int) string {
	sb := &bytes.Buffer{}
	spaces := strings.Repeat("        ", tab)
	spaces2 := ""
	if tab > 0 {
		spaces2 = strings.Repeat("        ", tab-1) + "+-------"
	}

	sb.WriteString(spaces2 + "+ START " + c.source + "\n")

	var cursor uint32
	readAddr := func(a uint16) string {
		if a == regA {
			return "$a"
		}
		return fmt.Sprintf("$%d$%d", a>>10, a&0x03ff)
	}
	readKAddr := func(a uint16) string {
		return fmt.Sprintf("k$%d(%+v)", a, c.consts[a])
	}

	oldpos := c.pos
MAIN:
	for {
		bop, a, b := op(crRead32(c.code, &cursor))
		sb.WriteString(spaces)

		if len(c.pos) > 0 {
			next, op, line, col := c.pos.readABC(0)
			// log.Println(cursor, op, unsafe.Pointer(&pos))
			for cursor > op {
				c.pos = c.pos[next:]
				if len(c.pos) == 0 {
					break
				}
				if next, op, line, col = c.pos.readABC(0); cursor <= op {
					break
				}
			}

			if op == cursor {
				x := fmt.Sprintf("%d:%d", line, col)
				sb.WriteString(fmt.Sprintf("|%-7s %d| ", x, cursor-1))
				c.pos = c.pos[next:]
			} else {
				sb.WriteString(fmt.Sprintf("|        %d| ", cursor-1))
			}
		} else {
			sb.WriteString(fmt.Sprintf("|      . %d| ", cursor-1))
		}

		switch bop {
		case OP_EOB:
			sb.WriteString("end\n")
			break MAIN
		case OP_SET:
			sb.WriteString(readAddr(a) + " = " + readAddr(b))
		case OP_SETK:
			sb.WriteString(readAddr(a) + " = " + readKAddr(uint16(b)))
		case OP_PUSH:
			sb.WriteString("push " + readAddr(a))
		case OP_RET:
			sb.WriteString("ret " + readAddr(a))
		case OP_YIELD:
			sb.WriteString("yield " + readAddr(a))
		case OP_LAMBDA:
			sb.WriteString("$a = closure:\n")
			cls := crReadClosure(c.code, &cursor, nil, a, b)
			sb.WriteString(cls.crPrettify(tab + 1))
		case OP_CALL:
			sb.WriteString("call " + readAddr(a))
			if b > 0 {
				sb.WriteString(" -> r" + strconv.Itoa(int(b)-1))
			}
		case OP_JMP:
			pos := int32(b) - 1<<12
			pos2 := uint32(int32(cursor) + pos)
			sb.WriteString("jmp " + strconv.Itoa(int(pos)) + " to " + strconv.Itoa(int(pos2)))
		case OP_IF, OP_IFNOT:
			addr := readAddr(a)
			pos := int32(b) - 1<<12
			pos2 := strconv.Itoa(int(int32(cursor) + pos))
			if bop == OP_IFNOT {
				sb.WriteString("if not " + addr + " jmp " + strconv.Itoa(int(pos)) + " to " + pos2)
			} else {
				sb.WriteString("if " + addr + " jmp " + strconv.Itoa(int(pos)) + " to " + pos2)
			}
		case OP_NOP:
			sb.WriteString("nop")
		case OP_INC:
			sb.WriteString("inc " + readAddr(a) + " " + readKAddr(uint16(b)))
		case OP_MAKEMAP:
			if a == 1 {
				sb.WriteString("make-array")
			} else {
				sb.WriteString("make-map")
			}
		default:
			if bs, ok := singleOp[bop]; ok {
				sb.WriteString(bs + " " + readAddr(a) + " " + readAddr(b))
			} else {
				sb.WriteString(fmt.Sprintf("? %02x", bop))
			}
		}

		sb.WriteString("\n")
	}

	c.pos = oldpos

	sb.WriteString(spaces2 + "+ END " + c.source)
	return sb.String()
}

func crReadClosure(code []uint32, cursor *uint32, env *Env, opa, opb uint16) *Closure {
	argsCount := byte(opa)
	options := byte(opb)
	constsLen := uint16(crRead32(code, cursor))
	consts := make([]Value, constsLen)
	for i := uint16(0); i < constsLen; i++ {
		x := crRead64(code, cursor)
		if x != math.MaxUint64 {
			consts[i] = NewNumberValue(math.Float64frombits(x))
			continue
		}
		x = crRead64(code, cursor)
		if x == math.MaxUint64 {
			consts[i] = NewNumberValue(math.Float64frombits(x))
			continue
		}
		consts[i] = NewStringValue(crReadStringLen(code, int(x), cursor))
	}

	xlen := crRead64(code, cursor)
	poslen, codelen, srclen := uint32(xlen>>38), uint32(xlen<<26>>38), uint16(xlen<<52>>52)
	src := crReadStringLen(code, int(srclen), cursor)
	pos := crReadBytesLen(code, int(poslen), cursor)
	clscode := crRead(code, cursor, int(codelen))
	cls := NewClosure(clscode, consts, env, byte(argsCount))
	cls.pos = posVByte(pos)
	cls.options = options
	cls.source = src
	return cls
}
