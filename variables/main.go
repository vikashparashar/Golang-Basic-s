// Date : 19/09/2022 00:55 IST
// Varibales And Constants In Golang

package main

import "fmt"

var a int = 54 // 1st method to decleare and initialize a variable in golang

func main() {
	var b = 54 // 2nd method to decleare and initialize a variable in golang
	b = 55     // re-assigning value to vaiable b
	// this is possible with variables only ,  not with constants because as per name varibales is en entity which can be changed if a programer or user want to change it during runtime
	// but this can not be possible with constants , they have fixed values once you decleare and store values to them
	c := 54 // 3rd method to decleare and initialize a variable in golang

	// Note :
	// Method No : 3 Is Called Short Hand Method Of Declearing A Variable And Storing Value In It
	// This Method Only Application In Function Level Varibale Declearation Not On Package Level

	// And Golang Has Only Short Hand Method For Variables Not For Constants

	fmt.Println(a)                                              // printing values stored inside variable a
	fmt.Printf(" Data Stored In Variable A is Of Type : %T", a) // printing type of value stored inside varible a
	fmt.Println(b)                                              // printing values stored inside variable b
	fmt.Printf(" Data Stored In Variable B is Of Type : %T", b) // printing type of value stored inside varible b
	fmt.Println(c)                                              // printing values stored inside variable c
	fmt.Printf(" Data Stored In Variable C is Of Type : %T", c) // printing type of value stored inside varible c

}
