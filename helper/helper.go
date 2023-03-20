package helper

import "strings"

func ValidateUserInput(fName string, lName string, email string, uTickets uint, rTickets uint) (bool, bool, bool) {
	isValidName := len(fName) >= 2 && len(lName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := uTickets > 0 && uTickets <= rTickets

	return isValidName, isValidEmail, isValidTicketNumber
}
