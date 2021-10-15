package main

import (
	"database/sql"
	"log"
	"os"
	"strings"

	"github.com/wcharczuk/go-chart/v2"
)

func main() {
	db := openDatabase()
	defer db.Close()
	startDatabase(db)

	var datas = getAlphabetFreq(db, 0)
	createChart(datas)
	log.Println("Chart created.")
}

func getAlphabetFreq(db *sql.DB, limit int) [27]float64 {
	if limit == 0 {
		limit = countValues(db)
	}

	alphabet := [27]float64{}

	rows, err := db.Query("SELECT title,links FROM doc LIMIT ?", limit)
	if err != nil {
		log.Fatal("Select fail - executing query:", err)
	}
	defer rows.Close()
	for rows.Next() {
		var title string
		var links int
		err = rows.Scan(&title, &links)
		if err != nil {
			log.Fatal("Select fail - scanning values:", err)
		}
		letter := parsingTitle(title)
		position := alphaToPos(letter)
		alphabet[position] = float64(links) + alphabet[position]
	}
	err = rows.Err()
	if err != nil {
		log.Fatal("Select fail - reading rows:", err)
	}

	return alphabet
}

func createChart(datas [27]float64) {

	values := []chart.Value{}

	for index, value := range datas {
		label := posToAlpha(index)
		adding := []chart.Value{{Label: label, Value: value}}
		values = append(values, adding...)
	}

	graph := chart.BarChart{
		Title:    "Stats - Links by Letter",
		Width:    1800,
		Height:   1000,
		BarWidth: 50,
		Bars:     values,
	}

	chartFile, _ := os.Create("output.png")
	defer chartFile.Close()
	graph.Render(chart.PNG, chartFile)
}

func parsingTitle(title string) string {
	withoutWikipedia := strings.Replace(title, "Wikipedia: ", "", -1)
	return withoutWikipedia[:1]
}

func alphaToPos(firstCharac string) int {
	lowerC := strings.ToLower(firstCharac)
	converRune := []rune(lowerC)
	interval := int(converRune[0]) - 96
	if interval >= 1 && interval <= 26 {
		return interval
	}
	return 0
}

func posToAlpha(pos int) string {
	if pos == 0 {
		return "spe"
	}
	ascii := pos + 64
	return string(rune(ascii))
}