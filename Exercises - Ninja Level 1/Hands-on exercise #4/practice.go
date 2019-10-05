package main

import "fmt"

type myOwnType int

func main() {
   var x myOwnType = 0
   fmt.Printf("Type of variable is %T \nvalue of variable is %v",x,x)
   x = 42
}
