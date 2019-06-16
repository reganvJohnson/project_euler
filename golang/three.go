

package main

//import "fmt"

func main() {
  const Size = 775147
  const Test = 600851475143
 // const Test = 400
  //const Size = 20
  var a[Size+1] bool
  for increment:=2; increment < Size; increment++ {
    for clearout:=increment*2; clearout <= Size; clearout+=increment {
        //fmt.Printf("doing %v\n", clearout)
        a[clearout] = true
    }
  }

//  for i:=0; i < Size; i++ {
//      fmt.Printf("%v %v\n", i, a[i])
//  }
  for i:=Size; i > 2; i-- {
    if a[i] == false && Test % i == 0 {
      println (i)
      break
    }
  }
}