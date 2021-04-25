package kata

import "strings"

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return result
}

func VertMirror(s string) string {
	lines := strings.Split(s, "\n")

	result := make([]string, len(lines))

	for ind, word := range lines {
		result[ind] = reverse(word)
	}

	return strings.Join(result, "\n")
}
func HorMirror(s string) string {
	lines := strings.Split(s, "\n")
	result := make([]string, len(lines))

	for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = lines[j], lines[i]
	}

	return strings.Join(result, "\n")
}

type FParam func(string) string

func Oper(f FParam, x string) string {
	return f(x)
}

// top

package kata

import "strings"

func VertMirror(s string) string {

  var b strings.Builder
  b.Grow(len(s) + 1)

  n := strings.Index(s, "\n")

  for i := 0; i < len(s); i += n + 1 {
    for j := i + n - 1; j >= i; j-- {
      b.WriteByte(s[j])
    }
    b.WriteByte('\n')
  }
  v := b.String()
  v = v[:b.Len()-1]

  return v
}

func HorMirror(s string) string {

  var b strings.Builder
  b.Grow(len(s) + 1)

  n := strings.Index(s, "\n")

  for i := len(s) - n; i >= 0; i -= (n + 1) {
    b.WriteString(s[i : i+n])
    b.WriteByte('\n')
  }
  v := b.String()
  v = v[:b.Len()-1]

  return v
}

type FParam func(string) string
func Oper(f FParam, x string) string {
    return f(x)
}
