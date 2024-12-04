package main

import "fmt"

func main() {
	
  var x,y int

  fmt.Scan(&x)
  fmt.Scan(&y)

  for cehck := false; !cehck; {
	x= x - y
	fmt.Println(x)
	cehck = x <= 0
}

fmt.Print(x == 0)
	
}