package kata

func Greet(name string) string {
  return "Hello, " + name + " how are you doing today?"
}

// top

package kata

import "fmt"

func Greet(name string) string {
  return fmt.Sprintf("Hello, %s how are you doing today?", name)
}
