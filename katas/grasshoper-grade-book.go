package kata

func GetGrade(a,b,c int) rune {
  moy := (a+b+c)/3
  switch {
case 90 <= moy:
	return []rune("A")[0]
case 80 <= moy:
	return []rune("B")[0]
case 70 <= moy:
	return []rune("C")[0]
case 60 <= moy:
	return []rune("D")[0]
default:
	return []rune("F")[0]
}
}

// top

package kata

func GetGrade(a,b,c int) rune {
    switch mean := (a+b+c)/3; {
    case mean < 60:
    return 'F'
    case mean < 70:
    return 'D'
    case mean < 80:
    return 'C'
    case mean < 90:
    return 'B'
    default:
    return 'A'
    }

}

// ---

package kata

func GetGrade(a, b, c int) rune {
    switch ((a + b + c) / 30) {
        case 10: return 'A'
        case 9: return 'A'
        case 8: return 'B'
        case 7: return 'C'
        case 6: return 'D'
        default: return 'F'
    }
}
