package kata

func MxDifLg(a1 []string, a2 []string) int {
	if len(a1) == 0 || len(a2) == 0 {
		return -1
	}
	a1len := make([]int, len(a1))
	a2len := make([]int, len(a2))

	for i1, s1 := range a1 {
		a1len[i1] = len(s1)
	}
	for i2, s2 := range a2 {
		a2len[i2] = len(s2)
	}

	var a1max int
	a1min := a1len[0]
	var a2max int
	a2min := a2len[0]

	for in1, n1 := range a1len {
		if a1max < a1len[in1] {
			a1max = n1
		}
		if a1min > a1len[in1] {
			a1min = n1
		}
	}
	for in2, n2 := range a2len {
		if a2max < a2len[in2] {
			a2max = n2
		}
		if a2min > a2len[in2] {
			a2min = n2
		}
	}

	right := a2max - a1min
	left := a1max - a2min
	if right > left {
		return right
	}
	return left
}

// top

package kata

import "math"

func findShortestAndLongest(a []string) (shortest, longest int) {
  lA := 0
  sA := math.MaxInt64
  for _, s := range a {
    if len(s) > lA {
      lA = len(s)
    }
    if len(s) < sA {
      sA = len(s)
    }
  }

  return sA, lA
}

func MxDifLg(a1 []string, a2 []string) int {
    if (len(a1) == 0 || len(a2) == 0) {
      return -1
    }

    sA1, lA1 := findShortestAndLongest(a1)
    sA2, lA2 := findShortestAndLongest(a2)

    res1 := math.Abs(float64(sA1 - lA2))
    res2 := math.Abs(float64(sA2 - lA1))

    if res1 > res2 {
      return int(res1)
    }

    return int(res2)
}

// ---

package kata

import "math"

func AbsDiff(a, b int) int {
  return int(math.Abs(float64(a) - float64(b)))
}

func MxDifLg(a, b []string) int {
  max, cur := -1, 0
  for _, x := range a {
    for _, y := range b {
      cur = AbsDiff(len(x), len(y))
      if cur > max {
        max = cur
      }
    }
  }
  return max
}
