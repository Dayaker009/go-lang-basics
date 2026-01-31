package helper

import "fmt"

// function should start with a capital letter to be exported
func ValidateUserInput(firstName string, lastName string, userTickets uint) (bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidTicketNumber := userTickets > 0
	return isValidName, isValidTicketNumber
}

func GetUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address:")
	fmt.Scan(&email)
	fmt.Println("Please Enter the number of tickets you want to book:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}