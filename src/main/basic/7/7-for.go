package main

func main() {
	normal()
	onlyCondition()
	//infinity()
	forRange()
	breakContinueGoto()
	breakGoto()
}

func normal() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += 1
	}
	println(sum)
}

func onlyCondition() {
	n := 1
	for n < 100 {
		n *= 2
	}
	println(n)
}

func infinity() {
	for {
		println("infinite loop")
	}
}

func forRange() {
	names := []string{"홍길동", "이순신", "강감찬"}
	for index, name := range names {
		println(index, name)
	}
}

func breakContinueGoto() {
	var a = 1
	for a < 15 {
		if a == 5 {
			a += a
			continue
		}
		a++
		if a > 10 {
			break
		}
	}
	if a == 11 {
		goto END
	}
	println(a)
END:
	println("END")
}

func breakGoto() {
	i := 0
L1:
	for {
		if i == 0 {
			break L1
		}
	}
	println("OK")
}
