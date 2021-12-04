package service

import "testing"

func TestParenthesesGeneration(t *testing.T) {
	tests := []struct {
		name       string
		length     int
		wantLength int
		wantErr    error
	}{
		{
			name:       "length:1",
			length:     1,
			wantLength: 1,
			wantErr:    nil,
		},
		{
			name:       "length:100",
			length:     100,
			wantLength: 100,
			wantErr:    nil,
		},
		{
			name:       "length:0",
			length:     0,
			wantLength: 0,
			wantErr:    nil,
		},
		{
			name:       "length:-1",
			length:     -1,
			wantLength: 0,
			wantErr:    ErrNegativeLength,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParenthesesGeneration(tt.length)
			if err != tt.wantErr {
				t.Errorf("ParenthesesGeneration() error = %v, wantErr %v", err, tt.wantErr)
			}
			if len(got) != tt.wantLength {
				t.Errorf("ParenthesesGeneration() = %v, want %v", got, tt.wantLength)
			}
		})
	}
}
