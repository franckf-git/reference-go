package kata

func Solve(s int, g int) []int {
	secondnum := s-g
  if s%g == 0 && secondnum%g == 0 {
    return []int{g,secondnum}
  }
  return []int{-1,-1}
}

// top

package kata

func Solve(sum int, gcd int) []int {
  if sum % gcd != 0 {
    return []int{-1, -1}
  }
  return []int{gcd, sum-gcd}
}

// ---

package kata

func Solve(s int, g int) []int {
    if s%g > 0 {
        return []int{-1,-1}
    }
    return []int{g,s-g}
}
