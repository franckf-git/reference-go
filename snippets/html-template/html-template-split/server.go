tmpl, _ := template.ParseFiles("header.html", "footer.html", "profile.html")
tmpl.Execute(w, User{Name: "philippta"})
