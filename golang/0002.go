package main

func main() {
	old1 := 1
	old2 := 1
	test := false
	testValue := 0
	sum := 0
    for ; test == false;  {
    	testValue = old1 + old2
    	old1 = old2
    	old2 = testValue
    	if testValue >= 4000000 {
      	    test = true
        }
   	    if testValue % 2 == 0 {
   		    sum += testValue
     	}
   	    println(sum)
    }
}