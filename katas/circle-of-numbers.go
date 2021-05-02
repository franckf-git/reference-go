package kata

func CircleOfNumbers(n int, firstNumber int) int {
	if firstNumber < n/2 {
		return n/2 + firstNumber
	}
	return Abs(n/2 - firstNumber)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// top

package kata

func CircleOfNumbers(n int, firstNumber int) int {
    return (firstNumber + n / 2) % n
}

// ---

package kata

func CircleOfNumbers(n int, firstNumber int) int {
  return (n >> 1 + firstNumber) % n
}
