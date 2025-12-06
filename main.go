package main

import (
	"fmt"
)

const ticket = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName   string
	lastName    string
	email       string
	userTickets uint
}

func main() {

	greet()

	for {
		firstName, lastName, email, userTickets := getUserData()

		isValidFirstName, isValidLastName, isValidEmail, isValidUserTickets :=
			validation(firstName, lastName, email, userTickets)

		if isValidFirstName && isValidLastName && isValidEmail && isValidUserTickets {

			remainingTickets -= userTickets

			var user = UserData{
				firstName:   firstName,
				lastName:    lastName,
				email:       email,
				userTickets: userTickets,
			}

			bookings = append(bookings, user)

			fmt.Println("Booking successful!")
			fmt.Printf("Hello %v %v, you booked %v tickets.\n", firstName, lastName, userTickets)
			fmt.Printf("Remaining tickets: %v\n", remainingTickets)

		} else {
			fmt.Println("Invalid details, please try again.")
		}

		if remainingTickets == 0 {
			fmt.Println("Tickets sold out! Please try next year")
			break
		}
	}
}

func greet() {
	fmt.Println("Welcome! Kindly proceed for booking...!")
}

func getUserData() (string, string, string, uint) {
	var firstName, lastName, email string
	var userTickets uint

	fmt.Print("Enter first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email: ")
	fmt.Scan(&email)

	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}
