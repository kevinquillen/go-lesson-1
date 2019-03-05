package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// This affects random generation in the rest of the program
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Spaceline    Days    Trip-Type    Price")
	fmt.Print("=====================================================\n\n")

	for total := 0; total < 10; total++ {
		ticket := RandomTicket()
		fmt.Printf("%v    %v    %v    $%vM\n", ticket.carrier, ticket.tripLength, ticket.tripType, ticket.tripPrice)
	}

	fmt.Print("\n=====================================================")
}

// RandomTicket - returns a randomized ticket.
func RandomTicket() *Ticket {
	const (
		departureDate  = "2020-10-21"
		layoutISO      = "2006-01-02"
		layoutUS       = "January 2, 2006"
		basePrice      = 10
		distanceToMars = 62100000
	)

	// is there a better / more normalized way to do this?
	carriers := [3]string{"Space Adventures", "Virgin Galactic", "Space X"}
	tripTypes := [2]string{"Round Trip", "One Way"}
	departs, _ := time.Parse(layoutISO, departureDate)
	tripLength := random(20, 50)
	price := basePrice

	selectedCarrier := carriers[rand.Intn(len(carriers))]
	selectedType := tripTypes[rand.Intn(len(tripTypes))]

	// Modify price based on trip length
	if tripLength < 25 {
		// Very speedy! Premium cost.
		price = (price * 4)
	}

	if tripLength >= 25 && tripLength <= 35 {
		// Average speed.
		price = (price * 2)
	}

	if tripLength > 35 {
		// Woof. Very slow!
	}

	// Round trip tickets are double.
	if selectedType == "Round Trip" {
		price = (price * 2)
	}

	return &Ticket{
		carrier:       selectedCarrier,
		tripType:      selectedType,
		tripLength:    tripLength,
		departureDate: departs.Format(layoutUS),
		tripPrice:     price,
	}
}

// This function does not exist in Go?
func random(min int, max int) int {
	// Should be a check here to throw an error if max is less than the minimum passed?
	return rand.Intn(max-min) + min
}

// Ticket holds information about a ticket.
type Ticket struct {
	carrier       string
	tripType      string
	tripLength    int
	departureDate string
	tripPrice     int
}
