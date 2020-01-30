package main

import (
	"fmt"

	"24lab.net/testlib"
	//mongo "other/mongo/db"
	//mysql "other/mysql/db"
)

func main() {
	// go package module, code reuse
	// package를 사용해 작은 단위의 컴포넌트를 작성
	// 이러한 작은 패키지를 활용해서 프로그램 작성을 권장
	// main package는 실행 프로그램으로 만든다
	// 패키지를 공유 라이브러리로 만들때는 main 패키지나 main 함수를 사용해서는 안된다.

	fmt.Println("HELLO")

	// import는 GO 컴파일러는 GOROOT 혹은 GOPATH 환경변수를 검색한다
	// 표준 패키지는 GOROOT/pkg 에서 사용자 패키지나 3rd party 패키지는 GOPATH/pkg

	// package scopre
	// 패키지 내 함수 구조체 인터페이스 메서드 등은 이름의 첫 문자를 대문자로 시작하게 되면 public
	// 소문자로 시작하면 non-public으로 패키지 내부에서만 사용됨

	// init은 패키지 사용시 실행되는 초기화 함수인데
	// init 함수만 호출하고자 하는 케이스인 경우
	// 패키지 import시 _ alias를 지정

	// package main
	// import _ "other/xlib"

	// mondb := mongo.Get()
	// mydb := mysql.Get()

	song := testlib.GetMusic("Alicia Keys")
	fmt.Println(song)
}
