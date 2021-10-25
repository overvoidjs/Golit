//Алгоритм поиска в ширину по графам
package main

import (
	"container/list"
	"fmt"
)

func isManager(name string) bool {
	//Тупо для теста, если имя заканчивается на м - то менеджер
	if name[len(name)-1:] == "m" {
		return true
	} else {
		return false
	}
}

func findManagerInGraph(graph map[string][]string) string {

	// инициализируем очередь
	queue := list.New()

	//Добавляем в очередь первичный графф, связи 1 уровня
	for _, value := range graph["you"] {
		queue.PushBack(value)
	}

	//Работаем с очередью
	for queue.Len() > 0 {
		e := queue.Front() // Первый элемент
		name := e.Value.(string)

		if isManager(name) {
			return name
			break
		} else {
			if len(graph[name]) > 0 {
				for _, value := range graph[name] {
					queue.PushBack(value)
				}
			}
		}
		queue.Remove(e)
	}

	return "Нету менеджеров"

}

func main() {

	graph := map[string][]string{}
	graph["you"] = []string{"alice", "bob", "claire"}
	graph["bob"] = []string{"anuj", "peggy"}
	graph["alice"] = []string{"peggy"}
	graph["claire"] = []string{"thom", "jonny"}
	graph["anuj"] = []string{}
	graph["peggy"] = []string{}
	graph["thom"] = []string{}
	graph["jonny"] = []string{}

	fmt.Println(findManagerInGraph(graph))

}
