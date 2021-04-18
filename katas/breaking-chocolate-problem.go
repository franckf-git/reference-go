package kata

func BreakChocolate(n, m int) int {
  if n == 1 && m == 1 {
    return 0
  }
  if n == 0 || m == 0 {
    return 0
  }
  return n*m-1
}

// top

package kata

func BreakChocolate(n, m int) int {
  if n < 1 || m < 1 {
    return 0
  }

  return (n * m) - 1
}
