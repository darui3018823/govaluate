package govaluate

import "github.com/darui3018823/govaluate/v3/token"

// TokenKind represents all valid types of expression tokens.
type TokenKind = token.TokenKind

const (
	UNKNOWN      = token.UNKNOWN
	PREFIX       = token.PREFIX
	NUMERIC      = token.NUMERIC
	BOOLEAN      = token.BOOLEAN
	STRING       = token.STRING
	PATTERN      = token.PATTERN
	TIME         = token.TIME
	VARIABLE     = token.VARIABLE
	FUNCTION     = token.FUNCTION
	SEPARATOR    = token.SEPARATOR
	ACCESSOR     = token.ACCESSOR
	COMPARATOR   = token.COMPARATOR
	LOGICALOP    = token.LOGICALOP
	MODIFIER     = token.MODIFIER
	CLAUSE       = token.CLAUSE
	CLAUSE_CLOSE = token.CLAUSE_CLOSE
	TERNARY      = token.TERNARY
)
