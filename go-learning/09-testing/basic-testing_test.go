// basic-testing_test.go - Go 기본 테스트 예제
package main

import (
	"reflect"
	"testing"
)

// 1. 기본 테스트 함수
func TestAdd(t *testing.T) {
	// 테스트 케이스
	result := Add(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("Add(2, 3) = %d; expected %d", result, expected)
	}
}

// 2. 테이블 기반 테스트 (Table-driven test)
func TestAddTableDriven(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"positive numbers", 2, 3, 5},
		{"negative numbers", -1, -1, -2},
		{"mixed numbers", 10, -5, 5},
		{"with zero", 0, 5, 5},
		{"both zero", 0, 0, 0},
		{"large numbers", 1000000, 2000000, 3000000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%d, %d) = %d; expected %d",
					tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// 3. 서브테스트 사용 (t.Run)
func TestMultiply(t *testing.T) {
	t.Run("positive numbers", func(t *testing.T) {
		result := Multiply(3, 4)
		if result != 12 {
			t.Errorf("Multiply(3, 4) = %d; expected 12", result)
		}
	})

	t.Run("negative numbers", func(t *testing.T) {
		result := Multiply(-2, 3)
		if result != -6 {
			t.Errorf("Multiply(-2, 3) = %d; expected -6", result)
		}
	})

	t.Run("zero multiplication", func(t *testing.T) {
		result := Multiply(5, 0)
		if result != 0 {
			t.Errorf("Multiply(5, 0) = %d; expected 0", result)
		}
	})
}

// 4. 에러 처리 테스트
func TestDivide(t *testing.T) {
	// 정상 케이스
	result, err := Divide(10, 2)
	if err != nil {
		t.Errorf("Divide(10, 2) returned error: %v", err)
	}
	if result != 5.0 {
		t.Errorf("Divide(10, 2) = %f; expected 5.0", result)
	}

	// 0으로 나누기 에러 케이스
	_, err = Divide(10, 0)
	if err == nil {
		t.Error("Divide(10, 0) should return an error")
	}
}

// 5. 불린 반환 함수 테스트
func TestIsEven(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		{2, true},
		{4, true},
		{1, false},
		{3, false},
		{0, true},
		{-2, true},
		{-1, false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("IsEven(%d)", tt.input), func(t *testing.T) {
			result := IsEven(tt.input)
			if result != tt.expected {
				t.Errorf("IsEven(%d) = %t; expected %t",
					tt.input, result, tt.expected)
			}
		})
	}
}

// 6. 문자열 처리 함수 테스트
func TestReverseString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"simple string", "hello", "olleh"},
		{"single character", "a", "a"},
		{"empty string", "", ""},
		{"palindrome", "racecar", "racecar"},
		{"numbers", "12345", "54321"},
		{"korean", "안녕", "녕안"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ReverseString(tt.input)
			if result != tt.expected {
				t.Errorf("ReverseString(%q) = %q; expected %q",
					tt.input, result, tt.expected)
			}
		})
	}
}

// 7. 수학 함수 테스트
func TestIsPrime(t *testing.T) {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	nonPrimes := []int{0, 1, 4, 6, 8, 9, 10, 12, 14, 15, 16, 18, 20}

	// 소수 테스트
	for _, prime := range primes {
		t.Run(fmt.Sprintf("prime_%d", prime), func(t *testing.T) {
			if !IsPrime(prime) {
				t.Errorf("IsPrime(%d) = false; expected true", prime)
			}
		})
	}

	// 비소수 테스트
	for _, nonPrime := range nonPrimes {
		t.Run(fmt.Sprintf("non_prime_%d", nonPrime), func(t *testing.T) {
			if IsPrime(nonPrime) {
				t.Errorf("IsPrime(%d) = true; expected false", nonPrime)
			}
		})
	}
}

