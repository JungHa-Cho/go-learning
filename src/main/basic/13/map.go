package main

import "fmt"

func main() {
	// nil map => becuse reference type
	var idMap map[int]string

	// init
	idMap = make(map[int]string)

	fmt.Println(idMap)

	// literal init
	tockers := map[string]string{
		"GOOG": "google Inc",
		"MSFT": "Microst",
		"FB":   "FaceBook",
	}

	fmt.Println(tockers)

	// using map
	// init map using make is no data
	var m map[int]string
	m = make(map[int]string)

	// add or replace
	m[901] = "Apple"
	m[134] = "Grape"
	m[777] = "Tomato"

	// read key value
	str := m[134]
	fmt.Println(str)

	noData := m[999]
	fmt.Println(noData) // return nil or zero

	delete(m, 777)

	tockers2 := map[string]string{
		"GOOG": "google Inc",
		"MSFT": "Microst",
		"FB":   "FaceBook",
		"AMZN": "Amazone",
	}

	val, exists := tockers2["MSFT"]

	if exists {
		fmt.Println(val, exists)
	} else if !exists {
		fmt.Println("No MSFT ticker")
	}

	myMap := map[string]string{
		"A": "Apple",
		"B": "Banana",
		"C": "Charlie",
	}

	for key, val := range myMap {
		fmt.Println(key, val)
	}
}
