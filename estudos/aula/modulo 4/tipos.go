package main

import "fmt"

func main() {
	fmt.Println("Type: %T - Value: %v\n", true, true)
	fmt.Println("Type: %T - Value: %v\n", "steph", "steph")
	fmt.Println("Type: %T - Value: %v\n", 1, 1)
	fmt.Println("Type: %T - Value: %v\n", "1", "1")
	fmt.Println("Type: %T - Value: %v\n", 1.233, 1.233)
}

//tipos:
//bool (true/false)
//string - sequencia de bytes
//int
//float (float64/float32) - decimal
