package main

import "fmt"

func main() {
	const Top = 100
    sum := uint64(0)
    sum_squares := uint64(0)

	for i:= uint64(1); i <= Top; i++ {
       sum += i
       sum_squares += i*i
	}
	println(sum_squares)
	println(sum*sum)
	diff :=  (sum*sum) - sum_squares 
	fmt.Printf("diff is %v", diff)	
}