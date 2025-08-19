package golang_book

import "fmt"

//package main

/**
 * 제어 구조
 */

func main() {
	fmt.Println(1)
	fmt.Println(2)
	fmt.Println(3)
	fmt.Println(4)
	fmt.Println(5)
	fmt.Println(6)
	fmt.Println(7)
	fmt.Println(8)
	fmt.Println(9)
	fmt.Println(10)

	fmt.Println(`1
2
3
4
5
6
7
8
9
10`)

	i := 1 // 단축 선언 (함수 안에서만)
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	for i := 1; i <= 10; i += 2 {
		fmt.Println(i)
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("짞쑤")
		} else {
			fmt.Println("홀쑤")
		}

		if i == 0 {
			fmt.Println("영")
		} else if i == 1 {
			fmt.Println("일")
		} else if i == 2 {
			fmt.Println("이")
		} else if i == 3 {
			fmt.Println("삼")
		} else if i == 4 {
			fmt.Println("사")
		} else if i == 5 {
			fmt.Println("오")
		}

		switch i {
		case 0:
			fmt.Println("영")
		case 1:
			fmt.Println("일")
		case 2:
			fmt.Println("이")
		case 3:
			fmt.Println("삼")
		case 4:
			fmt.Println("사")
		case 5:
			fmt.Println("오")
		default:
			fmt.Println("알 수 없는 숫자")
		}
	}
}
