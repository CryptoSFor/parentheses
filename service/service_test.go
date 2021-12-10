package service

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestParenthesesGeneration(t *testing.T) {
	tests := []struct {
		name       string
		length     int
		wantLength int
	}{
		{
			name:       "length:1",
			length:     1,
			wantLength: 1,
		},
		{
			name:       "length:100",
			length:     100,
			wantLength: 100,
		},
		{
			name:       "length:0",
			length:     0,
			wantLength: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParenthesesGeneration(tt.length)
			if len(got) != tt.wantLength {
				t.Errorf("ParenthesesGeneration() = %v, want %v", got, tt.wantLength)
			}
		})
	}
}

func TestGenerateHandler(t *testing.T) {
	tests := []struct {
		name       string
		length     int
		expLength  int
		httpStatus int
	}{
		{
			name:       "len of the sequence : 1",
			length:     1,
			expLength:  1,
			httpStatus: 200,
		},
		{
			name:       "len of the sequence : 0",
			length:     0,
			expLength:  0,
			httpStatus: 200,
		},
		{
			name:       "len of the sequence : -1",
			length:     -1,
			expLength:  0,
			httpStatus: 400,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := fmt.Sprintf("/generate?n=%v", tt.length)
			req := httptest.NewRequest(http.MethodGet, url, nil)
			w := httptest.NewRecorder()
			GenerateHandler(w, req)
			res := w.Result()

			defer res.Body.Close()
			data, err := ioutil.ReadAll(res.Body)

			if err != nil {
				t.Errorf("expected error to be nil got %v", err)
			}

			if tt.expLength != len(data) {
				t.Errorf("expected %v got %v", tt.expLength, len(data))
			}

			if status := w.Code; status != tt.httpStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.httpStatus)
			}
		})
	}
}
