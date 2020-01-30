package testlib2

var pop map[string]string

func init() { // 패키지 로드시 map 초기화
	pop = make(map[string]string)
}
