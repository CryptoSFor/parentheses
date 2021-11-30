package parentheses

import (
	"testing"
)

func TestBalancedString(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "balanced1",
			s:    "{([({}{}[][])])}",
			want: true,
		},
		{
			name: "balanced2",
			s:    "(те[с]т){}",
			want: true,
		},
		{
			name: "unbalanced1",
			s:    "()))",
			want: false,
		},
		{
			name: "unbalanced2",
			s:    "((()",
			want: false,
		},
		{
			name: "unbalanced3",
			s:    "(}",
			want: false,
		},
		{
			name: "unbalanced4",
			s:    "((1 + 2) * 3) - 4)/5",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BalancedString(tt.s); got != tt.want {
				t.Errorf("BalancedString() = %v, want %v", got, tt.want)
			}
		})
	}
}
