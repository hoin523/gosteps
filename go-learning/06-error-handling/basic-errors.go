// basic-errors.go - Go의 기본 에러 처리
package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	fmt.Println("=== Go 에러 처리 ===")

	// 1. 기본 에러 처리
	fmt.Println("\n1. 기본 에러 처리:")

	result, err := divide(10, 2)
	if err != nil {
		fmt.Printf("에러 발생: %v\n", err)
	} else {
		fmt.Printf("10 ÷ 2 = %.2f\n", result)
	}

	result, err = divide(10, 0)
	if err != nil {
		fmt.Printf("에러 발생: %v\n", err)
	} else {
		fmt.Printf("결과: %.2f\n", result)
	}

	// 2. 에러 생성 방법들
	fmt.Println("\n2. 에러 생성 방법:")

	// errors.New() 사용
	err1 := errors.New("간단한 에러 메시지")
	fmt.Printf("errors.New(): %v\n", err1)

	// fmt.Errorf() 사용 (서식화된 에러)
	userID := 123
	err2 := fmt.Errorf("사용자 ID %d를 찾을 수 없습니다", userID)
	fmt.Printf("fmt.Errorf(): %v\n", err2)

	// 3. 에러 체이닝 (Go 1.13+)
	fmt.Println("\n3. 에러 체이닝:")

	originalErr := errors.New("원본 에러")
	wrappedErr := fmt.Errorf("추가 정보: %w", originalErr)
	fmt.Printf("체인된 에러: %v\n", wrappedErr)

	// errors.Unwrap으로 원본 에러 추출
	unwrapped := errors.Unwrap(wrappedErr)
	fmt.Printf("언래핑된 에러: %v\n", unwrapped)

	// errors.Is로 특정 에러 확인
	if errors.Is(wrappedErr, originalErr) {
		fmt.Println("wrappedErr는 originalErr를 포함합니다")
	}

	// 4. 사용자 정의 에러 타입
	fmt.Println("\n4. 사용자 정의 에러 타입:")

	account := BankAccount{Balance: 1000}

	err = account.Withdraw(500)
	if err != nil {
		fmt.Printf("출금 에러: %v\n", err)
	} else {
		fmt.Printf("출금 성공, 잔액: %d\n", account.Balance)
	}

	err = account.Withdraw(800)
	if err != nil {
		// 에러 타입 확인
		var insufficientErr *InsufficientFundsError
		if errors.As(err, &insufficientErr) {
			fmt.Printf("잔액 부족: 현재 잔액 %d, 요청 금액 %d\n",
				insufficientErr.Balance, insufficientErr.Amount)
		}
	}

	// 5. 다중 반환값을 활용한 에러 처리
	fmt.Println("\n5. 다중 반환값을 활용한 에러 처리:")

	data, err := readFile("example.txt")
	if err != nil {
		fmt.Printf("파일 읽기 에러: %v\n", err)
	} else {
		fmt.Printf("파일 내용: %s\n", data)
	}

	// 6. 에러 처리 패턴들
	fmt.Println("\n6. 에러 처리 패턴들:")

	// 패턴 1: 빠른 실패 (Early Return)
	if err := processData("invalid_data"); err != nil {
		fmt.Printf("데이터 처리 실패: %v\n", err)
		// return // 실제 함수에서는 여기서 반환
	}

	// 패턴 2: 에러 무시 (언더스코어 사용)
	result, _ = divide(20, 4) // 에러 무시
	fmt.Printf("에러 무시하고 계산: %.2f\n", result)

	// 패턴 3: 에러 누적
	errors := validateUser("", "invalid-email", -1)
	if len(errors) > 0 {
		fmt.Println("사용자 검증 에러들:")
		for _, err := range errors {
			fmt.Printf("  - %v\n", err)
		}
	}

	// 7. defer와 함께 사용하는 에러 처리
	fmt.Println("\n7. defer와 함께 사용하는 에러 처리:")

	err = writeToFile("test.txt", "Hello, Go!")
	if err != nil {
		fmt.Printf("파일 쓰기 에러: %v\n", err)
	} else {
		fmt.Println("파일 쓰기 성공")
	}

	// 8. panic과 recover
	fmt.Println("\n8. panic과 recover:")

	fmt.Println("안전한 함수 호출:")
	safeCall(func() {
		fmt.Println("정상 함수 실행")
	})

	fmt.Println("패닉이 발생하는 함수 호출:")
	safeCall(func() {
		panic("의도적인 패닉!")
	})

	fmt.Println("프로그램 계속 실행됨")

	// 9. 에러 처리 모범 사례
	fmt.Println("\n9. 에러 처리 모범 사례:")

	// 구체적인 에러 메시지
	user, err := getUserByID(999)
	if err != nil {
		fmt.Printf("사용자 조회 실패: %v\n", err)
	} else {
		fmt.Printf("사용자: %+v\n", user)
	}

	// 에러 컨텍스트 추가
	err = performOperation("critical_task")
	if err != nil {
		fmt.Printf("작업 수행 실패: %v\n", err)
	}

	// 10. 실용적인 에러 처리 예제
	fmt.Println("\n10. 실용적인 에러 처리 예제:")

	// HTTP 요청 시뮬레이션
	response, err := makeHTTPRequest("https://api.example.com/data")
	if err != nil {
		handleHTTPError(err)
	} else {
		fmt.Printf("HTTP 응답: %s\n", response)
	}

	// 설정 파일 로딩
	config, err := loadConfig("app.config")
	if err != nil {
		fmt.Printf("설정 로딩 실패: %v\n", err)
		// 기본 설정 사용
		config = getDefaultConfig()
	}
	fmt.Printf("설정: %+v\n", config)

	// 11. 에러 로깅과 모니터링
	fmt.Println("\n11. 에러 로깅:")

	logger := &SimpleLogger{}

	err = performTaskWithLogging("데이터베이스 연결", logger)
	if err != nil {
		fmt.Printf("작업 실패: %v\n", err)
	}

	// 12. 에러 처리 유틸리티
	fmt.Println("\n12. 에러 처리 유틸리티:")

	// 재시도 로직
	err = retryOperation(func() error {
		return simulateUnstableOperation()
	}, 3, time.Second)

	if err != nil {
		fmt.Printf("재시도 후에도 실패: %v\n", err)
	} else {
		fmt.Println("재시도 성공!")
	}
}

