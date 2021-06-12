import "html/template"

tmpl, _ := template.New("").Parse("<h1>{{.}}</h1>") // from string
tmpl, _ := template.ParseFiles("demo.html")         // from file

tmpl.Execute(w, "Hello world")
// Output: <h1>Hello world</h1>
