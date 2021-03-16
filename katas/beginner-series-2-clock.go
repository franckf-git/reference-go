package kata

func Past(h, m, s int) int {
    var result int
	for i := 0; i < s; i++ {
		result = result + 1000
	}
	for i := 0; i < m; i++ {
		result = result + 60000
	}
	for i := 0; i < h; i++ {
		result = result + 3600000
	}
return result
}

// top

package kata

func Past(h, m, s int) int {
    return (h*3600000 + m*60000 + s*1000)
}

// ---

package kata

func Past(h, m, s int) int {
    return (h*60*60+m*60+s)*1000
}
