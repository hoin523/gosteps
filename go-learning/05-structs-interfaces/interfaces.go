// interfaces.go - Go의 인터페이스 (Interfaces)
package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("=== Go 인터페이스 ===")

	// 1. 기본 인터페이스 사용
	fmt.Println("\n1. 기본 인터페이스:")

	var animals []Animal = []Animal{
		Dog{Name: "멍멍이"},
		Cat{Name: "야옹이"},
		Bird{Name: "짹짹이"},
	}

	for _, animal := range animals {
		animal.Speak()
		animal.Move()
	}

	// 2. 빈 인터페이스 (interface{})
	fmt.Println("\n2. 빈 인터페이스:")

	var anything interface{}
	anything = 42
	fmt.Printf("정수: %v (타입: %T)\n", anything, anything)

	anything = "Hello, Go!"
	fmt.Printf("문자열: %v (타입: %T)\n", anything, anything)

	anything = []int{1, 2, 3}
	fmt.Printf("슬라이스: %v (타입: %T)\n", anything, anything)

	// 3. 타입 단언 (Type Assertion)
	fmt.Println("\n3. 타입 단언:")

	var value interface{} = "Hello, World!"

	// 안전한 타입 단언
	if str, ok := value.(string); ok {
		fmt.Printf("문자열 값: %s (길이: %d)\n", str, len(str))
	} else {
		fmt.Println("문자열이 아닙니다")
	}

	// 위험한 타입 단언 (panic 가능)
	// number := value.(int) // panic 발생

	// 4. 타입 스위치 (Type Switch)
	fmt.Println("\n4. 타입 스위치:")

	values := []interface{}{
		42,
		"Hello",
		3.14,
		true,
		[]int{1, 2, 3},
		Dog{Name: "타입스위치독"},
	}

	for i, v := range values {
		fmt.Printf("값 %d: ", i)

		switch val := v.(type) {
		case int:
			fmt.Printf("정수 %d (2배: %d)\n", val, val*2)
		case string:
			fmt.Printf("문자열 '%s' (대문자: %s)\n", val, strings.ToUpper(val))
		case float64:
			fmt.Printf("실수 %.2f (제곱: %.2f)\n", val, val*val)
		case bool:
			fmt.Printf("불린 %t (반대: %t)\n", val, !val)
		case []int:
			fmt.Printf("정수 슬라이스 %v (합계: %d)\n", val, sum(val))
		case Animal:
			fmt.Printf("동물: ")
			val.Speak()
		default:
			fmt.Printf("알 수 없는 타입: %T\n", val)
		}
	}

	// 5. 인터페이스 조합 (Interface Composition)
	fmt.Println("\n5. 인터페이스 조합:")

	var worker Worker = Developer{
		Name:     "김개발",
		Language: "Go",
	}

	worker.Work()
	worker.TakeBreak()
	worker.GetSalary()

	// 6. 표준 라이브러리 인터페이스 활용
	fmt.Println("\n6. 표준 라이브러리 인터페이스:")

	// fmt.Stringer 인터페이스 구현
	product := Product{
		Name:  "Go 프로그래밍 책",
		Price: 25000,
		Stock: 10,
	}

	fmt.Printf("제품 정보: %s\n", product) // String() 메서드 자동 호출

	// io.Writer 인터페이스 사용
	var writer io.Writer = os.Stdout
	writer.Write([]byte("io.Writer 인터페이스로 출력\n"))

	// 사용자 정의 Writer 구현
	var buffer Buffer
	buffer.Write([]byte("버퍼에 쓰기"))
	fmt.Printf("버퍼 내용: %s\n", buffer.Data)

	// 7. 인터페이스와 메서드 집합
	fmt.Println("\n7. 인터페이스와 메서드 집합:")

	shape := Circle{Radius: 5}

	// Shape 인터페이스로 사용
	printArea(shape)

	// Drawable 인터페이스로 사용 (Circle이 Draw 메서드도 구현)
	printDrawing(shape)

	// 8. 인터페이스를 활용한 다형성
	fmt.Println("\n8. 다형성:")

	vehicles := []Vehicle{
		Car{Brand: "현대", Model: "소나타"},
		Bicycle{Brand: "삼천리"},
		Motorcycle{Brand: "할리데이비슨", CC: 883},
	}

	fmt.Println("교통수단별 시작:")
	for _, vehicle := range vehicles {
		vehicle.Start()
		fmt.Printf("속도: %d km/h\n", vehicle.GetSpeed())
	}

	// 9. 인터페이스 값의 내부 구조
	fmt.Println("\n9. 인터페이스 값의 내부 구조:")

	demonstrateInterfaceValues()

	// 10. 함수형 인터페이스
	fmt.Println("\n10. 함수형 인터페이스:")

	// 함수를 인터페이스로 구현
	operations := []Operation{
		AddOperation{},
		SubtractOperation{},
		OperationFunc(func(a, b int) int { return a * b }), // 함수를 인터페이스로
	}

	for i, op := range operations {
		result := op.Execute(10, 5)
		fmt.Printf("연산 %d: %d\n", i+1, result)
	}

	// 11. 인터페이스를 활용한 의존성 주입
	fmt.Println("\n11. 의존성 주입:")

	// 데이터베이스 구현체들
	mysqlDB := &MySQLDatabase{}
	postgresDB := &PostgreSQLDatabase{}

	// 같은 서비스 코드로 다른 데이터베이스 사용
	userService1 := UserService{DB: mysqlDB}
	userService2 := UserService{DB: postgresDB}

	userService1.CreateUser("김유저")
	userService2.CreateUser("박유저")

	// 12. 실무에서 자주 사용되는 인터페이스 패턴들
	fmt.Println("\n12. 실무 인터페이스 패턴:")

	// Strategy 패턴
	context := PaymentContext{}

	context.SetStrategy(&CreditCardPayment{})
	context.ProcessPayment(100000)

	context.SetStrategy(&BankTransferPayment{})
	context.ProcessPayment(200000)

	// Observer 패턴
	newsletter := Newsletter{}
	newsletter.Subscribe(&EmailSubscriber{Email: "user1@example.com"})
	newsletter.Subscribe(&SMSSubscriber{Phone: "010-1234-5678"})
	newsletter.Publish("새로운 Go 튜토리얼 업데이트!")
}

