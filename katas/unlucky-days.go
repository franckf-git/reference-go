package kata
import "time"
func UnluckyDays(year int) int {

	friday13 := 0
	for i := 1; i <= 12; i++ {
		day := time.Date(year, time.Month(i), 13, 00, 00, 0, 0, time.UTC)
		if day.Weekday() == time.Weekday(5) {
			friday13++
		}
	}
	return friday13
}

// top

package kata 

import "time"

func UnluckyDays(year int) int {
  tot := 0
  for m:=1; m<=12; m++ {
    if time.Date(year, time.Month(m), 13, 0, 0, 0, 0, time.UTC).Weekday() == 5 {
      tot++
    }
  }
  return tot
}
