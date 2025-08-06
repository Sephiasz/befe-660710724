package main

import (
	"fmt"
)

// var gmail String = "Suebyart_p@su.ac.th"
func main() {
	// var name String = "Pongrapee"
	var age int = 20

	gmail := "Suebyart_p@su.ac.th"
	gpa := 3.17

	firstname, lastname := "Pongrapee", "Suebyart"

	fmt.Printf("Name %s %s, age %d, gmail %s, gpa %.2f\n", firstname, lastname, age, gmail, gpa)
}
