package main

import "fmt"

var (
	a      float64
	b      float64
	choice int
)

func main() {
	fmt.Println("Enter First Number :")
	fmt.Scanf("%f\n", &a)
	fmt.Println("Enter Second Number :")
	fmt.Scanf("%f\n", &b)

	fmt.Println("1 ----- Addition")
	fmt.Println("2 ----- Subtraction")
	fmt.Println("3 ----- Multiplication")
	fmt.Println("4 ----- Division")
	fmt.Println("Enter Your Choice : ")
	fmt.Scanf("%d\n", &choice)
	switch choice {
	case 1:
		Addition(a, b)
	case 2:
		Subtraction(a, b)
	case 3:
		Multiplication(a, b)
	case 4:
		Division(a, b)
	}

}

func Addition(x, y float64) {
	fmt.Println("Result :", a+b)
}

func Subtraction(x, y float64) {
	fmt.Println("Result :", a-b)
}
func Multiplication(x, y float64) {
	fmt.Println("Result :", a*b)
}
func Division(x, y float64) {
	fmt.Println("Result's")
	fmt.Println("Quotient Is : ", int(x)/int(y))
	fmt.Println("Remainder Is : ", int(x)%int(y))
}
