func Gps(s int, x []float64) int {
	if len(x) == 0 {
		return 0
	}
	delta := make([]float64, len(x)-1)
	for i, _ := range delta {
		delta[i] = x[i+1] - x[i]
	}
	speed := make([]int, len(delta))
	for id, vd := range delta {
		speed[id] = int(math.Floor((3600 * vd) / float64(s)))
	}
	result := 0
	for _, sp := range speed {
		if result < sp {
			result = sp
		}
	}
	return result
}

// top

package kata

import "math"

func Gps(s int, segments []float64) int {
  var maxSpeed float64
  for index := 1; index < len(segments); index++ {
    startSegment, endSegment := segments[index-1], segments[index]
    kmPerHour := 3600.0 * (endSegment - startSegment) / float64(s)
    maxSpeed = math.Max(maxSpeed, kmPerHour)
  }
  return int(maxSpeed)
}

// ---

package kata

import "math"

func Gps(s int, x []float64) int { //nolint
  if len(x) <= 1 {
    return 0
  }
  sect := sections(x)
  var maxHSpeed float64
  for _, v := range sect {
    if c := (3600 * v) / float64(s); c > maxHSpeed {
      maxHSpeed = c
    }
  }
  return int(math.Floor(maxHSpeed))
}

func sections(x []float64) (sect []float64) {
  for i := 1; i < len(x); i++ {
    sect = append(sect, x[i]-x[i-1])
  }
  return
}
