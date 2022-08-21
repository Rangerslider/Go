package main

import "fmt"

func main() {
	mynum := 23
	var ptr = &mynum
	fmt.Println("value of pointer  is :", *ptr)
	fmt.Println("value of location is :", ptr)

	*ptr = *ptr + 2
	fmt.Println("new sum pointer is :", mynum)
	*ptr = *ptr * 2
	fmt.Println("new multi pointer  is :", mynum)

}
