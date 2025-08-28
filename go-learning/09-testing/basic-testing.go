// basic-testing.go - Go 테스팅 기초
package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

// 테스트할 함수들 정의
func Add(a, b int) int {
	return a + b
}

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func IsEven(n int) bool {
	return n%2 == 0
}

func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// 구조체와 메서드
type Calculator struct {
	history []string
}

func NewCalculator() *Calculator {
	return &Calculator{
		history: make([]string, 0),
	}
}

func (c *Calculator) Add(a, b int) int {
	result := a + b
	c.history = append(c.history, fmt.Sprintf("%d + %d = %d", a, b, result))
	return result
}

func (c *Calculator) GetHistory() []string {
	return c.history
}

func (c *Calculator) Clear() {
	c.history = make([]string, 0)
}

// 에러를 반환하는 함수
type User struct {
	Name  string
	Age   int
	Email string
}

func CreateUser(name, email string, age int) (*User, error) {
	if name == "" {
		return nil, fmt.Errorf("name cannot be empty")
	}
	if age < 0 {
		return nil, fmt.Errorf("age cannot be negative")
	}
	if !strings.Contains(email, "@") {
		return nil, fmt.Errorf("invalid email format")
	}

	return &User{
		Name:  name,
		Age:   age,
		Email: email,
	}, nil
}

// 시간이 오래 걸리는 함수 (벤치마크용)
func SlowFunction(n int) int {
	time.Sleep(time.Millisecond * time.Duration(n))
	return n * n
}

// 복잡한 계산 함수 (벤치마크용)
func ComplexCalculation(data []int) int {
	sum := 0
	for _, v := range data {
		sum += v * v
	}
	return sum
}

// 팩토리얼 함수 (재귀)
func Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1)
}

// 문자열 처리 함수
func CountWords(text string) int {
	if text == "" {
		return 0
	}

	words := strings.Fields(text)
	return len(words)
}

func FindLongestWord(text string) string {
	if text == "" {
		return ""
	}

	words := strings.Fields(text)
	longest := ""

	for _, word := range words {
		if len(word) > len(longest) {
			longest = word
		}
	}

	return longest
}

// 슬라이스 관련 함수
func Sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func Max(numbers []int) (int, error) {
	if len(numbers) == 0 {
		return 0, fmt.Errorf("empty slice")
	}

	max := numbers[0]
	for _, num := range numbers[1:] {
		if num > max {
			max = num
		}
	}
	return max, nil
}

func Filter(numbers []int, predicate func(int) bool) []int {
	var result []int
	for _, num := range numbers {
		if predicate(num) {
			result = append(result, num)
		}
	}
	return result
}

// 맵 관련 함수
func CountCharacters(text string) map[rune]int {
	counts := make(map[rune]int)
	for _, char := range text {
		counts[char]++
	}
	return counts
}

// 동시성 관련 함수
func ProcessData(data []int) []int {
	result := make([]int, len(data))
	for i, v := range data {
		result[i] = v * 2
	}
	return result
}

// 파일 처리 시뮬레이션
func ReadConfig(filename string) (map[string]string, error) {
	// 실제 파일 읽기 대신 시뮬레이션
	if filename == "" {
		return nil, fmt.Errorf("filename cannot be empty")
	}

	if filename == "invalid.conf" {
		return nil, fmt.Errorf("file not found")
	}

	// 기본 설정 반환
	return map[string]string{
		"host":  "localhost",
		"port":  "8080",
		"debug": "true",
	}, nil
}

// 네트워크 요청 시뮬레이션
func FetchData(url string) ([]byte, error) {
	if url == "" {
		return nil, fmt.Errorf("URL cannot be empty")
	}

	if strings.HasPrefix(url, "http://invalid") {
		return nil, fmt.Errorf("invalid URL")
	}

	// 시뮬레이션된 데이터 반환
	return []byte(`{"message": "success", "data": [1, 2, 3]}`), nil
}

// 수학 관련 함수
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, fmt.Errorf("cannot calculate square root of negative number")
	}
	return math.Sqrt(x), nil
}

func Pow(base, exponent int) int {
	if exponent < 0 {
		return 0
	}
	if exponent == 0 {
		return 1
	}

	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result
}

