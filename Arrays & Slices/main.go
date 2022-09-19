// Date : 19/09/2022 00:55 IST
// Array & Slices In Golang

// Go program to illustrate how to create an array using shorthand declaration and accessing the elements of the array using for loop
package main

import "fmt"

func main() {

	// Shorthand declaration of array in golang
	a := [4]string{"Vikash", "Ritika", "Khushboo", "Niyati"} // a array of size 4 means it can store 4 string type values in it

	// Shorthand declaration of slice in golang
	s := []string{"Vikash", "Ritika", "Khushboo", "Niyati"} // but if did not given a size to an array then it become a slice

	// Accessing the elements of the array Using for loop
	fmt.Println("Elements of the array:")

	for i := 0; i < 3; i++ {
		fmt.Println(a[i])
	}

	// this is another method of accessing value stoed in a slice or an array using this for range method

	for i, v := range s {
		fmt.Println(" Index : ", i, "Value : ", v)
		fmt.Printf("Length Of Slice Is : %v", len(s)) // for finding size of an array or a slice in golang
	}

}
