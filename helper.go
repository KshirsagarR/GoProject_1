package main

import "strings"

func validation(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool, bool) {

	isValidFirstName := len(firstName) >= 2
	isValidLastName := len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidUserTickets := userTickets > 0 && userTickets <= remainingTickets

	return isValidFirstName, isValidLastName, isValidEmail, isValidUserTickets
}
