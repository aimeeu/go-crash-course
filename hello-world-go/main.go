package main

import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
	"strings"
)

func main() {
	/*
		var what string = "Bonjour!"
		blob := "Guten Tag"
		sayHelloWorld(what)
		sayHelloWorld(blob)
	*/
	reader := bufio.NewReader(os.Stdin)

	whatToSay := doctor.Intro()
	fmt.Println(whatToSay)

	for {
		fmt.Println("-> ")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1) //windows
		userInput = strings.Replace(userInput, "\n", "", -1)//mac, linux
		if (userInput == "quit"){
			break
		}else {
		  fmt.Println(doctor.Response(userInput))
		}		
	}

}

/*
func sayHelloWorld(what string) {
	fmt.Println(what)
}
*/
