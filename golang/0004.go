

package main

import "strconv"
import "fmt"

func main() {
    //const Start = 99
	const Start = 999
	var biggest = 0
    for left:=Start; left > 1; left-- {
    	for right := Start; right > 1; right-- {
    		var product = left*right
    		if (is_palindrome(product) && product > biggest) {
    			biggest = product
    		} 
    	}
    }
    fmt.Printf("biggest palindrome is %v\n", biggest)
}

func is_palindrome(product int) (bool)  {
    value := strconv.Itoa(product)
    for left, right:= -1,len(value)-1; right >= 0; right-- {
    	left++

   		if value[left] != value[right] {
   			return false
   		}
    }
    return true
}

