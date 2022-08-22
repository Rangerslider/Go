package main

import "fmt"

func main() {

	// sort using key  ... count the number of length(size) thn number thn  latter
	mp := make(map[string]string)
	mp["nuri0"] = "Ishmoth"
	mp["nuri1"] = "ura"
	mp["nuri"] = "nuri"
	mp["nuri2"] = "joy"

	fmt.Println(mp)

	// output ==> [nuri:nuri nuri0:Ishmoth nuri1:ura nuri2:joy]

	languages := make(map[string]string) // map  of string-key  string-value     make ...c t new die kaj kortam ekhane make cs eta non zero value dae
	// map[key]value

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	languages["GO"] = "Go"
	// insert in key-sorted  order

	fmt.Println("List of all languages: ", languages) //all print cmd at a time

	fmt.Println("JS shorts for: ", languages["JS"]) //accessing index by key

	// delete any element by key
	delete(languages, "RB")
	fmt.Println("After removing ruby List of all languages: ", languages)

	//inserting element
	languages["Abc"] = "BBBBb"
	fmt.Println("After insert Abc List of all languages: ", languages)

	// loops in map  in golang
	for key, value := range languages { // for key , value := range map_name
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
	//if we wnat to ignor key then ,syntax=>  _,value := range map_name
	for _, value := range languages {
		fmt.Printf("value is %v\n", value)
	}

	//if we wnat to ignor value then ,syntax=>  key,_ := range map_name
	for key, _ := range languages {
		fmt.Printf("key is %v\n", key)
	}

}

//output

// map[nuri:nuri nuri0:Ishmoth nuri1:ura nuri2:joy]
// List of all languages:  map[GO:Go JS:Javascript PY:Python RB:Ruby]
// JS shorts for:  Javascript
// After removing ruby List of all languages:  map[GO:Go JS:Javascript PY:Python]
// After insert Abc List of all languages:  map[Abc:BBBBb GO:Go JS:Javascript PY:Python]
// For key PY, value is Python
// For key GO, value is Go
// For key JS, value is Javascript
// For key Abc, value is BBBBb
// value is Javascript
// value is BBBBb
// value is Python
// value is Go
// key is Abc
// key is PY
// key is GO
// key is JS
