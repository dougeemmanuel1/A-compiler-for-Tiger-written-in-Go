// Code generated by goyacc tiger.y. DO NOT EDIT.

//line tiger.y:2

package main

import __yyfmt__ "fmt"

//line tiger.y:3

import (
	_ "fmt"
	"github.com/timtadh/lexmachine"
	"strconv"
)

var yylineno int = 1

func toInt(s string) int {
	num, _ := strconv.ParseInt(s, 10, 64)
	return int(num)
}

//line tiger.y:26
type yySymType struct {
	yys      int
	token    *lexmachine.Token
	ast      *Node
	NodeList []Node
	strLit   string
}

const OR = 57346
const AND = 57347
const EQUALS = 57348
const NEQ = 57349
const GT = 57350
const LT = 57351
const GEQ = 57352
const LEQ = 57353
const PLUS = 57354
const MINUS = 57355
const STAR = 57356
const FORWARDSLASH = 57357
const UNARY = 57358
const LPAREN = 57359
const RPAREN = 57360
const LBRACKET = 57361
const RBRACKET = 57362
const LCURLY = 57363
const RCURLY = 57364
const COLON = 57365
const COLONEQUALS = 57366
const DOT = 57367
const COMMA = 57368
const SEMICOLON = 57369
const DOUBLEARROW = 57370
const RARROW = 57371
const LARROW = 57372
const GREATERTHANEQ = 57373
const LESSTHANEQ = 57374
const BAR = 57375
const ARRAY = 57376
const BREAK = 57377
const DO = 57378
const ELSE = 57379
const END = 57380
const FOR = 57381
const FUNCTION = 57382
const IF = 57383
const IN = 57384
const LET = 57385
const NIL = 57386
const OF = 57387
const THEN = 57388
const TO = 57389
const TYPE = 57390
const VAR = 57391
const WHILE = 57392
const ID = 57393
const INTLIT = 57394
const STRINGLIT = 57395
const NEWLINE = 57396

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"OR",
	"AND",
	"EQUALS",
	"NEQ",
	"GT",
	"LT",
	"GEQ",
	"LEQ",
	"PLUS",
	"MINUS",
	"STAR",
	"FORWARDSLASH",
	"UNARY",
	"LPAREN",
	"RPAREN",
	"LBRACKET",
	"RBRACKET",
	"LCURLY",
	"RCURLY",
	"COLON",
	"COLONEQUALS",
	"DOT",
	"COMMA",
	"SEMICOLON",
	"DOUBLEARROW",
	"RARROW",
	"LARROW",
	"GREATERTHANEQ",
	"LESSTHANEQ",
	"BAR",
	"ARRAY",
	"BREAK",
	"DO",
	"ELSE",
	"END",
	"FOR",
	"FUNCTION",
	"IF",
	"IN",
	"LET",
	"NIL",
	"OF",
	"THEN",
	"TO",
	"TYPE",
	"VAR",
	"WHILE",
	"ID",
	"INTLIT",
	"STRINGLIT",
	"NEWLINE",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line tiger.y:193

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 63,
	6, 0,
	-2, 58,
}

const yyPrivate = 57344

const yyLast = 406

