package main // package main.

import (
	"fmt"                   // import fmt package
	"go-lang-basics/helper" // import helper package
	"strconv"			   // import strconv package
)

const conferenceName string = "Go Conference" // syntax for declaring and initializing a constant
var remainingTickets uint = 50
const conferenceTickets uint = 50
var bookings = make([]map[string]string, 0) // empty slice of boookings

// entry point of the program
func main() {

	// calling greetUsers function
	greetUsers()

	for {						
		// get user input using function
		firstName, lastName, email, userTickets := helper.GetUserInput()

		if !handleValidation(firstName, lastName, userTickets) {
			continue
		}

		remainingTickets = remainingTickets - userTickets

		bookTicket(userTickets, firstName, lastName, email)

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
		firstNames = append(firstNames, booking["firstName"]) // appending first name to firstNames slice
	}
	return firstNames
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {

	// make is a built-in function that allocates and initializes an object of type slice, map, or chan
	var userData = make(map[string]string) // create a map to store user data

	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["userTickets"] = strconv.FormatUint(uint64(userTickets), 10) // converting uint to string

	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, userData) // append userData to bookings map list
	fmt.Printf("Thank you %v for booking %v tickets. You will receive a confirmation email shortly.\n", firstName+" "+lastName, userTickets)
	fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)
}

func handleValidation(firstName string, lastName string, userTickets uint) bool {
    isValidName, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, userTickets)

    if !isValidName {
        fmt.Println("First name or last name you entered is too short. Please try again.")
        return false
    }

    if !isValidTicketNumber {
        fmt.Println("The number of tickets you entered is invalid. Please try again.")
        return false
    }

    if userTickets > remainingTickets {
        fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
        return false
    }

    return true
}