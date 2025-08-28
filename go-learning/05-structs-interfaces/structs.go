// structs.go - Go의 구조체 (Structs)
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Go 구조체 ===")

	// 1. 기본 구조체 정의와 사용
	fmt.Println("\n1. 기본 구조체:")

	// 구조체 인스턴스 생성 방법 1: 필드별 초기화
	var person1 Person
	person1.Name = "김철수"
	person1.Age = 30
	person1.Email = "kim@example.com"

	fmt.Printf("person1: %+v\n", person1)

	// 구조체 인스턴스 생성 방법 2: 구조체 리터럴
	person2 := Person{
		Name:  "이영희",
		Age:   25,
		Email: "lee@example.com",
	}

	fmt.Printf("person2: %+v\n", person2)

	// 구조체 인스턴스 생성 방법 3: 순서대로 초기화
	person3 := Person{"박민수", 35, "park@example.com"}
	fmt.Printf("person3: %+v\n", person3)

	// 2. 익명 구조체
	fmt.Println("\n2. 익명 구조체:")

	config := struct {
		Host string
		Port int
		SSL  bool
	}{
		Host: "localhost",
		Port: 8080,
		SSL:  true,
	}

	fmt.Printf("서버 설정: %+v\n", config)

	// 3. 중첩 구조체 (Nested Structs)
	fmt.Println("\n3. 중첩 구조체:")

	company := Company{
		Name: "테크 컴퍼니",
		Address: Address{
			Street:   "강남대로 123",
			City:     "서울",
			Country:  "대한민국",
			PostCode: "12345",
		},
		CEO: Person{
			Name:  "최대표",
			Age:   45,
			Email: "ceo@techcompany.com",
		},
	}

	fmt.Printf("회사 정보:\n")
	fmt.Printf("  이름: %s\n", company.Name)
	fmt.Printf("  주소: %s, %s, %s (%s)\n",
		company.Address.Street,
		company.Address.City,
		company.Address.Country,
		company.Address.PostCode)
	fmt.Printf("  CEO: %s (%s)\n", company.CEO.Name, company.CEO.Email)

	// 4. 임베디드 구조체 (Embedded Structs)
	fmt.Println("\n4. 임베디드 구조체:")

	employee := Employee{
		Person: Person{
			Name:  "김직원",
			Age:   28,
			Email: "employee@company.com",
		},
		ID:         "EMP001",
		Department: "개발팀",
		Salary:     5000000,
	}

	// 임베디드 필드에 직접 접근 가능
	fmt.Printf("직원 이름: %s\n", employee.Name) // employee.Person.Name과 같음
	fmt.Printf("직원 나이: %d\n", employee.Age)
	fmt.Printf("직원 ID: %s\n", employee.ID)
	fmt.Printf("부서: %s\n", employee.Department)

	// 5. 구조체 포인터
	fmt.Println("\n5. 구조체 포인터:")

	// new 함수로 포인터 생성
	personPtr := new(Person)
	personPtr.Name = "포인터 사람"
	personPtr.Age = 40

	fmt.Printf("포인터를 통한 접근: %s, %d세\n", personPtr.Name, personPtr.Age)

	// & 연산자로 포인터 생성
	person4 := Person{"주소 사람", 32, "addr@example.com"}
	personPtr2 := &person4

	fmt.Printf("원본: %+v\n", person4)
	fmt.Printf("포인터: %+v\n", personPtr2)

	// 포인터를 통한 수정
	personPtr2.Age = 33
	fmt.Printf("포인터로 수정 후: %+v\n", person4)

	// 6. 구조체 메서드 (값 리시버)
	fmt.Println("\n6. 구조체 메서드 (값 리시버):")

	rect := Rectangle{Width: 5, Height: 3}
	fmt.Printf("직사각형: 너비=%g, 높이=%g\n", rect.Width, rect.Height)
	fmt.Printf("넓이: %g\n", rect.Area())
	fmt.Printf("둘레: %g\n", rect.Perimeter())

	// 7. 구조체 메서드 (포인터 리시버)
	fmt.Println("\n7. 구조체 메서드 (포인터 리시버):")

	account := BankAccount{
		AccountNumber: "123-456-789",
		Balance:       100000,
		Owner:         "김계좌",
	}

	fmt.Printf("초기 잔액: %d원\n", account.Balance)

	account.Deposit(50000)
	fmt.Printf("입금 후 잔액: %d원\n", account.Balance)

	if account.Withdraw(30000) {
		fmt.Printf("출금 후 잔액: %d원\n", account.Balance)
	}

	if !account.Withdraw(200000) {
		fmt.Println("잔액 부족으로 출금 실패")
	}

	// 8. 구조체 태그 (Struct Tags)
	fmt.Println("\n8. 구조체 태그:")

	user := User{
		ID:       1,
		Username: "gopher",
		Email:    "gopher@golang.org",
		Created:  time.Now(),
	}

	// 구조체 태그는 주로 JSON, XML, 데이터베이스 매핑 등에 사용
	fmt.Printf("사용자 정보: %+v\n", user)

	// 9. 구조체 비교
	fmt.Println("\n9. 구조체 비교:")

	point1 := Point{X: 1, Y: 2}
	point2 := Point{X: 1, Y: 2}
	point3 := Point{X: 2, Y: 3}

	fmt.Printf("point1 == point2: %t\n", point1 == point2)
	fmt.Printf("point1 == point3: %t\n", point1 == point3)

	// 10. 구조체 복사
	fmt.Println("\n10. 구조체 복사:")

	original := Person{"원본", 25, "original@example.com"}
	copy := original // 값 복사 (shallow copy)

	copy.Name = "복사본"
	copy.Age = 26

	fmt.Printf("원본: %+v\n", original)
	fmt.Printf("복사본: %+v\n", copy)

	// 11. 구조체를 활용한 실제 예제
	fmt.Println("\n11. 실제 활용 예제:")

	// 학생 관리 시스템
	students := []Student{
		{"김학생", 20, "컴퓨터공학", []int{85, 90, 88}},
		{"이학생", 21, "수학과", []int{92, 87, 94}},
		{"박학생", 19, "물리학과", []int{78, 82, 85}},
	}

	fmt.Println("학생 성적표:")
	for _, student := range students {
		avg := student.Average()
		grade := student.Grade()
		fmt.Printf("이름: %s, 전공: %s, 평균: %.2f, 학점: %s\n",
			student.Name, student.Major, avg, grade)
	}

	// 도서관 시스템
	library := Library{
		Name: "중앙도서관",
		Books: []Book{
			{"Go 프로그래밍", "김고랭", 2023, true},
			{"알고리즘 기초", "박알고", 2022, false},
			{"데이터베이스 개론", "이디비", 2021, true},
		},
	}

	fmt.Printf("\n%s 도서 목록:\n", library.Name)
	availableBooks := library.GetAvailableBooks()
	for _, book := range availableBooks {
		fmt.Printf("- %s (저자: %s, %d년)\n", book.Title, book.Author, book.Year)
	}

	// 12. 구조체와 인터페이스 함께 사용
	fmt.Println("\n12. 구조체와 인터페이스:")

	shapes := []Shape{
		Circle{Radius: 5},
		Rectangle{Width: 4, Height: 6},
		Triangle{Base: 8, Height: 3},
	}

	fmt.Println("도형들의 넓이:")
	for i, shape := range shapes {
		fmt.Printf("%d. %T: 넓이 = %.2f\n", i+1, shape, shape.Area())
	}
}

