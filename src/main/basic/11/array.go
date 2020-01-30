package main

func main() {
	// normal array
	var a [3]int
	a[0] = 1
	a[1] = 2
	a[2] = 3

	println(a[1])

	// array init
	var a1 = [3]int{1, 2, 3}
	var a3 = [...]int{1, 2, 3}
	println(a1[0])
	println(a3[1])

	// multi array
	var multiArray [3][4][5]int
	multiArray[0][1][2] = 10

	// multi array init
	var b = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	println(b[1][2])
}
