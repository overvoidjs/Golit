package main

import(
  "fmt"
)

type Renderer interface {
  Render()
}

type Button struct {}

func (b Button) Render() {
  fmt.Println("Button")
}

type Window struct {}

func (b Window) Render() {
  fmt.Println("Window")
}

func render(r Renderer){
  r.Render()
}

func main(){
  render(Button{})
  render(Window{})
}
