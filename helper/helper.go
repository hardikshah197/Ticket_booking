package helper

import (
	"strings"
)

func ValidInputFields(name string, email string, ticket uint, remainingTickets uint) (bool, bool, bool) {

	isValidName := len(name) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTickets := ticket > 0 && ticket <= remainingTickets

	return isValidName, isValidEmail, isValidTickets
}
