package main

import (
	"fmt"
)

func main() {

	// if staement
	kondisi := 8

	if kondisi == 10 {
		fmt.Println("good")
	} else if kondisi > 6 {
		fmt.Println("lumayan")
	} else {
		fmt.Println("jelek")
	}

	// switche case stetment
	var nilai int = 10

	switch nilai {
	case 8:
		fmt.Println("good")
	case 10:
		fmt.Println("awesome")
	default:
		fmt.Println("okelah")
	}

	// nested if

	var point = 10

	if point > 7 {
		switch point {
		case 10:
			fmt.Println("perfect!")
		default:
			fmt.Println("nice!")
		}
	} else {
		if point == 5 {
			fmt.Println("not bad")
		} else if point == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if point == 0 {
				fmt.Println("try harder!")
			}
		}
	}

}
