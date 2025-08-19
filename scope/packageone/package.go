package packageone

import "fmt"

/*
// private variables start with lower case; only available within package
var privateVar = "I am private"

// public variables start with upper case and are exported
var PublicVar = "I am public (exported)"


func notExported() {

}

func Exported() string {
	return "Exported function"
}
*/

//variables
// declare package level variable named myVar for the main package

// declare a block variable names blockVar for the main function

// declare package-level variable named PackageVar in packageone

// create an exported function named PrintMe in packageone

// in main func, print myVar, blockVar, PackageVar on one line

var packageVar = "packageVar"

func PrintMe(var1, var2 string) {
 fmt.Println(var1, var2, packageVar)
}

