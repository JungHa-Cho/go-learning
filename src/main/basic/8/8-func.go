package main

func main() {
	msg := "Hello"
	say(&msg)
	println(msg)

	variableSay(msg, "is", "a", "book")

	total := sum(1, 7, 3, 5, 9)
	println(total)

	count, t := sum2(1, 7, 3, 5, 9)
	println(count, t)

	c, tt := sum3(1, 5, 3, 6, 8, 9)
	println(c, tt)
}

func say(msg *string) {
	println(*msg)
	*msg = "Changed"
}

func variableSay(msg ...string) {
	for _, s := range msg {
		println(s)
	}
}

func sum(nums ...int) int {
	s := 0
	for _, n := range nums {
		s += n
	}
	return s
}

func sum2(nums ...int) (int, int) {
	s := 0
	count := 0
	for _, n := range nums {
		s += n
		count++
	}
	return count, s
}

func sum3(nums ...int) (count int, total int) {
	for _, n := range nums {
		total += n
	}
	count = len(nums)
	return
}
