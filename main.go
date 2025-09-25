package main

import "fmt"

func main(){
	fmt.Println("hello world")
	fmt.Println(intDivision(10,2))
}

// func printme(printValue string){
// 	fmt.Println(printValue)
// }

func intDivision(numerator int, denominator int)(int, int){
	var result int = numerator/denominator
	var remainder int = numerator%denominator
	return result, remainder
}