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

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTicket := getUserInput()
		isValidName, isValidemail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTicket, remainingTickets)

		if isValidName && isValidemail && isValidTicketNumber {

			bookingTicket(userTicket, firstName, lastName, email)

			//Call function print first Name
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

func greetUsers() {
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still left\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames
}

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

func bookingTicket(userTicket uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTicket

	// create a map for a user
	userData := make(map[string]string)

	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTicket"] = strconv.FormatUint(uint64(userTicket), 10)

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v \n", bookings)

	fmt.Printf("Thank you %v %v for booking %v ticket(s). You will receive a confirmation email at %v \n", firstName, lastName, userTicket, email)
	fmt.Printf("%v ticket are remaining for conference %v \n", remainingTickets, conferenceName)

}
