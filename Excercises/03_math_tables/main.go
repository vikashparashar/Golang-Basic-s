package main

import "fmt"

func main() {
	Table(5)
	Table(100)
}

func Table(x int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%v * %v = %v\n", x, i, x*i)
	}
}
