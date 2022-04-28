package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func formatTime(totalSeconds int, includeSeconds bool) string {
	seconds := totalSeconds

	minutes := seconds / 60
	if minutes > 0 {
		seconds -= minutes * 60
	}

	hours := minutes / 60
	if hours > 0 {
		minutes -= hours * 60
	}

	var builder strings.Builder

	if hours > 0 {
		builder.WriteString(fmt.Sprintf("%02d:", hours))
	}

	if minutes > 0 || hours > 0 {
		builder.WriteString(fmt.Sprintf("%02d", minutes))

		if includeSeconds {
			builder.WriteByte(':')
		}
	}

	if includeSeconds {
		builder.WriteString(fmt.Sprintf("%02d", seconds))
	}

	return builder.String()
}

func uptime() {
	currentTime := time.Now().UTC()
	currentTimeTotalSeconds := currentTime.Second() +
		currentTime.Minute()*60 +
		currentTime.Hour()*60*60

	uptimeFileData, err := os.ReadFile("/proc/uptime")
	if err != nil {
		panic(err)
	}

	uptimeSecondsData := make([]byte, 0, len(uptimeFileData))

	for _, d := range uptimeFileData {
		if d != '.' && !unicode.IsDigit(rune(d)) {
			break
		}

		uptimeSecondsData = append(uptimeSecondsData, d)
	}

	uptimeSecondsFloat, err := strconv.ParseFloat(string(uptimeSecondsData), 10)
	if err != nil {
		panic(err)
	}

	fmt.Printf(
		"%s up %s\n",
		formatTime(currentTimeTotalSeconds, true),
		formatTime(int(uptimeSecondsFloat), false),
	)
}

func main() {
	uptime()
}