// 1. 기본 에러 반환 함수
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("0으로 나눌 수 없습니다")
	}
	return a / b, nil
}

// 4. 사용자 정의 에러 타입
type InsufficientFundsError struct {
	Balance int
	Amount  int
}

func (e *InsufficientFundsError) Error() string {
	return fmt.Sprintf("잔액 부족: 현재 잔액 %d, 요청 금액 %d", e.Balance, e.Amount)
}

type BankAccount struct {
	Balance int
}

func (ba *BankAccount) Withdraw(amount int) error {
	if ba.Balance < amount {
		return &InsufficientFundsError{
			Balance: ba.Balance,
			Amount:  amount,
		}
	}
	ba.Balance -= amount
	return nil
}

// 5. 파일 읽기 시뮬레이션
func readFile(filename string) (string, error) {
	if filename == "example.txt" {
		return "", fmt.Errorf("파일을 찾을 수 없습니다: %s", filename)
	}
	return "파일 내용", nil
}

// 6. 데이터 처리 함수
func processData(data string) error {
	if data == "invalid_data" {
		return fmt.Errorf("유효하지 않은 데이터: %s", data)
	}
	fmt.Printf("데이터 처리 성공: %s\n", data)
	return nil
}

// 사용자 검증 (다중 에러)
func validateUser(name, email string, age int) []error {
	var errors []error

	if name == "" {
		errors = append(errors, fmt.Errorf("이름은 필수입니다"))
	}

	if email == "" || !contains(email, "@") {
		errors = append(errors, fmt.Errorf("유효한 이메일이 필요합니다"))
	}

	if age < 0 || age > 150 {
		errors = append(errors, fmt.Errorf("나이는 0-150 사이여야 합니다"))
	}

	return errors
}

func contains(s, substr string) bool {
	return len(s) > 0 && len(substr) > 0 &&
		len(s) >= len(substr) &&
		s[0:len(substr)] == substr ||
		(len(s) > len(substr) && contains(s[1:], substr))
}

// 7. defer를 사용한 파일 처리
func writeToFile(filename, content string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("파일 생성 실패: %w", err)
	}

	// defer로 파일 닫기 보장
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			fmt.Printf("파일 닫기 에러: %v\n", closeErr)
		}
	}()

	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("파일 쓰기 실패: %w", err)
	}

	return nil
}

// 8. panic과 recover를 사용한 안전한 함수 호출
func safeCall(fn func()) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("패닉 복구: %v\n", r)
		}
	}()

	fn()
}

// 9. 사용자 조회 함수
type User struct {
	ID   int
	Name string
}

