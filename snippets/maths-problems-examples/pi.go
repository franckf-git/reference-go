package main

import (
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/wcharczuk/go-chart/v2"
)

func main() {
	pi := 3.14159265358979323846264338327950288419716939937510582097494459
	log.Printf("pi: %#+v\n", pi) // https://en.wikipedia.org/wiki/Double-precision_floating-point_format

	pistdlib := math.Pi
	log.Printf("pistdlib: %#+v\n", pistdlib)

	fractionnaire := 355.0 / 113.0 // The return value is of same datatype as that of first operand.
	log.Printf("fractionnaire: %#+v\n", fractionnaire)

	babylone := 3 + 1.0/8.0
	log.Printf("babylone: %#+v\n", babylone)

	arybhata := 62832.0 / 20000.0
	log.Printf("arybhata: %#+v\n", arybhata)

	fractionnaire2 := 103993.0 / 33102.0
	log.Printf("fractionnaire2: %#+v\n", fractionnaire2)

	pistring := "314159265358979323846264338327950288419716939937510582097494459"
	log.Printf("pistring: %#+v\n", pistring)

	piarray := strings.Split(pistring, "")
	var piarrayfloat []float64
	var piindexfloat []float64
	for i, v := range piarray {
		ifloat := float64(i)
		piindexfloat = append(piindexfloat, ifloat)
		vint, _ := strconv.Atoi(v)
		piarrayfloat = append(piarrayfloat, float64(vint))
	}
	graph := chart.Chart{
		Series: []chart.Series{
			chart.ContinuousSeries{
				Name: "pi graph",
				Style: chart.Style{
					Padding: chart.Box{
						Top:    10,
						Left:   10,
						Right:  10,
						Bottom: 10,
						IsSet:  false,
					},
				},
				XValues: piindexfloat,
				YValues: piarrayfloat,
			},
		},
	}

	chartFile, _ := os.Create("output.png")
	defer chartFile.Close()
	graph.Render(chart.PNG, chartFile)
}
