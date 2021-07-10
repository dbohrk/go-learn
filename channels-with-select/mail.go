package main
import (
	"fmt"
	"time"
)

var chan1 = make(chan string)	// Create two channels: chan1 and chan2
var chan2 = make(chan string)

func task1() {
	time.Sleep(1 * time.Second)
	chan1 <- "one"				// Send message to channel 1
}

func task2() {
	time.Sleep(2 * time.Second)
	chan2 <- "two"				// Send message to channel 2

}

func main() {
	go task1()					// Start task1 in the background
	go task2()					// Start task2 in the background

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <- chan1:	// Receive message from channel 1 (task 1)
			fmt.Println("received", msg1)
		case msg2 := <- chan2:
			fmt.Println("received", msg2)
		}
	}
}

