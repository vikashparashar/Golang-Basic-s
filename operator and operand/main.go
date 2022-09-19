// Date : 19/09/2022 00:55 IST
// Operator and Operand in Golang

package main

import "fmt"

func main() {
	var a = 14
	var b = 15
	fmt.Println("A + B : ", a+b) // in this case a and b are operand's and '+' is a operator

}

// Golnag has many type of perator's like

// 1. Arithmetic Operators + , - , * , / , %
/*
	These are used to perform arithmetic/mathematical operations on operands in Go language:

	Addition: The ‘+’ operator adds two operands. For example, x+y.
	Subtraction: The ‘-‘ operator subtracts two operands. For example, x-y.
	Multiplication: The ‘*’ operator multiplies two operands. For example, x*y.
	Division: The ‘/’ operator divides the first operand by the second. For example, x/y.
	Modulus: The ‘%’ operator returns the remainder when the first operand is divided by the second. For example, x%y.
*/
// 2. Relational Operators
/*
 Relational operators are used for the comparison of two values. Let’s see them one by one:
‘=='(Equal To) operator checks whether the two given operands are equal or not. If so, it returns true. Otherwise, it returns false. For example, 5==5 will return true.
‘!='(Not Equal To) operator checks whether the two given operands are equal or not. If not, it returns true. Otherwise, it returns false. It is the exact boolean complement of the ‘==’ operator. For example, 5!=5 will return false.
‘>'(Greater Than)operator checks whether the first operand is greater than the second operand. If so, it returns true. Otherwise, it returns false. For example, 6>5 will return true.
‘<‘(Less Than)operator checks whether the first operand is lesser than the second operand. If so, it returns true. Otherwise, it returns false. For example, 6<5 will return false.
‘>='(Greater Than Equal To)operator checks whether the first operand is greater than or equal to the second operand. If so, it returns true. Otherwise, it returns false. For example, 5>=5 will return true.
‘<='(Less Than Equal To)operator checks whether the first operand is lesser than or equal to the second operand. If so, it returns true. Otherwise, it returns false. For example, 5<=5 will also return true.
*/

// 3. Logical Operators
/*
They are used to combine two or more conditions/constraints or to complement the evaluation of the original condition in consideration.
Logical AND: The ‘&&’ operator returns true when both the conditions in consideration are satisfied. Otherwise it returns false. For example, a && b returns true when both a and b are true (i.e. non-zero).
Logical OR: The ‘||’ operator returns true when one (or both) of the conditions in consideration is satisfied. Otherwise it returns false. For example, a || b returns true if one of a or b is true (i.e. non-zero). Of course, it returns true when both a and b are true.
Logical NOT: The ‘!’ operator returns true the condition in consideration is not satisfied. Otherwise it returns false. For example, !a returns true if a is false, i.e. when a=0.
*/

// 4. Bitwise Operators
/*

In Go language, there are 6 bitwise operators which work at bit level or used to perform bit by bit operations. Following are the bitwise operators :

& (bitwise AND): Takes two numbers as operands and does AND on every bit of two numbers. The result of AND is 1 only if both bits are 1.
| (bitwise OR): Takes two numbers as operands and does OR on every bit of two numbers. The result of OR is 1 any of the two bits is 1.
^ (bitwise XOR): Takes two numbers as operands and does XOR on every bit of two numbers. The result of XOR is 1 if the two bits are different.
<< (left shift): Takes two numbers, left shifts the bits of the first operand, the second operand decides the number of places to shift.
>> (right shift): Takes two numbers, right shifts the bits of the first operand, the second operand decides the number of places to shift.
&^ (AND NOT): This is a bit clear operator.

*/

// 5. Assignment Operators & More !!   ++, --, +=, -=
