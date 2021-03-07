package kata

import "strings"

func AbbrevName(name string) string {
  namearr := strings.Split(name, " ")
  var result string
  for _, value := range namearr {
    result = result + value[:1] + "."
  }
  resUpp := strings.ToUpper(result[:3])
  return resUpp
}

// top

package kata

import "strings"

func AbbrevName(name string) string{
  words := strings.Split(name, " ")
  return strings.ToUpper(string(words[0][0])) + "." + strings.ToUpper(string(words[1][0]))
}

// ---

package kata
import "strings"
func AbbrevName(name string) string{
  var x = strings.Index(name, " ")
  return strings.ToUpper(string(name[0])+"."+string(name[x+1]))

}
