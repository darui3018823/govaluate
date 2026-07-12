// Package token defines the public token and operator model used by govaluate.
package token

// TokenKind represents all valid types of expression tokens.
type TokenKind int

const (
	UNKNOWN TokenKind = iota

	PREFIX
	NUMERIC
	BOOLEAN
	STRING
	PATTERN
	TIME
	VARIABLE
	FUNCTION
	SEPARATOR
	ACCESSOR

	COMPARATOR
	LOGICALOP
	MODIFIER

	CLAUSE
	CLAUSE_CLOSE

	TERNARY
)

func (kind TokenKind) String() string {
	switch kind {
	case PREFIX:
		return "PREFIX"
	case NUMERIC:
		return "NUMERIC"
	case BOOLEAN:
		return "BOOLEAN"
	case STRING:
		return "STRING"
	case PATTERN:
		return "PATTERN"
	case TIME:
		return "TIME"
	case VARIABLE:
		return "VARIABLE"
	case FUNCTION:
		return "FUNCTION"
	case SEPARATOR:
		return "SEPARATOR"
	case COMPARATOR:
		return "COMPARATOR"
	case LOGICALOP:
		return "LOGICALOP"
	case MODIFIER:
		return "MODIFIER"
	case CLAUSE:
		return "CLAUSE"
	case CLAUSE_CLOSE:
		return "CLAUSE_CLOSE"
	case TERNARY:
		return "TERNARY"
	case ACCESSOR:
		return "ACCESSOR"
	}

	return "UNKNOWN"
}

// ExpressionToken represents a single parsed token.
type ExpressionToken struct {
	Kind  TokenKind
	Value any
}

// OperatorSymbol represents the valid symbols for operators.
type OperatorSymbol int

const (
	VALUE OperatorSymbol = iota
	LITERAL
	NOOP
	EQ
	NEQ
	GT
	LT
	GTE
	LTE
	REQ
	NREQ
	IN

	AND
	OR

	PLUS
	MINUS
	BITWISE_AND
	BITWISE_OR
	BITWISE_XOR
	BITWISE_LSHIFT
	BITWISE_RSHIFT
	MULTIPLY
	DIVIDE
	MODULUS
	EXPONENT

	NEGATE
	INVERT
	BITWISE_NOT

	TERNARY_TRUE
	TERNARY_FALSE
	COALESCE

	FUNCTIONAL
	ACCESS
	SEPARATE
)

// IsModifierType returns true if this operator is contained by the given
// candidate symbols.
func (symbol OperatorSymbol) IsModifierType(candidate []OperatorSymbol) bool {
	for _, symbolType := range candidate {
		if symbol == symbolType {
			return true
		}
	}

	return false
}

func (symbol OperatorSymbol) String() string {
	switch symbol {
	case NOOP:
		return "NOOP"
	case VALUE:
		return "VALUE"
	case EQ:
		return "="
	case NEQ:
		return "!="
	case GT:
		return ">"
	case LT:
		return "<"
	case GTE:
		return ">="
	case LTE:
		return "<="
	case REQ:
		return "=~"
	case NREQ:
		return "!~"
	case AND:
		return "&&"
	case OR:
		return "||"
	case IN:
		return "in"
	case BITWISE_AND:
		return "&"
	case BITWISE_OR:
		return "|"
	case BITWISE_XOR:
		return "^"
	case BITWISE_LSHIFT:
		return "<<"
	case BITWISE_RSHIFT:
		return ">>"
	case PLUS:
		return "+"
	case MINUS:
		return "-"
	case MULTIPLY:
		return "*"
	case DIVIDE:
		return "/"
	case MODULUS:
		return "%"
	case EXPONENT:
		return "**"
	case NEGATE:
		return "-"
	case INVERT:
		return "!"
	case BITWISE_NOT:
		return "~"
	case TERNARY_TRUE:
		return "?"
	case TERNARY_FALSE:
		return ":"
	case COALESCE:
		return "??"
	}

	return ""
}
