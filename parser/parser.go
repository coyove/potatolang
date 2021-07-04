// Code generated by goyacc -o parser.go parser.go.y. DO NOT EDIT.

//line parser.go.y:2
package parser

import __yyfmt__ "fmt"

//line parser.go.y:2

//line parser.go.y:25
type yySymType struct {
	yys   int
	token Token
	expr  Node
}

const TDo = 57346
const TLocal = 57347
const TElseIf = 57348
const TThen = 57349
const TEnd = 57350
const TBreak = 57351
const TElse = 57352
const TFor = 57353
const TWhile = 57354
const TFunc = 57355
const TIf = 57356
const TReturn = 57357
const TReturnVoid = 57358
const TRepeat = 57359
const TUntil = 57360
const TNot = 57361
const TLabel = 57362
const TGoto = 57363
const TIn = 57364
const TLsh = 57365
const TRsh = 57366
const TURsh = 57367
const TOr = 57368
const TAnd = 57369
const TEqeq = 57370
const TNeq = 57371
const TLte = 57372
const TGte = 57373
const TIdent = 57374
const TNumber = 57375
const TString = 57376
const TIDiv = 57377
const ASSIGN = 57378
const FUNC = 57379
const UNARY = 57380

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"TDo",
	"TLocal",
	"TElseIf",
	"TThen",
	"TEnd",
	"TBreak",
	"TElse",
	"TFor",
	"TWhile",
	"TFunc",
	"TIf",
	"TReturn",
	"TReturnVoid",
	"TRepeat",
	"TUntil",
	"TNot",
	"TLabel",
	"TGoto",
	"TIn",
	"TLsh",
	"TRsh",
	"TURsh",
	"TOr",
	"TAnd",
	"TEqeq",
	"TNeq",
	"TLte",
	"TGte",
	"TIdent",
	"TNumber",
	"TString",
	"'{'",
	"'['",
	"'('",
	"'='",
	"'>'",
	"'<'",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"'^'",
	"'#'",
	"'.'",
	"'&'",
	"'|'",
	"'~'",
	"TIDiv",
	"'T'",
	"ASSIGN",
	"FUNC",
	"UNARY",
	"';'",
	"','",
	"')'",
	"']'",
	"'}'",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.go.y:350

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 24,
	38, 46,
	58, 46,
	-2, 80,
	-1, 100,
	38, 47,
	58, 47,
	-2, 80,
}

const yyPrivate = 57344

const yyLast = 1124

