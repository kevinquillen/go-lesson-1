package main

import (
	"fmt"
	"time"
)

func main() {
	const distanceToMars = 62100000

	fmt.Println("Spaceline    Days    Trip-Type    Departs    Price (in millions)")
	fmt.Println("=====================================================")

	for total := 0; total < 10; total++ {
		ticket := RandomTicket()
		fmt.Printf("%v    %v    %v    %v    100\n", ticket.carrier, ticket.tripLength, ticket.tripType, ticket.departureDate)
	}
}

// RandomTicket - factory method to return a new ticket.
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
		carrier:       carriers[0],
		tripType:      tripTypes[0],
		tripLength:    30,
		departureDate: departs.Format(layoutUS),
	}
}

// Ticket holds information about a ticket.
type Ticket struct {
	carrier       string
	tripType      string
	tripLength    int
	departureDate string
}
