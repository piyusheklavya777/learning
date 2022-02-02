package main

import "fmt"

func main() {

	play()

	return

	// var userAction string

	// fmt.Println("Please enter the user action intended: [newbooking, cancelBooking, updateBooking]")
	// fmt.Scan(&userAction)

	// for ticketBooking.TicketsRemaining > 0 {
	// 	switch userAction {
	// 	case "newvooking" :
	// 		ticketBooking.Main_booking()
	// 		break
	// 	case "cancelbooking" :
	// 		// cancel booking
	// 		break
	// 	case "updatebooking" :
	// 		// update booking
	// 		break
	// 	case "listbooking" :
	// 		// list bookings
	// 		break
	// 	default :
	// 		fmt.Println("Please choose a valid option")
	// 	}
	// }

}

// func makeFood(order string, c <- chan) {
// 	// randomTime := int(rand.Intn(5))
// 	time.Sleep(2 * time.Second)

// }

// func play() {

// 	bridge := make(chan int)

// 	orders := []string{"noodles", "soup", "sushi"}

// 	for _, val := range orders {
// 		go makeFood(val, bridge)
// 	}

// 	select {

// 	}

// }

type channel chan interface{}

func (c channel) wait(n int) {
	for i := 0; i < n; i++ {
		<-c
	}
}

func play() {

	done := make(channel)
	input := []int{1, 2, 3, 4, 5}
	output := make([]int, len(input))

	for i , _ := range input {
		// fmt.Printf("square %v ", inpi)
		go func() {
			output = append(output, square(in, done))
		} ()
	}

	done.wait(len(input))

	fmt.Println(output)

}

func square(n int, done channel) (squared int) {
	squared = n * n
	done <- 0
	return
}

//
