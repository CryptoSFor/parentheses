package parentheses

import (
	"testing"
)

func Test_pairOfParentheses(t *testing.T) {
	type args struct {
		x1 string
		x2 string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "comparison of } and {",
			args: args{x1: "{", x2: "}"},
			want: true,
		},
		{
			name: "comparison of ) and (",
			args: args{x1: "(", x2: ")"},
			want: true,
		},
		{
			name: "comparison of ] and [",
			args: args{x1: "[", x2: "]"},
			want: true,
		},
		{
			name: "comparison of { and }",
			args: args{x1: "}", x2: "{"},
			want: false,
		},
		{
			name: "comparison of empty string and [",
			args: args{x1: "", x2: "]"},
			want: false,
		},
		{
			name: "comparison of empty string and empty string",
			args: args{x1: "", x2: ""},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pairOfParentheses(tt.args.x1, tt.args.x2); got != tt.want {
				t.Errorf("isEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
			s:    "效{[复]制___()+{仿效_}}制",
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