// 8. 구조체와 메서드 테스트
func TestCalculator(t *testing.T) {
	calc := NewCalculator()

	// 초기 상태 테스트
	if len(calc.GetHistory()) != 0 {
		t.Error("New calculator should have empty history")
	}

	// Add 메서드 테스트
	result := calc.Add(5, 3)
	if result != 8 {
		t.Errorf("Calculator.Add(5, 3) = %d; expected 8", result)
	}

	// 히스토리 테스트
	history := calc.GetHistory()
	if len(history) != 1 {
		t.Errorf("History length = %d; expected 1", len(history))
	}

	expectedHistory := "5 + 3 = 8"
	if history[0] != expectedHistory {
		t.Errorf("History[0] = %q; expected %q", history[0], expectedHistory)
	}

	// 여러 연산 후 히스토리 테스트
	calc.Add(10, 2)
	calc.Add(1, 1)

	history = calc.GetHistory()
	if len(history) != 3 {
		t.Errorf("History length = %d; expected 3", len(history))
	}

	// Clear 메서드 테스트
	calc.Clear()
	if len(calc.GetHistory()) != 0 {
		t.Error("Calculator history should be empty after Clear()")
	}
}

// 9. 사용자 생성 함수 테스트
func TestCreateUser(t *testing.T) {
	// 정상 케이스
	user, err := CreateUser("John Doe", "john@example.com", 25)
	if err != nil {
		t.Errorf("CreateUser with valid data returned error: %v", err)
	}
	if user == nil {
		t.Fatal("CreateUser should return a user object")
	}
	if user.Name != "John Doe" {
		t.Errorf("User.Name = %q; expected %q", user.Name, "John Doe")
	}
	if user.Age != 25 {
		t.Errorf("User.Age = %d; expected 25", user.Age)
	}

	// 에러 케이스들
	errorCases := []struct {
		name        string
		userName    string
		email       string
		age         int
		expectError bool
	}{
		{"empty name", "", "test@example.com", 25, true},
		{"negative age", "John", "john@example.com", -1, true},
		{"invalid email", "John", "invalid-email", 25, true},
		{"valid data", "Jane", "jane@example.com", 30, false},
	}

	for _, tc := range errorCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := CreateUser(tc.userName, tc.email, tc.age)
			if tc.expectError && err == nil {
				t.Error("Expected error but got none")
			}
			if !tc.expectError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}

// 10. 슬라이스 관련 함수 테스트
func TestSum(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		expected int
	}{
		{"positive numbers", []int{1, 2, 3, 4, 5}, 15},
		{"negative numbers", []int{-1, -2, -3}, -6},
		{"mixed numbers", []int{-5, 5, -3, 3}, 0},
		{"single number", []int{42}, 42},
		{"empty slice", []int{}, 0},
		{"zeros", []int{0, 0, 0}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Sum(tt.numbers)
			if result != tt.expected {
				t.Errorf("Sum(%v) = %d; expected %d",
					tt.numbers, result, tt.expected)
			}
		})
	}
}

func TestMax(t *testing.T) {
	// 정상 케이스
	result, err := Max([]int{3, 1, 4, 1, 5, 9})
	if err != nil {
		t.Errorf("Max with valid slice returned error: %v", err)
	}
	if result != 9 {
		t.Errorf("Max([3,1,4,1,5,9]) = %d; expected 9", result)
	}

	// 빈 슬라이스 에러 케이스
	_, err = Max([]int{})
	if err == nil {
		t.Error("Max with empty slice should return error")
	}

	// 추가 테스트 케이스들
	tests := []struct {
		name     string
		numbers  []int
		expected int
		hasError bool
	}{
		{"single element", []int{42}, 42, false},
		{"negative numbers", []int{-5, -1, -10}, -1, false},
		{"all same", []int{3, 3, 3}, 3, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Max(tt.numbers)
			if tt.hasError {
				if err == nil {
					t.Error("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if result != tt.expected {
					t.Errorf("Max(%v) = %d; expected %d",
						tt.numbers, result, tt.expected)
				}
			}
		})
	}
}

// 11. 필터 함수 테스트
func TestFilter(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// 짝수 필터
	evens := Filter(numbers, func(n int) bool { return n%2 == 0 })
	expectedEvens := []int{2, 4, 6, 8, 10}

	if !reflect.DeepEqual(evens, expectedEvens) {
		t.Errorf("Filter evens = %v; expected %v", evens, expectedEvens)
	}

	// 5보다 큰 수 필터
	greaterThan5 := Filter(numbers, func(n int) bool { return n > 5 })
	expectedGreater := []int{6, 7, 8, 9, 10}

	if !reflect.DeepEqual(greaterThan5, expectedGreater) {
		t.Errorf("Filter >5 = %v; expected %v", greaterThan5, expectedGreater)
	}

	// 빈 결과 필터
	greaterThan100 := Filter(numbers, func(n int) bool { return n > 100 })
	if len(greaterThan100) != 0 {
		t.Errorf("Filter >100 should return empty slice, got %v", greaterThan100)
	}
}

