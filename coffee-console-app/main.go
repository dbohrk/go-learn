package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"strconv"
)

func main() {
	err := keyboard.Open
	if err != nil {
		// log.Fatal(err)
	}

	// Close keyboard on error
	defer func() {
		_ = keyboard.Close()
	}()

	coffees := make(map[int]string)
	coffees[1] = "Cappucino"
	coffees[2] = "Latté"
	coffees[3] = "Americano"
	coffees[4] = "Mocha"
	coffees[5] = "Macchiato"
	coffees[6] = "Espresso"

	fmt.Println("MENU")
	fmt.Println("====")
	fmt.Println("1 - Cappuchino")
	fmt.Println("2 - Latté")
	fmt.Println("3 - Arericano")
	fmt.Println("4 - Mocha")
	fmt.Println("5 - Macchiato")
	fmt.Println("6 - Espresso")
	fmt.Println("Q - Quit the program.")

	char := ' '

//	for char != 'q' && char != 'Q' {		// Clearer and easier to read then the "do while" below
//		char, _, _ = keyboard.GetSingleKey()
//		if err != nil {
//			log.Fatal(err)
//		}
//
//	if char == 'q' || char == 'Q' {
//		fmt.Println("Thank you using Dave's Coffee ")
//		break
//		}
//
//		i, _ := strconv.Atoi(string(char))
//
//		if _, ok := coffees[i]; ok {
//			fmt.Println(fmt.Sprintf("You chose %s", coffees[i]))
//		}
//	}

	for ok := true; ok; ok = char != 'q' && char != 'Q' {	// do while loop
		char, _, _ = keyboard.GetSingleKey()

		i, _ := strconv.Atoi(string(char))

		if _, ok := coffees[i]; ok {
			fmt.Println(fmt.Sprintf("You chose %s", coffees[i]))
		}
	}
	fmt.Println("Thank you using Dave's Coffee ")
}
