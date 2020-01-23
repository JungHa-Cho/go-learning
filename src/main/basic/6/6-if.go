package main

import "fmt"

func main() {
	// if
	k := 1
	if k == 1 {
		fmt.Println("ONE")
	}

	if k == 0 {
		fmt.Println("ZERO")
	} else if k == 1 {
		fmt.Println("ONE")
	} else {
		fmt.Println("OTHER")
	}

	// switch
	var name string
	var category = 1

	switch category {
	case 1:
		name = "Paper Book"
	case 2:
		name = "eBook"
	case 3, 4:
		name = "Blog"
	default:
		name = "Other"
	}
	println(name)

	// For Expression
	switch x := category << 2; x - 1 {
	// ...
	}

	grade(100)
	grade(70)
	checkType(21)
	checkType("hello")
	checkType(true)
	checkType(21.230)
}

func checkType(i interface{}) {
	// type check
	switch v := i.(type) {
	case int:
		println("int %v is %v", v, v*2)
	case bool:
		println("%v", v)
	case string:
		println("string %q is %v", v, len(v))
	default:
		println("unknown")
	}
}

func grade(score int) {
	switch {
	case score >= 90:
		println("A")
	case score >= 80:
		println("B")
	case score >= 70:
		println("C")
	case score >= 70:
		println("C")
	default:
		println(" NO HOPE")
	}
}
