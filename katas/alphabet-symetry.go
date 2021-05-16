package main

import (
	"fmt"
	"strings"
)

func solve(slice []string) []int {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	countOK := make([]int, len(slice))
	for i, v := range slice {
		for pos, let := range v {
			letS := strings.ToLower(string(let))
			if letS == alphabet[pos:pos+1] {
				countOK[i] = countOK[i] + 1
			}

		}
	}
	return countOK
}

func main() {
	fmt.Println(solve([]string{"abode", "ABc", "xyzD"}))
}

// top
package kata

func solve(slice []string) []int {
  result := make([]int, len(slice))

  for i, str := range slice {
    count := 0

    for i := 0; i < len(str); i++ {
      if int(str[i]) & 31 == i+1 {
        count++
      }
    }

    result[i] = count
  }

  return result
}

// ---
package kata

import "unicode"

func solve(slice []string) []int {
  alphabet := "abcdefghijklmnopqrstuvwxyz"
  results := make([]int, len(slice))
  // Loop strings
  for sliceIndex, str := range slice {
    // Loop characters
    for charIndex, character := range str {
      // avoid match case
      if unicode.ToLower(character) == rune(alphabet[charIndex]) {
        results[sliceIndex]++
      }
    }
  }
  return results
}

