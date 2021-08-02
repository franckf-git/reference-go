package xkcd

import (
	"bufio"
	"fmt"
	"strings"
	"time"
)

// A Comic holds all the metadata associated to an xkcd comic.
type Comic struct {
	Num        int    `json:"num"`
	Img        string `json:"img"`
	Year       string `json:"year"`
	Month      string `json:"month"`
	Day        string `json:"day"`
	Title      string `json:"title"`
	SafeTitle  string `json:"safe_title"`
	Alt        string `json:"alt"`
	Transcript string `json:"transcript"`
	News       string `json:"news"`
	Link       string `json:"link"`
}

// URL returns the URL of the xkcd page for the comic
func (c *Comic) URL() string {
	return fmt.Sprintf("https://xkcd.com/%d/", c.Num)
}

// Date returns the comic's publication date
func (c *Comic) Date() (time.Time, error) {
	return time.Parse("2006 1 2", fmt.Sprintf("%s %s %s", c.Year, c.Month, c.Day))
}

// Script parses the transcript and returns all visible text from the image
func (c *Comic) Script() string {
	var w strings.Builder
	scanner := bufio.NewScanner(strings.NewReader(c.Transcript))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "{{") || strings.HasPrefix(line, "[[") {
			continue
		}
		split := strings.SplitN(line, ":", 2)
		if len(split) == 2 {
			line = split[1]
		}
		fmt.Fprintln(&w, strings.TrimSpace(line))
	}
	return w.String()
}
