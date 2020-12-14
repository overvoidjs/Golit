package modules

import ("fmt"; "net/http")

func routingMaster() {
  http.HandleFunc("/", home_page)
  http.HandleFunc("/about/", about_page)

  // Запускаем локальный сервер
  fmt.Println("Сервер запушен на http://localhost:7829")
  http.ListenAndServe(":7829", nil)
}
