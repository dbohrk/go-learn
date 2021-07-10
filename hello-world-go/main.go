package main

import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	whatToSay := doctor.Intro()

	fmt.Println(whatToSay)

	for {
		fmt.Print("-> ")
		userInput, _ := reader.ReadString('\n')

		userInput = strings.Replace(userInput, "\r\n", "", -1) // For Windows systems
		userInput = strings.Replace(userInput, "\n", "", -1)   // For non-Windows systems

		if userInput == "quit" {
			break
		} else {
			fmt.Println(doctor.Response(userInput))
		}

	}
}
