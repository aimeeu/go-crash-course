package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

type User struct {
	UserName       string
	Age            int
	FavoriteNumber float64
	OwnsDog        bool
}

func (u *User) String() string {
	return fmt.Sprintf("Your name is %s. You are %d years old. %g is your favorite number. Owns a dog: %t\n", u.UserName, u.Age, u.FavoriteNumber, u.OwnsDog)
}

func main() {
	reader = bufio.NewReader(os.Stdin)
	var user = &User{}
	user.FavoriteNumber = 123456

	printQuestion("What is your name?")
	user.UserName = readString()

	printQuestion("What is your age?")
	user.Age = readInt()

	printQuestion("What is your favorite number?")
	user.FavoriteNumber = readFloat()

	printQuestion("Do you own a dog? (y/n)")
	user.OwnsDog = readBool()

	fmt.Println(user)
}

func printPrompt() {
	fmt.Println("-> ")
}

func convertToInt(input string) (int, error) {
	age, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}
	return age, err
}

func convertToFloat(input string) (float64, error) {
	favNum, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return favNum, err
}

func printQuestion(s string) {
	fmt.Println(s)
	printPrompt()
}

func readString() string {
	for {
		userInput := getInput()
		if userInput == "" {
			fmt.Println("Enter your name")
		} else {
			return userInput
		}
	}
}

func readInt() int {
	for {
		userInput := getInput()
		if userInput == "" {
			fmt.Println("Enter your age")
		} else {
			age, _ := convertToInt(userInput)
			if age == 0 {
				fmt.Println("Enter a whole number.")
			} else {
				return age
			}
		}
	}
}

func readFloat() float64 {
	for {
		userInput := getInput()
		if userInput == "" {
			fmt.Println("Enter your favorite number")
		} else {
			favNum, _ := convertToFloat(userInput)
			if favNum == 0 {
				fmt.Println("Enter your favorite number as a whole number or decimal number.")
			} else {
				return favNum
			}
		}
	}
}

func readBool() bool {
	for {
		userInput := getInput()
		if userInput == "" {
			fmt.Println("Do you own a dog? (y/n)")
		} else {
			//userInput not nil
			switch userInput {
			case "y", "Y":
				return true
			case "n", "N":
				return false
			default:
				fmt.Println("Do you own a dog? (y/n)")
			}
		}
	}
}

func getInput() string {
	userInput, _ := reader.ReadString('\n')
	// read Windows
	userInput = strings.Replace(userInput, "\r\n", "", -1)
	// read Linux, Mac
	userInput = strings.Replace(userInput, "\n", "", -1)
	return userInput
}
