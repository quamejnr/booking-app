package main 

import (
	"fmt" 
	"strings" 
)

const conferenceTickets = 50
var conferenceName = "Go conference"
var remainingTickets uint = 50 
var bookings = []string{}

func main() {

	greetUser()

	for remainingTickets > 0{
		var userName string
		var userEmail string
		var userTickets uint

		userName, userEmail, userTickets = getUserInfo(userName, userEmail, userTickets)
		if isInvalidEmail(userEmail) {
			fmt.Println("Email is invalid!!!")
			continue
		}
		if isInvalidNumber(userTickets) {
			fmt.Println("Number of tickets is invalid!!!")
			continue
		}
		if remainingTicketsIsNotEnough(userTickets) {
			fmt.Printf("Sorry, we only have %v tickets available.\n", remainingTickets)
			continue
		}
		remainingTickets = buyTickets(userTickets)
		printBookingInfo(userName, userTickets)
	}

	fmt.Println("Sorry we're fully booked, See you next year")

}

func greetUser() {
	fmt.Printf("Welcome to %v\n", conferenceName)
	fmt.Println("We have", conferenceTickets, "tickets")
	fmt.Println("Grab your tickets now!!!")
	fmt.Println("")
}

func getUserInfo(userName string, userEmail string, userTickets uint) (string, string, uint){
	fmt.Println("Kindly provide your name:")
	fmt.Scan(&userName)

	fmt.Println("Kindly provide your email:")
	fmt.Scan(&userEmail)

	fmt.Println("How many tickets are you buying: ")
	fmt.Scan(&userTickets)

	return userName, userEmail, userTickets
}

func isInvalidEmail(userEmail string) bool {
	isValidEmail := strings.Contains(userEmail, "@")
	return !isValidEmail
}

func isInvalidNumber(userTickets uint) bool {
	return userTickets < 1
}

func remainingTicketsIsNotEnough(userTickets uint) bool {
	return remainingTickets < userTickets
}

func buyTickets(userTickets uint) uint {
	remainingTickets -= userTickets
	return remainingTickets
}

func printBookingInfo(userName string, userTickets uint) {
	bookings = append(bookings, userName)
	fmt.Println("")
	fmt.Printf("User %s has booked %v tickets.\n", userName, userTickets)
	fmt.Println("Bookings list:", bookings)
	fmt.Println("Remaining tickets:", remainingTickets)
	fmt.Println("")		
}