var yyAct = [...]int{
	30, 31, 166, 16, 131, 152, 86, 29, 151, 46,
	45, 98, 192, 98, 162, 46, 186, 159, 157, 134,
	97, 50, 62, 139, 53, 47, 132, 196, 16, 171,
	169, 135, 42, 83, 40, 44, 177, 89, 90, 91,
	98, 92, 143, 85, 106, 48, 41, 73, 74, 76,
	101, 96, 95, 16, 163, 99, 75, 153, 201, 145,
	191, 154, 176, 109, 110, 111, 112, 113, 114, 115,
	116, 117, 118, 119, 120, 121, 122, 123, 124, 125,
	126, 127, 128, 129, 161, 39, 102, 24, 142, 136,
	25, 133, 71, 72, 73, 74, 76, 103, 93, 52,
	138, 140, 49, 75, 46, 141, 147, 148, 28, 27,
	16, 61, 24, 144, 180, 59, 17, 168, 6, 26,
	9, 167, 22, 20, 15, 23, 13, 12, 21, 14,
	43, 11, 10, 109, 100, 18, 155, 24, 54, 2,
	109, 51, 57, 25, 4, 56, 158, 3, 16, 1,
	58, 16, 5, 0, 0, 170, 0, 0, 0, 0,
	0, 0, 16, 172, 0, 0, 179, 0, 60, 182,
	183, 0, 185, 0, 178, 0, 16, 16, 0, 108,
	0, 0, 0, 16, 0, 0, 0, 0, 0, 0,
	0, 16, 16, 0, 24, 203, 0, 205, 0, 0,
	0, 16, 16, 0, 16, 37, 16, 211, 0, 0,
	16, 0, 0, 0, 0, 16, 0, 0, 87, 33,
	34, 35, 88, 32, 146, 0, 0, 149, 38, 0,
	0, 0, 24, 0, 0, 24, 0, 36, 0, 0,
	59, 17, 0, 0, 214, 9, 24, 22, 20, 0,
	23, 13, 12, 21, 0, 0, 11, 10, 0, 0,
	24, 24, 0, 160, 0, 0, 0, 24, 25, 0,
	0, 0, 0, 0, 0, 24, 24, 0, 0, 0,
	0, 174, 175, 0, 0, 24, 24, 181, 24, 0,
	24, 37, 0, 60, 24, 0, 189, 190, 0, 24,
	0, 0, 0, 0, 25, 33, 34, 35, 0, 32,
	0, 199, 200, 202, 38, 204, 0, 0, 0, 0,
	0, 208, 0, 36, 0, 0, 0, 0, 0, 0,
	213, 80, 81, 82, 63, 64, 69, 70, 68, 67,
	0, 0, 0, 0, 0, 0, 0, 65, 66, 71,
	72, 73, 74, 76, 79, 0, 0, 77, 78, 0,
	75, 0, 0, 0, 0, 0, 0, 0, 184, 80,
	81, 82, 63, 64, 69, 70, 68, 67, 0, 0,
	0, 0, 0, 0, 0, 65, 66, 71, 72, 73,
	74, 76, 79, 0, 0, 77, 78, 0, 75, 0,
	0, 0, 0, 0, 0, 0, 156, 80, 81, 82,
	63, 64, 69, 70, 68, 67, 0, 0, 0, 80,
	81, 82, 0, 65, 66, 71, 72, 73, 74, 76,
	79, 193, 0, 77, 78, 0, 75, 71, 72, 73,
	74, 76, 79, 0, 137, 77, 78, 0, 75, 0,
	80, 81, 82, 63, 64, 69, 70, 68, 67, 0,
	0, 0, 0, 0, 0, 0, 65, 66, 71, 72,
	73, 74, 76, 79, 0, 0, 77, 78, 0, 75,
	0, 0, 0, 0, 0, 194, 80, 81, 82, 63,
	64, 69, 70, 68, 67, 0, 0, 0, 0, 0,
	0, 0, 65, 66, 71, 72, 73, 74, 76, 79,
	0, 0, 77, 78, 0, 75, 0, 0, 0, 0,
	0, 0, 130, 80, 81, 82, 63, 64, 69, 70,
	68, 67, 0, 0, 0, 80, 81, 82, 0, 65,
	66, 71, 72, 73, 74, 76, 79, 210, 0, 77,
	78, 0, 75, 71, 72, 73, 74, 76, 165, 0,
	0, 0, 0, 0, 75, 0, 80, 81, 82, 63,
	64, 69, 70, 68, 67, 0, 0, 0, 0, 0,
	195, 0, 65, 66, 71, 72, 73, 74, 76, 79,
	0, 0, 77, 78, 0, 75, 80, 81, 82, 63,
	64, 69, 70, 68, 67, 0, 0, 0, 0, 0,
	107, 0, 65, 66, 71, 72, 73, 74, 76, 79,
	0, 0, 77, 78, 0, 75, 80, 81, 82, 63,
	64, 69, 70, 68, 67, 0, 0, 104, 0, 0,
	0, 0, 65, 66, 71, 72, 73, 74, 76, 79,
	0, 0, 77, 78, 0, 75, 80, 81, 82, 63,
	64, 69, 70, 68, 67, 0, 0, 0, 0, 0,
	0, 0, 65, 66, 71, 72, 73, 74, 76, 79,
	0, 0, 77, 78, 0, 75, 80, 81, 82, 63,
	64, 69, 70, 68, 67, 0, 0, 0, 0, 0,
	0, 0, 65, 66, 71, 72, 73, 74, 76, 79,
	0, 0, 77, 78, 0, 75, 59, 17, 0, 0,
	212, 9, 0, 22, 20, 0, 23, 13, 12, 21,
	59, 17, 11, 10, 209, 9, 0, 22, 20, 0,
	23, 13, 12, 21, 25, 0, 11, 10, 0, 59,
	17, 0, 0, 207, 9, 0, 22, 20, 25, 23,
	13, 12, 21, 59, 17, 11, 10, 206, 9, 60,
	22, 20, 0, 23, 13, 12, 21, 25, 0, 11,
	10, 0, 0, 60, 0, 0, 0, 0, 0, 0,
	0, 25, 0, 59, 17, 0, 0, 198, 9, 0,
	22, 20, 60, 23, 13, 12, 21, 59, 17, 11,
	10, 197, 9, 0, 22, 20, 60, 23, 13, 12,
	21, 25, 0, 11, 10, 0, 59, 17, 0, 0,
	188, 9, 0, 22, 20, 25, 23, 13, 12, 21,
	59, 17, 11, 10, 187, 9, 60, 22, 20, 0,
	23, 13, 12, 21, 25, 0, 11, 10, 0, 0,
	60, 0, 0, 0, 0, 0, 0, 0, 25, 0,
	59, 17, 0, 0, 173, 9, 0, 22, 20, 60,
	23, 13, 12, 21, 59, 17, 11, 10, 164, 9,
	0, 22, 20, 60, 23, 13, 12, 21, 25, 0,
	11, 10, 0, 0, 0, 0, 0, 0, 0, 80,
	81, 82, 25, 64, 69, 70, 68, 67, 0, 0,
	0, 0, 0, 60, 0, 65, 66, 71, 72, 73,
	74, 76, 79, 0, 0, 77, 78, 60, 75, 59,
	17, 0, 0, 150, 9, 0, 22, 20, 0, 23,
	13, 12, 21, 0, 0, 11, 10, 0, 59, 17,
	0, 0, 0, 9, 37, 22, 20, 25, 23, 13,
	12, 21, 105, 0, 11, 10, 0, 87, 33, 34,
	35, 88, 32, 0, 0, 0, 25, 38, 0, 0,
	0, 0, 60, 0, 0, 0, 36, 59, 17, 0,
	0, 55, 9, 0, 22, 20, 84, 23, 13, 12,
	21, 60, 0, 11, 10, 0, 0, 0, 0, 7,
	17, 0, 0, 0, 9, 25, 22, 20, 19, 23,
	13, 12, 21, 0, 0, 11, 10, 0, 0, 0,
	0, 59, 17, 0, 0, 0, 9, 25, 22, 20,
	60, 23, 13, 12, 21, 0, 0, 11, 10, 80,
	81, 82, 0, 0, 69, 70, 68, 67, 0, 25,
	0, 0, 8, 0, 0, 65, 66, 71, 72, 73,
	74, 76, 79, 37, 0, 77, 78, 0, 75, 0,
	0, 0, 0, 0, 60, 0, 87, 33, 34, 35,
	88, 32, 0, 0, 0, 0, 38, 0, 0, 0,
	0, 0, 0, 0, 0, 36, 0, 0, 0, 0,
	0, 0, 0, 94,
}