// 12. 맵 반환 함수 테스트
func TestCountCharacters(t *testing.T) {
	result := CountCharacters("hello")
	expected := map[rune]int{
		'h': 1,
		'e': 1,
		'l': 2,
		'o': 1,
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("CountCharacters('hello') = %v; expected %v", result, expected)
	}

	// 빈 문자열 테스트
	emptyResult := CountCharacters("")
	if len(emptyResult) != 0 {
		t.Errorf("CountCharacters('') should return empty map, got %v", emptyResult)
	}

	// 한글 문자열 테스트
	koreanResult := CountCharacters("안녕")
	expectedKorean := map[rune]int{
		'안': 1,
		'녕': 1,
	}

	if !reflect.DeepEqual(koreanResult, expectedKorean) {
		t.Errorf("CountCharacters('안녕') = %v; expected %v",
			koreanResult, expectedKorean)
	}
}

// 13. 설정 읽기 함수 테스트
func TestReadConfig(t *testing.T) {
	// 정상 케이스
	config, err := ReadConfig("app.conf")
	if err != nil {
		t.Errorf("ReadConfig with valid filename returned error: %v", err)
	}
	if config == nil {
		t.Fatal("ReadConfig should return config map")
	}

	expectedHost := "localhost"
	if config["host"] != expectedHost {
		t.Errorf("Config host = %q; expected %q", config["host"], expectedHost)
	}

	// 에러 케이스들
	errorCases := []struct {
		filename    string
		expectError bool
	}{
		{"", true},
		{"invalid.conf", true},
		{"valid.conf", false},
	}

	for _, tc := range errorCases {
		t.Run(tc.filename, func(t *testing.T) {
			_, err := ReadConfig(tc.filename)
			if tc.expectError && err == nil {
				t.Error("Expected error but got none")
			}
			if !tc.expectError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}

// 14. 헬퍼 함수 사용 예제
func TestFactorial(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
		{5, 120},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("factorial_%d", tt.input), func(t *testing.T) {
			result := Factorial(tt.input)
			if result != tt.expected {
				t.Errorf("Factorial(%d) = %d; expected %d",
					tt.input, result, tt.expected)
			}
		})
	}
}

// 15. Skip과 Helper 사용 예제
func TestSkipExample(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	// 시간이 오래 걸리는 테스트
	t.Log("Running long test...")
}

// 헬퍼 함수 정의
func assertEqual(t *testing.T, actual, expected interface{}) {
	t.Helper() // 이 함수를 헬퍼로 표시
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v; expected %v", actual, expected)
	}
}

func TestWithHelper(t *testing.T) {
	result := Add(2, 3)
	assertEqual(t, result, 5) // 헬퍼 함수 사용
}

/*
Go 테스팅 모범 사례:

1. 테스트 함수 명명:
   - TestXxx 형태로 명명
   - 테스트 대상을 명확히 표현
   - 서브테스트로 세분화

2. 테이블 기반 테스트:
   - 여러 입력/출력 조합을 효율적으로 테스트
   - 테스트 케이스 추가가 쉬움
   - 가독성과 유지보수성 향상

3. 에러 처리 테스트:
   - 정상 케이스와 에러 케이스 모두 테스트
   - 예상되는 에러 타입과 메시지 검증
   - nil 체크 포함

4. 테스트 격리:
   - 각 테스트는 독립적으로 실행
   - 전역 상태 변경 지양
   - 필요시 setup/teardown 사용

5. 어서션 패턴:
   - 명확한 에러 메시지 제공
   - 예상값과 실제값 모두 표시
   - reflect.DeepEqual로 복합 타입 비교

6. 테스트 커버리지:
   - 모든 코드 경로 테스트
   - 경계값 테스트 포함
   - 에러 케이스도 빠뜨리지 말 것

7. 테스트 성능:
   - 빠른 테스트 작성
   - 외부 의존성 최소화
   - Mock/Stub 활용 고려
*/
