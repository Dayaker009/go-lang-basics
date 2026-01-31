package main // package main.

import (
	"fmt" // import fmt package
	"strings" // import strings package
	"go-lang-basics/helper" // import helper package
)

const conferenceName string = "Go Conference" // syntax for declaring and initializing a constant
var remainingTickets uint = 50
const conferenceTickets uint = 50
var bookings []string // empty slice of strings

// entry point of the program
func main() {

	// calling greetUsers function
	greetUsers()

	for {						
		// get user input using function
		firstName, lastName, userTickets := helper.GetUserInput()

		// validate user input
		isValidName, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, userTickets)

		if !isValidName {
			fmt.Println("First name or last name you entered is too short. Please try again.")
			continue
		}

		if !isValidTicketNumber {
			fmt.Println("The number of tickets you entered is invalid. Please try again.")
			continue
		}

		if userTickets > remainingTickets {
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
			continue
		}

		remainingTickets = remainingTickets - userTickets
		fmt.Printf("Thank you %v for booking %v tickets. You will receive a confirmation email shortly.\n", firstName+" "+lastName, userTickets)
		fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)

		bookings = append(bookings, firstName+" "+lastName) // adding firstName and lastName to bookings slice
		fmt.Printf("Type of bookings is %T\n", bookings)
		fmt.Printf("The first booking is %v\n", bookings[0])
		fmt.Printf("The number of bookings is %v\n", len(bookings))

		firstNames := getFirstNames()
		fmt.Printf("The bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			println("Our conference is booked out. Come back next year.")
			break
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to our %v booking Application!!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}	
	for _, booking := range bookings {
		var names = strings.Fields(booking) // assigning booking to names
		firstNames = append(firstNames, names[0]) // appending first name to firstNames slice
	}
	return firstNames
}

