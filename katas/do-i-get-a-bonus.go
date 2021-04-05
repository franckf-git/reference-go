package kata

import ("strconv")

func BonusTime(salary int, bonus bool) string {
  if bonus {
    salary = salary * 10
    return "£" + strconv.Itoa(salary)
  }
    return "£" + strconv.Itoa(salary)
}

// top

package kata
import ( "fmt" )


func BonusTime(salary int, bonus bool) string {
  if bonus {
    salary = salary * 10
  }
  return fmt.Sprintf("£%d", salary)
}
