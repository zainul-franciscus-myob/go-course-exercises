package main

import "fmt"

type myOwnType int
var x1 myOwnType
var y1 int

func main() {
	fmt.Printf("value of x1 %v\n",x1)
	fmt.Printf("type of x1 %T",x1)

	x1 = 42
	fmt.Printf("value of x1 %v\n",x1)

	y1 = int(x1)
	fmt.Printf("value of y1 %v\n",y1)
}
