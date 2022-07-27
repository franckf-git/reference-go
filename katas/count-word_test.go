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
				input: strings.NewReader("un Deux, Trois : deux TROIS trois. l'avion"),
			},
			wantResults: map[string]int{"un": 1, "deux": 2, "trois": 3, `l'avion`: 1},
		},
		{
			name: "multilines",
			args: args{
				input: strings.NewReader(`
				Maître Renard par l’odeur alléché
				Lui tint à peu près ce langage :
				`),
			},
			wantResults: map[string]int{"alléché": 1, "ce": 1, "langage": 1, "lui": 1, "l’odeur": 1, "maître": 1, "par": 1, "peu": 1, "près": 1, "renard": 1, "tint": 1, "à": 1},
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

func Test_sortByValues(t *testing.T) {
	type args struct {
		inputs map[string]int
	}
	tests := []struct {
		name       string
		args       args
		wantSorted []Counter
	}{
		{
			name: "simple sort",
			args: args{
				inputs: map[string]int{"un": 1, "deux": 2, "trois": 3},
			},
			wantSorted: []Counter{{
				word:  "trois",
				count: 3,
			}, {
				word:  "deux",
				count: 2,
			}, {
				word:  "un",
				count: 1,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSorted := sortByValues(tt.args.inputs); !reflect.DeepEqual(gotSorted, tt.wantSorted) {
				t.Errorf("sortByValues() = %v, want %v", gotSorted, tt.wantSorted)
			}
		})
	}
}
