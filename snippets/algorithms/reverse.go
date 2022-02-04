package main

func main() {
	//start := "😎⚽"
	start := "123456789"
	end := reverseRuneMiddle(start)
	print(end)
	print("\n")
}

func reverseRune(in string) string {
	runes := []rune(in)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[j], runes[i] = runes[i], runes[j]
	}
	return string(runes)
}

func reverseRune2(in string) string {
	runes := []rune(in)
	for i := 0; i < len(runes)-1; i = i + 1 {
		for j := len(runes) - 1; i < j; j = j - 1 {
			runes[j], runes[i] = runes[i], runes[j]
		}
	}
	return string(runes)
}

func reverseRuneSplit(in string) string {
	runes := []rune(in)
	for i := 0; i < len(runes)/2; i++ {
		j := len(runes) - i - 1
		runes[j], runes[i] = runes[i], runes[j]
	}
	return string(runes)
}

func reverseRuneMiddle(in string) string {
	runes := []rune(in)
	start := 0
	end := len(runes) - 1
	for start < end {
		runes[start], runes[end] = runes[end], runes[start]
		start++
		end--
	}
	return string(runes)
}

func reverseRuneInefficient(in string) (out string) {
	runes := []rune(in)
	var runesOut []rune
	for i := len(runes) - 1; i >= 0; i-- {
		runesOut = append(runesOut, runes[i])
	}
	out = string(runesOut)
	return
}

func reverseInefficient(in string) (out string) {
	for _, v := range in {
		out = string(v) + out
	}
	return
}
