// conditionals.go - Go의 조건문 (if, if-else, switch)
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("=== Go 조건문 ===")

	// 1. 기본 if 문
	fmt.Println("\n1. 기본 if 문:")

	age := 25
	if age >= 18 {
		fmt.Printf("나이 %d세: 성인입니다\n", age)
	}

	// 조건식에서 변수 선언과 동시에 검사 (짧은 구문)
	if currentHour := time.Now().Hour(); currentHour < 12 {
		fmt.Println("좋은 아침입니다!")
	} else if currentHour < 17 {
		fmt.Println("좋은 오후입니다!")
	} else {
		fmt.Println("좋은 저녁입니다!")
	}

	// 2. if-else 문
	fmt.Println("\n2. if-else 문:")

	temperature := 28.5
	if temperature > 30 {
		fmt.Printf("온도 %.1f°C: 매우 덥습니다\n", temperature)
	} else if temperature > 20 {
		fmt.Printf("온도 %.1f°C: 쾌적한 날씨입니다\n", temperature)
	} else if temperature > 10 {
		fmt.Printf("온도 %.1f°C: 조금 쌀쌀합니다\n", temperature)
	} else {
		fmt.Printf("온도 %.1f°C: 매우 춥습니다\n", temperature)
	}

	// 3. 복합 조건문
	fmt.Println("\n3. 복합 조건문:")

	score := 85
	attendance := 95

	// AND 연산자 (&&)
	if score >= 80 && attendance >= 90 {
		fmt.Printf("점수 %d점, 출석률 %d%%: 우수 학생입니다\n", score, attendance)
	}

	// OR 연산자 (||)
	if score < 60 || attendance < 70 {
		fmt.Println("보충 학습이 필요합니다")
	} else {
		fmt.Println("양호한 성과입니다")
	}

	// NOT 연산자 (!)
	isWeekend := false
	if !isWeekend {
		fmt.Println("평일입니다. 열심히 일하세요!")
	}

	// 4. switch 문 (기본)
	fmt.Println("\n4. switch 문 (기본):")

	dayOfWeek := time.Now().Weekday()
	switch dayOfWeek {
	case time.Monday:
		fmt.Println("월요일: 새로운 한 주의 시작!")
	case time.Tuesday:
		fmt.Println("화요일: 힘내세요!")
	case time.Wednesday:
		fmt.Println("수요일: 주의 중간점!")
	case time.Thursday:
		fmt.Println("목요일: 거의 다 왔어요!")
	case time.Friday:
		fmt.Println("금요일: 불금이에요!")
	case time.Saturday, time.Sunday: // 여러 케이스를 하나로 처리
		fmt.Println("주말: 쉬는 날!")
	default:
		fmt.Println("알 수 없는 요일")
	}

	// 5. switch 문 (표현식 없는 switch)
	fmt.Println("\n5. switch 문 (표현식 없는 switch):")

	number := rand.Intn(100) // 0-99 사이의 랜덤 숫자
	fmt.Printf("랜덤 숫자: %d\n", number)

	switch {
	case number < 10:
		fmt.Println("한 자리 수입니다")
	case number < 50:
		fmt.Println("50 미만입니다")
	case number < 80:
		fmt.Println("50 이상 80 미만입니다")
	default:
		fmt.Println("80 이상입니다")
	}

	// 6. switch 문 (초기화문과 함께)
	fmt.Println("\n6. switch 문 (초기화문과 함께):")

	switch grade := getGrade(88); grade {
	case "A":
		fmt.Println("최우수 성적입니다!")
	case "B":
		fmt.Println("우수 성적입니다!")
	case "C":
		fmt.Println("보통 성적입니다")
	case "D":
		fmt.Println("미흡한 성적입니다")
	case "F":
		fmt.Println("재시험이 필요합니다")
	default:
		fmt.Printf("알 수 없는 등급: %s\n", grade)
	}

	// 7. fallthrough 키워드
	fmt.Println("\n7. fallthrough 사용:")

	month := int(time.Now().Month())
	fmt.Printf("현재 월: %d월\n", month)

	switch month {
	case 12, 1, 2:
		fmt.Println("겨울입니다")
		fmt.Println("따뜻하게 입으세요")
	case 3, 4, 5:
		fmt.Println("봄입니다")
		fmt.Println("꽃이 피는 계절이에요")
	case 6, 7, 8:
		fmt.Println("여름입니다")
		fmt.Println("시원하게 지내세요")
	case 9, 10, 11:
		fmt.Println("가을입니다")
		fmt.Println("단풍이 아름다운 계절이에요")
	}

	// 8. 타입 switch (인터페이스 타입 판별)
	fmt.Println("\n8. 타입 switch:")

	var items []interface{} = []interface{}{
		42,
		"Hello",
		3.14,
		true,
		[]int{1, 2, 3},
	}

	for i, item := range items {
		fmt.Printf("항목 %d: ", i)
		switch v := item.(type) {
		case int:
			fmt.Printf("정수 %d\n", v)
		case string:
			fmt.Printf("문자열 '%s' (길이: %d)\n", v, len(v))
		case float64:
			fmt.Printf("실수 %.2f\n", v)
		case bool:
			fmt.Printf("불린 %t\n", v)
		case []int:
			fmt.Printf("정수 슬라이스 %v (길이: %d)\n", v, len(v))
		default:
			fmt.Printf("알 수 없는 타입: %T\n", v)
		}
	}

	// 9. 조건문 활용 예제
	fmt.Println("\n9. 조건문 활용 예제:")

	// 윤년 판별
	year := time.Now().Year()
	fmt.Printf("올해 (%d년)는 ", year)
	if isLeapYear(year) {
		fmt.Println("윤년입니다")
	} else {
		fmt.Println("평년입니다")
	}

	// 숫자 분류
	numbers := []int{-5, 0, 3, 8, -2, 15}
	positive, negative, zero := classifyNumbers(numbers)
	fmt.Printf("숫자들 %v을 분류한 결과:\n", numbers)
	fmt.Printf("양수: %v\n", positive)
	fmt.Printf("음수: %v\n", negative)
	fmt.Printf("0: %v\n", zero)
}

