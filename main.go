package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

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

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	fmt.Println("Carrier     Speed         Days     Type     Price")
	fmt.Print("=====================================================\n\n")

	for total := 0; total < 10; total++ {
		ticket := RandomTicket()
		fmt.Printf("%v     %v km/s     %v     %v     $%vM\n", ticket.ship.carrier, ticket.ship.speed, ticket.tripLength, ticket.tripType, ticket.tripPrice)
	}

	fmt.Print("\n=====================================================")
}

// RandomTicket - returns a randomized ticket.
func RandomTicket() *Ticket {
	const (
		departureDate  = "2020-10-13"
		layoutISO      = "2006-01-02"
		layoutUS       = "January 2, 2006"
		distanceToMars = 62100000
	)

	var price int
	var err error

	// is there a better / more normalized way to do this (pick a random value)?
	tripTypes := [2]string{"Round Trip", "One Way"}
	departs, _ := time.Parse(layoutISO, departureDate)
	ship := GenerateShip()

	// Slower ships are cheaper, but this is kludgey trying to learn how to handle an error.
	if ship.speed == 16 {
		price, err = randomRange(15, 36)
	} else {
		price, err = randomRange(36, 50)
	}

	// Exception needed? I don't see how we can continue otherwise, if the programmer passed unexpected argument values.
	if err != nil {
		panic(err)
	}

	// Now that we have Ship with designated speed, we should be able to calculate
	// a price based on speed, and distanceToMars
	tripLength := distanceToMars / ((ship.speed * 3600) * 24)

	selectedType := tripTypes[rand.Intn(2)]

	// Round trip tickets are double price.
	if selectedType == "Round Trip" {
		price = (price * 2)
	}

	return &Ticket{
		tripType:      tripTypes[rand.Intn(2)],
		tripLength:    tripLength,
		departureDate: departs.Format(layoutUS),
		tripPrice:     price,
		ship:          ship,
	}
}

// GenerateShip generate a random ship.
func GenerateShip() *Ship {
	carriers := [3]string{"Space Adventures", "Virgin Galactic", "Space X"}
	speeds := [2]int{16, 30}
	randomCarrier := carriers[rand.Intn(3)]
	randomSpeed := speeds[rand.Intn(2)]

	return &Ship{
		carrier: randomCarrier,
		speed:   randomSpeed,
	}
}

// This function does not exist in Go?
func randomRange(min int, max int) (int, error) {
	if max <= min {
		return -1, errors.New("random func: maximum value passed should be greater than the minimum")
	}

	return (rand.Intn(max-min) + min), nil
}
