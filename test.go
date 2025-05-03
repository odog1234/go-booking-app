// ============================================
// Go Booking Application
// Author: Owen Yesuf
// Description:
//
//	A terminal-based ticket booking system for the Go Conference.
//	It allows users to input their information and reserve tickets.
//	The app performs basic validation, tracks remaining tickets,
//	and stores each booking using maps and slices.
//
//	Features:
//	- Input validation (names, email, ticket count)
//	- Booking confirmation with user details
//	- Remaining tickets tracking
//	- First names display from all bookings
//
// ============================================
package main

import (
	"awesomeProject/helper"
	"fmt"
	"strconv"
)

var conferenceName = "Go Conference"

const conferenceTickets = 50

var remainingTickets uint = 50
var bookings = make([]map[string]string, 0)

// main runs the main booking loop, handles user input and validation,
// and manages ticket booking and display logic.
func main() {

	greetUsers()

	for {

		// Get user input
		firstName, lastName, email, userTicket := getUserInput()

		// Validate the input
		isValidName, isValidemail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTicket, remainingTickets)

		if isValidName && isValidemail && isValidTicketNumber {

			// Book the ticket and store user data
			bookingTicket(userTicket, firstName, lastName, email)

			// Display the list of first names of all bookings
			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Printf("Our conference is booked out. Come back next year")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			} else if !isValidemail {
				fmt.Println("email address you entered doesn't contain @ sign")
			} else if !isValidTicketNumber {
				fmt.Println("number of ticket you entered is invalid")
			}

		}

	}
}

// greetUsers prints the initial welcome message and ticket availability.
func greetUsers() {
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still left\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets to attend")
}

// getFirstNames returns a slice of first names from all current bookings.
func getFirstNames() []string {
	firstNames := []string{}
	// Iterate over each booking and extract the first name
	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames
}

// getUserInput prompts the user for their first name, last name, email,
// and number of tickets, then returns those values.
func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTicket uint
	//Ask user for name
	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets you want")
	fmt.Scan(&userTicket)

	return firstName, lastName, email, userTicket
}

// bookingTicket deducts the booked tickets from remainingTickets,
// creates a map of user data, appends it to the bookings slice,
// and prints a confirmation message.
func bookingTicket(userTicket uint, firstName string, lastName string, email string) {
	// Reduce the number of remaining tickets
	remainingTickets = remainingTickets - userTicket

	// create a map for a user
	userData := make(map[string]string)

	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTicket"] = strconv.FormatUint(uint64(userTicket), 10)

	// Append user data to bookings slice
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v \n", bookings)

	// Print confirmation message
	fmt.Printf("Thank you %v %v for booking %v ticket(s). You will receive a confirmation email at %v \n", firstName, lastName, userTicket, email)
	fmt.Printf("%v ticket are remaining for conference %v \n", remainingTickets, conferenceName)

}
