package main

import (
	"testing"
)

func Test_parsingTitle(t *testing.T) {
	type args struct {
		title string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"one word", args{title: "Wikipedia: Albedo"}, "A"},
		{"two words", args{title: "Wikipedia: Julie Doucet"}, "J"},
		{"one letter", args{title: "Wikipedia: A"}, "A"},
		{"sentence", args{title: "Wikipedia: An American in Paris"}, "A"},
		{"long", args{title: "Wikipedia: Academy Award for Best Production Design"}, "A"},
		{"not only A", args{title: "Wikipedia: Mitochondrion"}, "M"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parsingTitle(tt.args.title); got != tt.want {
				t.Errorf("parsingTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_posAlphabet(t *testing.T) {
	type args struct {
		firstCharac string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "uppercase", args: args{firstCharac: "A"}, want: 1},
		{name: "lowercase", args: args{firstCharac: "p"}, want: 16},
		{name: "special1", args: args{firstCharac: "%"}, want: 0},
		{name: "special2", args: args{firstCharac: "&"}, want: 0},
		{name: "special3", args: args{firstCharac: " "}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := posAlphabet(tt.args.firstCharac); got != tt.want {
				t.Errorf("posAlphabet() = %v, want %v", got, tt.want)
			}
		})
	}
}
