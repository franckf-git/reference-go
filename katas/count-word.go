package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	var osStdin = os.Stdin
	visualize(sortByValues(countWord(osStdin)))
}

func countWord(input io.Reader) map[string]int {
	var results = make(map[string]int)
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanWords) // feel like cheating
	for scanner.Scan() {
		currentWord := scanner.Text()
		if validateWord(&currentWord) {
			results[currentWord] = results[currentWord] + 1
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return results
}

func validateWord(word *string) bool {
	*word = strings.ToLower(*word)
	matched, err := regexp.Match(`^[[:alnum:]]+$`, []byte(*word))
	if err != nil {
		log.Fatal(err)
	}
	if !matched {
		if len(*word) == 1 {
			return false
		}
		re := regexp.MustCompile(`[^\w|àâçéèêëîïôûùüÿñæœ]$`)
		*word = re.ReplaceAllString(*word, "")
	}
	if *word == "" {
		return false
	}
	return true
}

func sortByValues(inputs map[string]int) []Counter {
	unsorted := []Counter{}
	for word, count := range inputs {
		unsorted = append(unsorted, Counter{
			word:  word,
			count: count,
		})
	}
	sorted := quickSort(unsorted)
	return sorted
}

type Counter struct {
	word  string
	count int
}

// this is really cheating, get one of my old quick sort
func quickSort(datas []Counter) []Counter {
	if len(datas) == 0 {
		return []Counter{}
	}
	pivot := datas[len(datas)-1]
	smallerValues := []Counter{}
	largerValues := []Counter{}
	for index := 0; index < len(datas)-1; index++ {
		if datas[index].count > pivot.count {
			smallerValues = append(smallerValues, datas[index])
		} else {
			largerValues = append(largerValues, datas[index])
		}
	}
	qsSmallerPivot := append(quickSort(smallerValues), pivot)
	return append(qsSmallerPivot, quickSort(largerValues)...)
}

func visualize(datas []Counter) {
	var biggerSize int
	for _, data := range datas {
		if len(data.word) > biggerSize {
			biggerSize = len(data.word)
		}
	}

	fmt.Print("+", "--", strings.Repeat("-", biggerSize), "--", "+", "-------", "+", "\n")
	fmt.Print("+ ", "word", strings.Repeat(" ", biggerSize-2), " + ", "count", " +", "\n")
	for _, data := range datas {
		fmt.Print("+ ", data.word, strings.Repeat(" ", biggerSize-len(data.word)), "   +   ", data.count, "   +", "\n")
	}
	fmt.Print("+", "--", strings.Repeat("-", biggerSize), "--", "+", "-------", "+", "\n")
}