// 1. 기본 인터페이스 정의
type Animal interface {
	Speak()
	Move()
}

// Animal 인터페이스 구현체들
type Dog struct {
	Name string
}

func (d Dog) Speak() {
	fmt.Printf("%s: 멍멍!\n", d.Name)
}

func (d Dog) Move() {
	fmt.Printf("%s가 뛰어다닙니다.\n", d.Name)
}

type Cat struct {
	Name string
}

func (c Cat) Speak() {
	fmt.Printf("%s: 야옹~\n", c.Name)
}

func (c Cat) Move() {
	fmt.Printf("%s가 슬금슬금 걸어갑니다.\n", c.Name)
}

type Bird struct {
	Name string
}

func (b Bird) Speak() {
	fmt.Printf("%s: 짹짹!\n", b.Name)
}

func (b Bird) Move() {
	fmt.Printf("%s가 날아다닙니다.\n", b.Name)
}

// 4. 슬라이스 합계 계산 도우미 함수
func sum(numbers []int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// 5. 인터페이스 조합
type Employee interface {
	Work()
	GetSalary()
}

type Person interface {
	TakeBreak()
}

// 인터페이스 조합 (임베딩)
type Worker interface {
	Employee
	Person
}

type Developer struct {
	Name     string
	Language string
}

func (d Developer) Work() {
	fmt.Printf("%s가 %s로 개발 중입니다.\n", d.Name, d.Language)
}

func (d Developer) GetSalary() {
	fmt.Printf("%s가 월급을 받습니다.\n", d.Name)
}

func (d Developer) TakeBreak() {
	fmt.Printf("%s가 휴식을 취합니다.\n", d.Name)
}

// 6. 표준 라이브러리 인터페이스 구현
type Product struct {
	Name  string
	Price int
	Stock int
}

// fmt.Stringer 인터페이스 구현
func (p Product) String() string {
	return fmt.Sprintf("%s (가격: %d원, 재고: %d개)", p.Name, p.Price, p.Stock)
}

// 사용자 정의 Writer
type Buffer struct {
	Data string
}

func (b *Buffer) Write(data []byte) (int, error) {
	b.Data += string(data)
	return len(data), nil
}

// 7. 도형 관련 인터페이스들
type Shape interface {
	Area() float64
}

type Drawable interface {
	Draw()
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Draw() {
	fmt.Printf("반지름 %.1f인 원을 그립니다.\n", c.Radius)
}

func printArea(s Shape) {
	fmt.Printf("도형의 넓이: %.2f\n", s.Area())
}

func printDrawing(d Drawable) {
	d.Draw()
}

// 8. 다형성을 위한 Vehicle 인터페이스
type Vehicle interface {
	Start()
	GetSpeed() int
}

type Car struct {
	Brand string
	Model string
}

func (c Car) Start() {
	fmt.Printf("%s %s 시동을 겁니다.\n", c.Brand, c.Model)
}

func (c Car) GetSpeed() int {
	return 120 // km/h
}

type Bicycle struct {
	Brand string
}

func (b Bicycle) Start() {
	fmt.Printf("%s 자전거 페달을 밟습니다.\n", b.Brand)
}

func (b Bicycle) GetSpeed() int {
	return 25 // km/h
}

type Motorcycle struct {
	Brand string
	CC    int
}

func (m Motorcycle) Start() {
	fmt.Printf("%s %dcc 오토바이 시동을 겁니다.\n", m.Brand, m.CC)
}

func (m Motorcycle) GetSpeed() int {
	return 200 // km/h
}

// 9. 인터페이스 값의 내부 구조 데모
func demonstrateInterfaceValues() {
	var animal Animal

	fmt.Printf("nil 인터페이스: %v (타입: %T)\n", animal, animal)

	animal = Dog{Name: "내부구조독"}
	fmt.Printf("Dog 할당 후: %v (타입: %T)\n", animal, animal)

	if animal != nil {
		animal.Speak()
	}
}

// 10. 함수형 인터페이스
type Operation interface {
	Execute(a, b int) int
}

type AddOperation struct{}

func (AddOperation) Execute(a, b int) int {
	return a + b
}

type SubtractOperation struct{}

func (SubtractOperation) Execute(a, b int) int {
	return a - b
}

// 함수 타입을 인터페이스로 구현
type OperationFunc func(int, int) int

func (f OperationFunc) Execute(a, b int) int {
	return f(a, b)
}

// 11. 의존성 주입을 위한 인터페이스
type Database interface {
	Save(data string) error
	Load(id string) (string, error)
}

type MySQLDatabase struct{}

func (db *MySQLDatabase) Save(data string) error {
	fmt.Printf("MySQL에 데이터 저장: %s\n", data)
	return nil
}

func (db *MySQLDatabase) Load(id string) (string, error) {
	fmt.Printf("MySQL에서 데이터 로드: %s\n", id)
	return "MySQL 데이터", nil
}

type PostgreSQLDatabase struct{}

func (db *PostgreSQLDatabase) Save(data string) error {
	fmt.Printf("PostgreSQL에 데이터 저장: %s\n", data)
	return nil
}

func (db *PostgreSQLDatabase) Load(id string) (string, error) {
	fmt.Printf("PostgreSQL에서 데이터 로드: %s\n", id)
	return "PostgreSQL 데이터", nil
}

type UserService struct {
	DB Database
}

func (us *UserService) CreateUser(name string) {
	userData := "사용자: " + name
	us.DB.Save(userData)
}

// 12. 실무 패턴들

// Strategy 패턴
type PaymentStrategy interface {
	Pay(amount int) error
}

type CreditCardPayment struct{}

func (c *CreditCardPayment) Pay(amount int) error {
	fmt.Printf("신용카드로 %s원 결제\n", addCommas(amount))
	return nil
}

type BankTransferPayment struct{}

func (b *BankTransferPayment) Pay(amount int) error {
	fmt.Printf("계좌이체로 %s원 결제\n", addCommas(amount))
	return nil
}

type PaymentContext struct {
	strategy PaymentStrategy
}

func (pc *PaymentContext) SetStrategy(strategy PaymentStrategy) {
	pc.strategy = strategy
}

func (pc *PaymentContext) ProcessPayment(amount int) {
	if pc.strategy != nil {
		pc.strategy.Pay(amount)
	}
}

// Observer 패턴
type Subscriber interface {
	Notify(message string)
}

type EmailSubscriber struct {
	Email string
}

func (e *EmailSubscriber) Notify(message string) {
	fmt.Printf("이메일 (%s)로 알림: %s\n", e.Email, message)
}

type SMSSubscriber struct {
	Phone string
}

func (s *SMSSubscriber) Notify(message string) {
	fmt.Printf("SMS (%s)로 알림: %s\n", s.Phone, message)
}

type Newsletter struct {
	subscribers []Subscriber
}

func (n *Newsletter) Subscribe(subscriber Subscriber) {
	n.subscribers = append(n.subscribers, subscriber)
}

func (n *Newsletter) Publish(message string) {
	for _, subscriber := range n.subscribers {
		subscriber.Notify(message)
	}
}

// 유틸리티 함수
func addCommas(num int) string {
	str := strconv.Itoa(num)
	for i := len(str) - 3; i > 0; i -= 3 {
		str = str[:i] + "," + str[i:]
	}
	return str
}

/*
Go 인터페이스 특징:
1. 메서드 집합을 정의하는 타입
2. 암시적 구현 (implements 키워드 없음)
3. 덕 타이핑 (Duck Typing) - "오리처럼 걷고 울면 오리다"
4. 빈 인터페이스 interface{}는 모든 타입을 수용
5. 인터페이스 조합으로 새로운 인터페이스 생성
6. 값과 타입 정보를 함께 저장

인터페이스 사용 원칙:
1. 인터페이스는 작게 만들어라 (Single Method Interface 선호)
2. 소비자(사용하는 쪽)가 인터페이스를 정의하라
3. 필요한 곳에서만 인터페이스를 사용하라
4. 구체 타입을 받고 인터페이스를 반환하라

실무 활용:
1. 의존성 주입으로 테스트 용이성 향상
2. 플러그인 아키텍처 구현
3. 표준 라이브러리 인터페이스 활용 (io.Reader, io.Writer 등)
4. 디자인 패턴 구현 (Strategy, Observer, Factory 등)
5. 다형성을 통한 유연한 코드 작성
*/
