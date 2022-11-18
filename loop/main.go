package main

import "fmt"

func main() {

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	for i := range days { // for (auto i: vector ) this kind of format
		fmt.Println(days[i])
	}
	//syntx of range
	// for index, value := range anyDS {
	// 	fmt.Println(value)
	// 	}

	for _, day := range days { // we can ignor index , print only value
		fmt.Printf("value is %v\n", day)
	}

	// continue goto break

	num := 1
	for num < 10 {

		if num == 4 {
			goto nuri
		}
		if num == 5 {
			num++
			continue
		}
		fmt.Println("Value is: ", num)
		number++
	}

nuri:
	fmt.Printf("Hey nuri your number is %v\n", num)

}
