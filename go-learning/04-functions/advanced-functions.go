// advanced-functions.go - Go의 고급 함수 기법
package main

import (
	"fmt"
	"reflect"
	"runtime"
	"sort"
	"time"
)

func main() {
	fmt.Println("=== Go 고급 함수 기법 ===")

	// 1. 고차 함수 (Higher-Order Functions)
	fmt.Println("\n1. 고차 함수:")

	// 함수를 반환하는 함수
	multiplier := createMultiplier(5)
	fmt.Printf("5 x 3 = %d\n", multiplier(3))

	// 함수를 매개변수로 받는 함수
	numbers := []int{1, 2, 3, 4, 5}
	doubled := mapFunction(numbers, func(x int) int { return x * 2 })
	fmt.Printf("원본: %v, 2배: %v\n", numbers, doubled)

	// 2. 함수 조합 (Function Composition)
	fmt.Println("\n2. 함수 조합:")

	add5 := func(x int) int { return x + 5 }
	multiply3 := func(x int) int { return x * 3 }

	// 함수 조합: (x + 5) * 3
	composed := compose(multiply3, add5)
	result := composed(2) // (2 + 5) * 3 = 21
	fmt.Printf("(2 + 5) * 3 = %d\n", result)

	// 3. 커링 (Currying)
	fmt.Println("\n3. 커링:")

	// 다중 매개변수 함수를 단일 매개변수 함수들의 연쇄로 변환
	curriedAdd := curryAdd(10)
	fmt.Printf("10 + 5 = %d\n", curriedAdd(5))

	// 부분 적용 (Partial Application)
	addTo100 := curryAdd(100)
	fmt.Printf("100 + 25 = %d\n", addTo100(25))

	// 4. 메모이제이션 (Memoization)
	fmt.Println("\n4. 메모이제이션:")

	// 느린 함수 (피보나치)
	fmt.Println("메모이제이션 없이:")
	start := time.Now()
	result1 := slowFibonacci(35)
	duration1 := time.Since(start)
	fmt.Printf("fibonacci(35) = %d, 시간: %v\n", result1, duration1)

	// 메모이제이션된 함수
	fmt.Println("메모이제이션 사용:")
	memoFib := memoize(slowFibonacci)
	start = time.Now()
	result2 := memoFib(35)
	duration2 := time.Since(start)
	fmt.Printf("fibonacci(35) = %d, 시간: %v\n", result2, duration2)

	// 두 번째 호출 (캐시에서)
	start = time.Now()
	result3 := memoFib(35)
	duration3 := time.Since(start)
	fmt.Printf("fibonacci(35) = %d, 시간: %v (캐시에서)\n", result3, duration3)

	// 5. 함수 데코레이터 패턴
	fmt.Println("\n5. 함수 데코레이터 패턴:")

	// 실행 시간 측정 데코레이터
	timedFunction := timeDecorator(expensiveOperation)
	timedFunction("데이터 처리")

	// 로깅 데코레이터
	loggedFunction := loggingDecorator(businessLogic)
	loggedFunction("중요한 작업")

	// 6. 제네릭 함수 (Go 1.18+)
	fmt.Println("\n6. 제네릭 함수:")

	// 정수 슬라이스
	intSlice := []int{3, 1, 4, 1, 5, 9}
	intMax := findMax(intSlice)
	fmt.Printf("정수 최대값: %d\n", intMax)

	// 문자열 슬라이스
	stringSlice := []string{"apple", "banana", "cherry"}
	stringMax := findMax(stringSlice)
	fmt.Printf("문자열 최대값: %s\n", stringMax)

	// 7. 함수 파이프라인
	fmt.Println("\n7. 함수 파이프라인:")

	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	pipeline := NewPipeline(data).
		Filter(func(x int) bool { return x%2 == 0 }). // 짝수만
		Map(func(x int) int { return x * x }).        // 제곱
		Filter(func(x int) bool { return x > 10 }).   // 10보다 큰 것만
		Sort()

	fmt.Printf("파이프라인 결과: %v\n", pipeline.Result())

	// 8. 이벤트 핸들링 시스템
	fmt.Println("\n8. 이벤트 핸들링 시스템:")

	eventBus := NewEventBus()

	// 이벤트 리스너 등록
	eventBus.Subscribe("user_login", func(data interface{}) {
		fmt.Printf("사용자 로그인: %v\n", data)
	})

	eventBus.Subscribe("user_login", func(data interface{}) {
		fmt.Printf("로그인 로그 저장: %v\n", data)
	})

	// 이벤트 발생
	eventBus.Emit("user_login", "김철수")

	// 9. 함수 체이닝
	fmt.Println("\n9. 함수 체이닝:")

	calculator := NewCalculator().
		Add(10).
		Multiply(2).
		Subtract(5).
		Divide(3)

	fmt.Printf("계산 결과: %.2f\n", calculator.Result())

	// 10. 동적 함수 생성
	fmt.Println("\n10. 동적 함수 생성:")

	// 다양한 검증 함수 생성
	isPositive := createValidator(func(x int) bool { return x > 0 }, "양수가 아님")
	isEven := createValidator(func(x int) bool { return x%2 == 0 }, "짝수가 아님")
	isLessThan100 := createValidator(func(x int) bool { return x < 100 }, "100 이상임")

	testValue := 42
	fmt.Printf("값 %d 검증:\n", testValue)
	fmt.Println(isPositive(testValue))
	fmt.Println(isEven(testValue))
	fmt.Println(isLessThan100(testValue))

	// 11. 함수 리플렉션
	fmt.Println("\n11. 함수 리플렉션:")

	demonstrateFunctionReflection()
}

