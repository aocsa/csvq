// Code generated by goyacc -p jq -o query_parser.go -v query_parser.output query_parser.y. DO NOT EDIT.

//line query_parser.y:2
package json

import __yyfmt__ "fmt"

//line query_parser.y:2

import "strconv"

//line query_parser.y:7
type jqSymType struct {
	yys        int
	expression QueryExpression
	element    Element
	field      FieldExpr
	fields     []FieldExpr
	token      QueryToken
}

const PATH_IDENTIFIER = 57346
const PATH_INDEX = 57347
const AS = 57348

var jqToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"PATH_IDENTIFIER",
	"PATH_INDEX",
	"AS",
	"'.'",
	"'['",
	"']'",
	"'{'",
	"'}'",
	"','",
}
var jqStatenames = [...]string{}

const jqEofCode = 1
const jqErrCode = 2
const jqInitialStackSize = 16

//line query_parser.y:185

func ParseQuery(src string) (QueryExpression, error) {
	l := new(QueryLexer)
	l.Init(src)
	jqParse(l)
	return l.query, l.err
}

//line yacctab:1
var jqExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const jqPrivate = 57344

const jqLast = 50

var jqAct = [...]int{

	18, 3, 23, 16, 26, 6, 5, 30, 8, 4,
	9, 40, 20, 13, 12, 10, 8, 11, 9, 27,
	7, 25, 29, 34, 8, 35, 9, 33, 32, 38,
	36, 31, 39, 14, 41, 24, 19, 15, 28, 24,
	22, 24, 43, 42, 21, 7, 37, 17, 2, 1,
}
var jqPact = [...]int{

	16, -1000, -1000, -1000, -1000, -1000, -1000, 8, 28, 32,
	41, -1000, -1000, -1000, 35, 33, 10, -8, 13, 31,
	-1000, 0, 32, -1000, 20, -1000, 32, 42, 32, -1000,
	41, -1000, -1000, -1000, -1000, 2, -1000, -1000, -1000, -1000,
	27, 32, -1000, -1000,
}
var jqPgo = [...]int{

	0, 49, 48, 1, 0, 9, 2, 6, 5, 47,
	3,
}
var jqR1 = [...]int{

	0, 1, 1, 2, 2, 2, 2, 3, 3, 3,
	3, 3, 4, 4, 4, 5, 5, 5, 5, 5,
	6, 6, 6, 7, 7, 7, 8, 9, 9, 10,
	10, 10,
}
var jqR2 = [...]int{

	0, 0, 1, 1, 1, 1, 1, 1, 3, 2,
	2, 2, 1, 3, 2, 3, 5, 4, 4, 4,
	3, 5, 4, 2, 4, 3, 3, 1, 3, 0,
	1, 3,
}
var jqChk = [...]int{

	-1000, -1, -2, -3, -5, -7, -8, 4, 8, 10,
	7, -5, -7, -8, 5, 9, -10, -9, -4, 4,
	-3, 9, 7, -6, 8, 11, 12, 6, 7, -6,
	7, -5, -7, -8, -4, 5, -10, 4, -4, -3,
	9, 7, -6, -4,
}
var jqDef = [...]int{

	1, -2, 2, 3, 4, 5, 6, 7, 0, 29,
	0, 9, 10, 11, 0, 23, 0, 30, 27, 12,
	8, 15, 0, 25, 0, 26, 29, 0, 0, 14,
	0, 17, 18, 19, 24, 0, 31, 28, 13, 16,
	20, 0, 22, 21,
}
var jqTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 12, 3, 7, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 8, 3, 9, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 10, 3, 11,
}
var jqTok2 = [...]int{

	2, 3, 4, 5, 6,
}
var jqTok3 = [...]int{
	0,
}

var jqErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	jqDebug        = 0
	jqErrorVerbose = false
)

type jqLexer interface {
	Lex(lval *jqSymType) int
	Error(s string)
}

type jqParser interface {
	Parse(jqLexer) int
	Lookahead() int
}

type jqParserImpl struct {
	lval  jqSymType
	stack [jqInitialStackSize]jqSymType
	char  int
}

func (p *jqParserImpl) Lookahead() int {
	return p.char
}

func jqNewParser() jqParser {
	return &jqParserImpl{}
}

const jqFlag = -1000

