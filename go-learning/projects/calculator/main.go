// main.go - 계산기 프로젝트
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Calculator 구조체
type Calculator struct {
	history []string
	memory  float64
}

// NewCalculator 생성자
func NewCalculator() *Calculator {
	return &Calculator{
		history: make([]string, 0),
		memory:  0.0,
	}
}

// 기본 연산 메서드들
func (c *Calculator) Add(a, b float64) float64 {
	result := a + b
	c.addHistory(fmt.Sprintf("%.2f + %.2f = %.2f", a, b, result))
	return result
}

func (c *Calculator) Subtract(a, b float64) float64 {
	result := a - b
	c.addHistory(fmt.Sprintf("%.2f - %.2f = %.2f", a, b, result))
	return result
}

func (c *Calculator) Multiply(a, b float64) float64 {
	result := a * b
	c.addHistory(fmt.Sprintf("%.2f × %.2f = %.2f", a, b, result))
	return result
}

func (c *Calculator) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("0으로 나눌 수 없습니다")
	}
	result := a / b
	c.addHistory(fmt.Sprintf("%.2f ÷ %.2f = %.2f", a, b, result))
	return result, nil
}

// 고급 연산 메서드들
func (c *Calculator) Power(base, exponent float64) float64 {
	result := 1.0
	if exponent >= 0 {
		for i := 0; i < int(exponent); i++ {
			result *= base
		}
	} else {
		for i := 0; i < int(-exponent); i++ {
			result /= base
		}
	}
	c.addHistory(fmt.Sprintf("%.2f ^ %.2f = %.2f", base, exponent, result))
	return result
}

func (c *Calculator) Sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0, fmt.Errorf("음수의 제곱근을 구할 수 없습니다")
	}

	// 뉴턴-랩슨 방법으로 제곱근 계산
	if n == 0 {
		return 0, nil
	}

	x := n
	for {
		root := 0.5 * (x + n/x)
		if abs(root-x) < 0.000001 {
			c.addHistory(fmt.Sprintf("√%.2f = %.6f", n, root))
			return root, nil
		}
		x = root
	}
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

// 메모리 관련 메서드들
func (c *Calculator) MemoryStore(value float64) {
	c.memory = value
	c.addHistory(fmt.Sprintf("메모리에 %.2f 저장", value))
}

func (c *Calculator) MemoryRecall() float64 {
	c.addHistory(fmt.Sprintf("메모리에서 %.2f 불러옴", c.memory))
	return c.memory
}

func (c *Calculator) MemoryClear() {
	c.memory = 0.0
	c.addHistory("메모리 지움")
}

func (c *Calculator) MemoryAdd(value float64) {
	c.memory += value
	c.addHistory(fmt.Sprintf("메모리에 %.2f 더함 (결과: %.2f)", value, c.memory))
}

// 히스토리 관련 메서드들
func (c *Calculator) addHistory(operation string) {
	c.history = append(c.history, operation)
	// 히스토리가 너무 길어지지 않도록 제한
	if len(c.history) > 100 {
		c.history = c.history[1:]
	}
}

func (c *Calculator) GetHistory() []string {
	return c.history[:]
}

func (c *Calculator) ClearHistory() {
	c.history = make([]string, 0)
	fmt.Println("히스토리가 지워졌습니다.")
}

func main() {
	calculator := NewCalculator()
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("=== Go 계산기 ===")
	fmt.Println("명령어 목록:")
	fmt.Println("  기본 연산: +, -, *, /")
	fmt.Println("  고급 연산: pow (거듭제곱), sqrt (제곱근)")
	fmt.Println("  메모리: ms (저장), mr (불러오기), mc (지우기), m+ (더하기)")
	fmt.Println("  기타: history (히스토리), clear (히스토리 지우기), help (도움말), quit (종료)")
	fmt.Println("  사용 예: 5 + 3, 2 pow 3, sqrt 16")
	fmt.Println()

	for {
		fmt.Print("계산기> ")

		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())

		if input == "" {
			continue
		}

		// 명령어 처리
		switch input {
		case "quit", "exit", "q":
			fmt.Println("계산기를 종료합니다.")
			return
		case "help", "h":
			showHelp()
			continue
		case "history":
			showHistory(calculator)
			continue
		case "clear":
			calculator.ClearHistory()
			continue
		case "mc":
			calculator.MemoryClear()
			continue
		case "mr":
			result := calculator.MemoryRecall()
			fmt.Printf("메모리 값: %.6f\n", result)
			continue
		}

		// 표현식 처리
		result, err := processExpression(calculator, input)
		if err != nil {
			fmt.Printf("오류: %v\n", err)
		} else {
			fmt.Printf("결과: %.6f\n", result)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("입력 읽기 오류: %v\n", err)
	}
}

