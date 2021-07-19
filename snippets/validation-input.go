// <input type="text" name="foo">

// Required inputs
if r.Form.Get("foo") == "" {
    fmt.Println("error: foo is required")
}

// Blank text
import "strings"
···
if strings.TrimSpace(r.Form.Get("foo")) == "" {
    fmt.Println("error: foo must not be blank")
}

// Min and max length (bytes)
l := len(r.Form.Get("foo"))
if l < 5 || l > 10 {
    fmt.Println("error: foo must be between 5 and 10 bytes long")
}

// Min and max length (number of characters)
import "unicode/utf8"
···
l := utf8.RuneCountInString(r.Form.Get("foo"))
if l < 5 || l > 10 {
    fmt.Println("error: foo must be between 5 and 10 characters long")
}

// Starts with, ends with and contains
import "strings"
···
// Check that the field value starts with 'abc'.
if !strings.HasPrefix(r.Form.Get("foo"), "abc") {
    fmt.Println("error: foo does not start with 'abc'")
}

// Check that the field value ends with 'abc'.
if !strings.HasSuffix(r.Form.Get("foo"), "abc") {
    fmt.Println("error: foo does not end with 'abc'")
}

// Check that the field value contains 'abc' anywhere in it.
if !strings.Contains(r.Form.Get("foo"), "abc") {
    fmt.Println("error: foo does not contain 'abc'")
}

// Matches regular expression pattern
import "regexp"
···
// Pre-compiling the regular expression and storing it in a variable is more efficient
// if you're going to use it multiple times. The regexp.MustCompile function will
// panic on failure.
var rxPat = regexp.MustCompile(`^[a-z]{4}.[0-9]{2}$`)

if !rxPat.MatchString(r.Form.Get("foo")) {
    fmt.Println("error: foo does not match the required pattern")
}

// Unicode character range
import "regexp"
···
// Use an interpreted string and the \u escape notation to create a regular
// expression matching the range of characters in the two unicode code blocks.
var rxCyrillic = regexp.MustCompile("^[\u0400-\u04FF\u0500-\u052F]+$")

if !rxCyrillic.MatchString(r.Form.Get("foo")) {
    fmt.Println("error: foo must only contain Cyrillic characters")
}

// <input type="email" name="foo">
// Email validation
import "regexp"
···
var rxEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

e := r.Form.Get("foo")
if len(e) > 254 || !rxEmail.MatchString(e) {
    fmt.Println("error: foo is not a valid email address")
}

// <input type="url" name="foo">
// URL validation
import "net/url"
···
// If there are any major problems with the format of the URL, url.Parse() will
// return an error.
u, err := url.Parse(r.Form.Get("foo"))
if err != nil {
    fmt.Println("error: foo is not a valid URL")
} else if u.Scheme == "" || u.Host == "" {
    fmt.Println("error: foo must be an absolute URL")
} else if u.Scheme != "http" && u.Scheme != "https" {
    fmt.Println("error: foo must begin with http or https")
}

// <input type="number" name="foo" min="0" max="100" step="5">
// Integers
import "strconv"
···
n, err := strconv.Atoi(r.Form.Get("foo"))
if err != nil {
    fmt.Println("error: foo must be an integer")
} else if n < 0 || n > 10  {
    fmt.Printf("error: foo must be between 0 and 100")
} else if  n%5 != 0 {
    fmt.Println("error: foo must be an multiple of 5")
}

// <input type="number" name="foo" min="0" max="1" step="0.01">
// Floats
import "strconv"

n, err := strconv.ParseFloat(r.Form.Get("foo"), 64)
if err != nil {
    fmt.Println("error: foo must be a float")
} else if n < 0 || n > 1  {
    fmt.Printf("error: foo must be between 0 and 1")
}

// <input type="date" name="foo" min="2017-01-01" max="2017-12-31">
// Date
import "time"
···
d, err := time.Parse("2006-01-02", r.Form.Get("foo"))
if err != nil {
    fmt.Printf("error: foo is not a valid date")
} else if d.Year() != 2017 {
    fmt.Printf("error: foo is not between 2017-01-01 and 2017-12-31")
}

// <input type="datetime-local" name="foo" min="2017-01-01" max="2017-12-31">
// Datetime-local
import "time"
···

// Load the users local time zone. This accepts a location name corresponding
// to a file in your IANA Time Zone database.
loc, err := time.LoadLocation("Europe/Vienna")
if err != nil {
    ···
}

d, err := time.ParseInLocation("2006-01-02T15:04:05", r.Form.Get("foo"), loc)
if err != nil {
    fmt.Printf("error: foo is not a valid datetime")
} else if d.Year() != 2017 {
    fmt.Printf("error: foo is not between 2017-01-01 00:00:00 and 2017-12-31 23:59:00")
}

// <input type="radio" name="foo" value="wibble"> Wibble
// <input type="radio" name="foo" value="wobble"> Wobble
// <input type="radio" name="foo" value="wubble"> Wubble
// Radio, Select and Datalist (one-in-a-set) validation
set := map[string]bool{"wibble": true, "wobble": true, "wubble": true}

if !set[r.Form.Get("foo")] {
    fmt.Printf("error: foo not match 'wibble', 'wobble' or 'wubble'")
}

// <input type="checkbox" name="foo" value="wibble"> Wibble
// <input type="checkbox" name="foo" value="wobble"> Wobble
// <input type="checkbox" name="foo" value="wubble"> Wubble
// Checkboxes (many-in-a-set) validation
set := map[string]bool{"wibble": true, "wobble": true, "wubble": true}

for _, f := range r.Form["foo"] {
    if !set[f] {
        fmt.Printf("error: foo does not match 'wibble', 'wobble' or 'wubble'")
        break
    }
}

// <input type="checkbox" name="foo" value="checked"> I accept the terms
// Single checkboxes
if r.Form.Get("foo") != "checked" {
    fmt.Println("foo must be checked")
}

