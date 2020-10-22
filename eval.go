package potatolang

import (
	"bytes"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"sync/atomic"
	"unsafe"
)

type stacktrace struct {
	cursor      uint32
	stackOffset uint32
	cls         *Func
}

// ExecError represents the runtime error
type ExecError struct {
	r      interface{}
	stacks []stacktrace
}

func (e *ExecError) Error() string {
	msg := bytes.Buffer{}
	msg.WriteString("stacktrace:\n")
	for i := len(e.stacks) - 1; i >= 0; i-- {
		r := e.stacks[i]
		src := "<unknown>"
		for i := 0; i < len(r.cls.Pos); {
			var op, line uint32
			var opx uint32 = math.MaxUint32
			i, op, line = r.cls.Pos.read(i)
			if i < len(r.cls.Pos)-1 {
				_, opx, _ = r.cls.Pos.read(i)
			}
			if r.cursor >= op && r.cursor < opx {
				src = fmt.Sprintf("%s:%d", r.cls.Source, line)
				break
			}
		}
		// the recorded cursor was advanced by 1 already
		msg.WriteString(fmt.Sprintf("at %d in %s\n", r.cursor-1, src))
	}
	msg.WriteString("root panic:\n")
	msg.WriteString(fmt.Sprintf("%v\n", e.r))
	return msg.String()
}

func kodeaddr(code []uint32) uintptr { return (*reflect.SliceHeader)(unsafe.Pointer(&code)).Data }

func konstaddr(consts []Value) uintptr { return (*reflect.SliceHeader)(unsafe.Pointer(&consts)).Data }

func returnVararg(a Value, b []Value) (Value, []Value) {
	flag := a.Type() == STK
	if len(b) == 0 {
		if flag {
			u := a.unpackedStack().a
			if len(u) == 0 {
				return Value{}, nil
			}
			return u[0], u[1:]
		}
		return a, nil
	}

	for _, b := range b {
		flag = flag || b.Type() == STK
	}

	if !flag {
		// both 'a' and 'b' are not (neither containing) unpacked values
		return a, b
	}

	var b2 []Value
	if a.Type() == STK {
		b2 = append(b2, a.unpackedStack().a...)
	} else {
		b2 = append(b2, a)
	}
	for _, b := range b {
		if b.Type() == STK {
			b2 = append(b2, b.unpackedStack().a...)
		} else {
			b2 = append(b2, b)
		}
	}
	if len(b2) == 0 {
		return Value{}, nil
	}
	return b2[0], b2[1:]
}

