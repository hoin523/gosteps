// basic-functions.go - Go의 기본 함수 사용법
package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println("=== Go 함수 ===")

	// 1. 기본 함수 호출
	fmt.Println("\n1. 기본 함수 호출:")

	greet("Go 개발자")
	sayHello()

	// 2. 매개변수가 있는 함수
	fmt.Println("\n2. 매개변수가 있는 함수:")

	result := add(10, 20)
	fmt.Printf("10 + 20 = %d\n", result)

	area := calculateRectangleArea(5.0, 3.0)
	fmt.Printf("직사각형 넓이 (5 x 3) = %.2f\n", area)

	// 3. 여러 개의 반환값
	fmt.Println("\n3. 여러 개의 반환값:")

	quotient, remainder := divide(17, 3)
	fmt.Printf("17 ÷ 3 = %d 나머지 %d\n", quotient, remainder)

	// 일부 반환값 무시
	sum, _ := calculate(10, 5)
	fmt.Printf("10과 5의 합: %d (차이는 무시)\n", sum)

	// 4. 명명된 반환값 (Named Return Values)
	fmt.Println("\n4. 명명된 반환값:")

	perimeter, area2 := circleProperties(5.0)
	fmt.Printf("반지름 5인 원 - 둘레: %.2f, 넓이: %.2f\n", perimeter, area2)

	// 5. 가변 매개변수 (Variadic Parameters)
	fmt.Println("\n5. 가변 매개변수:")

	total1 := sum(1, 2, 3, 4, 5)
	fmt.Printf("1+2+3+4+5 = %d\n", total1)

	total2 := sum(10, 20)
	fmt.Printf("10+20 = %d\n", total2)

	// 슬라이스를 가변 매개변수로 전달
	numbers := []int{2, 4, 6, 8, 10}
	total3 := sum(numbers...) // ... 연산자로 슬라이스 확장
	fmt.Printf("2+4+6+8+10 = %d\n", total3)

	// 6. 함수를 변수로 사용
	fmt.Println("\n6. 함수를 변수로 사용:")

	// 함수를 변수에 할당
	var operation func(int, int) int = add
	result1 := operation(15, 25)
	fmt.Printf("함수 변수로 계산: %d\n", result1)

	// 함수를 매개변수로 전달
	result2 := applyOperation(10, 5, multiply)
	fmt.Printf("10 x 5 = %d\n", result2)

	result3 := applyOperation(10, 5, func(a, b int) int {
		return a - b // 익명 함수
	})
	fmt.Printf("10 - 5 = %d\n", result3)

	// 7. 클로저 (Closure)
	fmt.Println("\n7. 클로저:")

	// 카운터 클로저
	counter := createCounter()
	fmt.Printf("카운터: %d\n", counter()) // 1
	fmt.Printf("카운터: %d\n", counter()) // 2
	fmt.Printf("카운터: %d\n", counter()) // 3

	// 다른 카운터 인스턴스
	counter2 := createCounter()
	fmt.Printf("카운터2: %d\n", counter2()) // 1 (독립적)

	// 8. 재귀 함수
	fmt.Println("\n8. 재귀 함수:")

	fact := factorial(5)
	fmt.Printf("5! = %d\n", fact)

	fib := fibonacci(10)
	fmt.Printf("피보나치(10) = %d\n", fib)

	// 9. defer 키워드
	fmt.Println("\n9. defer 키워드:")

	demonstrateDefer()

	// 10. panic과 recover
	fmt.Println("\n10. panic과 recover:")

	fmt.Println("안전한 나눗셈 테스트:")
	result4 := safeDivide(10, 2)
	fmt.Printf("10 ÷ 2 = %.2f\n", result4)

	result5 := safeDivide(10, 0)
	fmt.Printf("10 ÷ 0 = %.2f\n", result5)

	// 11. 메서드와 함수의 차이
	fmt.Println("\n11. 메서드 예제:")

	p := Person{Name: "김철수", Age: 30}
	p.introduce()
	p.haveBirthday()
	p.introduce()

	// 12. 함수 실용 예제들
	fmt.Println("\n12. 실용 예제들:")

	// 문자열 처리
	text := "Hello, Go World!"
	reversed := reverseString(text)
	fmt.Printf("원본: %s, 뒤집기: %s\n", text, reversed)

	// 슬라이스 필터링
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evens := filter(nums, isEven)
	odds := filter(nums, isOdd)
	fmt.Printf("숫자들: %v\n", nums)
	fmt.Printf("짝수: %v\n", evens)
	fmt.Printf("홀수: %v\n", odds)

	// 맵 적용
	squares := mapInts(nums, square)
	fmt.Printf("제곱: %v\n", squares)

	// 리듀스 (합계)
	totalSum := reduce(nums, 0, addInt)
	fmt.Printf("총합: %d\n", totalSum)
}

