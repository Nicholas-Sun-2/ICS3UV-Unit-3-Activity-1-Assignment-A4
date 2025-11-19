/**
 * @author Nicholas Sun
 * @version 1.0.0
 * @date 2025-11-18
 * @fileoverview This program prints the average for the following numbers, 56.9, 89.7, 90.2.
 */

package main

import "fmt"

func main(){
	// Declare variables
	var num1 float32 = 56.9;
	var num2 float32 = 89.7;
	var num3 float32 = 90.2;
	var average float32 = (num1 + num2 + num3) / 3.0;

	// Print the output
	fmt.Println("The average of " + fmt.Sprint(num1) + ", " + fmt.Sprint(num2) + " and " + fmt.Sprint(num3) + " is " + fmt.Sprint(average) + ".");
}