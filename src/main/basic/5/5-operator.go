package main

import "fmt"

func main() {
	// 산술
	var a = 10
	var b = 5
	var c = (a + b) / 5
	c++

	fmt.Println(c)

	// 관계
	var d bool = true
	var e bool = false
	fmt.Println(d && e, d || e, d || !(e && d))

	// 비트 연산
	var f = 100
	fmt.Println(f << 5)

	// 축약 할당
	var g = 100
	g *= 10
	g >>= 2
	g |= 1
	fmt.Println(g)

	// 포인터
	var h int = 10
	var i = &h
	fmt.Println(*i)
}
