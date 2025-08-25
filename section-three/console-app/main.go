package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eiannone/keyboard"
)

func main() {
	/*
		reader := bufio.NewReader(os.Stdin)

		for {
			fmt.Println("-> ")
			// ReadString returns 2 values, so the _ means to disgard the second return value
			userInput, _ := reader.ReadString('\n')

			userInput = strings.Replace(userInput, "\n", "", -1)
			if userInput == "quit" {
				break
			} else {
				fmt.Println(userInput)
			}
		}
	*/
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	// defer means do not execute immediately
	defer func() {
		_ = keyboard.Close()
	}()

	coffees := make(map[int]string)
	coffees[1] = "Cappucino"
	coffees[2] = "Latte"
	coffees[3] = "Americano"
	coffees[4] = "Mocha"
	coffees[5] = "Macchiato"
	coffees[6] = "Espresso"
	
	fmt.Println("MENU")
	fmt.Println("----")
	for i := 1; i < len(coffees) + 1; i++ {
		fmt.Println(i, coffees[i])
	}
  fmt.Println("Q - Quit")

	for {
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		i, _ := strconv.Atoi(string(char))
		//fmt.Println(fmt.Sprintf("You chose %s", coffees[i]))
		fmt.Printf("You chose %s", coffees[i])
		fmt.Println("\n")

		if char == 'q' || char == 'Q'  {
			break
		}
	}

	fmt.Println("Program exiting...")
}
