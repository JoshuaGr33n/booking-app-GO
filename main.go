package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
	// "strconv"
	// "strings"
)

const conferenceTickets = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50

// var bookings = []string{}   slice
// var bookings = make([]map[string]string, 0) map
var bookings = make([]UserData, 0) // struct

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	fmt.Printf("ConferenceName DataType is %T. ConferenceTickets DataType is %T. RemainingTickets DataType is %T\n", conferenceName, conferenceTickets, remainingTickets)

	// for {
	firstName, lastName, email, userTickets := getUserInput()

	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		if userTickets > uint(remainingTickets) {
			fmt.Printf("We only have %v tickets remaining. So you cant book %v tickets\n", remainingTickets, userTickets)
			//continue
		}

		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Printf("No tickets remaining\n")
			//break
		}
	} else {
		if !isValidName {
			fmt.Println("Invalid first name or last name")
		}
		if !isValidEmail {
			fmt.Println("Invalid email")
		}
		if !isValidTicketNumber {
			fmt.Println("Invalid ticket number")
		}
	}
	// }
	wg.Wait()
}

func greetUsers() {

	fmt.Println("Welcome to our", conferenceName, "booking application")
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		// var names = strings.Fields(booking)
		// firstNames = append(firstNames, names[0])
		// firstNames = append(firstNames, booking["firstName"])
		firstNames = append(firstNames, booking.firstName) // struct
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// ask user for their name

	// firstName = "Tom"
	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - uint(userTickets)

	//create a map for a user
	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	// bookings[0] = firstName + " " + lastName
	// bookings = append(bookings, firstName+" "+lastName)
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is: %v \n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. Check your email %v for confirmation\n", firstName, lastName, userTickets, email)
	fmt.Printf("Remaining Tickets is now %v for %v \n", remainingTickets, conferenceName)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("################")
	fmt.Printf("Sending ticket:\n %v to email address %v\n", ticket, email)
	fmt.Println("################")

	wg.Done()

}