func processExpression(calc *Calculator, input string) (float64, error) {
	tokens := strings.Fields(input)

	if len(tokens) == 0 {
		return 0, fmt.Errorf("빈 표현식")
	}

	// 단일 연산 처리
	switch {
	case len(tokens) == 2 && tokens[0] == "sqrt":
		// 제곱근 계산
		n, err := strconv.ParseFloat(tokens[1], 64)
		if err != nil {
			return 0, fmt.Errorf("잘못된 숫자: %s", tokens[1])
		}
		return calc.Sqrt(n)

	case len(tokens) == 2 && tokens[0] == "ms":
		// 메모리 저장
		value, err := strconv.ParseFloat(tokens[1], 64)
		if err != nil {
			return 0, fmt.Errorf("잘못된 숫자: %s", tokens[1])
		}
		calc.MemoryStore(value)
		return value, nil

	case len(tokens) == 2 && tokens[0] == "m+":
		// 메모리 더하기
		value, err := strconv.ParseFloat(tokens[1], 64)
		if err != nil {
			return 0, fmt.Errorf("잘못된 숫자: %s", tokens[1])
		}
		calc.MemoryAdd(value)
		return calc.memory, nil

	case len(tokens) == 3:
		// 이항 연산 처리
		return processBinaryOperation(calc, tokens)

	default:
		return 0, fmt.Errorf("올바르지 않은 표현식 형식")
	}
}

func processBinaryOperation(calc *Calculator, tokens []string) (float64, error) {
	// 첫 번째 피연산자
	a, err := parseOperand(calc, tokens[0])
	if err != nil {
		return 0, err
	}

	operator := tokens[1]

	// 두 번째 피연산자
	b, err := parseOperand(calc, tokens[2])
	if err != nil {
		return 0, err
	}

	// 연산 수행
	switch operator {
	case "+", "add":
		return calc.Add(a, b), nil
	case "-", "sub":
		return calc.Subtract(a, b), nil
	case "*", "mul":
		return calc.Multiply(a, b), nil
	case "/", "div":
		return calc.Divide(a, b)
	case "pow", "^", "**":
		return calc.Power(a, b), nil
	default:
		return 0, fmt.Errorf("알 수 없는 연산자: %s", operator)
	}
}

func parseOperand(calc *Calculator, token string) (float64, error) {
	// 메모리 값 참조
	if token == "mr" || token == "mem" {
		return calc.MemoryRecall(), nil
	}

	// 숫자 파싱
	return strconv.ParseFloat(token, 64)
}

func showHelp() {
	fmt.Println("\n=== 계산기 도움말 ===")
	fmt.Println("기본 연산:")
	fmt.Println("  5 + 3     : 덧셈")
	fmt.Println("  10 - 4    : 뺄셈")
	fmt.Println("  6 * 7     : 곱셈")
	fmt.Println("  15 / 3    : 나눗셈")
	fmt.Println()
	fmt.Println("고급 연산:")
	fmt.Println("  2 pow 3   : 거듭제곱 (2³)")
	fmt.Println("  sqrt 16   : 제곱근 (√16)")
	fmt.Println()
	fmt.Println("메모리 기능:")
	fmt.Println("  ms 5      : 메모리에 5 저장")
	fmt.Println("  mr        : 메모리 값 불러오기")
	fmt.Println("  mc        : 메모리 지우기")
	fmt.Println("  m+ 3      : 메모리 값에 3 더하기")
	fmt.Println("  5 + mr    : 메모리 값을 연산에 사용")
	fmt.Println()
	fmt.Println("기타 명령:")
	fmt.Println("  history   : 계산 히스토리 보기")
	fmt.Println("  clear     : 히스토리 지우기")
	fmt.Println("  help      : 이 도움말")
	fmt.Println("  quit      : 프로그램 종료")
	fmt.Println()
}

func showHistory(calc *Calculator) {
	history := calc.GetHistory()
	if len(history) == 0 {
		fmt.Println("계산 히스토리가 없습니다.")
		return
	}

	fmt.Println("\n=== 계산 히스토리 ===")
	for i, operation := range history {
		fmt.Printf("%3d. %s\n", i+1, operation)
	}
	fmt.Printf("총 %d개의 연산 기록\n\n", len(history))
}