var yyPact = [...]int{
	-1000, 1015, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	77, 76, -1000, 272, -1000, -1000, -2, 72, -13, 70,
	272, -1000, 67, 272, -1000, -1000, 993, -1000, 91, -36,
	663, -2, 272, -1000, -1000, 945, 272, 272, 272, -1000,
	272, 66, -1000, -1000, 1064, -18, -1000, 272, 58, 49,
	633, 954, 6, 603, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 272, 272, 272, 272, 272, 272, 272, 272,
	272, 272, 272, 272, 272, 272, 272, 272, 272, 272,
	272, 272, 272, 463, -1000, -32, -39, -7, 272, -1000,
	-1000, -1000, 384, -1000, -1000, -35, -39, 272, 56, -36,
	-1000, -2, -17, 27, -1000, 272, 272, -1000, 935, 663,
	886, 1036, 396, 396, 396, 396, 396, 396, 4, 4,
	-1000, -1000, -1000, -1000, 512, 512, 512, 51, 51, 51,
	-1000, -53, 272, -56, 25, 272, 346, -1000, -41, 186,
	-42, -36, -1000, 50, -45, 17, 880, 663, 500, 111,
	-1000, -1000, -1000, -8, 272, 663, -9, -1000, -39, -1000,
	866, -1000, 28, -23, -1000, 272, 106, -1000, 272, 272,
	308, 272, -43, -1000, 836, 822, -1000, 26, -47, 427,
	-1000, 1037, 573, 663, -11, 663, -1000, -1000, -1000, 803,
	789, -1000, 24, -1000, 272, -1000, 272, -1000, -1000, 759,
	745, -1000, 726, 543, 111, 663, -1000, -1000, 712, -1000,
	-1000, -1000, -1000, 236, -1000,
}

