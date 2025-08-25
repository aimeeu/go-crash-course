package main

import (
	"scope/packageone"
)

/*
import (
	"fmt"
	"scope/packageone"
)

var one = "One"

func main() {
	myFunc()
	newString := packageone.PublicVar
	fmt.Println("From packageone: ", newString)
	fmt.Println(packageone.Exported())
}

func myFunc() {
	fmt.Println(one)
}
*/

//variables
// declare package level variable named myVar for the main package

// declare a block variable names blockVar for the main function

// declare package-level variable named PackageVar in packageone

// create an exported function named PrintMe in packageone

// in main func, print myVar, blockVar, PackageVar on one line. Use PrintMe to print

	var myVar = "myVar"

func main() {
  var blockVar = "blockVar"

	packageone.PrintMe(myVar, blockVar)
}
