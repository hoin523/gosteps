// syntax.go - Go의 기본 문법
package main

import "fmt"

func main() {
	// 1. 주석 (Comments)
	// 단일 행 주석은 //로 시작합니다.
	/*
		여러 줄 주석은
		이렇게 작성할 수 있습니다.
	*/

	// 2. 세미콜론 (Semicolons)
	// Go에서는 세미콜론이 자동으로 삽입됩니다.
	// 명시적으로 작성할 수도 있지만 권장하지 않습니다.
	fmt.Println("세미콜론은 자동으로 삽입됩니다")

	// 3. 중괄호와 들여쓰기
	// 여는 중괄호 {는 같은 줄에 있어야 합니다.
	if true {
		fmt.Println("올바른 중괄호 사용법")
	}

	// 4. 변수 명명 규칙
	// - camelCase 사용 (myVariable)
	// - 첫 글자가 대문자면 public, 소문자면 private
	var myVariable int = 42
	fmt.Printf("변수 값: %d\n", myVariable)

	// 5. 상수 선언
	const PI = 3.14159
	const MESSAGE = "상수 메시지"
	fmt.Printf("원주율: %f\n", PI)
	fmt.Println(MESSAGE)

	// 6. 여러 변수 선언
	var (
		name   string = "Go"
		age    int    = 14 // Go는 2009년에 출시
		active bool   = true
	)

	fmt.Printf("언어: %s, 나이: %d년, 활성: %t\n", name, age, active)

	// 7. 짧은 변수 선언 (:=)
	// 함수 내부에서만 사용 가능
	language := "Go"
	version := 1.21
	fmt.Printf("%s 버전 %g\n", language, version)

	// 8. 빈 식별자 (_)
	// 사용하지 않는 값을 무시할 때 사용
	value, _ := 42, "ignored"
	fmt.Printf("사용된 값: %d\n", value)

	// 9. 패키지 레벨 변수
	fmt.Printf("글로벌 카운터: %d\n", globalCounter)
}

// 패키지 레벨 변수 (전역 변수)
// 함수 외부에서 선언되며 패키지 전체에서 사용 가능
var globalCounter int = 100

/*
Go 코딩 스타일 가이드:
1. gofmt 도구를 사용하여 코드를 자동 포맷팅
2. 변수명은 간결하고 의미있게 작성
3. 패키지명은 소문자로, 간단하고 명확하게
4. 인터페이스명은 보통 -er로 끝남 (예: Reader, Writer)
5. 에러 처리는 명시적으로
*/
