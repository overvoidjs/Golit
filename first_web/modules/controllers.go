package modules

import ("fmt"; "net/http"; "html/template")


func home_page(w http.ResponseWriter, r *http.Request) {
   bob := User{"Boby", 25, -50, 4.2, 0.8, []string{"Football", "Skate"}}
  // fmt.Fprintf(w, bob.getAllInfo())

  // fmt.Fprintf(w, "<b>Main Text</b>")

  templ, _ := template.ParseFiles("templates/home_page.html")
  templ.Execute(w, bob)
}

func about_page(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "It's all about someone. <br> <a href='/'>go to home page</a>")
}