var yyAct = [...]int{

	74, 2, 132, 123, 81, 43, 52, 151, 150, 147,
	135, 134, 139, 140, 122, 120, 82, 133, 58, 133,
	87, 44, 80, 48, 49, 90, 56, 57, 119, 59,
	60, 61, 62, 63, 64, 65, 66, 67, 68, 69,
	70, 71, 135, 73, 135, 116, 77, 89, 88, 72,
	50, 82, 20, 58, 128, 108, 19, 131, 86, 97,
	85, 56, 57, 98, 106, 105, 75, 145, 33, 34,
	35, 36, 37, 39, 16, 130, 92, 114, 24, 94,
	22, 143, 25, 3, 100, 101, 102, 124, 107, 23,
	21, 4, 5, 103, 136, 46, 109, 45, 144, 47,
	111, 153, 40, 110, 28, 29, 121, 42, 41, 125,
	152, 137, 20, 126, 127, 104, 19, 96, 33, 34,
	35, 36, 37, 39, 129, 95, 99, 1, 54, 53,
	27, 26, 141, 18, 16, 17, 142, 146, 24, 148,
	22, 20, 25, 3, 149, 19, 76, 15, 14, 23,
	21, 4, 5, 154, 155, 75, 13, 118, 117, 115,
	20, 55, 12, 16, 19, 11, 10, 24, 9, 22,
	8, 25, 3, 7, 6, 79, 51, 78, 23, 21,
	4, 5, 16, 0, 0, 0, 24, 0, 22, 0,
	25, 3, 0, 0, 38, 32, 0, 23, 21, 4,
	5, 30, 31, 28, 29, 38, 32, 0, 0, 0,
	0, 0, 30, 31, 28, 29, 0, 33, 34, 35,
	36, 37, 39, 0, 0, 0, 0, 0, 33, 34,
	35, 36, 37, 39, 38, 32, 113, 0, 0, 0,
	0, 30, 31, 28, 29, 0, 83, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 33, 34, 35,
	36, 37, 39, 38, 32, 0, 112, 0, 0, 0,
	30, 31, 28, 29, 38, 32, 0, 0, 0, 0,
	0, 30, 31, 28, 29, 0, 33, 34, 35, 36,
	37, 39, 0, 0, 138, 0, 0, 33, 34, 35,
	36, 37, 39, 38, 32, 84, 0, 0, 0, 0,
	30, 31, 28, 29, 0, 0, 0, 0, 93, 0,
	0, 0, 0, 0, 0, 0, 33, 34, 35, 36,
	37, 39, 38, 32, 0, 0, 0, 0, 0, 30,
	31, 28, 29, 38, 32, 0, 0, 91, 0, 0,
	30, 31, 28, 29, 0, 33, 34, 35, 36, 37,
	39, 0, 0, 0, 0, 0, 33, 34, 35, 36,
	37, 39, 32, 0, 0, 0, 0, 0, 30, 31,
	28, 29, 0, 0, 30, 31, 28, 29, 0, 0,
	0, 0, 0, 0, 33, 34, 35, 36, 37, 39,
	33, 34, 35, 36, 37, 39,
}
var yyPact = [...]int{

	147, -1000, 338, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 83, -1000,
	147, 78, 147, 147, -1, 13, -1000, -1000, 147, 147,
	147, 147, 147, 147, 147, 147, 147, 147, 147, 147,
	147, -2, 147, 128, 40, 147, -1000, 0, 200, 269,
	36, -22, -1000, -1000, -1000, -1000, -3, -4, -26, 40,
	40, 90, 90, 372, 338, 338, 338, 338, 338, 366,
	338, 327, -1000, 338, 338, 147, -1000, 298, 99, 37,
	-1000, -1000, 120, 147, 147, 147, -1000, -1000, 109, 41,
	71, -1000, 338, 10, 338, 147, -1000, -1000, -35, 147,
	229, 338, 189, 39, -6, 147, -37, 69, 147, 338,
	-1000, 338, 147, 147, -1000, -1000, -1000, -1000, -1000, 9,
	53, 338, 33, -7, 88, 338, 338, 258, -39, -9,
	-1000, 147, -1000, -41, 75, 44, 147, -42, 147, -1000,
	-1000, 338, -1000, 147, -43, -44, 338, 104, 338, 338,
	95, -1000, 147, 147, 338, 338,
}
var yyPgo = [...]int{

	0, 5, 177, 176, 3, 175, 0, 174, 173, 170,
	168, 166, 165, 162, 161, 159, 158, 157, 2, 4,
	156, 148, 147, 135, 133, 131, 130, 6, 129, 128,
	127,
}
var yyR1 = [...]int{

	0, 30, 27, 27, 27, 28, 15, 15, 15, 16,
	17, 17, 4, 4, 4, 18, 14, 14, 14, 14,
	29, 29, 25, 25, 11, 24, 24, 24, 26, 1,
	1, 1, 2, 2, 2, 6, 6, 6, 6, 6,
	6, 6, 6, 6, 6, 6, 6, 6, 6, 6,
	6, 7, 8, 9, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 10, 12, 12, 5, 5,
	19, 13, 20, 20, 21, 22, 3, 3, 23,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 4, 1, 1, 1, 3,
	3, 2, 0, 2, 3, 3, 7, 6, 9, 8,
	4, 6, 4, 4, 6, 1, 1, 1, 3, 0,
	2, 3, 0, 2, 3, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 3, 2, 4, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 4, 3, 3, 1,
	3, 3, 6, 4, 4, 8, 2, 1, 5,
}
var yyChk = [...]int{

	-1000, -30, -6, 44, 52, 53, -7, -8, -9, -10,
	-11, -12, -13, -20, -21, -22, 35, -23, -24, 17,
	13, 51, 41, 50, 39, 43, -25, -26, 14, 15,
	12, 13, 6, 28, 29, 30, 31, 32, 5, 33,
	19, 25, 24, -1, -6, 19, 17, 21, -6, -6,
	51, -3, -27, -28, -29, -14, 48, 49, 40, -6,
	-6, -6, -6, -6, -6, -6, -6, -6, -6, -6,
	-6, -6, 51, -6, -6, 27, 18, -6, -2, -5,
	22, -19, 51, 46, 36, 24, -27, 42, 51, 51,
	51, 20, -6, 20, -6, 26, 18, 22, 26, 6,
	-6, -6, -6, -1, 6, 24, 23, 17, 45, -6,
	-19, -6, 37, 47, 38, -15, 51, -16, -17, 34,
	21, -6, 51, -4, 18, -6, -6, -6, 45, -4,
	22, 24, -18, 26, 18, 51, 6, 23, 36, 51,
	22, -6, -18, 6, 23, 23, -6, 51, -6, -6,
	51, 51, 6, 6, -6, -6,
}
var yyDef = [...]int{

	0, -2, 1, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 29,
	0, 25, 0, 0, 0, 0, 26, 27, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 52, 0, 32, 0, 0, 0,
	0, 0, 77, 2, 3, 4, 0, 0, 0, 54,
	55, 56, 57, -2, 59, 60, 61, 62, 63, 64,
	65, 0, 28, 71, 30, 0, 51, 0, 0, 0,
	67, 69, 0, 0, 0, 0, 76, 29, 0, 0,
	0, 22, 31, 23, 33, 0, 53, 66, 0, 0,
	73, 74, 0, 0, 0, 0, 0, 12, 0, 34,
	68, 70, 0, 0, 78, 5, 6, 7, 8, 0,
	12, 20, 0, 0, 0, 24, 72, 0, 0, 0,
	11, 0, 13, 0, 0, 0, 0, 0, 0, 9,
	10, 21, 14, 0, 0, 0, 17, 0, 75, 16,
	0, 15, 0, 0, 19, 18,
}
var yyTok1 = [...]int{

	1,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54,
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
		yyDollar = yyS[yypt-1 : yypt+1]
//line tiger.y:52
		{
			yylex.(*golex).stmts = append(yylex.(*golex).stmts, yyDollar[1].ast)
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
//line tiger.y:55
		{
			yyVAL.ast = yyDollar[1].ast
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
//line tiger.y:56
		{
			yyVAL.ast = yyDollar[1].ast
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
//line tiger.y:57
		{
			yyVAL.ast = yyDollar[1].ast
		}
	case 5:
		yyDollar = yyS[yypt-4 : yypt+1]
//line tiger.y:60
		{
			yyVAL.ast = NewNode("typeDec", nil, NewTypeDeclaration(string(yyDollar[2].token.Lexeme), yyDollar[4].ast, yyDollar[2].token.StartLine))
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
//line tiger.y:63
		{
			yyVAL.ast = NewNode("ID", yyDollar[1].token, NewIdentifier(string(yyDollar[1].token.Lexeme), yyDollar[1].token.StartLine))
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
//line tiger.y:64
		{
			yyVAL.ast = yyDollar[1].ast
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
//line tiger.y:65
		{
			yyVAL.ast = yyDollar[1].ast
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
//line tiger.y:68
		{
			yyVAL.ast = NewNode("arrTy", nil, NewArrayType(string(yyDollar[3].token.Lexeme), yyDollar[3].token.StartLine))
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
//line tiger.y:71
		{
			yyVAL.ast = NewNode("sad", nil, NewRecordType(yyDollar[2].NodeList, yyDollar[1].token.StartLine))
		}
	case 11:
		yyDollar = yyS[yypt-2 : yypt+1]
//line tiger.y:72
		{
			yyVAL.ast = NewNode("recTy", nil, NewRecordType([]Node{}, yyDollar[1].token.StartLine))
		}
	case 12:
		yyDollar = yyS[yypt-0 : yypt+1]
//line tiger.y:75
		{
			yyVAL.NodeList = []Node{}
		}
	case 13:
		yyDollar = yyS[yypt-2 : yypt+1]
//line tiger.y:76
		{
			yyVAL.NodeList = append(yyVAL.NodeList, *yyDollar[2].ast)
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
//line tiger.y:77
		{
			yyVAL.NodeList = append(yyVAL.NodeList, *yyDollar[3].ast)
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
//line tiger.y:80
		{
			yyVAL.ast = NewNode("Param", nil, NewParam(string(yyDollar[1].token.Lexeme), string(yyDollar[3].token.Lexeme), yyDollar[1].token.StartLine))
		}
	case 16:
		yyDollar = yyS[yypt-7 : yypt+1]
//line tiger.y:83
		{
			yyVAL.ast = NewNode("funDec", nil, NewFuncDeclaration(string(yyDollar[2].token.Lexeme), yyDollar[4].NodeList, "", *yyDollar[7].ast, yyDollar[2].token.StartLine))
		}
	case 17:
		yyDollar = yyS[yypt-6 : yypt+1]
//line tiger.y:84
		{
			yyVAL.ast = NewNode("funDec", nil, NewFuncDeclaration(string(yyDollar[2].token.Lexeme), []Node{}, "", *yyDollar[6].ast, yyDollar[2].token.StartLine))
		}
	case 18:
		yyDollar = yyS[yypt-9 : yypt+1]
//line tiger.y:85
		{
			yyVAL.ast = NewNode("funDec", nil, NewFuncDeclaration(string(yyDollar[2].token.Lexeme), yyDollar[4].NodeList, string(yyDollar[7].token.Lexeme), *yyDollar[9].ast, yyDollar[2].token.StartLine))
		}
	case 19:
		yyDollar = yyS[yypt-8 : yypt+1]
//line tiger.y:86
		{
			yyVAL.ast = NewNode("funDec", nil, NewFuncDeclaration(string(yyDollar[2].token.Lexeme), []Node{}, string(yyDollar[6].token.Lexeme), *yyDollar[8].ast, yyDollar[2].token.StartLine))
		}
	case 20:
		yyDollar = yyS[yypt-4 : yypt+1]
//line tiger.y:89
		{
			yyVAL.ast = NewNode("varDec", nil, NewVariable(string(yyDollar[2].token.Lexeme), "", yyDollar[4].ast, yyDollar[2].token.StartLine))
		}
	case 21:
		yyDollar = yyS[yypt-6 : yypt+1]
//line tiger.y:90
		{
			yyVAL.ast = NewNode("varDec", nil, NewVariable(string(yyDollar[2].token.Lexeme), string(yyDollar[4].token.Lexeme), yyDollar[6].ast, yyDollar[2].token.StartLine))
		}
	case 22:
		yyDollar = yyS[yypt-4 : yypt+1]
//line tiger.y:94
		{
			yyVAL.ast = NewNode("subscript", nil, NewSubscriptExpression("", yyDollar[1].ast, *yyDollar[3].ast, yyDollar[2].token.StartLine))
		}
	case 23:
		yyDollar = yyS[yypt-4 : yypt+1]
//line tiger.y:95
		{
			yyVAL.ast = NewNode("subscript", nil, NewSubscriptExpression(string(yyDollar[1].token.Lexeme), nil, *yyDollar[3].ast, yyDollar[2].token.StartLine))
		}
	case 24:
		yyDollar = yyS[yypt-6 : yypt+1]
//line tiger.y:98
		{
			yyVAL.ast = NewNode("arrExp", nil, NewArrayExp(string(yyDollar[1].token.Lexeme), *yyDollar[3].ast, *yyDollar[6].ast))
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
//line tiger.y:101
		{
			yyVAL.ast = NewNode("ID", yyDollar[1].token, NewIdentifier(string(yyDollar[1].token.Lexeme), yyDollar[1].token.StartLine))
		}
	case 26:
		yyDollar = yyS[yypt-1 : yypt+1]
//line tiger.y:102
		{
			yyVAL.ast = yyDollar[1].ast
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
//line tiger.y:103
		{
			yyVAL.ast = yyDollar[1].ast
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
//line tiger.y:106
		{
			yyVAL.ast = NewNode("fieldExp", nil, NewMemberExp(*yyDollar[1].ast, string(yyDollar[3].token.Lexeme), yyDollar[3].token.StartLine))
		}
	case 29:
		yyDollar = yyS[yypt-0 : yypt+1]
//line tiger.y:109
		{
			yyVAL.NodeList = []Node{}
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
//line tiger.y:110
		{
			yyVAL.NodeList = append(yyVAL.NodeList, *yyDollar[2].ast)
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
//line tiger.y:111
		{
			yyVAL.NodeList = append(yyVAL.NodeList, *yyDollar[3].ast)
		}
	case 32:
		yyDollar = yyS[yypt-0 : yypt+1]
//line tiger.y:114
		{
			yyVAL.NodeList = []Node{}
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
//line tiger.y:115
		{
			yyVAL.NodeList = append(yyVAL.NodeList, *yyDollar[2].ast)
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
//line tiger.y:116
		{
			yyVAL.NodeList = append(yyVAL.NodeList, *yyDollar[3].ast)
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
//line tiger.y:120
		{
			yyVAL.ast = NewNode("nil", nil, NewNil(yyDollar[1].token.StartLine))
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
//line tiger.y:121
		{
			yyVAL.ast = NewNode("INTLIT", nil, NewInteger(toInt(string(yyDollar[1].token.Lexeme)), yyDollar[1].token.StartLine))
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
//line tiger.y:122
		{
			yyVAL.ast = NewNode("STRINGLIT", nil, NewStringPrimitive(string(yyDollar[1].token.Lexeme), yyDollar[1].token.StartLine))
		}
	case 38:
		yyDollar = yyS[yypt-1 : yypt+1]
//line tiger.y:123
		{
			yyVAL.ast = yyDollar[1].ast
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
//line tiger.y:124
		{
			yyVAL.ast = yyDollar[1].ast
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
//line tiger.y:125
		{
			yyVAL.ast = yyDollar[1].ast
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
//line tiger.y:126
		{
			yyVAL.ast = yyDollar[1].ast
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
//line tiger.y:127
		{
			yyVAL.ast = yyDollar[1].ast
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
//line tiger.y:128
		{
			yyVAL.ast = yyDollar[1].ast
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
//line tiger.y:129
		{
			yyVAL.ast = yyDollar[1].ast
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
//line tiger.y:130
		{
			yyVAL.ast = yyDollar[1].ast
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
//line tiger.y:131
		{
			yyVAL.ast = yyDollar[1].ast
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
//line tiger.y:132
		{
			yyVAL.ast = yyDollar[1].ast
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
//line tiger.y:133
		{
			yyVAL.ast = yyDollar[1].ast
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
//line tiger.y:134
		{
			yyVAL.ast = yyDollar[1].ast
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
//line tiger.y:135
		{
			yyVAL.ast = yyDollar[1].ast
		}
	case 51:
		yyDollar = yyS[yypt-3 : yypt+1]
//line tiger.y:138
		{
			yyVAL.ast = NewNode("seqexp", nil, NewSeqExpression(yyDollar[2].NodeList, yyDollar[1].token.StartLine))
		}
	case 52:
		yyDollar = yyS[yypt-2 : yypt+1]
//line tiger.y:141
		{
			yyVAL.ast = NewNode("NEG", nil, NewNegation(yyDollar[2].ast, yyDollar[1].token.StartLine))
		}
	case 53:
		yyDollar = yyS[yypt-4 : yypt+1]
//line tiger.y:144
		{
			yyVAL.ast = NewNode("CALLEXP", nil, NewCallExpression(string(yyDollar[1].token.Lexeme), yyDollar[3].NodeList, yyDollar[1].token.StartLine))
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
//line tiger.y:147
		{
			yyVAL.ast = NewNode("MUL", nil, NewInfixExpression(Op_MUL, *yyDollar[1].ast, *yyDollar[3].ast, yyDollar[2].token.StartLine))
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
//line tiger.y:148
		{
			yyVAL.ast = NewNode("DIV", nil, NewInfixExpression(Op_DIV, *yyDollar[1].ast, *yyDollar[3].ast, yyDollar[2].token.StartLine))
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
//line tiger.y:149
		{
			yyVAL.ast = NewNode("DIV", nil, NewInfixExpression(Op_PLUS, *yyDollar[1].ast, *yyDollar[3].ast, yyDollar[2].token.StartLine))
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
//line tiger.y:150
		{
			yyVAL.ast = NewNode("DIV", nil, NewInfixExpression(Op_MINUS, *yyDollar[1].ast, *yyDollar[3].ast, yyDollar[2].token.StartLine))
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
//line tiger.y:151
		{
			yyVAL.ast = NewNode("DIV", nil, NewInfixExpression(Op_EQUALS, *yyDollar[1].ast, *yyDollar[3].ast, yyDollar[2].token.StartLine))
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
//line tiger.y:152
		{
			yyVAL.ast = NewNode("DIV", nil, NewInfixExpression(Op_NEQ, *yyDollar[1].ast, *yyDollar[3].ast, yyDollar[2].token.StartLine))
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
//line tiger.y:153
		{
			yyVAL.ast = NewNode("DIV", nil, NewInfixExpression(Op_GT, *yyDollar[1].ast, *yyDollar[3].ast, yyDollar[2].token.StartLine))
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
//line tiger.y:154
		{
			yyVAL.ast = NewNode("DIV", nil, NewInfixExpression(Op_LT, *yyDollar[1].ast, *yyDollar[3].ast, yyDollar[2].token.StartLine))
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
//line tiger.y:155
		{
			yyVAL.ast = NewNode("DIV", nil, NewInfixExpression(Op_GTE, *yyDollar[1].ast, *yyDollar[3].ast, yyDollar[2].token.StartLine))
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
//line tiger.y:156
		{
			yyVAL.ast = NewNode("DIV", nil, NewInfixExpression(Op_LTE, *yyDollar[1].ast, *yyDollar[3].ast, yyDollar[2].token.StartLine))
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
//line tiger.y:157
		{
			yyVAL.ast = NewNode("DIV", nil, NewInfixExpression(Op_AND, *yyDollar[1].ast, *yyDollar[3].ast, yyDollar[2].token.StartLine))
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
//line tiger.y:158
		{
			yyVAL.ast = NewNode("DIV", nil, NewInfixExpression(Op_OR, *yyDollar[1].ast, *yyDollar[3].ast, yyDollar[2].token.StartLine))
		}
	case 66:
		yyDollar = yyS[yypt-4 : yypt+1]
//line tiger.y:161
		{
			yyVAL.ast = NewNode("recordExp", nil, NewRecordExp(string(yyDollar[1].token.Lexeme), yyDollar[3].NodeList, yyDollar[1].token.StartLine))
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
//line tiger.y:162
		{
			yyVAL.ast = NewNode("recordExp", nil, NewRecordExp(string(yyDollar[1].token.Lexeme), []Node{}, yyDollar[1].token.StartLine))
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
//line tiger.y:165
		{
			yyVAL.NodeList = append(yyVAL.NodeList, *yyDollar[3].ast)
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
//line tiger.y:166
		{
			yyVAL.NodeList = append(yyVAL.NodeList, *yyDollar[1].ast)
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
//line tiger.y:169
		{
			yyVAL.ast = NewNode("Binding", nil, NewBinding(string(yyDollar[1].token.Lexeme), *yyDollar[3].ast, yyDollar[1].token.StartLine))
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
//line tiger.y:172
		{
			yyVAL.ast = NewNode("assignment", nil, NewAssignment(*yyDollar[1].ast, *yyDollar[3].ast, yyDollar[2].token.StartLine))
		}
	case 72:
		yyDollar = yyS[yypt-6 : yypt+1]
//line tiger.y:175
		{
			yyVAL.ast = NewNode("Iftheneelse", nil, NewIfThenElseExpression(*yyDollar[2].ast, *yyDollar[4].ast, yyDollar[6].ast, yyDollar[1].token.StartLine))
		}
	case 73:
		yyDollar = yyS[yypt-4 : yypt+1]
//line tiger.y:176
		{
			yyVAL.ast = NewNode("Iftheneelse", nil, NewIfThenElseExpression(*yyDollar[2].ast, *yyDollar[4].ast, nil, yyDollar[1].token.StartLine))
		}
	case 74:
		yyDollar = yyS[yypt-4 : yypt+1]
//line tiger.y:179
		{
			yyVAL.ast = NewNode("whileExp", nil, NewWhileExpression(*yyDollar[2].ast, *yyDollar[4].ast, yyDollar[1].token.StartLine))
		}
	case 75:
		yyDollar = yyS[yypt-8 : yypt+1]
//line tiger.y:182
		{
			yyVAL.ast = NewNode("Forexp", nil, NewForExpression(string(yyDollar[2].token.Lexeme), *yyDollar[4].ast, *yyDollar[6].ast, *yyDollar[8].ast, yyDollar[3].token.StartLine))
		}
	case 76:
		yyDollar = yyS[yypt-2 : yypt+1]
//line tiger.y:185
		{
			yyVAL.NodeList = append(yyVAL.NodeList, *yyDollar[2].ast)
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
//line tiger.y:186
		{
			yyVAL.NodeList = append(yyVAL.NodeList, *yyDollar[1].ast)
		}
	case 78:
		yyDollar = yyS[yypt-5 : yypt+1]
//line tiger.y:189
		{
			yyVAL.ast = NewNode("letexp", nil, NewLetExpression(yyDollar[2].NodeList, yyDollar[4].NodeList, yyDollar[1].token.StartLine))
		}
	}
	goto yystack /* stack new state and value */
}