// 1. 기본 구조체 정의
type Person struct {
	Name  string
	Age   int
	Email string
}

// 3. 중첩 구조체를 위한 타입들
type Address struct {
	Street   string
	City     string
	Country  string
	PostCode string
}

type Company struct {
	Name    string
	Address Address
	CEO     Person
}

// 4. 임베디드 구조체
type Employee struct {
	Person     // 임베디드 필드 (상속과 유사)
	ID         string
	Department string
	Salary     int
}

// 6. 메서드를 가진 구조체 (값 리시버)
type Rectangle struct {
	Width  float64
	Height float64
}

// 값 리시버 메서드 (구조체를 복사하여 사용)
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// 7. 포인터 리시버 메서드
type BankAccount struct {
	AccountNumber string
	Balance       int
	Owner         string
}

// 포인터 리시버 메서드 (원본 구조체를 수정 가능)
func (ba *BankAccount) Deposit(amount int) {
	ba.Balance += amount
}

func (ba *BankAccount) Withdraw(amount int) bool {
	if ba.Balance >= amount {
		ba.Balance -= amount
		return true
	}
	return false
}

// 8. 구조체 태그
type User struct {
	ID       int       `json:"id" db:"user_id"`
	Username string    `json:"username" db:"username" validate:"required,min=3"`
	Email    string    `json:"email" db:"email" validate:"required,email"`
	Created  time.Time `json:"created_at" db:"created_at"`
}

// 9. 구조체 비교용
type Point struct {
	X, Y int
}

// 11. 실제 예제용 구조체들
type Student struct {
	Name   string
	Age    int
	Major  string
	Scores []int
}

func (s Student) Average() float64 {
	if len(s.Scores) == 0 {
		return 0
	}

	sum := 0
	for _, score := range s.Scores {
		sum += score
	}
	return float64(sum) / float64(len(s.Scores))
}

func (s Student) Grade() string {
	avg := s.Average()
	switch {
	case avg >= 90:
		return "A"
	case avg >= 80:
		return "B"
	case avg >= 70:
		return "C"
	case avg >= 60:
		return "D"
	default:
		return "F"
	}
}

type Book struct {
	Title     string
	Author    string
	Year      int
	Available bool
}

type Library struct {
	Name  string
	Books []Book
}

func (l Library) GetAvailableBooks() []Book {
	var available []Book
	for _, book := range l.Books {
		if book.Available {
			available = append(available, book)
		}
	}
	return available
}

// 12. 인터페이스 (다음 파일에서 자세히 다룰 예정)
type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

/*
Go 구조체 특징:
1. 값 타입 (Value Type) - 할당 시 복사됨
2. 메서드를 가질 수 있음 (값 리시버 또는 포인터 리시버)
3. 상속은 없지만 임베딩(embedding)으로 유사한 효과
4. 구조체 태그로 메타데이터 제공
5. 모든 필드가 comparable하면 구조체도 비교 가능
6. 익명 구조체로 임시 데이터 구조 생성 가능

구조체 사용 팁:
1. 필드명은 대문자로 시작하면 public, 소문자면 private
2. 메서드가 구조체를 변경해야 하면 포인터 리시버 사용
3. 읽기만 하는 메서드는 값 리시버 사용 (복사 비용 고려)
4. 임베딩을 통해 코드 재사용성 향상
5. 구조체 태그를 활용한 직렬화/역직렬화
6. new() 또는 &로 포인터 생성
7. 빈 구조체 struct{}는 메모리를 사용하지 않음
*/
