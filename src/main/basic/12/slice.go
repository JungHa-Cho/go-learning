package main

import "fmt"

func main() {
	//slice int
	var a []int
	a = []int{1, 2, 3}
	a[1] = 10
	fmt.Println(a)

	// make
	// make는 길이 용량을 임의 지정 가능
	// make첫번째 파라미터에 슬라이스 타입 지정, 두번째는 길이 세번째는 용량
	// 모든 요소가 Zero Value인 슬라이스를 만든다
	// Capacity 파라미터를 생략하면 Length와 같은 값을 가진다
	s := make([]int, 5, 10)
	fmt.Println(len(s), cap(s))
	fmt.Println(s)

	// slice zero value = nil slice
	var f []int
	if f == nil {
		fmt.Println("NIL SLICE")
	}

	fmt.Println(len(f), cap(f))

	// sub slice
	c := []int{0, 1, 2, 3, 4, 5}
	c = c[2:5]
	fmt.Println(c)

	// index slice
	q := []int{0, 1, 2, 3, 4, 5}
	q = q[2:5]
	q = q[1:]
	fmt.Println(q)

	// append and copy
	p := []int{0, 1}
	p = append(p, 2)
	fmt.Println(p)

	p = append(p, 2, 3, 4, 5, 6)
	fmt.Println(p)

	// append movement
	// 슬라이스 용량 capacity가 남이 있을 경우 그 용량 내에서 슬라이스의 길이 length를 변경하여 데이타를 추가
	// 용량을 초과하는 경우 현재 용량의 2배에 해당하는 새로운 underlying array를 생성하고 기존 배열 값들을 모두 새 배열에 복제한 후 다시 슬라이스를 할당한다.

	// len = 0, cap = 3 slice
	sliceA := make([]int, 0, 3)

	// keep add
	for i := 1; i <= 15; i++ {
		sliceA = append(sliceA, i)
		// slice len cap check
		fmt.Println(len(sliceA), cap(sliceA))
	}
	fmt.Println(sliceA)

	// slice to slice add
	// 뒤에 병합됨
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}

	sliceTotal := append(slice1, slice2...)
	fmt.Println(sliceTotal)

	// slice는 실제 배열을 가리키는 포인터 정보만 가진다.
	// 따라서 슬라이스 복사는 소스 슬라이스가 갖는 배열의 데이타를 타겟 슬라이스가 갖는 배열로 복제하는 것
	source := []int{0, 1, 2}
	target := make([]int, len(source), cap(source)*2)
	copy(target, source)
	fmt.Println(target)
	fmt.Println(len(target), cap(target))

	// slice는 배열 부분 영역인 세그먼트에 대한 메타 정보를 가지고 있음
	// 슬라이스는 크게 3개의 필드로 구성
	// 1. 내부적으로 사용하는 배열의 포인터 정보
	// 2. 세그먼트 길이 length
	// 3. 세그먼트의 최대 용량 capacity

	v := []int{0, 1, 2, 3, 4, 5}
	// slice meta [pointer is array [0] pointing, length 6, cap 6]
	// internal array [P0][1][2][3][4][5]
	r := v[2:5]
	fmt.Println(r)
	// slice meta [pointer is array[2] pointing, length 3, capacity 4]
	// internal array [0][1][P2][3][4][5]
}
