package kata
import "strings"
func MakeUpperCase(str string) string {
    return strings.ToUpper(str)
}

// not top but something

package kata

func MakeUpperCase(str string) string {
    STR := ""
    for _, i := range str {
      if i >= 97 && i <= 123 {
        STR += string(i-32)
      } else {
        STR += string(i)
      }
    }
    return STR 
}
