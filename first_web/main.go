package main

import ("fmt"; "net/http"; "html/template"; "./modules")


func home_page(w http.ResponseWriter, r *http.Request) {
   bob := modules.User{"Boby", 25, -50, 4.2, 0.8, []string{"Football", "Skate"}}
  // fmt.Fprintf(w, bob.getAllInfo())

  // fmt.Fprintf(w, "<b>Main Text</b>")

  templ, _ := template.ParseFiles("templates/home_page.html")
  templ.Execute(w, bob)
}

func about_page(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "It's all about someone. <br> <a href='/'>go to home page</a>")
}

func routingMaster() {
  http.HandleFunc("/", home_page)
  http.HandleFunc("/about/", about_page)

  // Запускаем локальный сервер
  fmt.Println("Сервер запушен на http://localhost:7829")
  http.ListenAndServe(":7829", nil)
}

func main() {
  // var bob User = ....
  // bob := User{name: "Bob", age: 25, money: -50, avg_grades: 4.2, happiness: 0.8}
  // bob := User{"Bob", 25, -50, 4.2, 0.8}

  routingMaster()
}
