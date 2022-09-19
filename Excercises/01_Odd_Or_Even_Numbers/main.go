package main

import (
	"fmt"
)

func main() {
	Odd_Or_Even(1, 100000)
}

func Odd_Or_Even(from, upto int) {
	for i := from; i <= upto; i++ {
		if i%2 == 0 {
			fmt.Printf("%v is Even\n", i)
		} else {
			fmt.Printf("%v is Odd\n", i)
		}
		// time.Sleep(1 * time.Second)
	}

}
