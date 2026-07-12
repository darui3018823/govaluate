package token

import "testing"

func TestTokenKindString(t *testing.T) {
	if NUMERIC.String() != "NUMERIC" {
		t.Fatalf("expected NUMERIC string, got %q", NUMERIC.String())
	}

	if TokenKind(-1).String() != "UNKNOWN" {
		t.Fatalf("expected unknown token kind fallback")
	}
}

func TestOperatorSymbolString(t *testing.T) {
	if EXPONENT.String() != "**" {
		t.Fatalf("expected exponent string, got %q", EXPONENT.String())
	}

	if OperatorSymbol(-1).String() != "" {
		t.Fatalf("expected unknown operator fallback")
	}
}

func TestOperatorSymbolIsModifierType(t *testing.T) {
	if !PLUS.IsModifierType([]OperatorSymbol{MINUS, PLUS}) {
		t.Fatalf("expected PLUS to match candidate set")
	}

	if AND.IsModifierType([]OperatorSymbol{MINUS, PLUS}) {
		t.Fatalf("did not expect AND to match candidate set")
	}
}
