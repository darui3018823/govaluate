package datetime

import "testing"

func TestParseSupportedFormat(t *testing.T) {
	parsed, ok := Parse("2014-01-02 15:04:05")
	if !ok {
		t.Fatalf("expected supported date format to parse")
	}

	if parsed.Year() != 2014 || parsed.Month() != 1 || parsed.Day() != 2 {
		t.Fatalf("unexpected parsed date: %v", parsed)
	}
}

func TestParseUnsupportedFormat(t *testing.T) {
	if parsed, ok := Parse("not a date"); ok {
		t.Fatalf("expected unsupported date to fail, got %v", parsed)
	}
}
