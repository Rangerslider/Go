package main

import "fmt"

/*type struct_name_1 struct{
	// Fields
}
type struct_name_2 struct{
	variable_name  struct_name_1
}*/

type Person struct {
	name string
	age  int
}

//create function thn call struct
func NewPerson(pass_Name string) *Person {

	// age and name local cilo tobuo amra main block a eita use korty parbo karon amra pointer return korsi

	p := Person{name: pass_Name}
	p.age = 22
	p.name = "nuri" // we dont use the passing name
	return &p
}

type User struct { // 1st latter is capital bcz user may needs to be exported out side of main
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {

	// This syntax creates a new struct.
	fmt.Println(Person{"Bob", 20})

	per := Person{"nuri", 22}
	fmt.Printf("created variable info is ==> %+v\n", per)

	// You can name the fields when initializing a struct.
	fmt.Println(Person{name: "Alice", age: 30})

	// Omitted fields will be zero-valued.
	fmt.Println(Person{name: "Fred"})

	// An `&` prefix yields a pointer to the struct.
	fmt.Println(&Person{name: "Ann", age: 40})

	// It's idiomatic to encapsulate new struct creation in constructor functions
	fmt.Println(NewPerson("Jon"))

	// Access struct fields with a dot.
	s := Person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// You can also use dots with struct pointers - the
	// pointers are automatically dereferenced.
	sp := &s
	fmt.Println(sp.age)

	// Structs are mutable.
	sp.age = 51
	fmt.Println(sp.age)

	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent

	nuri := User{"nuri", "nuri@go.dev", true, 16} // create user joy
	fmt.Println(nuri)
	fmt.Printf("details of nuri are: %+v\n", nuri) // %+v for detalied information
	fmt.Printf("Name is %v and email is %v.", nuri.Name, nuri.Email)

}

/*output
{Bob 20}
created variable info is ==> {name:nuri age:22}
{Alice 30}
{Fred 0}
&{Ann 40}
&{nuri 22}
Sean
50
51
Structs in golang
{nuri nuri@go.dev true 16}
details of nuri are: {Name:nuri Email:nuri@go.dev Status:true Age:16}
Name is nuri and email is nuri@go.dev.*/

/*type Employee struct {
	firstName string
	lastName  string
	age       int
	salary    int
}
func main() {

	//emp part  ekhane pointer use kore kora hoise
	emp8 := &Employee{
		firstName: "Sam",
		lastName:  "Anderson",
		age:       55,
		salary:    6000,
	}
	fmt.Println("\nFirst Name:", (*emp8).firstName)
	fmt.Println("Age:", (*emp8).age)

}*/
