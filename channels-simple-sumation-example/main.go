package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)	// The length (6) divided by 2 (3) from the end forward
	go sum(s[len(s)/2:], c)	// The length (6) divided by 2 (3) from the beginning
	x, y := <-c, <-c// receive from each concurrent instance of sum() (channel 'c')

	fmt.Println(x, y, x+y)	// Print each return from channel 'c' and sum to two values
}
