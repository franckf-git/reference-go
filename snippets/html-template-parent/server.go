tmpl, _ := template.ParseFiles("layout.html", "profile.html")
tmpl.Execute(w, User{Name: "philippta"})
