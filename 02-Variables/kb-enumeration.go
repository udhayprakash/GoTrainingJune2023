package main

import "fmt"

const (
	_  = iota             // 0
	KB = 1 << (iota * 10) // 1 << (1 * 10)
	MB = 1 << (iota * 10) // 1 << (2 * 10)
	GB = 1 << (iota * 10) // 1 << (3 * 10)
	TB = 1 << (iota * 10) // 1 << (4 * 10)
)

func main() {
	fmt.Println("binary\t\tdecimal")
	fmt.Printf("%40b\t", KB)
	fmt.Printf("%40d\n", KB)
	fmt.Printf("%40b\t", MB)
	fmt.Printf("%40d\n", MB)
	fmt.Printf("%40b\t", GB)
	fmt.Printf("%40d\n", GB)
	fmt.Printf("%40b\t", TB)
	fmt.Printf("%40d\n", TB)
}