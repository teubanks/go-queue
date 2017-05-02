package main

import (
	"log"

	"github.com/teubanks/go-queue"
)

func main() {
	q := queue.NewQueue()
	dat := struct {
		Name    string
		Address string
		Age     int
	}{
		Name:    "Henry Atkins",
		Address: "475 L'Enfant Plaza, Washington DC 20260",
		Age:     46,
	}

	q.Push(dat)

	fetchedDat, valid := q.Pop()
	if !valid {
		log.Printf("Error fetching element from go-queue")
	}

	log.Printf("fetchedDat: %+v\n", fetchedDat)
}
