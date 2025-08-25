package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = "and press ENTER when ready."

func main() {
	// seed random number generator
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var firstNumber = r.Intn(8) + 2
	var secondNumber = r.Intn(8) + 2
	var subtraction = r.Intn(8) + 2

	playTheGame(firstNumber, secondNumber, subtraction)

}

func playTheGame(first int, second int, subtraction int) {
	var answer = first*second - subtraction
	reader := bufio.NewReader(os.Stdin)

	// display welcome and instructions
	fmt.Println("Guess the Number Game")
	fmt.Println("-------------------")
	fmt.Println("")
	fmt.Println("Think of a number between 1 and 10", prompt)
	reader.ReadString('\n')

	// take user through game
	fmt.Println("Multiply your number by", first, prompt)
	reader.ReadString('\n')
	fmt.Println("Now multiply the result by", second, prompt)
	fmt.Println("Divide the result by the number you originally thought of", prompt)
	reader.ReadString('\n')
	fmt.Println("Now subtract", subtraction, prompt)
	reader.ReadString('\n')

	// give answer
	fmt.Println("The answer is", answer)
}
