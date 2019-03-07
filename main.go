package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// This affects random generation in the rest of the program
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Carrier     Speed         Days     Type     Price")
	fmt.Print("=====================================================\n\n")

	for total := 0; total < 10; total++ {
		ticket := RandomTicket()
		fmt.Printf("%v     %v km/h     %v     %v     $%vM\n", ticket.ship.carrier, ticket.ship.speed, ticket.tripLength, ticket.tripType, ticket.tripPrice)
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
	tripTypes := [2]string{"Round Trip", "One Way"}
	departs, _ := time.Parse(layoutISO, departureDate)
	tripLength, err := random(20, 50)
	price := basePrice

	if err != nil {
		panic(err)
	}

	// selectedCarrier := carriers[rand.Intn(len(carriers))]
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
		tripType:      selectedType,
		tripLength:    tripLength,
		departureDate: departs.Format(layoutUS),
		tripPrice:     price,
		ship:          GenerateShip(),
	}
}

// GenerateShip generate a random ship.
func GenerateShip() *Ship {
	carriers := [3]string{"Space Adventures", "Virgin Galactic", "Space X"}
	randomCarrier := carriers[rand.Intn(len(carriers))]
	var speed = 0

	switch randomCarrier {
	case "Space X":
		speed = 32
		break
	default:
		speed = 16
		break
	}

	return &Ship{
		carrier: randomCarrier,
		speed:   speed,
	}
}

// This function does not exist in Go?
func random(min int, max int) (int, error) {
	if max <= min {
		return -1, errors.New("random func: maximum value passed should be greater than the minimum")
	}

	return (rand.Intn(max-min) + min), nil
}

// Ticket holds information about a ticket.
type Ticket struct {
	tripType      string
	tripLength    int
	departureDate string
	tripPrice     int
	ship          *Ship
}

// Ship holds information about a ship.
type Ship struct {
	carrier string
	speed   int
}
