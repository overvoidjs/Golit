// http://golang-book.ru/chapter-09-structs-and-interfaces.html
package main

import(
  "fmt"
  "math"
)

func distance(x1, y1, x2, y2 float64) float64 {
    a := x2 - x1
    b := y2 - y1
    return math.Sqrt(a*a + b*b)
}

type Rectangle struct {
  x1, y1, x2, y2 float64
}

type Circle struct {
    x, y, r float64
}

func (c *Circle) perimeter() float64 {
  return 2.0 * math.Pi * c.r
}

func (c *Circle) area() float64 {
  return math.Pi * c.r * c.r
}

func (r *Rectangle) area() float64 {
  l := distance(r.x1, r.y1, r.x1, r.y2)
  w := distance(r.x1, r.y1, r.x2, r.y1)

  return l * w
}

func (r *Rectangle) perimeter() float64 {
  l := distance(r.x1, r.y1, r.x1, r.y2)
  w := distance(r.x1, r.y1, r.x2, r.y1)

  return 2.0 * (w + l)
}

type Shape interface {
    area() float64
    perimeter() float64
}

func totalArea(shapes ...Shape) float64 {
    var area float64
    for _, s := range shapes {
        area += s.area()
    }
    return area
}

type MultiShape struct {
    shapes []Shape
}

func (m *MultiShape) area() float64 {
    var area float64
    for _, s := range m.shapes {
        area += s.area()
    }
    return area
}

func main()  {

  r := Rectangle{0, 0, 10, 10}
  fmt.Println(r.area())

  c := Circle{0, 0, 5}
  fmt.Println(c.area())

  fmt.Println(totalArea(&c, &r))



}
