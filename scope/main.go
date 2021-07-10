package main

import "myapp/packageone"

var myVar = "Package Level Variable: myVar"

func main() {
	// variables
	// declare a package level variable for the main package
	// called myVar

	// then declare a block variable for the main function
	// called blockVar
	var blockVar = "Block Level Variable: blockVar"
	// declare a parkage level variable in packageone
	// named PackageVar

	// create an exported function in packageone
	// names PrintMe

	// in the main function, print out the values of
	// myVar, blockVar, and ParkageVar on one line
	// using PrintMe

	packageone.PrintMe(myVar, blockVar, packageone.PackageVar)
}
