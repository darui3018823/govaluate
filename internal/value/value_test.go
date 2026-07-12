package value

import "testing"

func TestNormalizeNumber(t *testing.T) {
	tests := []struct {
		name  string
		input any
		want  any
	}{
		{name: "int", input: int(7), want: float64(7)},
		{name: "uint64", input: uint64(9), want: float64(9)},
		{name: "float32", input: float32(1.5), want: float64(1.5)},
		{name: "float64 unchanged", input: float64(2.5), want: float64(2.5)},
		{name: "string unchanged", input: "value", want: "value"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NormalizeNumber(tt.input); got != tt.want {
				t.Fatalf("NormalizeNumber(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
