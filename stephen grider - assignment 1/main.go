package main

import "fmt"

func main() {
	myInts := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := range myInts {
		if i == 0 {
			fmt.Printf("%v is neither even nor odd number.\n", i)
			continue
		}
		if i % 2 == 0 {
			fmt.Printf("%v is even number.\n", i)
		} else {
			fmt.Printf("%v is odd number.\n", i)
		}
	} 
}