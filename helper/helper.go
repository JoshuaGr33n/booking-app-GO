package helper

import "strings"


func ValidateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {

	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0

	return isValidName, isValidEmail, isValidTicketNumber

	// isValidCity := city == "Singapore" || city == "London"
}
