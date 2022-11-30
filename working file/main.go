package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("working with files")
	content := "nuri need a go file"

	file, err := os.Create("./mylocgofile.txt") //file err use korsi cs os pacakge file creart korte jie errror aste pare jate seta show kore tai

	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("length is :", length)
	defer file.Close()
	readFile("./mylocgofile.txt")
}
func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)

	}
	fmt.Println("text data inside the file is \n", databyte)         /// data file showing
	fmt.Println("text data inside the file is \n", string(databyte)) ///file content show for adding string
}

//output
/*working with files
length is : 19
text data inside the file is
 [110 117 114 105 32 110 101 101 100 32 97 32 103 111 32 102 105 108 101]
text data inside the file is
 nuri need a go file*/

//error function calling style method
// package main

// import (
// 	"fmt"
// 	"io"
// 	"io/ioutil"
// 	"os"
// )

// func main() {
// 	fmt.Println("working with files")
// 	content := "nuri need a go file"

// 	file, err := os.Create("./mylocgofile.txt")
// 	checkNilErr(err)
// 	length, err := io.WriteString(file, content)
// 	checkNilErr(err)
// 	fmt.Println("length is :", length)
// 	defer file.Close()
// 	readFile("./mylocgofile.txt")
// }
// func readFile(filename string) {
// 	databyte, err := ioutil.ReadFile(filename)
// 	checkNilErr(err)
// 	fmt.Println("text data inside the file is \n", databyte)
// 	fmt.Println("text data inside the file is \n", string(databyte))
// }
// func checkNilErr(err error) {  //here define the erorr func 59,61,68 no line call the error funtion
// 	if err != nil {
// 		panic(err)
// 	}
// }
