package main

import (
	"bufio"
	"fmt"
	"github.com/eiannone/keyboard"
	"log"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

type User struct {
	UserName string
	Age int
	FavoriteNumber float64
	OwnsADog	bool
}

func main() {
	reader = bufio.NewReader(os.Stdin)

	var user User

	user.UserName = readString("What is you name?")
	user.Age = readInt("How old are you?")
	user.FavoriteNumber = readFloat("What is your favorite number?")
	user.OwnsADog = readBoolean("Do you own a dog y/n?")

	fmt.Printf("Your name is %s. You are %d years old. Your favorite number is %.2f. And you own a dog is %t\n",
		user.UserName,
		user.Age,
		user.FavoriteNumber,
		user.OwnsADog)
}

func prompt() {
	fmt.Print("-> ")
}

func readString(s string) string {
	for {
		fmt.Println(s)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)
		if userInput == "" {
			fmt.Println("Please enter a name")
		} else {
			return userInput
		}
	}
}
func readInt(s string) int {
	for {
		fmt.Println(s)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)
		num, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("Please enter a whole number")
		} else {
			return num
		}
	}
}

func readFloat( s string) float64 {
	for {
		fmt.Println(s)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)
		num, err := strconv.ParseFloat(userInput, 64)
		if err != nil {
			fmt.Println("Please enter a number")
		} else {
			return num
		}
	}
}

func readBoolean(q string) bool {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	for {
		fmt.Println(q)
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}
		if char == 'n' || char == 'N' {
			return false
		}
		if char == 'y' || char == 'Y' {
			return true
		}
	}
}