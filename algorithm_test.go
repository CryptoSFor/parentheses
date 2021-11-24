package parentheses

import (
	"reflect"
	"testing"
)

func Test_isOpen(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  string
	}{
		{
			name:  "return {",
			value: "{",
			want:  "{",
		},
		{
			name:  "return [",
			value: "[",
			want:  "[",
		},
		{
			name:  "return {",
			value: "(",
			want:  "(",
		},
		{
			name:  "return empty string",
			value: ")",
			want:  "",
		},
		{
			name:  "return empty string",
			value: "0",
			want:  "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isOpen(tt.value); got != tt.want {
				t.Errorf("isOpen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isClose(t *testing.T) {
	tests := []struct {
		name string
		v    string
		want string
	}{
		{
			name: "return }",
			v:    "}",
			want: "}",
		},
		{
			name: "return ]",
			v:    "]",
			want: "]",
		},
		{
			name: "return )",
			v:    ")",
			want: ")",
		},
		{
			name: "return empty string",
			v:    "{",
			want: "",
		},
		{
			name: "return empty string",
			v:    "1",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isClose(tt.v); got != tt.want {
				t.Errorf("isClose() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
			args: args{x1: "}", x2: "{"},
			want: true,
		},
		{
			name: "comparison of ) and (",
			args: args{x1: ")", x2: "("},
			want: true,
		},
		{
			name: "comparison of ] and [",
			args: args{x1: "]", x2: "["},
			want: true,
		},
		{
			name: "comparison of empty string and [",
			args: args{x1: "", x2: "["},
			want: false,
		},
		{
			name: "comparison of { and }",
			args: args{x1: "{", x2: "}"},
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

func Test_pop(t *testing.T) {
	tests := []struct {
		name      string
		v         *[]string
		want      string
		wantErr   error
		wantSlice []string
	}{
		{
			name:      "pop from []string{'{', '}'}",
			v:         &[]string{"{", "}"},
			want:      "}",
			wantErr:   nil,
			wantSlice: []string{"{"},
		},
		{
			name:      "pop from []string{'{'}",
			v:         &[]string{"{"},
			want:      "{",
			wantErr:   nil,
			wantSlice: []string{},
		},
		{
			name:      "pop from []string{}",
			v:         &[]string{},
			want:      "the slice is empty",
			wantErr:   ErrEmptySlice,
			wantSlice: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := pop(tt.v); got != tt.want || !reflect.DeepEqual(*tt.v, tt.wantSlice) || err != tt.wantErr {
				t.Errorf("pop() = %v, want = %v, slice = %v wantSlice = %v", got, tt.want, *tt.v, tt.wantSlice)
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
			s:    "(){",
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
