package golang_book

//package main

import "fmt"

func main() {
	variable()
}

func variable() {

	/**
	선언 후 사용하지 않으면 컴파일 에러
	var x string = "Hello World" < 그냥 이렇게만 쓰면 에러
	_ = x < 이 선언으로 컴파일 에러를 임시로 막기 가능
	*/
	// 1. 기본 선언과 초기화
	var x string = "Hello World"
	fmt.Println(x) // Hello World

	// 2. 선언만 하고 나중에 대입
	var y string
	y = "Hello World2"
	fmt.Println(y) // Hello World2

	// 3. 타입 생략 (Go가 타입 추론)
	var z = "Type inference works!"
	fmt.Println(z)

	// 4. := 단축 선언 (함수 안에서만 가능)
	msg := "Short declaration"
	fmt.Println(msg)

	// 5. 여러 개 한 번에 선언
	var a, b, c int = 1, 2, 3
	fmt.Println(a, b, c) // 1 2 3

	// 6. 여러 타입 한 번에 선언
	var (
		name string = "Alice"
		age  int    = 30
		ok   bool   = true
	)
	fmt.Println(name, age, ok) // Alice 30 true

	// 7. 기본값(Zero value) 확인
	var num int
	var str string
	var bln bool
	fmt.Println(num, str, bln) // 0 "" false

	// 8. 변수 재할당
	count := 10
	fmt.Println("count:", count)
	count = 20
	fmt.Println("count after reassignment:", count)

	// 9. 상수와의 차이
	const pi = 3.14
	// pi = 3.1415
	fmt.Println("pi:", pi)

	// 10. 타입 변환
	var i int = 42
	var f float64 = float64(i)
	fmt.Println("int:", i, "float64:", f) // int: 42 float64: 42
}
