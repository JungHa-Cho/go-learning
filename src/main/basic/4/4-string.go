package main

import "fmt"

func main() {
	rawLiteral := `가나다\n
  라마바\n
  사아차타카파하`

	interLiteral := "가나다\n라마바사"
	interLiteral2 := "가나다" +
		"\n라마바사"

	fmt.Println(rawLiteral, interLiteral, interLiteral2)

	var i int = 100
	var u uint = uint(i)
	var f float32 = float32(i)
	fmt.Println(i, u, f)
}
