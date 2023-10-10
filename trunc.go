package main

import "fmt"

func main() {
	var x float32
	fmt.Println("Enter float number: ")
	fmt.Scan(&x)
	fmt.Println("Truncate number is: ", int(x))
}