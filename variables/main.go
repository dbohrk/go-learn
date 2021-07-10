package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = "and press enter when ready."

func main() {

	// Seed the random number
	rand.Seed(time.Now().UnixNano())

	// one way - declare, then assign (two steps)
	var firstNumber int
	firstNumber = rand.Intn(8) + 2 // Returns 0 to 8, so +2 avoids divide by 0

	// another way, dedclare type and name and assign value
	var secondNumber = rand.Intn(8) + 2 // Returns 0 to 8, so +2 avoids divide by 0

	// one step variable: declare name, assign value, and let GO determine the type
	subtraction := rand.Intn(8) + 2 // Returns 0 to 8, so +2 avoids divide by 0

	var answer = firstNumber*secondNumber - subtraction

	// Call game to display answers
	playTheGame(firstNumber, secondNumber, subtraction, answer)
}

func playTheGame(firstNumber int, secondNumber int, subtraction int, answer int) {

	reader := bufio.NewReader(os.Stdin)

	// display a welcom message/instructions
	fmt.Println("Guess the number game")
	fmt.Println("----------------------")
	fmt.Println("")

	fmt.Println("Think of a number between 1 and 10", prompt)
	reader.ReadString('\n')

	// thake them through the games

	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiply the result by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you originally thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now substract", subtraction, prompt)
	reader.ReadString('\n')

	// give them the answer
	fmt.Println("The answwer is", answer)
}
