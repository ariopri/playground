package main

import "fmt"

// Check Point:
// Diamond Pattern
// - Input: Size: 5
// - Output:
//         *
//        ***
//       *****
//      *******
//     *********
//    ***********
//     *********
//      *******
//       *****
//        ***
//         *

func main() {
	var size int
	fmt.Print("Size: ")
	fmt.Scanf("%d", &size)

	// TODO: answer here
	for i := 0; i < size; i++ {
		for j := size; j > i; j-- {
			fmt.Printf(" ")
		}
		for k := 0; k <= i; k++ {
			fmt.Printf("*")
		}
		for l := 0; l < i; l++ {
			fmt.Printf("*")
		}
		fmt.Println("")
	}
	for p := 0; p <= size; p++ {
		for  q:= 0; q < p; q++ {
			fmt.Printf(" ")
		}
		for r := size; r >= p; r-- {
			fmt.Printf("*")
		}
		for s := size; s > p; s-- {
			fmt.Printf("*")
		}
		fmt.Println("")
	}

}