// execCursorLoop executes 'K' under 'Env' from the given start 'cursor'
func execCursorLoop(env Env, K *Func, cursor uint32) (result Value, resultV []Value, nextCursor uint32, yielded bool) {
	var stackEnv = env
	var retStack []stacktrace

	stackEnv.stackOffset = len(*env.stack)

	defer func() {
		if r := recover(); r != nil {
			rr := stacktrace{
				cursor: cursor,
				cls:    K,
			}

			if re, ok := r.(*ExecError); ok {
				retStack = append(retStack, rr)
				re.stacks = append(retStack, re.stacks...)
				panic(re)
			} else {
				e := &ExecError{}
				e.r = r // root panic
				e.stacks = make([]stacktrace, len(retStack)+1)
				copy(e.stacks, retStack)
				e.stacks[len(e.stacks)-1] = rr
				panic(e)
			}
		}
	}()

	returnUpperWorld := func(v Value) {
		r := retStack[len(retStack)-1]
		cursor = r.cursor
		K = r.cls

		env.stackOffset = int(r.stackOffset)
		env.A, env.V = returnVararg(v, env.V)
		*env.stack = (*env.stack)[:env.stackOffset+int(r.cls.stackSize)]
		stackEnv.stackOffset = len(*env.stack)
		retStack = retStack[:len(retStack)-1]
	}

MAIN:
	for {
		if env.global.Deadline != 0 {
			if atomic.LoadInt64(&now) > env.global.Deadline {
				panicf("timeout")
			}
		}

		v := K.Code[cursor]
		cursor++
		bop, opa, opb := op(v)

		switch bop {
		case OpEOB:
			break MAIN
		case OpSet:
			env._set(opa, env._get(opb, K))
		case OpPushV:
			if opb != 0 {
				env.V = make([]Value, 0, opb)
			}
			env.V = append(env.V, env._get(opa, K))
		case OpPopVClear:
			env.V = nil
		case OpPopVAll:
			if opa == 1 {
				// popv-all-with-a, e.g.: local ... = foo()
				env.A = unpackedStack(&unpacked{append([]Value{env.A}, env.V...)})
			} else {
				// popv-all, e.g.: local a, ... = foo()
				env.A = unpackedStack(&unpacked{env.V})
			}
			env.V = nil
		case OpPopV:
			if len(env.V) == 0 {
				env.A = Value{}
			} else {
				env.A, env.V = env.V[0], env.V[1:]
			}
		case OpInc:
			vaf, vai, vaIsInt := env._get(opa, K).Expect(NUM).Num()
			vbf, vbi, vbIsInt := env._get(opb, K).Expect(NUM).Num()
			if vaIsInt && vbIsInt {
				env.A = Int(vai + vbi)
			} else {
				env.A = Num(vaf + vbf)
			}
			env._set(opa, env.A)
		case OpConcat:
			var x string
			switch va, vb := env._get(opa, K), env._get(opb, K); va.Type() + vb.Type() {
			case StrStr:
				x = va.Str() + vb.Str()
			default:
				if va.Type() == STR && vb.Type() == NUM {
					vbf, vbi, vbIsInt := vb.Num()
					if vbIsInt {
						x = va.Str() + strconv.FormatInt(vbi, 10)
					} else {
						x = va.Str() + strconv.FormatFloat(vbf, 'f', 0, 64)
					}
				} else {
					va, vb = va.Expect(STR), vb.Expect(STR)
				}
			}
			if env.global.MaxStringSize > 0 && int64(len(x)) > env.global.MaxStringSize {
				panicf("string overflow, max: %d", env.global.MaxStringSize)
			}
			env.A = Str(x)
		case OpAdd:
			switch va, vb := env._get(opa, K), env._get(opb, K); va.Type() + vb.Type() {
			case NumNum:
				vaf, vai, vaIsInt := va.Num()
				vbf, vbi, vbIsInt := vb.Num()
				if vaIsInt && vbIsInt {
					env.A = Int(vai + vbi)
				} else {
					env.A = Num(vaf + vbf)
				}
			default:
				if va.Type() == NUM && vb.Type() == STR {
					vaf, vai, vaIsInt := va.Num()
					if vaIsInt {
						vbi, _ := strconv.ParseInt(vb.Str(), 0, 64)
						env.A = Int(vai + vbi)
					} else {
						vbf, _ := strconv.ParseFloat(vb.Str(), 64)
						env.A = Num(vaf + vbf)
					}
				} else {
					va, vb = va.Expect(NUM), vb.Expect(NUM)
				}
			}
		case OpSub:
			switch va, vb := env._get(opa, K), env._get(opb, K); va.Type() + vb.Type() {
			case NumNum:
				vaf, vai, vaIsInt := va.Num()
				vbf, vbi, vbIsInt := vb.Num()
				if vaIsInt && vbIsInt {
					env.A = Int(vai - vbi)
				} else {
					env.A = Num(vaf - vbf)
				}
			default:
				va, vb = va.Expect(NUM), vb.Expect(NUM)
			}
		case OpMul:
			switch va, vb := env._get(opa, K), env._get(opb, K); va.Type() + vb.Type() {
			case NumNum:
				vaf, vai, vaIsInt := va.Num()
				vbf, vbi, vbIsInt := vb.Num()
				if vaIsInt && vbIsInt {
					env.A = Int(vai * vbi)
				} else {
					env.A = Num(vaf * vbf)
				}
			default:
				va, vb = va.Expect(NUM), vb.Expect(NUM)
			}
		case OpDiv:
			switch va, vb := env._get(opa, K), env._get(opb, K); va.Type() + vb.Type() {
			case NumNum:
				vaf, vai, vaIsInt := va.Num()
				vbf, vbi, vbIsInt := vb.Num()
				if vaIsInt && vbIsInt {
					env.A = Int(vai / vbi)
				} else {
					env.A = Num(vaf / vbf)
				}
			default:
				va, vb = va.Expect(NUM), vb.Expect(NUM)
			}
		case OpMod:
			switch va, vb := env._get(opa, K), env._get(opb, K); va.Type() + vb.Type() {
			case NumNum:
				vaf, vai, vaIsInt := va.Num()
				vbf, vbi, vbIsInt := vb.Num()
				if vaIsInt && vbIsInt {
					env.A = Int(vai % vbi)
				} else {
					env.A = Num(math.Remainder(vaf, vbf))
				}
			default:
				va, vb = va.Expect(NUM), vb.Expect(NUM)
			}
		case OpEq:
			env.A = NumBool(env._get(opa, K).Equal(env._get(opb, K)))
		case OpNeq:
			env.A = NumBool(!env._get(opa, K).Equal(env._get(opb, K)))
		case OpLess:
			switch va, vb := env._get(opa, K), env._get(opb, K); va.Type() + vb.Type() {
			case NumNum:
				vaf, vai, vaIsInt := va.Num()
				vbf, vbi, vbIsInt := vb.Num()
				if vaIsInt && vbIsInt {
					env.A = NumBool(vai < vbi)
				} else {
					env.A = NumBool(vaf < vbf)
				}
			case StrStr:
				env.A = NumBool(va.Str() < vb.Str())
			default:
				va, vb = va.Expect(NUM), vb.Expect(NUM)
			}
		case OpLessEq:
			switch va, vb := env._get(opa, K), env._get(opb, K); va.Type() + vb.Type() {
			case NumNum:
				vaf, vai, vaIsInt := va.Num()
				vbf, vbi, vbIsInt := vb.Num()
				if vaIsInt && vbIsInt {
					env.A = NumBool(vai <= vbi)
				} else {
					env.A = NumBool(vaf <= vbf)
				}
			case StrStr:
				env.A = NumBool(va.Str() <= vb.Str())
			default:
				va, vb = va.Expect(NUM), vb.Expect(NUM)
			}
		case OpNot:
			env.A = NumBool(env._get(opa, K).IsFalse())
		case OpPow:
			switch va, vb := env._get(opa, K), env._get(opb, K); va.Type() + vb.Type() {
			case NumNum:
				vaf, vai, vaIsInt := va.Num()
				vbf, vbi, vbIsInt := vb.Num()
				if vaIsInt && vbIsInt && vbi >= 1 {
					env.A = Int(ipow(vai, vbi))
				} else {
					env.A = Num(math.Pow(vaf, vbf))
				}
			default:
				va, vb = va.Expect(NUM), vb.Expect(NUM)
			}
		case OpLen:
			switch v := env._get(opa, K); v.Type() {
			case STR:
				env.A = Num(float64(len(v.Str())))
			case STK:
				env.A = Num(float64(v.unpackedStack().Len()))
			case FUN:
				env.A = Num(float64(v.Fun().NumParam))
			default:
				env.A = Int(int64(reflectLen(v.Any())))
			}
		case OpStore:
			subject, v := env._get(opa, K), env._get(opb, K)
			switch subject.Type() {
			case STK:
				subject.unpackedStack().Put(env.A.ExpectMsg(NUM, "store").Int(), v)
			case ANY:
				reflectStore(subject.Any(), env.A, v)
			default:
				subject = subject.Expect(STK)
			}
			env.A = v
		case OpLoad:
			switch a := env._get(opa, K); a.Type() {
			case STK:
				env.A = a.unpackedStack().Get(env._get(opb, K).ExpectMsg(NUM, "load").Int())
			case ANY:
				env.A = reflectLoad(a.Any(), env._get(opb, K))
			case STR:
				if idx, s := env._get(opb, K).ExpectMsg(NUM, "load").Int(), a.Str(); idx >= 1 && idx <= int64(len(s)) {
					env.A = Int(int64(s[idx]))
				}
			default:
				a = a.Expect(STK)
			}
		case OpPush:
			if v := env._get(opa, K); v.Type() == STK {
				*stackEnv.stack = append(*stackEnv.stack, v.unpackedStack().a...)
			} else {
				stackEnv.Push(v)
			}
			if opa == regA && len(env.V) > 0 {
				*stackEnv.stack = append(*stackEnv.stack, env.V...)
			}
			if env.global.MaxStackSize > 0 && int64(len(*stackEnv.stack)) > env.global.MaxStackSize {
				panicf("stack overflow, max: %d", env.global.MaxStackSize)
			}
		case OpRet:
			v := env._get(opa, K)
			if len(retStack) == 0 {
				v, env.V = returnVararg(v, env.V)
				return v, env.V, 0, false
			}
			returnUpperWorld(v)
		case OpYield:
			v := env._get(opa, K)
			v, env.V = returnVararg(v, env.V)
			return v, env.V, cursor, true
		case OpLoadFunc:
			env.A = Fun(K.Funcs[opa])
		case OpCall:
			cls := env._get(opa, K).ExpectMsg(FUN, "call").Fun()
			if cls.yEnv.stack != nil { // resume yielded coroutine
				env.A, env.V = cls.exec(Env{})
				stackEnv.Clear()
			} else {
				if cls.Is(FuncVararg) && cls.native == nil {
					var varg []Value
					if stackEnv.Size() > int(cls.NumParam) {
						varg = append([]Value{}, stackEnv.Stack()[cls.NumParam:]...)
					}
					if stackEnv.Size() <= int(cls.NumParam) {
						stackEnv.grow(int(cls.NumParam) + 1)
					}
					stackEnv._set(uint16(cls.NumParam), unpackedStack(&unpacked{a: varg}))
				}

				if env.global == nil {
					panic("nil global")
				}
				stackEnv.global = env.global

				if cls.Is(FuncYield) {
					x := stackEnv
					tmp := append([]Value{}, x.Stack()...)
					stackEnv.Clear()
					x.stack = &tmp
					x.stackOffset = 0
					x.grow(int(cls.stackSize))
					env.A, env.V = cls.exec(x)
				} else if cls.native != nil {
					env.A, env.V = cls.exec(stackEnv)
					stackEnv.Clear()
				} else {
					last := stacktrace{
						cls:         K,
						cursor:      cursor,
						stackOffset: uint32(env.stackOffset),
					}

					// Switch to the Env of cls
					cursor = 0
					K = cls
					env.stackOffset = stackEnv.stackOffset
					env.global = stackEnv.global

					if opb == 0 {
						retStack = append(retStack, last)
					}

					if cls.stackSize > 0 {
						env.grow(int(cls.stackSize))
					}

					stackEnv.stackOffset = len(*env.stack)
				}
			}
		case OpJmp:
			cursor = uint32(int32(cursor) + int32(opb) - 1<<12)
		case OpIfNot:
			if cond := env._get(opa, K); cond.IsFalse() {
				cursor = uint32(int32(cursor) + int32(opb) - 1<<12)
			}
		case OpIf:
			if cond := env._get(opa, K); !cond.IsFalse() {
				cursor = uint32(int32(cursor) + int32(opb) - 1<<12)
			}
		}
	}

	if len(retStack) > 0 {
		returnUpperWorld(Value{})
		goto MAIN
	}
	return Value{}, nil, 0, false
}