var yyPgo = [...]int{
	0, 149, 119, 139, 138, 85, 135, 10, 0, 7,
	6, 1, 130, 150, 129, 124, 2, 145, 142, 118,
	4,
}

var yyR1 = [...]int{
	0, 1, 1, 2, 2, 3, 3, 3, 3, 3,
	3, 4, 4, 4, 4, 4, 18, 18, 13, 13,
	13, 13, 14, 14, 14, 14, 15, 16, 16, 16,
	19, 19, 19, 19, 19, 19, 19, 19, 17, 17,
	17, 17, 17, 5, 5, 5, 6, 6, 7, 7,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	11, 11, 11, 12, 12, 12, 12, 9, 9, 10,
	10, 10, 10, 20, 20,
}

var yyR2 = [...]int{
	0, 0, 2, 0, 2, 1, 1, 1, 1, 3,
	1, 1, 1, 1, 3, 1, 1, 1, 1, 2,
	4, 3, 5, 4, 9, 11, 6, 0, 2, 5,
	6, 7, 7, 8, 8, 9, 9, 10, 1, 2,
	3, 1, 2, 1, 4, 3, 1, 3, 1, 3,
	1, 3, 1, 1, 2, 4, 4, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 2, 2, 2,
	1, 2, 2, 2, 4, 4, 6, 1, 3, 3,
	5, 5, 7, 0, 1,
}

var yyChk = [...]int{
	-1000, -1, -3, -17, -18, -13, -19, 4, 57, 9,
	21, 20, 16, 15, -14, -15, -11, 5, -6, 13,
	12, 17, 11, 14, -5, 32, -2, 32, 32, -9,
	-8, -11, 37, 33, 34, 35, 51, 19, 42, -5,
	36, 48, 34, -12, 37, -7, 32, 38, 58, 32,
	-8, -2, 32, -8, -4, 8, -17, -18, -13, 4,
	57, 20, 58, 26, 27, 39, 40, 31, 30, 28,
	29, 41, 42, 43, 44, 52, 45, 49, 50, 46,
	23, 24, 25, -8, 61, -9, -10, 32, 36, -8,
	-8, -8, -8, 32, 59, -9, -10, 38, 58, -9,
	-5, -11, 37, 48, 4, 18, 38, 7, -2, -8,
	-8, -8, -8, -8, -8, -8, -8, -8, -8, -8,
	-8, -8, -8, -8, -8, -8, -8, -8, -8, -8,
	59, -20, 58, -20, 58, 38, -8, 60, -20, 58,
	-20, -9, 32, 59, -7, 32, -2, -8, -8, -2,
	8, 61, 61, 32, 36, -8, 60, 59, -10, 59,
	-2, 34, 59, 37, 8, 58, -16, 10, 6, 38,
	-8, 38, -20, 8, -2, -2, 34, 59, -7, -8,
	8, -2, -8, -8, 60, -8, 59, 8, 8, -2,
	-2, 34, 59, 4, 58, 7, 38, 8, 8, -2,
	-2, 34, -2, -8, -2, -8, 8, 8, -2, 8,
	4, -16, 8, -2, 8,
}

var yyDef = [...]int{
	1, -2, 2, 5, 6, 7, 8, 3, 10, 38,
	0, 0, 41, 0, 16, 17, 18, 0, 0, 0,
	0, 3, 0, 0, -2, 43, 0, 39, 0, 42,
	87, 50, 0, 52, 53, 0, 0, 0, 0, 80,
	0, 0, 81, 82, 0, 19, 48, 0, 0, 0,
	0, 0, 0, 0, 4, 9, 11, 12, 13, 3,
	15, 40, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 54, 93, 93, 43, 0, 77,
	78, 79, 0, 45, 83, 93, 93, 0, 0, 21,
	-2, 0, 0, 0, 3, 0, 0, 3, 0, 88,
	57, 58, 59, 60, 61, 62, 63, 64, 65, 66,
	67, 68, 69, 70, 71, 72, 73, 74, 75, 76,
	51, 0, 94, 0, 94, 0, 0, 44, 0, 94,
	0, 20, 49, 3, 0, 0, 0, 23, 0, 27,
	14, 55, 56, 0, 0, 89, 0, 84, 93, 85,
	0, 3, 3, 0, 22, 0, 0, 3, 0, 0,
	0, 0, 0, 30, 0, 0, 3, 3, 0, 0,
	26, 28, 0, 91, 0, 90, 86, 32, 31, 0,
	0, 3, 3, 3, 0, 3, 0, 33, 34, 0,
	0, 3, 0, 0, 27, 92, 36, 35, 0, 24,
	3, 29, 37, 0, 25,
}

