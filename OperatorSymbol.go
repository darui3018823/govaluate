package govaluate

import "github.com/darui3018823/govaluate/token"

// OperatorSymbol represents the valid symbols for operators.
type OperatorSymbol = token.OperatorSymbol

const (
	VALUE          = token.VALUE
	LITERAL        = token.LITERAL
	NOOP           = token.NOOP
	EQ             = token.EQ
	NEQ            = token.NEQ
	GT             = token.GT
	LT             = token.LT
	GTE            = token.GTE
	LTE            = token.LTE
	REQ            = token.REQ
	NREQ           = token.NREQ
	IN             = token.IN
	AND            = token.AND
	OR             = token.OR
	PLUS           = token.PLUS
	MINUS          = token.MINUS
	BITWISE_AND    = token.BITWISE_AND
	BITWISE_OR     = token.BITWISE_OR
	BITWISE_XOR    = token.BITWISE_XOR
	BITWISE_LSHIFT = token.BITWISE_LSHIFT
	BITWISE_RSHIFT = token.BITWISE_RSHIFT
	MULTIPLY       = token.MULTIPLY
	DIVIDE         = token.DIVIDE
	MODULUS        = token.MODULUS
	EXPONENT       = token.EXPONENT
	NEGATE         = token.NEGATE
	INVERT         = token.INVERT
	BITWISE_NOT    = token.BITWISE_NOT
	TERNARY_TRUE   = token.TERNARY_TRUE
	TERNARY_FALSE  = token.TERNARY_FALSE
	COALESCE       = token.COALESCE
	FUNCTIONAL     = token.FUNCTIONAL
	ACCESS         = token.ACCESS
	SEPARATE       = token.SEPARATE
)

type operatorPrecedence int

const (
	noopPrecedence operatorPrecedence = iota
	valuePrecedence
	functionalPrecedence
	prefixPrecedence
	exponentialPrecedence
	additivePrecedence
	bitwisePrecedence
	bitwiseShiftPrecedence
	multiplicativePrecedence
	comparatorPrecedence
	ternaryPrecedence
	logicalAndPrecedence
	logicalOrPrecedence
	separatePrecedence
)

func findOperatorPrecedenceForSymbol(symbol OperatorSymbol) operatorPrecedence {

	switch symbol {
	case NOOP:
		return noopPrecedence
	case VALUE:
		return valuePrecedence
	case EQ:
		fallthrough
	case NEQ:
		fallthrough
	case GT:
		fallthrough
	case LT:
		fallthrough
	case GTE:
		fallthrough
	case LTE:
		fallthrough
	case REQ:
		fallthrough
	case NREQ:
		fallthrough
	case IN:
		return comparatorPrecedence
	case AND:
		return logicalAndPrecedence
	case OR:
		return logicalOrPrecedence
	case BITWISE_AND:
		fallthrough
	case BITWISE_OR:
		fallthrough
	case BITWISE_XOR:
		return bitwisePrecedence
	case BITWISE_LSHIFT:
		fallthrough
	case BITWISE_RSHIFT:
		return bitwiseShiftPrecedence
	case PLUS:
		fallthrough
	case MINUS:
		return additivePrecedence
	case MULTIPLY:
		fallthrough
	case DIVIDE:
		fallthrough
	case MODULUS:
		return multiplicativePrecedence
	case EXPONENT:
		return exponentialPrecedence
	case BITWISE_NOT:
		fallthrough
	case NEGATE:
		fallthrough
	case INVERT:
		return prefixPrecedence
	case COALESCE:
		fallthrough
	case TERNARY_TRUE:
		fallthrough
	case TERNARY_FALSE:
		return ternaryPrecedence
	case ACCESS:
		fallthrough
	case FUNCTIONAL:
		return functionalPrecedence
	case SEPARATE:
		return separatePrecedence
	}

	return valuePrecedence
}

/*
Map of all valid comparators, and their string equivalents.
Used during parsing of expressions to determine if a symbol is, in fact, a comparator.
Also used during evaluation to determine exactly which comparator is being used.
*/
var comparatorSymbols = map[string]OperatorSymbol{
	"==": EQ,
	"!=": NEQ,
	">":  GT,
	">=": GTE,
	"<":  LT,
	"<=": LTE,
	"=~": REQ,
	"!~": NREQ,
	"in": IN,
}

var logicalSymbols = map[string]OperatorSymbol{
	"&&": AND,
	"||": OR,
}

var bitwiseSymbols = map[string]OperatorSymbol{
	"^": BITWISE_XOR,
	"&": BITWISE_AND,
	"|": BITWISE_OR,
}

var bitwiseShiftSymbols = map[string]OperatorSymbol{
	">>": BITWISE_RSHIFT,
	"<<": BITWISE_LSHIFT,
}

var additiveSymbols = map[string]OperatorSymbol{
	"+": PLUS,
	"-": MINUS,
}

var multiplicativeSymbols = map[string]OperatorSymbol{
	"*": MULTIPLY,
	"/": DIVIDE,
	"%": MODULUS,
}

var exponentialSymbolsS = map[string]OperatorSymbol{
	"**": EXPONENT,
}

var prefixSymbols = map[string]OperatorSymbol{
	"-": NEGATE,
	"!": INVERT,
	"~": BITWISE_NOT,
}

var ternarySymbols = map[string]OperatorSymbol{
	"?":  TERNARY_TRUE,
	":":  TERNARY_FALSE,
	"??": COALESCE,
}

// this is defined separately from additiveSymbols et al because it's needed for parsing, not stage planning.
var modifierSymbols = map[string]OperatorSymbol{
	"+":  PLUS,
	"-":  MINUS,
	"*":  MULTIPLY,
	"/":  DIVIDE,
	"%":  MODULUS,
	"**": EXPONENT,
	"&":  BITWISE_AND,
	"|":  BITWISE_OR,
	"^":  BITWISE_XOR,
	">>": BITWISE_RSHIFT,
	"<<": BITWISE_LSHIFT,
}

var separatorSymbols = map[string]OperatorSymbol{
	",": SEPARATE,
}
