// hello.go - Go의 첫 번째 프로그램
// 모든 Go 프로그램은 package 선언으로 시작합니다.
// main 패키지는 실행 가능한 프로그램을 만들 때 사용됩니다.
package main

// import 문을 사용하여 다른 패키지를 가져올 수 있습니다.
// fmt는 형식화된 입출력을 위한 표준 라이브러리입니다.
import "fmt"

// main 함수는 프로그램의 진입점입니다.
// 모든 실행 가능한 Go 프로그램에는 main 함수가 있어야 합니다.
func main() {
	// fmt.Println()은 텍스트를 출력하고 새 줄을 추가합니다.
	fmt.Println("안녕하세요, Go 세계!")

	// fmt.Printf()는 형식화된 출력을 제공합니다.
	fmt.Printf("Go 버전: %s\n", "1.21")

	// fmt.Print()는 새 줄 없이 출력합니다.
	fmt.Print("Go는 ")
	fmt.Print("정말 ")
	fmt.Println("멋진 언어입니다!")
}

/*
실행 방법:
1. go run hello.go     // 직접 실행
2. go build hello.go   // 실행 파일 생성 후 실행
   ./hello (Linux/Mac) 또는 hello.exe (Windows)

Go 프로그램의 기본 구조:
- package 선언
- import 문
- 함수, 변수, 상수, 타입 선언
*/