var yyTok1 = [...]int{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 47, 3, 45, 49, 3,
	37, 59, 43, 41, 58, 42, 48, 44, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 57,
	40, 38, 39, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 53, 3, 3, 3, 3, 3,
	3, 36, 3, 60, 46, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 35, 50, 61, 51,
}

var yyTok2 = [...]int{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 52, 54, 55, 56,
}

var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:54
		{
			yyVAL.expr = __chain()
			if l, ok := yylex.(*Lexer); ok {
				l.Stmts = yyVAL.expr
			}
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:60
		{
			yyVAL.expr = yyDollar[1].expr.append(yyDollar[2].expr)
			if l, ok := yylex.(*Lexer); ok {
				l.Stmts = yyVAL.expr
			}
		}
	case 3:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:68
		{
			yyVAL.expr = __chain()
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:71
		{
			yyVAL.expr = yyDollar[1].expr.append(yyDollar[2].expr)
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:76
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:77
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:78
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:79
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:80
		{
			yyVAL.expr = __do(yyDollar[2].expr)
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:81
		{
			yyVAL.expr = emptyNode
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:84
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:85
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:86
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:87
		{
			yyVAL.expr = __do(yyDollar[2].expr)
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:88
		{
			yyVAL.expr = emptyNode
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:91
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:92
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:95
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 19:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:98
		{
			yyVAL.expr = __chain()
			for _, v := range yyDollar[2].expr.Nodes {
				yyVAL.expr = yyVAL.expr.append(__set(v, NewSymbol(ANil)).SetPos(yyDollar[1].token.Pos))
			}
		}
	case 20:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:104
		{
			if len(yyDollar[4].expr.Nodes) == 1 && len(yyDollar[2].expr.Nodes) > 1 {
				tmp := randomVarname()
				yyVAL.expr = __chain(__local([]Node{tmp}, yyDollar[4].expr.Nodes, yyDollar[1].token.Pos))
				for i, ident := range yyDollar[2].expr.Nodes {
					yyVAL.expr = yyVAL.expr.append(__local([]Node{ident}, []Node{__load(tmp, NewNumberFromInt(int64(i))).SetPos(yyDollar[1].token.Pos)}, yyDollar[1].token.Pos))
				}
			} else {
				yyVAL.expr = __local(yyDollar[2].expr.Nodes, yyDollar[4].expr.Nodes, yyDollar[1].token.Pos)
			}
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:115
		{
			if len(yyDollar[3].expr.Nodes) == 1 && len(yyDollar[1].expr.Nodes) > 1 {
				tmp := randomVarname()
				yyVAL.expr = __chain(__local([]Node{tmp}, yyDollar[3].expr.Nodes, yyDollar[2].token.Pos))
				for i, decl := range yyDollar[1].expr.Nodes {
					x := decl.moveLoadStore(__move, __load(tmp, NewNumberFromInt(int64(i))).SetPos(yyDollar[2].token.Pos)).SetPos(yyDollar[2].token.Pos)
					yyVAL.expr = yyVAL.expr.append(x)
				}
			} else {
				yyVAL.expr = __moveMulti(yyDollar[1].expr.Nodes, yyDollar[3].expr.Nodes, yyDollar[2].token.Pos)
			}
		}
	case 22:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.go.y:129
		{
			yyVAL.expr = __loop(__if(yyDollar[2].expr, yyDollar[4].expr, breakNode).SetPos(yyDollar[1].token.Pos)).SetPos(yyDollar[1].token.Pos)
		}
	case 23:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:132
		{
			yyVAL.expr = __loop(
				__chain(
					yyDollar[2].expr,
					__if(yyDollar[4].expr, breakNode, emptyNode).SetPos(yyDollar[1].token.Pos),
				).SetPos(yyDollar[1].token.Pos),
			).SetPos(yyDollar[1].token.Pos)
		}
	case 24:
		yyDollar = yyS[yypt-9 : yypt+1]
//line parser.go.y:140
		{
			forVar, forEnd := NewSymbolFromToken(yyDollar[2].token), randomVarname()
			yyVAL.expr = __do(
				__set(forVar, yyDollar[4].expr).SetPos(yyDollar[1].token.Pos),
				__set(forEnd, yyDollar[6].expr).SetPos(yyDollar[1].token.Pos),
				__loop(
					__if(
						__less(forVar, forEnd),
						__chain(yyDollar[8].expr, __inc(forVar, oneNode).SetPos(yyDollar[1].token.Pos)),
						breakNode,
					).SetPos(yyDollar[1].token.Pos),
				).SetPos(yyDollar[1].token.Pos),
			)
		}
	case 25:
		yyDollar = yyS[yypt-11 : yypt+1]
//line parser.go.y:154
		{
			forVar, forEnd, forStep := NewSymbolFromToken(yyDollar[2].token), randomVarname(), randomVarname()
			body := __chain(yyDollar[10].expr, __inc(forVar, forStep))
			yyVAL.expr = __do(
				__set(forVar, yyDollar[4].expr).SetPos(yyDollar[1].token.Pos),
				__set(forEnd, yyDollar[6].expr).SetPos(yyDollar[1].token.Pos),
				__set(forStep, yyDollar[8].expr).SetPos(yyDollar[1].token.Pos))

			if yyDollar[8].expr.IsNumber() { // step is a static number, easy case
				if yyDollar[8].expr.IsNegativeNumber() {
					yyVAL.expr = yyVAL.expr.append(__loop(__if(__less(forEnd, forVar), body, breakNode).SetPos(yyDollar[1].token.Pos)).SetPos(yyDollar[1].token.Pos))
				} else {
					yyVAL.expr = yyVAL.expr.append(__loop(__if(__less(forVar, forEnd), body, breakNode).SetPos(yyDollar[1].token.Pos)).SetPos(yyDollar[1].token.Pos))
				}
			} else {
				yyVAL.expr = yyVAL.expr.append(__loop(
					__if(
						__less(zeroNode, forStep).SetPos(yyDollar[1].token.Pos),
						// +step
						__if(__lessEq(forEnd, forVar), breakNode, body).SetPos(yyDollar[1].token.Pos),
						// -step
						__if(__lessEq(forVar, forEnd), breakNode, body).SetPos(yyDollar[1].token.Pos),
					).SetPos(yyDollar[1].token.Pos),
				).SetPos(yyDollar[1].token.Pos))
			}
		}
	case 26:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.go.y:182
		{
			yyVAL.expr = __if(yyDollar[2].expr, yyDollar[4].expr, yyDollar[5].expr).SetPos(yyDollar[1].token.Pos)
		}
	case 27:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:187
		{
			yyVAL.expr = NewComplex()
		}
	case 28:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:190
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 29:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.go.y:193
		{
			yyVAL.expr = __if(yyDollar[2].expr, yyDollar[4].expr, yyDollar[5].expr).SetPos(yyDollar[1].token.Pos)
		}
	case 30:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.go.y:198
		{
			yyVAL.expr = __func(yyDollar[2].token, emptyNode, "", yyDollar[5].expr)
		}
	case 31:
		yyDollar = yyS[yypt-7 : yypt+1]
//line parser.go.y:199
		{
			yyVAL.expr = __func(yyDollar[2].token, yyDollar[4].expr, "", yyDollar[6].expr)
		}
	case 32:
		yyDollar = yyS[yypt-7 : yypt+1]
//line parser.go.y:200
		{
			yyVAL.expr = __func(yyDollar[2].token, emptyNode, yyDollar[5].token.Str, yyDollar[6].expr)
		}
	case 33:
		yyDollar = yyS[yypt-8 : yypt+1]
//line parser.go.y:201
		{
			yyVAL.expr = __func(yyDollar[2].token, yyDollar[4].expr, yyDollar[6].token.Str, yyDollar[7].expr)
		}
	case 34:
		yyDollar = yyS[yypt-8 : yypt+1]
//line parser.go.y:202
		{
			yyVAL.expr = __store(NewSymbolFromToken(yyDollar[2].token), NewString(yyDollar[4].token.Str), __func(__markupFuncName(yyDollar[4].token), emptyNode, "", yyDollar[7].expr))
		}
	case 35:
		yyDollar = yyS[yypt-9 : yypt+1]
//line parser.go.y:205
		{
			yyVAL.expr = __store(NewSymbolFromToken(yyDollar[2].token), NewString(yyDollar[4].token.Str), __func(__markupFuncName(yyDollar[4].token), yyDollar[6].expr, "", yyDollar[8].expr))
		}
	case 36:
		yyDollar = yyS[yypt-9 : yypt+1]
//line parser.go.y:208
		{
			yyVAL.expr = __store(NewSymbolFromToken(yyDollar[2].token), NewString(yyDollar[4].token.Str), __func(__markupFuncName(yyDollar[4].token), emptyNode, yyDollar[7].token.Str, yyDollar[8].expr))
		}
	case 37:
		yyDollar = yyS[yypt-10 : yypt+1]
//line parser.go.y:211
		{
			yyVAL.expr = __store(NewSymbolFromToken(yyDollar[2].token), NewString(yyDollar[4].token.Str), __func(__markupFuncName(yyDollar[4].token), yyDollar[6].expr, yyDollar[8].token.Str, yyDollar[9].expr))
		}
	case 38:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:216
		{
			yyVAL.expr = NewComplex(NewSymbol(ABreak)).SetPos(yyDollar[1].token.Pos)
		}
	case 39:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:219
		{
			yyVAL.expr = NewComplex(NewSymbol(AGoto), NewSymbolFromToken(yyDollar[2].token)).SetPos(yyDollar[1].token.Pos)
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:222
		{
			yyVAL.expr = NewComplex(NewSymbol(ALabel), NewSymbolFromToken(yyDollar[2].token))
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:225
		{
			yyVAL.expr = NewComplex(NewSymbol(AReturn), NewSymbol(ANil)).SetPos(yyDollar[1].token.Pos)
		}
	case 42:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:228
		{
			if len(yyDollar[2].expr.Nodes) == 1 {
				a := yyDollar[2].expr.Nodes[0]
				if len(a.Nodes) == 3 && a.Nodes[0].SymbolValue() == ACall {
					// return call(...) -> return tailcall(...)
					a.Nodes[0].strSym = ATailCall
				}
				yyVAL.expr = NewComplex(NewSymbol(AReturn), a).SetPos(yyDollar[1].token.Pos)
			} else {
				yyVAL.expr = NewComplex(NewSymbol(AReturn), NewComplex(NewSymbol(AArray), yyDollar[2].expr)).SetPos(yyDollar[1].token.Pos)
			}
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:242
		{
			yyVAL.expr = NewSymbolFromToken(yyDollar[1].token)
		}
	case 44:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:245
		{
			yyVAL.expr = __load(yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:248
		{
			yyVAL.expr = __load(yyDollar[1].expr, NewString(yyDollar[3].token.Str)).SetPos(yyDollar[2].token.Pos)
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:253
		{
			yyVAL.expr = NewComplex(yyDollar[1].expr)
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:256
		{
			yyVAL.expr = yyDollar[1].expr.append(yyDollar[3].expr)
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:261
		{
			yyVAL.expr = NewComplex(NewSymbolFromToken(yyDollar[1].token))
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:264
		{
			yyVAL.expr = yyDollar[1].expr.append(NewSymbolFromToken(yyDollar[3].token))
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:269
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 51:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:270
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:271
		{
			yyVAL.expr = NewNumberFromString(yyDollar[1].token.Str)
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:272
		{
			yyVAL.expr = NewString(yyDollar[1].token.Str)
		}
	case 54:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:273
		{
			yyVAL.expr = NewComplex(NewSymbol(AArrayMap), emptyNode).SetPos(yyDollar[1].token.Pos)
		}
	case 55:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:274
		{
			yyVAL.expr = NewComplex(NewSymbol(AArray), yyDollar[2].expr).SetPos(yyDollar[1].token.Pos)
		}
	case 56:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:275
		{
			yyVAL.expr = NewComplex(NewSymbol(AArrayMap), yyDollar[2].expr).SetPos(yyDollar[1].token.Pos)
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:276
		{
			yyVAL.expr = NewComplex(NewSymbol(AOr), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:277
		{
			yyVAL.expr = NewComplex(NewSymbol(AAnd), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:278
		{
			yyVAL.expr = NewComplex(NewSymbol(ALess), yyDollar[3].expr, yyDollar[1].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:279
		{
			yyVAL.expr = NewComplex(NewSymbol(ALess), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:280
		{
			yyVAL.expr = NewComplex(NewSymbol(ALessEq), yyDollar[3].expr, yyDollar[1].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:281
		{
			yyVAL.expr = NewComplex(NewSymbol(ALessEq), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:282
		{
			yyVAL.expr = NewComplex(NewSymbol(AEq), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:283
		{
			yyVAL.expr = NewComplex(NewSymbol(ANeq), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:284
		{
			yyVAL.expr = NewComplex(NewSymbol(AAdd), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:285
		{
			yyVAL.expr = NewComplex(NewSymbol(ASub), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:286
		{
			yyVAL.expr = NewComplex(NewSymbol(AMul), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:287
		{
			yyVAL.expr = NewComplex(NewSymbol(ADiv), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:288
		{
			yyVAL.expr = NewComplex(NewSymbol(AIDiv), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:289
		{
			yyVAL.expr = NewComplex(NewSymbol(AMod), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:290
		{
			yyVAL.expr = NewComplex(NewSymbol(ABitAnd), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:291
		{
			yyVAL.expr = NewComplex(NewSymbol(ABitOr), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:292
		{
			yyVAL.expr = NewComplex(NewSymbol(ABitXor), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:293
		{
			yyVAL.expr = NewComplex(NewSymbol(ABitLsh), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 75:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:294
		{
			yyVAL.expr = NewComplex(NewSymbol(ABitRsh), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:295
		{
			yyVAL.expr = NewComplex(NewSymbol(ABitURsh), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 77:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:296
		{
			yyVAL.expr = NewComplex(NewSymbol(ABitNot), yyDollar[2].expr).SetPos(yyDollar[1].token.Pos)
		}
	case 78:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:297
		{
			yyVAL.expr = NewComplex(NewSymbol(ANot), yyDollar[2].expr).SetPos(yyDollar[1].token.Pos)
		}
	case 79:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:298
		{
			yyVAL.expr = NewComplex(NewSymbol(ASub), zeroNode, yyDollar[2].expr).SetPos(yyDollar[1].token.Pos)
		}
	case 80:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:301
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 81:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:304
		{
			yyVAL.expr = __call(yyDollar[1].expr, NewComplex(NewString(yyDollar[2].token.Str))).SetPos(yyDollar[1].expr.Pos())
		}
	case 82:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:307
		{
			yyDollar[2].expr.Nodes[1] = yyDollar[1].expr
			yyVAL.expr = yyDollar[2].expr
		}
	case 83:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:313
		{
			yyVAL.expr = __call(emptyNode, emptyNode).SetPos(yyDollar[1].token.Pos)
		}
	case 84:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:316
		{
			yyVAL.expr = __call(emptyNode, yyDollar[2].expr).SetPos(yyDollar[1].token.Pos)
		}
	case 85:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:319
		{
			yyVAL.expr = __callMap(emptyNode, emptyNode, yyDollar[2].expr).SetPos(yyDollar[1].token.Pos)
		}
	case 86:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.go.y:322
		{
			yyVAL.expr = __callMap(emptyNode, yyDollar[2].expr, yyDollar[4].expr).SetPos(yyDollar[1].token.Pos)
		}
	case 87:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:327
		{
			yyVAL.expr = NewComplex(yyDollar[1].expr)
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:330
		{
			yyVAL.expr = yyDollar[1].expr.append(yyDollar[3].expr)
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:335
		{
			yyVAL.expr = NewComplex(NewString(yyDollar[1].token.Str), yyDollar[3].expr)
		}
	case 90:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.go.y:338
		{
			yyVAL.expr = NewComplex(yyDollar[2].expr, yyDollar[5].expr)
		}
	case 91:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.go.y:341
		{
			yyVAL.expr = yyDollar[1].expr.append(NewString(yyDollar[3].token.Str)).append(yyDollar[5].expr)
		}
	case 92:
		yyDollar = yyS[yypt-7 : yypt+1]
//line parser.go.y:344
		{
			yyVAL.expr = yyDollar[1].expr.append(yyDollar[4].expr).append(yyDollar[7].expr)
		}
	case 93:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:348
		{
			yyVAL.expr = emptyNode
		}
	case 94:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:348
		{
			yyVAL.expr = emptyNode
		}
	}
	goto yystack /* stack new state and value */
}
