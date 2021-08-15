package main

import "fmt"

func repeater(s string, n int, c chan string) {
	defer close(c)
	repeat := ""

	for i := 0; i < n; i++ {
		repeat += s
	}

	//Я вообще не понимаю как функцией вернуть в какой то определененый канал если он заранее не известен функции
	c <- repeat
}

func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	go repeater("x", 50, ch1)
	go repeater("y", 25, ch2)

	x := <- ch1
	y := <- ch2
	fmt.Println(x, y)

}
