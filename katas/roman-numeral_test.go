package katas

import (
	"fmt"
	"testing"
)

func Test_num2rom(t *testing.T) {
	tests := []struct {
		inNum     int
		wantRoman string
	}{
		{1, "I"},
		{2, "II"},
		{3, "III"},
		{4, "IV"},
		{5, "V"},
		{6, "VI"},
		{7, "VII"},
		{8, "VIII"},
		{9, "IX"},
		{10, "X"},
		{14, "XIV"},
		{18, "XVIII"},
		{20, "XX"},
		{39, "XXXIX"},
		{40, "XL"},
		{47, "XLVII"},
		{49, "XLIX"},
		{50, "L"},
		{69, "LXIX"},
		{100, "C"},
		{90, "XC"},
		{400, "CD"},
		{500, "D"},
		{900, "CM"},
		{1000, "M"},
		{1984, "MCMLXXXIV"},
		{3999, "MMMCMXCIX"},
		{2014, "MMXIV"},
		{1006, "MVI"},
		{798, "DCCXCVIII"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.inNum), func(t *testing.T) {
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
