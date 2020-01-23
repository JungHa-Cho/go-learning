package main

import "fmt"

func main() {
	var a int
	var f float32 = 11.0

	// 일반적인 변수 할당
	fmt.Println(a, f)

	a = 10
	f = 12.32

	// 변수 값 변경
	fmt.Println(a, f)

	var i, j, k int = 1, 2, 3

	// 다중 변수 할당
	fmt.Println(i, j, k)

	const c int = 10
	const s string = "Hi"

	// 상수 할당
	fmt.Println(c, s)

	const cc = 10
	const ss = "HI WORLD"

	// 타입 추론
	fmt.Println(cc, ss)

	const (
		VISA   = "visa"
		MASTER = "MasterCard"
		AMEX   = "American Express"
	)

	fmt.Println(VISA, MASTER, AMEX)

	const (
		Apple  = iota // 0
		Grape         // 1
		Orange        // 2
	)

	fmt.Println(Apple, Grape, Orange)
}
