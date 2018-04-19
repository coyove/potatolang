package base

const REG_A = -1

const (
	OP_ASSERT = iota
	OP_TYPEOF
	OP_NIL
	OP_TRUE
	OP_FALSE
	OP_LIST
	OP_MAP
	OP_BYTES
	OP_LEN
	OP_EXPAND
	OP_SET
	OP_SET_NUM
	OP_SET_STR
	OP_STORE
	OP_LOAD
	OP_INC
	OP_INC_NUM

	OP_ADD
	OP_SUB
	OP_MUL
	OP_DIV
	OP_MOD
	OP_EQ
	OP_NEQ
	OP_LESS
	OP_LESS_EQ
	OP_MORE
	OP_MORE_EQ
	OP_NOT
	OP_AND
	OP_OR
	OP_XOR
	OP_BIT_NOT
	OP_BIT_AND
	OP_BIT_OR
	OP_BIT_XOR
	OP_BIT_LSH
	OP_BIT_RSH

	OP_IF
	OP_JMP
	OP_LAMBDA
	OP_CALL
	OP_PUSHF
	OP_PUSHF_NUM
	OP_PUSHF_STR
	OP_PUSH
	OP_PUSH_NUM
	OP_PUSH_STR
	OP_RET
	OP_RET_NUM
	OP_RET_STR
	OP_LIB_CALL
	OP_LIB_CALL_EX

	OP_EXT_IMM_F = 0xF9
	OP_EXT_F_IMM = 0xFA
	OP_EXT_F_F   = 0xFB
	OP_NOP       = 0xFE
	OP_EOB       = 0xFF
)