// 1. 고차 함수들
func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func mapFunction(slice []int, fn func(int) int) []int {
	result := make([]int, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

// 2. 함수 조합
func compose(f, g func(int) int) func(int) int {
	return func(x int) int {
		return f(g(x))
	}
}

// 3. 커링
func curryAdd(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}

// 4. 메모이제이션
func slowFibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return slowFibonacci(n-1) + slowFibonacci(n-2)
}

func memoize(fn func(int) int) func(int) int {
	cache := make(map[int]int)
	return func(n int) int {
		if result, exists := cache[n]; exists {
			return result
		}
		result := fn(n)
		cache[n] = result
		return result
	}
}

// 5. 데코레이터 패턴
func timeDecorator(fn func(string)) func(string) {
	return func(arg string) {
		start := time.Now()
		fn(arg)
		duration := time.Since(start)
		fmt.Printf("실행 시간: %v\n", duration)
	}
}

func loggingDecorator(fn func(string)) func(string) {
	return func(arg string) {
		fmt.Printf("함수 실행 시작: %s\n", getFunctionName(fn))
		fn(arg)
		fmt.Printf("함수 실행 완료: %s\n", getFunctionName(fn))
	}
}

func getFunctionName(fn interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
}

func expensiveOperation(data string) {
	fmt.Printf("비용이 많이 드는 작업 수행: %s\n", data)
	time.Sleep(100 * time.Millisecond)
}

func businessLogic(task string) {
	fmt.Printf("비즈니스 로직 실행: %s\n", task)
}

// 6. 제네릭 함수
func findMax[T comparable](slice []T) T {
	if len(slice) == 0 {
		var zero T
		return zero
	}

	max := slice[0]
	for _, v := range slice[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

// 7. 함수 파이프라인
type Pipeline struct {
	data []int
}

func NewPipeline(data []int) *Pipeline {
	return &Pipeline{data: data}
}

func (p *Pipeline) Filter(predicate func(int) bool) *Pipeline {
	var filtered []int
	for _, v := range p.data {
		if predicate(v) {
			filtered = append(filtered, v)
		}
	}
	p.data = filtered
	return p
}

func (p *Pipeline) Map(transform func(int) int) *Pipeline {
	for i, v := range p.data {
		p.data[i] = transform(v)
	}
	return p
}

func (p *Pipeline) Sort() *Pipeline {
	sort.Ints(p.data)
	return p
}

func (p *Pipeline) Result() []int {
	return p.data
}

// 8. 이벤트 버스
type EventBus struct {
	listeners map[string][]func(interface{})
}

func NewEventBus() *EventBus {
	return &EventBus{
		listeners: make(map[string][]func(interface{})),
	}
}

func (eb *EventBus) Subscribe(event string, handler func(interface{})) {
	eb.listeners[event] = append(eb.listeners[event], handler)
}

func (eb *EventBus) Emit(event string, data interface{}) {
	if handlers, exists := eb.listeners[event]; exists {
		for _, handler := range handlers {
			handler(data)
		}
	}
}

// 9. 함수 체이닝
type Calculator struct {
	value float64
}

func NewCalculator() *Calculator {
	return &Calculator{value: 0}
}

func (c *Calculator) Add(x float64) *Calculator {
	c.value += x
	return c
}

func (c *Calculator) Multiply(x float64) *Calculator {
	c.value *= x
	return c
}

func (c *Calculator) Subtract(x float64) *Calculator {
	c.value -= x
	return c
}

func (c *Calculator) Divide(x float64) *Calculator {
	if x != 0 {
		c.value /= x
	}
	return c
}

func (c *Calculator) Result() float64 {
	return c.value
}

// 10. 동적 함수 생성
func createValidator(predicate func(int) bool, errorMsg string) func(int) string {
	return func(value int) string {
		if predicate(value) {
			return fmt.Sprintf("✓ 검증 통과: %d", value)
		}
		return fmt.Sprintf("✗ 검증 실패: %d (%s)", value, errorMsg)
	}
}

// 11. 함수 리플렉션
func demonstrateFunctionReflection() {
	fn := func(a int, b string) (int, error) {
		return len(b) + a, nil
	}

	fnType := reflect.TypeOf(fn)
	fmt.Printf("함수 타입: %v\n", fnType)
	fmt.Printf("매개변수 개수: %d\n", fnType.NumIn())
	fmt.Printf("반환값 개수: %d\n", fnType.NumOut())

	for i := 0; i < fnType.NumIn(); i++ {
		fmt.Printf("매개변수 %d 타입: %v\n", i, fnType.In(i))
	}

	for i := 0; i < fnType.NumOut(); i++ {
		fmt.Printf("반환값 %d 타입: %v\n", i, fnType.Out(i))
	}

	// 동적 함수 호출
	fnValue := reflect.ValueOf(fn)
	args := []reflect.Value{
		reflect.ValueOf(10),
		reflect.ValueOf("Hello"),
	}
	results := fnValue.Call(args)
	fmt.Printf("동적 호출 결과: %v, %v\n", results[0].Interface(), results[1].Interface())
}

/*
Go 고급 함수 기법:
1. 고차 함수: 함수를 매개변수로 받거나 함수를 반환
2. 함수 조합: 여러 함수를 결합하여 새로운 함수 생성
3. 커링: 다중 매개변수 함수를 단일 매개변수 함수들의 연쇄로 변환
4. 메모이제이션: 함수 결과를 캐싱하여 성능 향상
5. 데코레이터 패턴: 함수의 기능을 확장
6. 제네릭 함수: 타입에 독립적인 함수 작성
7. 함수 파이프라인: 함수들을 연결하여 데이터 처리
8. 이벤트 핸들링: 함수를 이벤트 리스너로 사용
9. 함수 체이닝: 메서드 체이닝 패턴 구현
10. 동적 함수 생성: 런타임에 함수 생성
11. 함수 리플렉션: 런타임에 함수 정보 분석

활용 팁:
1. 함수형 프로그래밍 기법으로 코드 재사용성 향상
2. 클로저를 활용한 상태 캡슐화
3. 고차 함수로 추상화 레벨 향상
4. 메모이제이션으로 비용이 큰 계산 최적화
5. 데코레이터 패턴으로 관심사 분리
6. 파이프라인으로 데이터 처리 과정 명확화
*/
