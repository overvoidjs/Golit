package main

import(
  "fmt"
  "time"
)

type Printer interface {
  Print()
}

type User struct {
  name string
  age int
  lastName string
}

type Document struct {
  name string
  documentType string
  date time.Time
}

func (d Document) Print() {
  fmt.Printf("Document name: %s, type: %s, date:%s \n", d.name, d.documentType, d.date)
}

func (u User) Print() {
  fmt.Printf("Hi I am %s %s and I am %d years old \n", u.name, u.lastName, u.age)
}

func Process(obj Printer) {
  obj.Print()
}

func main(){

  u := User{name: "Vladimir", age: 27, lastName:"Makarov"}
  doc := Document{name: "123.doc", documentType:"doc", date: time.Now()}

  Process(u)
  Process(doc)

}
