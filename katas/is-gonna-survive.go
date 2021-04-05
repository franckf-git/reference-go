package kata

func Hero(bullets, dragons int) bool {
  if bullets / dragons >= 2 {
    return true
  }
  return false
}

// top

package kata

func Hero(bullets, dragons int) bool {
  return bullets >= 2*dragons
}