// 점수를 학점으로 변환하는 함수
func getGrade(score int) string {
	switch {
	case score >= 90:
		return "A"
	case score >= 80:
		return "B"
	case score >= 70:
		return "C"
	case score >= 60:
		return "D"
	default:
		return "F"
	}
}

// 윤년 판별 함수
func isLeapYear(year int) bool {
	// 4로 나누어떨어지고, 100으로 나누어떨어지지 않거나, 400으로 나누어떨어지면 윤년
	return (year%4 == 0 && year%100 != 0) || year%400 == 0
}

// 숫자들을 양수, 음수, 0으로 분류하는 함수
func classifyNumbers(numbers []int) (positive, negative, zero []int) {
	for _, num := range numbers {
		if num > 0 {
			positive = append(positive, num)
		} else if num < 0 {
			negative = append(negative, num)
		} else {
			zero = append(zero, num)
		}
	}
	return // named return values 사용
}

/*
Go 조건문 특징:
1. 조건식에 괄호 () 없어도 됨 (하지만 중괄호 {} 필수)
2. 조건식은 반드시 bool 타입이어야 함
3. if문에서 초기화문 사용 가능 (변수의 스코프는 if 블록으로 제한)
4. switch문에 break 불필요 (자동으로 break됨)
5. fallthrough 키워드로 다음 case로 진행 가능
6. switch문에서 여러 값을 하나의 case로 처리 가능
7. 타입 switch로 인터페이스의 실제 타입 판별 가능

조건문 사용 팁:
1. 복잡한 조건은 함수로 분리하여 가독성 향상
2. switch문은 긴 if-else if 체인보다 읽기 쉬움
3. 초기화문을 활용하여 변수 스코프 최소화
4. 타입 switch는 인터페이스 처리에 유용
*/