// 1. 기본 함수들
func greet(name string) {
	fmt.Printf("안녕하세요, %s님!\n", name)
}

func sayHello() {
	fmt.Println("Hello, World!")
}

// 2. 반환값이 있는 함수
func add(a, b int) int {
	return a + b
}

func calculateRectangleArea(width, height float64) float64 {
	return width * height
}

// 3. 여러 반환값
func divide(dividend, divisor int) (int, int) {
	quotient := dividend / divisor
	remainder := dividend % divisor
	return quotient, remainder
}

func calculate(a, b int) (int, int) {
	sum := a + b
	diff := a - b
	return sum, diff
}

// 4. 명명된 반환값
func circleProperties(radius float64) (perimeter, area float64) {
	perimeter = 2 * math.Pi * radius
	area = math.Pi * radius * radius
	return // 명명된 반환값들이 자동으로 반환됨
}

// 5. 가변 매개변수
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// 6. 함수를 매개변수로 받는 함수
func multiply(a, b int) int {
	return a * b
}

func applyOperation(a, b int, op func(int, int) int) int {
	return op(a, b)
}

// 7. 클로저를 반환하는 함수
func createCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// 8. 재귀 함수
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// 9. defer 데모
func demonstrateDefer() {
	fmt.Println("함수 시작")

	defer fmt.Println("defer 1 - 마지막에 실행")
	defer fmt.Println("defer 2 - 두 번째로 실행")
	defer fmt.Println("defer 3 - 첫 번째로 실행")

	fmt.Println("함수 중간")
	fmt.Println("함수 끝")
	// defer들이 역순으로 실행됨: defer 3, defer 2, defer 1
}

// 10. panic과 recover
func safeDivide(a, b float64) (result float64) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("패닉 발생: %v\n", r)
			result = 0 // 기본값 설정
		}
	}()

	if b == 0 {
		panic("0으로 나눌 수 없습니다")
	}

	result = a / b
	return
}

// 11. 메서드 (구조체에 연결된 함수)
type Person struct {
	Name string
	Age  int
}

// 값 리시버 메서드
func (p Person) introduce() {
	fmt.Printf("안녕하세요, 저는 %s이고 %d세입니다.\n", p.Name, p.Age)
}

// 포인터 리시버 메서드 (값을 변경할 수 있음)
func (p *Person) haveBirthday() {
	p.Age++
	fmt.Printf("%s님이 생일을 맞아 %d세가 되었습니다!\n", p.Name, p.Age)
}

// 12. 함수형 프로그래밍 스타일 함수들
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func filter(slice []int, predicate func(int) bool) []int {
	var result []int
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

func isEven(n int) bool {
	return n%2 == 0
}

func isOdd(n int) bool {
	return n%2 != 0
}

func mapInts(slice []int, fn func(int) int) []int {
	result := make([]int, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

func square(n int) int {
	return n * n
}

func reduce(slice []int, initial int, fn func(int, int) int) int {
	result := initial
	for _, v := range slice {
		result = fn(result, v)
	}
	return result
}

func addInt(a, b int) int {
	return a + b
}

/*
Go 함수 특징:
1. func 키워드로 선언
2. 여러 개의 반환값 지원
3. 명명된 반환값 지원
4. 가변 매개변수 지원 (...type)
5. 함수는 일급 객체 (변수, 매개변수, 반환값으로 사용 가능)
6. 클로저 지원
7. defer 키워드로 지연 실행
8. panic/recover로 예외 처리
9. 메서드 (구조체에 연결된 함수)

함수 사용 팁:
1. 단일 책임 원칙: 함수는 하나의 일만 수행
2. 명명된 반환값으로 코드 가독성 향상
3. defer로 리소스 정리 보장
4. 클로저로 상태를 캡슐화
5. 함수형 프로그래밍 기법 활용
6. 재귀는 적절히 사용 (스택 오버플로 주의)
7. panic/recover는 예외적인 상황에만 사용
*/
