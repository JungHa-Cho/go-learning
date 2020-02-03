package main

import "fmt"

type person struct {
	name string
	age  int
}

type dict struct {
	data map[int]string
}

func main() {
	// case 1
	p := person{}

	p.name = "Lee"
	p.age = 10

	fmt.Println(p)

	// case 2
	var p1 person
	p1 = person{"Bob", 20}
	p2 := person{name: "Sean", age: 50}

	fmt.Println(p1, p2)

	// case 3
	p3 := new(person)
	p3.name = "Lee"

	fmt.Println(p3)

	// dic := newDict()
	// dic.data[1] = "A"
}

//type newDict() *dict {
//	d := dict{}
//	d.data = map[int]string{}
//	return &d
//}
