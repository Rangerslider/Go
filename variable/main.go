package main

import "fmt"

func main() {
	var a float64 //float num  print
	a = 15.155151

	var b, c int = 1, 2
	var d = true
	var e int
	f := "nuri"
	g := 52.71

	fmt.Println(a)
	fmt.Println(b, c)
	fmt.Println(d)
	fmt.Println(e) //e t kisu input kori nai tai output 0
	fmt.Println(f) //dynamic variable also g
	fmt.Println(g)
	fmt.Printf("g is %T", g) //type define kore dibe

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable type is : %T \n ", isLoggedIn)

	var smallfloat float64 = 250.36405
	fmt.Printf("Variable type is : %T \n ", smallfloat)
	fmt.Println(smallfloat)

}
