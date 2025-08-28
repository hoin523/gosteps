// variables.go - Go의 변수 선언과 초기화
package main

import "fmt"

// 패키지 레벨 변수 선언
var (
	globalInt    int    = 100
	globalString string = "전역 변수"
	globalBool   bool   = true
)

// 상수 선언
const (
	PI       = 3.14159
	APP_NAME = "Go Learning App"
	MAX_SIZE = 1000
)

// iota를 사용한 상수 열거
const (
	Sunday    = iota // 0
	Monday           // 1
	Tuesday          // 2
	Wednesday        // 3
	Thursday         // 4
	Friday           // 5
	Saturday         // 6
)

func main() {
	fmt.Println("=== Go 변수와 상수 ===")

	// 1. 변수 선언 방법들
	fmt.Println("\n1. 변수 선언 방법:")

	// 방법 1: var 키워드 사용 (타입 명시)
	var name string
	var age int
	var height float64

	// 초기값을 주지 않으면 제로값으로 초기화
	fmt.Printf("초기화 전 - 이름: '%s', 나이: %d, 키: %f\n", name, age, height)

	// 값 할당
	name = "김고랭"
	age = 25
	height = 175.5
	fmt.Printf("초기화 후 - 이름: %s, 나이: %d, 키: %.1fcm\n", name, age, height)

	// 방법 2: var 키워드 사용 (선언과 동시에 초기화)
	var city string = "서울"
	var temperature float64 = 23.5
	var isRaining bool = false

	fmt.Printf("도시: %s, 온도: %.1f°C, 비오는중: %t\n", city, temperature, isRaining)

	// 방법 3: var 키워드 사용 (타입 추론)
	var language = "Go" // string으로 추론
	var version = 1.21  // float64로 추론
	var popular = true  // bool로 추론

	fmt.Printf("언어: %s, 버전: %g, 인기: %t\n", language, version, popular)

	// 방법 4: 짧은 변수 선언 (:=) - 함수 내부에서만 사용 가능
	country := "대한민국"
	population := 51_780_000 // 언더스코어로 숫자 가독성 향상
	area := 100_210.0

	fmt.Printf("국가: %s, 인구: %d명, 면적: %.1fkm²\n", country, population, area)

	// 2. 여러 변수 동시 선언
	fmt.Println("\n2. 여러 변수 동시 선언:")

	// var를 사용한 여러 변수 선언
	var x, y, z int = 1, 2, 3
	fmt.Printf("x=%d, y=%d, z=%d\n", x, y, z)

	// 타입이 다른 여러 변수 선언
	var (
		username string  = "gopher"
		userid   int     = 12345
		active   bool    = true
		balance  float64 = 1250.75
	)
	fmt.Printf("사용자: %s (ID: %d), 활성: %t, 잔액: $%.2f\n", username, userid, active, balance)

	// 짧은 선언으로 여러 변수
	firstName, lastName, fullName := "길동", "홍", "홍길동"
	fmt.Printf("이름: %s %s (%s)\n", firstName, lastName, fullName)

	// 3. 변수 값 교체
	fmt.Println("\n3. 변수 값 교체:")

	a, b := 10, 20
	fmt.Printf("교체 전: a=%d, b=%d\n", a, b)

	// Go에서는 간단히 값을 교체할 수 있습니다
	a, b = b, a
	fmt.Printf("교체 후: a=%d, b=%d\n", a, b)

	// 4. 빈 식별자 (_) 사용
	fmt.Println("\n4. 빈 식별자 사용:")

	// 함수에서 여러 값을 반환할 때 일부만 사용하고 싶은 경우
	result, _ := divideAndRemainder(17, 3)
	fmt.Printf("17 ÷ 3의 몫: %d (나머지는 무시)\n", result)

	_, remainder := divideAndRemainder(17, 3)
	fmt.Printf("17 ÷ 3의 나머지: %d (몫은 무시)\n", remainder)

	// 5. 상수 사용
	fmt.Println("\n5. 상수 사용:")

	fmt.Printf("원주율: %f\n", PI)
	fmt.Printf("앱 이름: %s\n", APP_NAME)
	fmt.Printf("최대 크기: %d\n", MAX_SIZE)

	// iota 상수 사용
	fmt.Printf("오늘은 %s입니다\n", getDayName(Tuesday))

	// 6. 타입 변환
	fmt.Println("\n6. 타입 변환:")

	var intValue int = 42
	var floatValue float64 = float64(intValue)           // 명시적 타입 변환
	var stringValue string = fmt.Sprintf("%d", intValue) // 문자열로 변환

	fmt.Printf("정수: %d, 실수: %f, 문자열: %s\n", intValue, floatValue, stringValue)

	// 7. 전역 변수 사용
	fmt.Println("\n7. 전역 변수:")
	fmt.Printf("전역 정수: %d\n", globalInt)
	fmt.Printf("전역 문자열: %s\n", globalString)
	fmt.Printf("전역 불린: %t\n", globalBool)

	// 전역 변수 수정
	globalInt = 200
	globalString = "수정된 전역 변수"
	fmt.Printf("수정 후 - 전역 정수: %d, 전역 문자열: %s\n", globalInt, globalString)
}

// 두 개의 값을 반환하는 함수 (몫과 나머지)
func divideAndRemainder(dividend, divisor int) (int, int) {
	quotient := dividend / divisor
	remainder := dividend % divisor
	return quotient, remainder
}

// 요일 이름을 반환하는 함수
func getDayName(day int) string {
	days := []string{"일요일", "월요일", "화요일", "수요일", "목요일", "금요일", "토요일"}
	if day >= 0 && day < len(days) {
		return days[day]
	}
	return "알 수 없는 요일"
}

/*
변수 선언 규칙:
1. 변수명은 문자나 밑줄(_)로 시작
2. 이후에는 문자, 숫자, 밑줄 사용 가능
3. 대소문자 구분
4. Go 키워드는 변수명으로 사용 불가
5. 첫 글자가 대문자면 exported (다른 패키지에서 접근 가능)
6. 첫 글자가 소문자면 unexported (같은 패키지 내에서만 접근 가능)

제로값 (Zero Values):
- bool: false
- 숫자 타입: 0
- string: ""
- 포인터, 슬라이스, 맵, 채널, 함수: nil

변수 스코프:
- 패키지 레벨: 패키지 전체에서 접근 가능
- 함수 레벨: 함수 내부에서만 접근 가능
- 블록 레벨: {} 내부에서만 접근 가능
*/