func getUserByID(id int) (*User, error) {
	if id <= 0 {
		return nil, fmt.Errorf("유효하지 않은 사용자 ID: %d", id)
	}

	if id == 999 {
		return nil, fmt.Errorf("사용자를 찾을 수 없습니다 (ID: %d)", id)
	}

	return &User{ID: id, Name: "테스트 사용자"}, nil
}

func performOperation(taskName string) error {
	// 작업 수행 시뮬레이션
	if taskName == "critical_task" {
		err := errors.New("데이터베이스 연결 실패")
		return fmt.Errorf("작업 '%s' 수행 중 에러 발생: %w", taskName, err)
	}
	return nil
}

// 10. HTTP 요청 시뮬레이션
type HTTPError struct {
	StatusCode int
	Message    string
}

func (e HTTPError) Error() string {
	return fmt.Sprintf("HTTP %d: %s", e.StatusCode, e.Message)
}

func makeHTTPRequest(url string) (string, error) {
	// HTTP 요청 시뮬레이션
	if url == "https://api.example.com/data" {
		return "", HTTPError{
			StatusCode: 404,
			Message:    "Not Found",
		}
	}
	return "응답 데이터", nil
}

func handleHTTPError(err error) {
	var httpErr HTTPError
	if errors.As(err, &httpErr) {
		switch httpErr.StatusCode {
		case 404:
			fmt.Println("리소스를 찾을 수 없습니다")
		case 500:
			fmt.Println("서버 내부 에러입니다")
		default:
			fmt.Printf("HTTP 에러: %v\n", httpErr)
		}
	} else {
		fmt.Printf("네트워크 에러: %v\n", err)
	}
}

// 설정 관련
type Config struct {
	Host string
	Port int
}

func loadConfig(filename string) (*Config, error) {
	if filename == "app.config" {
		return nil, errors.New("설정 파일을 읽을 수 없습니다")
	}
	return &Config{Host: "localhost", Port: 8080}, nil
}

func getDefaultConfig() *Config {
	return &Config{Host: "127.0.0.1", Port: 3000}
}

// 11. 간단한 로거
type SimpleLogger struct{}

func (l *SimpleLogger) LogError(err error) {
	fmt.Printf("[ERROR] %s: %v\n", time.Now().Format("15:04:05"), err)
}

func performTaskWithLogging(taskName string, logger *SimpleLogger) error {
	fmt.Printf("작업 시작: %s\n", taskName)

	// 작업 시뮬레이션
	if taskName == "데이터베이스 연결" {
		err := errors.New("연결 시간 초과")
		logger.LogError(fmt.Errorf("작업 '%s' 실패: %w", taskName, err))
		return err
	}

	fmt.Printf("작업 완료: %s\n", taskName)
	return nil
}

// 12. 재시도 유틸리티
func retryOperation(operation func() error, maxRetries int, delay time.Duration) error {
	var lastErr error

	for i := 0; i < maxRetries; i++ {
		err := operation()
		if err == nil {
			return nil // 성공
		}

		lastErr = err
		fmt.Printf("재시도 %d/%d 실패: %v\n", i+1, maxRetries, err)

		if i < maxRetries-1 {
			time.Sleep(delay)
		}
	}

	return fmt.Errorf("최대 재시도 횟수 초과 (%d회): %w", maxRetries, lastErr)
}

var attemptCount = 0

func simulateUnstableOperation() error {
	attemptCount++
	if attemptCount < 3 {
		return fmt.Errorf("시도 %d 실패", attemptCount)
	}
	attemptCount = 0 // 리셋
	return nil
}

/*
Go 에러 처리 철학:
1. 에러는 값이다 - error 타입은 인터페이스
2. 명시적 에러 처리 - 예외가 아닌 반환값으로 처리
3. 조기 반환 - 에러 발생 시 빠르게 반환
4. 에러 체이닝 - 에러에 컨텍스트 추가

에러 처리 모범 사례:
1. 에러를 무시하지 마라 (_, err := func(); err != nil 체크)
2. 구체적이고 유용한 에러 메시지 작성
3. 에러 타입을 활용한 세분화된 처리
4. defer를 활용한 리소스 정리
5. panic은 복구 불가능한 상황에서만 사용
6. 에러 체이닝으로 디버깅 정보 보강
7. 로깅과 모니터링으로 에러 추적

주요 함수들:
- errors.New(): 간단한 에러 생성
- fmt.Errorf(): 포맷된 에러 생성 (%w로 래핑)
- errors.Is(): 에러 체인에서 특정 에러 확인
- errors.As(): 에러 타입 변환
- errors.Unwrap(): 래핑된 에러 추출
*/
