package main

import ( 
	"fmt"
    "strings"
)

func main ()  {

	var conferneceName = "Go confernece"
	const totalTicket = 50
	var remainingTicket uint = 50
	var booking []string
    var maxTicket uint = 10

/*	fmt.Printf("Welcome to the %v booking application\n", conferneceName)
	fmt.Printf("...And we have total number of booking ticket is %v from which %v are still available\n", totalTicket, remainingTicket)
	fmt.Println("Get you tickets here to attend")
*/
// Calling Greet User function

    greetConference(conferneceName, totalTicket, remainingTicket)
    for{
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("How many tickets you want to book? ")
		fmt.Scan(&userTickets)

		// User input validation

		isValidName := len(firstName) > 2 && len(lastName) > 2
		isValidEmail := strings.Contains(email, "@")
		isValidTickets := userTickets > 0 && userTickets < maxTicket && userTickets < totalTicket

		if isValidName && isValidEmail && isValidTickets {
			if userTickets > maxTicket {
				fmt.Printf("Sorry you can't book more than %v tickets at once\nMaximam of %v tickets can be booked per user\n", userTickets, maxTicket)
				continue
			}
	
			remainingTicket= remainingTicket - userTickets
			booking = append(booking,firstName + " " + lastName)
		
			if remainingTicket ==0 {
				fmt.Printf("All %v tickets are booked. Better Luck Next Time!!!!\n", totalTicket)
				break
			}
			
			fmt.Printf("Thank you %v %v for booking  %v tickets, you will get confirmation mail shortly at your email id %v\n", firstName, lastName, userTickets, email)
			fmt.Printf(" reamining tickets available is %v\n", remainingTicket)
			fmt.Printf("These are all bookings %v\n", booking)

		} else{
			if !isValidName {
				fmt.Println("Your First Name or Last Name is not correct, Please provide vaild First Name!!")
			}
			if !isValidEmail {
				fmt.Println("Your email is not in correct format, Please provide valid email!!")
			}
			if !isValidTickets {
				fmt.Println("Number of tickets you entered is incorrect")
			}
		}
	}
}


func greetConference (conferneceName string, totalTicket int, remainingTicket uint ) {

	fmt.Printf("Welcome to the %v booking application\n", conferneceName)
	fmt.Printf("...And we have total number of booking ticket %v from which %v are still available\n", totalTicket, remainingTicket)
	fmt.Println("Get you tickets here to attend")
}