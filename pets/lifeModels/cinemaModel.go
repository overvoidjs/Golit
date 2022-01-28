package main

import(
    "fmt"
    "strconv"
    "math/rand"
    "time"
)

const maxSeats = 100

func maker(makerType string) []string {
    makerMap := []string{}
    for index := 0; index < maxSeats; index++ {
        makerElement := makerType + strconv.Itoa(index+1)
        makerMap = append(makerMap, makerElement)
    }
    return makerMap
}

func makePeoples() []string {
    return maker("U")
}

func makeScene() []string {
    return maker("S")
}

func takeRandomSeat() string {
    rand.Seed(time.Now().UnixNano())
    return "S" + strconv.Itoa(rand.Intn(maxSeats-1))
}

func takeOut(storage []string, target string) []string {
    for n, value := range storage {
        if value == target {
            storage = append(storage[:n], storage[n+1:]...)
            break
        }
    }
    return storage
}

func isPeopleSeatFree(seats []string, ticket string) bool {
    for _, seat := range seats {
        if seat == ticket {
            return true
        }
    }
    return false
}

func getTicket(people string) string {
    return "S" + people[1:]
}

func liveModel(peoples []string, seats []string) bool {
    randomSeat := takeRandomSeat()
    seats = takeOut(seats, randomSeat)

    lastResult := false

    for _, people := range peoples {

        ticket := getTicket(people)

        if isPeopleSeatFree(seats, ticket) {
            seats = takeOut(seats, ticket)
            lastResult = true
        } else {
            if len(seats) == 1 {
                lastResult = false
                break
            }
            found := false
            for sndx := 0; sndx < maxSeats; sndx++ {
                randomSeat := takeRandomSeat()

                if isPeopleSeatFree(seats, randomSeat) {
                    seats = takeOut(seats, randomSeat)
                    lastResult = false
                    found = true
                    break
                }
            }
            if found == false {
                lastResult = false
                break
            }
        }
    }
    return lastResult
}

func main() {
    hit := 0
    itterations := 1000

    for fdex := 0; fdex < itterations; fdex++ {
        peoples := makePeoples()
        seats := makeScene()
        if liveModel(peoples, seats) {
            hit++
        }
    }

    ver := (float64(hit) / float64(itterations)) * 100

    fmt.Println( ver )
}
