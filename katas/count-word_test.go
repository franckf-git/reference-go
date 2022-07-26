package main

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func Test_countWord(t *testing.T) {
	type args struct {
		input io.Reader
	}
	tests := []struct {
		name        string
		args        args
		wantResults map[string]int
	}{
		{
			name: "one line",
			args: args{
				input: strings.NewReader("un deux trois deux trois trois"),
			},
			wantResults: map[string]int{"un": 1, "deux": 2, "trois": 3},
		},
		{
			name: "uppercases",
			args: args{
				input: strings.NewReader("un Deux Trois deux TROIS trois"),
			},
			wantResults: map[string]int{"un": 1, "deux": 2, "trois": 3},
		},
		{
			name: "simple punctuation",
			args: args{
				input: strings.NewReader("un Deux, Trois : deux TROIS trois."),
			},
			wantResults: map[string]int{"un": 1, "deux": 2, "trois": 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResults := countWord(tt.args.input); !reflect.DeepEqual(gotResults, tt.wantResults) {
				t.Errorf("countWord() = %v, want %v", gotResults, tt.wantResults)
			}
		})
	}
}
