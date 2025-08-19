package main

import "fmt"

func main() {
	//fmt.Println("1 + 1 =", 1+1)
	//
	//fmt.Println("1 + 1 =", 1.1+1.0)
	//fmt.Println("1 + 1 =", 1.1+1.0)
	//str()
	//funcBool()
}

func str() {
	fmt.Println("str()")
	fmt.Println(len("Hello World"))       // 11
	fmt.Println("Hello World"[1])         // 아스키로 변환해서 101 이라는 값이 나옴
	fmt.Println(string("Hello World"[1])) // e
	fmt.Println("Hello " + "World")       // Hello World
}

func funcBool() {
	fmt.Println("bool()")
	fmt.Println(true && true)  // true
	fmt.Println(true && false) // false
	fmt.Println(true || true)  // true
	fmt.Println(true || false) // true
	fmt.Println(!true)         // false
}
