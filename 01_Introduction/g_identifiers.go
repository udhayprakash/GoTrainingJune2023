package main

import "fmt"

func main() {
	emptyString1 := ""
	var emptyString2 string // default initial value for string is ""
	var emptyString3 = ""
	var emptyString4 string = ""

	// To get the address location where an object is stored
	fmt.Println("emptyString1:", emptyString1, "&emptyString1:", &emptyString1)
	fmt.Println("emptyString2:", emptyString2, "&emptyString2:", &emptyString2)
	fmt.Println("emptyString3:", emptyString3, "&emptyString3:", &emptyString3)
	fmt.Println("emptyString4:", emptyString4, "&emptyString4:", &emptyString4)

	fmt.Println()

	fmt.Println("emptyString1 == emptyString2 :", emptyString1 == emptyString2) // true
	fmt.Println("emptyString1 == emptyString3 :", emptyString1 == emptyString3) // true
	fmt.Println("emptyString1 == emptyString4 :", emptyString1 == emptyString4) // true
}
