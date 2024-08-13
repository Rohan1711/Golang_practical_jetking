// 1) Write a Golang program for adding, subtracting,
// multiplying, dividing and finding a modulus of twonumbers given.

package main

import "fmt"

func main() {
	var num1 int
	var num2 int

	fmt.Println("Enter the value of number one :")
	fmt.Scan(&num1)
	fmt.Println("Enter the value of number two :")
	fmt.Scan(&num2)

	// addition
	result1 := num1 + num2
	fmt.Printf("\nresult of num1 + num2 = %d\n", result1)
	// subtraction
	result2 := num1 - num2
	fmt.Printf("\nresult of num1 - num2 = %d\n", result2)
	// multiflication
	result3 := num1 * num2
	fmt.Printf("\nresult of num1 * num2 = %d\n", result3)
	// division
	result4 := num1 / num2
	fmt.Printf("\nresult of num1 / num2 = %d\n", result4)
	// modulus
	result5 := num1 % num2
	fmt.Printf("\nresult of num1 %% num2 = %d\n", result5)

}
