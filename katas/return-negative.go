package kata

func MakeNegative(x int) int {
	if x < 0 {
		return x
	} else {
		return -x
	}
}

// top

package kata

func MakeNegative(x int) int {
  if x >= 0 {
    return -x
  }
  return x
}