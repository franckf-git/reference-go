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

func reverseBatch(input []int) []int {
	const batchSize = 1000
	inputLen := len(input)
	output := make([]int, inputLen)
	var wg sync.WaitGroup

	for i := 0; i < inputLen; i += batchSize {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			l := i + batchSize
			if l > inputLen {
				l = inputLen
			}
			for ; i < l; i++ {
				j := inputLen - i - 1
				n := input[i]
				output[j] = n
			}
		}(i)
	}

	wg.Wait()
	return output
}

func reverseBatchMiddle(input []int) {
	const batchSize = 1000
	inputLen := len(input)
	inputMid := inputLen / 2
	var wg sync.WaitGroup

	for i := 0; i < inputMid; i += batchSize {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			l := i + batchSize
			if l > inputMid {
				l = inputMid
			}
			for ; i < l; i++ {
				j := inputLen - i - 1
				input[i], input[j] = input[j], input[i]
			}
		}(i)
	}

	wg.Wait()
}

func reverseDefer(input string) (output string) {
	var outputBuilder strings.Builder
	defer func() {
	output = outputBuilder.String()
	}()
	for _, r := range input {
		defer func(r rune) {
			outputBuilder.WriteRune(r)
		}(r)
	}
	return
}
