// Date : 19/09/2022 00:55 IST
// Function And Constants In Golang

/*

Note : Funtion's Are First Class Citizen's In Golang
They Can Be Used Mostly EveryWhere

Function Is A Method Of Doing Task Repeatly As We Want Without Repeating Code So Many Time

*/

// Let's See Funtion's In Action's

// Syntex For Funtion's In Golang

// func (func is a keyword ) (argument\s -  if any ! )function name(parameter's -  if any ! )(return's -  if any !){
// 	// code in action
// }

// A function can take 0 or more arguments
// A function can take 0 or more paramenters
// A function can give's us 0 or more returns values
// A Function can be decleared By giving a name to it or without name - these type of functions are called anonymous function

// you can decleare a function anywhere as you want - in a package level or inside a function also

package main

import "fmt"

// let's add 2 numbers in a simple form
func main() {
	var a int = 164
	var b int = 100
	var total = a + b
	fmt.Println("Addition Of A + B Is :", total)

	// This Is An Example Of anonymous function - a function without an identification / name
	func() {
		fmt.Println(a + b)
	}() // here we are calling that anonymous function ,
}

// what if we want add 2 number multiple times in our code ?
// it's so annoying to write same code repeatly to do that task

// so here comes function

// so lets create a function which will take 2 numbers and add them

func Add(x, y int) { // func - keyword , Add - Name Of Function , x&y Are The Parameters For This Function

	fmt.Println(x + y)
}

// Now Lets Call This Add Function In Our Main Function And See The Magic !

// Try It

// this is a function with 2 arguments and 1 return value
func Addition(x, y int) int { // func - keyword , Add - Name Of Function , x&y Are The Parameters For This Function
	return x + y
}