// 타입 변환 함수
func ParseInt(s string) (int, error) {
	if s == "" {
		return 0, fmt.Errorf("empty string")
	}

	result := 0
	sign := 1
	i := 0

	if s[0] == '-' {
		sign = -1
		i = 1
	} else if s[0] == '+' {
		i = 1
	}

	for i < len(s) {
		if s[i] < '0' || s[i] > '9' {
			return 0, fmt.Errorf("invalid character: %c", s[i])
		}
		result = result*10 + int(s[i]-'0')
		i++
	}

	return result * sign, nil
}

func main() {
	fmt.Println("=== Go 테스팅 기초 ===")
	fmt.Println()

	fmt.Println("이 파일은 테스트 대상 함수들을 포함합니다.")
	fmt.Println("실제 테스트는 다음 파일들을 참고하세요:")
	fmt.Println("- basic-testing_test.go : 기본 테스트")
	fmt.Println("- advanced-testing_test.go : 고급 테스트")
	fmt.Println("- benchmark_test.go : 벤치마크 테스트")
	fmt.Println()

	fmt.Println("테스트 실행 방법:")
	fmt.Println("go test                    # 모든 테스트 실행")
	fmt.Println("go test -v                 # 상세 출력")
	fmt.Println("go test -run TestAdd       # 특정 테스트만 실행")
	fmt.Println("go test -bench=.           # 벤치마크 테스트 실행")
	fmt.Println("go test -cover             # 코드 커버리지 측정")
	fmt.Println("go test -race              # 레이스 컨디션 탐지")
	fmt.Println()

	// 함수들 간단 테스트
	fmt.Println("함수 동작 확인:")
	fmt.Printf("Add(2, 3) = %d\n", Add(2, 3))
	fmt.Printf("IsEven(4) = %t\n", IsEven(4))
	fmt.Printf("ReverseString('hello') = %s\n", ReverseString("hello"))
	fmt.Printf("IsPrime(17) = %t\n", IsPrime(17))

	calc := NewCalculator()
	calc.Add(5, 3)
	calc.Add(10, 2)
	fmt.Printf("Calculator history: %v\n", calc.GetHistory())

	user, err := CreateUser("John", "john@example.com", 25)
	if err != nil {
		fmt.Printf("Error creating user: %v\n", err)
	} else {
		fmt.Printf("Created user: %+v\n", user)
	}
}

/*
Go 테스팅 시스템 개요:

1. 테스트 파일 명명 규칙:
   - *_test.go 형태
   - 같은 패키지에 위치
   - 테스트 대상 파일과 함께 빌드되지 않음

2. 테스트 함수 규칙:
   - func TestXxx(*testing.T) 형태
   - 함수명은 Test로 시작
   - *testing.T 매개변수 필수

3. 벤치마크 함수 규칙:
   - func BenchmarkXxx(*testing.B) 형태
   - 함수명은 Benchmark로 시작
   - *testing.B 매개변수 필수

4. 예제 함수 규칙:
   - func ExampleXxx() 형태
   - 함수명은 Example로 시작
   - Output: 주석으로 예상 결과 명시

5. 주요 테스팅 메서드:
   - t.Error(), t.Errorf() : 테스트 실패 기록 후 계속
   - t.Fatal(), t.Fatalf() : 테스트 실패 후 즉시 종료
   - t.Log(), t.Logf() : 로그 출력 (-v 옵션 시)
   - t.Skip(), t.Skipf() : 테스트 건너뛰기

6. 테스트 실행 옵션:
   - go test : 모든 테스트 실행
   - go test -v : 상세 출력
   - go test -run=pattern : 패턴에 맞는 테스트만
   - go test -bench=pattern : 벤치마크 실행
   - go test -cover : 코드 커버리지
   - go test -race : 레이스 컨디션 탐지

7. 테스트 구조 (AAA 패턴):
   - Arrange : 테스트 준비
   - Act : 실행
   - Assert : 결과 검증

8. 테스트 모범 사례:
   - 독립적인 테스트 작성
   - 명확한 테스트 이름 사용
   - 테이블 기반 테스트 활용
   - 에러 케이스도 테스트
   - 경계값 테스트 포함
*/
