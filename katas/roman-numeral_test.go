package katas

import (
	"testing"
)

func Test_num2rom(t *testing.T) {
	tests := []struct {
		inNum     int
		wantRoman string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{69, "LXIX"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if gotRoman := num2rom(tt.inNum); gotRoman != tt.wantRoman {
				t.Errorf("num2rom() = %v, want %v", gotRoman, tt.wantRoman)
			}
		})
	}
}

func Test_rom2num(t *testing.T) {
	tests := []struct {
		inRoman string
		wantNum int
	}{
		{"I", 1},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if gotNum := rom2num(tt.inRoman); gotNum != tt.wantNum {
				t.Errorf("num2rom() = %v, want %v", gotNum, tt.wantNum)
			}
		})
	}
}

func Test_bothWay(t *testing.T) {
	tests := []int{1}
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			n2r := num2rom(test)
			r2n := rom2num(n2r)
			if test != r2n {
				t.Errorf("both way for %v: num2rom=%v,rom2num=%v", test, n2r, r2n)
			}
		})
	}
}