func jqTokname(c int) string {
	if c >= 1 && c-1 < len(jqToknames) {
		if jqToknames[c-1] != "" {
			return jqToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func jqStatname(s int) string {
	if s >= 0 && s < len(jqStatenames) {
		if jqStatenames[s] != "" {
			return jqStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func jqErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !jqErrorVerbose {
		return "syntax error"
	}

	for _, e := range jqErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + jqTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := jqPact[state]
	for tok := TOKSTART; tok-1 < len(jqToknames); tok++ {
		if n := base + tok; n >= 0 && n < jqLast && jqChk[jqAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if jqDef[state] == -2 {
		i := 0
		for jqExca[i] != -1 || jqExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; jqExca[i] >= 0; i += 2 {
			tok := jqExca[i]
			if tok < TOKSTART || jqExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if jqExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += jqTokname(tok)
	}
	return res
}

func jqlex1(lex jqLexer, lval *jqSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = jqTok1[0]
		goto out
	}
	if char < len(jqTok1) {
		token = jqTok1[char]
		goto out
	}
	if char >= jqPrivate {
		if char < jqPrivate+len(jqTok2) {
			token = jqTok2[char-jqPrivate]
			goto out
		}
	}
	for i := 0; i < len(jqTok3); i += 2 {
		token = jqTok3[i+0]
		if token == char {
			token = jqTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = jqTok2[1] /* unknown char */
	}
	if jqDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", jqTokname(token), uint(char))
	}
	return char, token
}

func jqParse(jqlex jqLexer) int {
	return jqNewParser().Parse(jqlex)
}

func (jqrcvr *jqParserImpl) Parse(jqlex jqLexer) int {
	var jqn int
	var jqVAL jqSymType
	var jqDollar []jqSymType
	_ = jqDollar // silence set and not used
	jqS := jqrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	jqstate := 0
	jqrcvr.char = -1
	jqtoken := -1 // jqrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		jqstate = -1
		jqrcvr.char = -1
		jqtoken = -1
	}()
	jqp := -1
	goto jqstack

ret0:
	return 0

ret1:
	return 1

jqstack:
	/* put a state and value onto the stack */
	if jqDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", jqTokname(jqtoken), jqStatname(jqstate))
	}

	jqp++
	if jqp >= len(jqS) {
		nyys := make([]jqSymType, len(jqS)*2)
		copy(nyys, jqS)
		jqS = nyys
	}
	jqS[jqp] = jqVAL
	jqS[jqp].yys = jqstate

jqnewstate:
	jqn = jqPact[jqstate]
	if jqn <= jqFlag {
		goto jqdefault /* simple state */
	}
	if jqrcvr.char < 0 {
		jqrcvr.char, jqtoken = jqlex1(jqlex, &jqrcvr.lval)
	}
	jqn += jqtoken
	if jqn < 0 || jqn >= jqLast {
		goto jqdefault
	}
	jqn = jqAct[jqn]
	if jqChk[jqn] == jqtoken { /* valid shift */
		jqrcvr.char = -1
		jqtoken = -1
		jqVAL = jqrcvr.lval
		jqstate = jqn
		if Errflag > 0 {
			Errflag--
		}
		goto jqstack
	}

jqdefault:
	/* default state action */
	jqn = jqDef[jqstate]
	if jqn == -2 {
		if jqrcvr.char < 0 {
			jqrcvr.char, jqtoken = jqlex1(jqlex, &jqrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if jqExca[xi+0] == -1 && jqExca[xi+1] == jqstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			jqn = jqExca[xi+0]
			if jqn < 0 || jqn == jqtoken {
				break
			}
		}
		jqn = jqExca[xi+1]
		if jqn < 0 {
			goto ret0
		}
	}
	if jqn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			jqlex.Error(jqErrorMessage(jqstate, jqtoken))
			Nerrs++
			if jqDebug >= 1 {
				__yyfmt__.Printf("%s", jqStatname(jqstate))
				__yyfmt__.Printf(" saw %s\n", jqTokname(jqtoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for jqp >= 0 {
				jqn = jqPact[jqS[jqp].yys] + jqErrCode
				if jqn >= 0 && jqn < jqLast {
					jqstate = jqAct[jqn] /* simulate a shift of "error" */
					if jqChk[jqstate] == jqErrCode {
						goto jqstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if jqDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", jqS[jqp].yys)
				}
				jqp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if jqDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", jqTokname(jqtoken))
			}
			if jqtoken == jqEofCode {
				goto ret1
			}
			jqrcvr.char = -1
			jqtoken = -1
			goto jqnewstate /* try again in the same state */
		}
	}

	/* reduction by production jqn */
	if jqDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", jqn, jqStatname(jqstate))
	}

	jqnt := jqn
	jqpt := jqp
	_ = jqpt // guard against "declared and not used"

	jqp -= jqR2[jqn]
	// jqp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if jqp+1 >= len(jqS) {
		nyys := make([]jqSymType, len(jqS)*2)
		copy(nyys, jqS)
		jqS = nyys
	}
	jqVAL = jqS[jqp+1]

	/* consult goto table to find next state */
	jqn = jqR1[jqn]
	jqg := jqPgo[jqn]
	jqj := jqg + jqS[jqp].yys + 1

	if jqj >= jqLast {
		jqstate = jqAct[jqg]
	} else {
		jqstate = jqAct[jqj]
		if jqChk[jqstate] != -jqn {
			jqstate = jqAct[jqg]
		}
	}
	// dummy call; replaced with literal code
	switch jqnt {

	case 1:
		jqDollar = jqS[jqpt-0 : jqpt+1]
		//line query_parser.y:33
		{
			jqVAL.expression = nil
			jqlex.(*QueryLexer).query = jqVAL.expression
		}
	case 2:
		jqDollar = jqS[jqpt-1 : jqpt+1]
		//line query_parser.y:38
		{
			jqVAL.expression = jqDollar[1].expression
			jqlex.(*QueryLexer).query = jqVAL.expression
		}
	case 3:
		jqDollar = jqS[jqpt-1 : jqpt+1]
		//line query_parser.y:45
		{
			jqVAL.expression = jqDollar[1].element
		}
	case 4:
		jqDollar = jqS[jqpt-1 : jqpt+1]
		//line query_parser.y:49
		{
			jqVAL.expression = jqDollar[1].expression
		}
	case 5:
		jqDollar = jqS[jqpt-1 : jqpt+1]
		//line query_parser.y:53
		{
			jqVAL.expression = jqDollar[1].expression
		}
	case 6:
		jqDollar = jqS[jqpt-1 : jqpt+1]
		//line query_parser.y:57
		{
			jqVAL.expression = jqDollar[1].expression
		}
	case 7:
		jqDollar = jqS[jqpt-1 : jqpt+1]
		//line query_parser.y:63
		{
			jqVAL.element = Element{Label: jqDollar[1].token.Literal}
		}
	case 8:
		jqDollar = jqS[jqpt-3 : jqpt+1]
		//line query_parser.y:67
		{
			jqVAL.element = Element{Label: jqDollar[1].token.Literal, Child: jqDollar[3].element}
		}
	case 9:
		jqDollar = jqS[jqpt-2 : jqpt+1]
		//line query_parser.y:71
		{
			jqVAL.element = Element{Label: jqDollar[1].token.Literal, Child: jqDollar[2].expression}
		}
	case 10:
		jqDollar = jqS[jqpt-2 : jqpt+1]
		//line query_parser.y:75
		{
			jqVAL.element = Element{Label: jqDollar[1].token.Literal, Child: jqDollar[2].expression}
		}
	case 11:
		jqDollar = jqS[jqpt-2 : jqpt+1]
		//line query_parser.y:79
		{
			jqVAL.element = Element{Label: jqDollar[1].token.Literal, Child: jqDollar[2].expression}
		}
	case 12:
		jqDollar = jqS[jqpt-1 : jqpt+1]
		//line query_parser.y:85
		{
			jqVAL.element = Element{Label: jqDollar[1].token.Literal}
		}
	case 13:
		jqDollar = jqS[jqpt-3 : jqpt+1]
		//line query_parser.y:89
		{
			jqVAL.element = Element{Label: jqDollar[1].token.Literal, Child: jqDollar[3].element}
		}
	case 14:
		jqDollar = jqS[jqpt-2 : jqpt+1]
		//line query_parser.y:93
		{
			jqVAL.element = Element{Label: jqDollar[1].token.Literal, Child: jqDollar[2].expression}
		}
	case 15:
		jqDollar = jqS[jqpt-3 : jqpt+1]
		//line query_parser.y:99
		{
			i, _ := strconv.Atoi(jqDollar[2].token.Literal)
			jqVAL.expression = ArrayItem{Index: i}
		}
	case 16:
		jqDollar = jqS[jqpt-5 : jqpt+1]
		//line query_parser.y:104
		{
			i, _ := strconv.Atoi(jqDollar[2].token.Literal)
			jqVAL.expression = ArrayItem{Index: i, Child: jqDollar[5].element}
		}
	case 17:
		jqDollar = jqS[jqpt-4 : jqpt+1]
		//line query_parser.y:109
		{
			i, _ := strconv.Atoi(jqDollar[2].token.Literal)
			jqVAL.expression = ArrayItem{Index: i, Child: jqDollar[4].expression}
		}
	case 18:
		jqDollar = jqS[jqpt-4 : jqpt+1]
		//line query_parser.y:114
		{
			i, _ := strconv.Atoi(jqDollar[2].token.Literal)
			jqVAL.expression = ArrayItem{Index: i, Child: jqDollar[4].expression}
		}
	case 19:
		jqDollar = jqS[jqpt-4 : jqpt+1]
		//line query_parser.y:119
		{
			i, _ := strconv.Atoi(jqDollar[2].token.Literal)
			jqVAL.expression = ArrayItem{Index: i, Child: jqDollar[4].expression}
		}
	case 20:
		jqDollar = jqS[jqpt-3 : jqpt+1]
		//line query_parser.y:126
		{
			i, _ := strconv.Atoi(jqDollar[2].token.Literal)
			jqVAL.expression = ArrayItem{Index: i}
		}
	case 21:
		jqDollar = jqS[jqpt-5 : jqpt+1]
		//line query_parser.y:131
		{
			i, _ := strconv.Atoi(jqDollar[2].token.Literal)
			jqVAL.expression = ArrayItem{Index: i, Child: jqDollar[5].element}
		}
	case 22:
		jqDollar = jqS[jqpt-4 : jqpt+1]
		//line query_parser.y:136
		{
			i, _ := strconv.Atoi(jqDollar[2].token.Literal)
			jqVAL.expression = ArrayItem{Index: i, Child: jqDollar[4].expression}
		}
	case 23:
		jqDollar = jqS[jqpt-2 : jqpt+1]
		//line query_parser.y:143
		{
			jqVAL.expression = RowValueExpr{}
		}
	case 24:
		jqDollar = jqS[jqpt-4 : jqpt+1]
		//line query_parser.y:147
		{
			jqVAL.expression = RowValueExpr{Child: jqDollar[4].element}
		}
	case 25:
		jqDollar = jqS[jqpt-3 : jqpt+1]
		//line query_parser.y:151
		{
			jqVAL.expression = RowValueExpr{Child: jqDollar[3].expression}
		}
	case 26:
		jqDollar = jqS[jqpt-3 : jqpt+1]
		//line query_parser.y:157
		{
			jqVAL.expression = TableExpr{Fields: jqDollar[2].fields}
		}
	case 27:
		jqDollar = jqS[jqpt-1 : jqpt+1]
		//line query_parser.y:163
		{
			jqVAL.field = FieldExpr{Element: jqDollar[1].element}
		}
	case 28:
		jqDollar = jqS[jqpt-3 : jqpt+1]
		//line query_parser.y:167
		{
			jqVAL.field = FieldExpr{Element: jqDollar[1].element, Alias: jqDollar[3].token.Literal}
		}
	case 29:
		jqDollar = jqS[jqpt-0 : jqpt+1]
		//line query_parser.y:173
		{
			jqVAL.fields = nil
		}
	case 30:
		jqDollar = jqS[jqpt-1 : jqpt+1]
		//line query_parser.y:177
		{
			jqVAL.fields = []FieldExpr{jqDollar[1].field}
		}
	case 31:
		jqDollar = jqS[jqpt-3 : jqpt+1]
		//line query_parser.y:181
		{
			jqVAL.fields = append([]FieldExpr{jqDollar[1].field}, jqDollar[3].fields...)
		}
	}
	goto jqstack /* stack new state and value */
}
