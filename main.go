package main // package main.

import "fmt" // import fmt package

// entry point of the program
func main() {

	conferenceName := "Go Conference" // syntax for declaring and initializing a variable
	const conferenceTickets uint = 50
	var remainingTickets uint = 50

	fmt.Printf("Type of conferenceName is %T, conferenceTickets is %T \n", conferenceName, conferenceTickets)

	fmt.Printf("Welcome to our %v booking Application!!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var bookings []string

	var userName string
	fmt.Println("Enter your name:")
	fmt.Scan(&userName)

	var userTickets uint
	fmt.Println("Please Enter the number of tickets you want to book:")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets
	fmt.Printf("Thank you %v for booking %v tickets. You will receive a confirmation email shortly.\n", userName, userTickets)
	fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)

	bookings = append(bookings, userName) // adding userName to bookings slice
	fmt.Printf("Type of bookings is %T\n", bookings)
	fmt.Printf("The first booking is %v\n", bookings[0])
	fmt.Printf("The number of bookings is %v\n", len(bookings))
	
	fmt.Printf("The bookings are: %v\n", bookings)

	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)
}
