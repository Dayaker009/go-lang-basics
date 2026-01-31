package main // package main.

import (
	"fmt"                   // import fmt package
	"go-lang-basics/helper" // import helper package
	"sync" 				// import sync package
	"time"
)

const conferenceName string = "Go Conference" // syntax for declaring and initializing a constant
var remainingTickets uint = 50
const conferenceTickets uint = 50
var bookings = make([]UserData, 0) // empty slice of boookings

// wait group to wait for go routines to finish
var waitGroup = sync.WaitGroup{}

// defining a custom data type using struct
type UserData  = struct {
	firstName   string
	lastName    string
	email       string
	userTickets uint
}

type Rectangle struct {
	length float64
	width  float64
}

// entry point of the program
func main() {

	rect := Rectangle{length: 10.0, width: 5.0}
	area := rect.Area()
	println("Area of Rectangle:", area)

	// calling greetUsers function
	greetUsers()
						
	// get user input using function
	firstName, lastName, email, userTickets := helper.GetUserInput()
	handleValidation(firstName, lastName, userTickets)
	bookTicket(userTickets, firstName, lastName, email)

	waitGroup.Add(1) // add a wait group counter
	// go routine to send email, runs concurrently with the main function
	go sendEmail(userTickets, firstName, lastName, email)

	// firstNames := getFirstNames()
	fmt.Printf("The bookings are: %v\n", bookings)
	if remainingTickets == 0 {
		println("Our conference is booked out. Come back next year.")
	}

	waitGroup.Wait() // wait for all go routines to finish
}


func greetUsers() {
	fmt.Printf("Welcome to our %v booking Application!!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}	
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName) // appending first name to firstNames slice
	}
	return firstNames
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {

	// make is a built-in function that allocates and initializes an object of type slice, map, or chan
	var userData = UserData{
		firstName: firstName, 
		lastName: lastName, 
		email: email, 
		userTickets: userTickets,
	} // create a map to store user data

	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, userData) // append userData to bookings map list
	fmt.Printf("Thank you %v for booking %v tickets. You will receive a confirmation email shortly.\n", firstName+" "+lastName, userTickets)
	fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)
}

func sendEmail(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var emailContent = fmt.Sprintf("Sending %v tickets to %v %v at %v\n", userTickets, firstName, lastName, email)
	fmt.Println("####################")
	fmt.Println("Sending email...")
	fmt.Print(emailContent)
	fmt.Println("####################")
	waitGroup.Done() // decrement the wait group counter
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

// method to calculate area of rectangle
func (r Rectangle) Area() float64 {
	return r.length * r.width
}