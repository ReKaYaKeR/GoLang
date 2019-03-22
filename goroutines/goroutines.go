package main

import (
	"fmt"
	"time"
)

// func goroutine(val string) {
// 	fmt.Println(val)
// }

// func manyRoutines() {
// 	for i := 0; i < 10; i++ {
// 		go goroutine(strconv.Itoa(i))
// 	}
// }

func main() {
	// go goroutine("world") //will run at same time as remaining code
	// fmt.Println("hello")  //wont wait for goroutine to finish

	// //delay and wait for input before terminating
	// var input string
	// fmt.Print("Print your name: ")
	// fmt.Scanln(&input)

	// //make many routines each will finish in its own time
	// manyRoutines()
	// fmt.Println("delay... Hit enter to end.")
	// fmt.Scanln(&input)

	//---------------Intro to Channels---------------
	//Prints ping every 1 second
	var c = make(chan string) //initialize channel

	go pinger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input) //hit enter to end program

	//More info later on Select control, More Channels and Buffered Channels
}

//---------------Channels---------------
func pinger(c chan string) { //chan is keyword
	for {
		c <- "ping" //sends ping to c channel
	}
}

func printer(c chan string) {
	for {
		msg := <-c //receives info from channel c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}
