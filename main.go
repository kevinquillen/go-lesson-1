package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	const distanceToMars = 62100000
	const basePrice = 1000

	// This affects random generation in the rest of the program
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Spaceline    Days    Trip-Type    Price (in millions)")
	fmt.Print("=====================================================\n\n")

	for total := 0; total < 10; total++ {
		ticket := RandomTicket()
		var ticketPrice = basePrice

		// logic here to determine price based on distance, trip type, carrier speed
		// price should likely be part of the ticket struct itself based on other properties
		if ticket.tripType == "Round Trip" {
			ticketPrice = (ticketPrice * 5)
		}

		fmt.Printf("%v    %v    %v    $%v\n", ticket.carrier, ticket.tripLength, ticket.tripType, ticketPrice)
	}

	fmt.Print("\n=====================================================")
}

// RandomTicket - returns a randomized ticket.
func RandomTicket() *Ticket {
	const (
		departureDate = "2020-10-21"
		layoutISO     = "2006-01-02"
		layoutUS      = "January 2, 2006"
	)

	// is there a better / more normalized way to do this?
	carriers := [3]string{"Space Adventures", "Virgin Galactic", "Space X"}
	tripTypes := [2]string{"Round Trip", "One Way"}
	departs, _ := time.Parse(layoutISO, departureDate)

	return &Ticket{
		carrier:       carriers[rand.Intn(len(carriers))],
		tripType:      tripTypes[rand.Intn(len(tripTypes))],
		tripLength:    random(20, 50),
		departureDate: departs.Format(layoutUS),
	}
}

// This does not exist in Go?
func random(min, max int) int {
	return rand.Intn(max-min) + min
}

// Ticket holds information about a ticket.
type Ticket struct {
	carrier       string
	tripType      string
	tripLength    int
	departureDate string
}
