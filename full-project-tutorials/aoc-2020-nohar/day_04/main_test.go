package main

import "testing"

func TestHeightValidator(t *testing.T) {
	cases := []struct {
		Input    string
		Expected bool
	}{
		{"170", false},
		{"170in", false},
		{"170cm", true},
		{"cm", false},
		{"60", false},
		{"60in", true},
		{"60cm", false},
	}
	for _, c := range cases {
		if isValidHeight(c.Input) != c.Expected {
			t.Errorf("Input: %q, Expected: %t", c.Input, c.Expected)
		}
	}
}

func TestYearValidator(t *testing.T) {
	isValidYear := yearValidator(1000, 2000)
	cases := []struct {
		Input    string
		Expected bool
	}{
		{"1984", true},
		{"abc", false},
		{"984", false},
		{"2984", false},
	}

	for _, c := range cases {
		if isValidYear(c.Input) != c.Expected {
			t.Errorf("Input: %q, Expected: %t", c.Input, c.Expected)
		}
	}
}

func TestPart2(t *testing.T) {
	cases := []struct {
		Passport Passport
		Expected bool
	}{
		{
			Passport{"eyr": "1972", "cid": "100", "hcl": "#18171d", "ecl": "amb", "hgt": "170", "pid": "186cm", "iyr": "2018", "byr": "1926"},
			false,
		},
		{
			Passport{"iyr": "2019", "hcl": "#602927", "eyr": "1967", "hgt": "170cm", "ecl": "grn", "pid": "012533040", "byr": "1946"},
			false,
		},
		{
			Passport{"iyr": "2019", "hcl": "#602927", "eyr": "1967", "hgt": "170cm", "ecl": "grn", "pid": "012533040", "byr": "1946"},
			false,
		},
		{
			Passport{"iyr": "2019", "hcl": "#602927", "eyr": "1967", "hgt": "170cm", "ecl": "grn", "pid": "012533040", "byr": "1946"},
			false,
		},
		{
			Passport{"hcl": "dab227", "iyr": "2012", "ecl": "brn", "hgt": "182cm", "pid": "021572410", "eyr": "2020", "byr": "1992", "cid": "277"},
			false,
		},
		{
			Passport{"hgt": "59cm", "ecl": "zzz", "eyr": "2038", "hcl": "74454a", "iyr": "2023", "pid": "3556412378", "byr": "2007"},
			false,
		},
		{
			Passport{"pid": "087499704", "hgt": "74in", "ecl": "grn", "iyr": "2012", "eyr": "2030", "byr": "1980", "hcl": "#623a2f"},
			true,
		},
		{
			Passport{"eyr": "2029", "ecl": "blu", "cid": "129", "byr": "1989", "iyr": "2014", "pid": "896056539", "hcl": "#a97842", "hgt": "165cm"},
			true,
		},
		{
			Passport{"hcl": "#888785", "hgt": "164cm", "byr": "2001", "iyr": "2015", "cid": "88", "pid": "545766238", "ecl": "hzl", "eyr": "2022"},
			true,
		},
		{
			Passport{"iyr": "2010", "hgt": "158cm", "hcl": "#b6652a", "ecl": "blu", "byr": "1944", "eyr": "2021", "pid": "093154719"},
			true,
		},
	}

	for _, c := range cases {
		if isValidPart2(c.Passport) != c.Expected {
			t.Errorf("Input: %v, Expected: %t", c.Passport, c.Expected)
		}
	}
}
