// loops.go - Go의 반복문 (for 루프)
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("=== Go 반복문 ===")

	// Go에는 for 루프만 존재합니다 (while, do-while 없음)

	// 1. 기본 for 루프 (C 스타일)
	fmt.Println("\n1. 기본 for 루프:")

	fmt.Print("카운트다운: ")
	for i := 5; i >= 1; i-- {
		fmt.Printf("%d ", i)
		time.Sleep(500 * time.Millisecond) // 0.5초 대기
	}
	fmt.Println("발사! 🚀")

	// 2. while 루프처럼 사용 (조건만 있는 for)
	fmt.Println("\n2. while 루프 스타일:")

	number := 1
	fmt.Print("2의 거듭제곱: ")
	for number <= 100 {
		fmt.Printf("%d ", number)
		number *= 2
	}
	fmt.Println()

	// 3. 무한 루프
	fmt.Println("\n3. 무한 루프 (break로 종료):")

	count := 0
	for {
		count++
		fmt.Printf("반복 %d회\n", count)

		if count == 3 {
			fmt.Println("3회 반복 완료, 루프 종료")
			break
		}
	}

	// 4. continue 사용 (홀수만 출력)
	fmt.Println("\n4. continue 사용 (홀수만 출력):")

	fmt.Print("1부터 10까지 홀수: ")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue // 짝수는 건너뛰기
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// 5. 중첩 루프와 레이블
	fmt.Println("\n5. 중첩 루프와 레이블:")

	fmt.Println("구구단 (2~4단):")
OuterLoop:
	for i := 2; i <= 9; i++ {
		if i > 4 { // 4단까지만
			break OuterLoop
		}

		fmt.Printf("%d단: ", i)
		for j := 1; j <= 9; j++ {
			if j > 5 { // 5까지만
				fmt.Println() // 다음 단으로
				continue OuterLoop
			}
			fmt.Printf("%dx%d=%d ", i, j, i*j)
		}
	}

	// 6. 슬라이스 순회 (range 사용)
	fmt.Println("\n6. 슬라이스 순회:")

	fruits := []string{"사과", "바나나", "오렌지", "포도", "딸기"}

	// 인덱스와 값 모두 사용
	fmt.Println("과일 목록:")
	for index, fruit := range fruits {
		fmt.Printf("%d. %s\n", index+1, fruit)
	}

	// 값만 사용 (인덱스 무시)
	fmt.Print("과일들: ")
	for _, fruit := range fruits {
		fmt.Printf("%s ", fruit)
	}
	fmt.Println()

	// 인덱스만 사용 (값 무시)
	fmt.Print("인덱스들: ")
	for index := range fruits {
		fmt.Printf("%d ", index)
	}
	fmt.Println()

	// 7. 맵 순회
	fmt.Println("\n7. 맵 순회:")

	studentGrades := map[string]int{
		"김철수": 85,
		"이영희": 92,
		"박민수": 78,
		"최지원": 96,
	}

	fmt.Println("학생 성적:")
	for name, grade := range studentGrades {
		fmt.Printf("%s: %d점\n", name, grade)
	}

	// 키만 순회
	fmt.Print("학생 이름들: ")
	for name := range studentGrades {
		fmt.Printf("%s ", name)
	}
	fmt.Println()

	// 8. 문자열 순회 (룬 단위)
	fmt.Println("\n8. 문자열 순회:")

	koreanText := "안녕하세요"
	fmt.Printf("문자열 '%s' 분석:\n", koreanText)

	// 룬(문자) 단위로 순회
	for index, runeValue := range koreanText {
		fmt.Printf("인덱스 %d: %c (유니코드: %d)\n", index, runeValue, runeValue)
	}

	// 바이트 단위로 순회
	fmt.Printf("\n바이트 단위 순회:\n")
	for i := 0; i < len(koreanText); i++ {
		fmt.Printf("바이트 %d: %d\n", i, koreanText[i])
	}

	// 9. 채널 순회 (채널이 닫힐 때까지)
	fmt.Println("\n9. 채널 순회:")

	numbers := make(chan int)

	// 고루틴에서 숫자 전송
	go func() {
		for i := 1; i <= 5; i++ {
			numbers <- i * i // 제곱수 전송
			time.Sleep(200 * time.Millisecond)
		}
		close(numbers) // 채널 닫기
	}()

	fmt.Print("제곱수들: ")
	for num := range numbers {
		fmt.Printf("%d ", num)
	}
	fmt.Println()

	// 10. 실용적인 루프 예제들
	fmt.Println("\n10. 실용적인 예제들:")

	// 최대값 찾기
	values := []int{23, 45, 12, 67, 34, 89, 56}
	max := findMax(values)
	fmt.Printf("배열 %v에서 최대값: %d\n", values, max)

	// 팩토리얼 계산
	n := 5
	factorial := calculateFactorial(n)
	fmt.Printf("%d! = %d\n", n, factorial)

	// 소수 찾기
	limit := 20
	primes := findPrimes(limit)
	fmt.Printf("%d 이하의 소수: %v\n", limit, primes)

	// 피보나치 수열 생성
	fibCount := 10
	fibonacci := generateFibonacci(fibCount)
	fmt.Printf("피보나치 수열 (첫 %d개): %v\n", fibCount, fibonacci)

	// 배열 뒤집기
	original := []string{"A", "B", "C", "D", "E"}
	reversed := reverseSlice(original)
	fmt.Printf("원본: %v, 뒤집힌 배열: %v\n", original, reversed)

	// 단어 개수 세기
	text := "Go는 정말 멋진 프로그래밍 언어입니다. Go로 많은 것을 만들 수 있어요!"
	wordCount := countWords(text)
	fmt.Printf("텍스트: %s\n", text)
	fmt.Printf("단어 개수: %d개\n", wordCount)

	// 11. 성능을 고려한 루프
	fmt.Println("\n11. 성능 고려사항:")

	// 큰 슬라이스에서 range vs 인덱스 비교
	largeSlice := make([]int, 1000000)
	for i := range largeSlice {
		largeSlice[i] = rand.Intn(1000)
	}

	// range 사용 (값 복사 방지)
	sum1 := 0
	start := time.Now()
	for _, value := range largeSlice {
		sum1 += value
	}
	duration1 := time.Since(start)

	// 인덱스 사용
	sum2 := 0
	start = time.Now()
	for i := 0; i < len(largeSlice); i++ {
		sum2 += largeSlice[i]
	}
	duration2 := time.Since(start)

	fmt.Printf("range 사용: 합계 %d, 시간 %v\n", sum1, duration1)
	fmt.Printf("인덱스 사용: 합계 %d, 시간 %v\n", sum2, duration2)
}

// 최대값을 찾는 함수
func findMax(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	max := numbers[0]
	for _, num := range numbers[1:] {
		if num > max {
			max = num
		}
	}
	return max
}

// 팩토리얼을 계산하는 함수
func calculateFactorial(n int) int {
	if n <= 1 {
		return 1
	}

	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

// 소수를 찾는 함수
func findPrimes(limit int) []int {
	var primes []int

	for num := 2; num <= limit; num++ {
		isPrime := true

		// 제곱근까지만 확인하면 충분
		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			primes = append(primes, num)
		}
	}

	return primes
}

// 피보나치 수열을 생성하는 함수
func generateFibonacci(count int) []int {
	if count <= 0 {
		return []int{}
	}

	fibonacci := make([]int, count)

	if count >= 1 {
		fibonacci[0] = 0
	}
	if count >= 2 {
		fibonacci[1] = 1
	}

	for i := 2; i < count; i++ {
		fibonacci[i] = fibonacci[i-1] + fibonacci[i-2]
	}

	return fibonacci
}

// 슬라이스를 뒤집는 함수
func reverseSlice(slice []string) []string {
	reversed := make([]string, len(slice))

	for i, value := range slice {
		reversed[len(slice)-1-i] = value
	}

	return reversed
}

// 단어 개수를 세는 함수
func countWords(text string) int {
	count := 0
	inWord := false

	for _, char := range text {
		if char == ' ' || char == '\t' || char == '\n' {
			inWord = false
		} else {
			if !inWord {
				count++
				inWord = true
			}
		}
	}

	return count
}

/*
Go 반복문 특징:
1. for 루프만 존재 (while, do-while 없음)
2. 세 가지 형태: C 스타일, while 스타일, 무한 루프
3. range 키워드로 배열, 슬라이스, 맵, 채널, 문자열 순회
4. break, continue 키워드 지원
5. 레이블을 사용한 중첩 루프 제어
6. 고루틴과 채널을 활용한 동시성 처리

루프 사용 팁:
1. range를 사용할 때 불필요한 값 복사 방지 (포인터나 인덱스 사용)
2. 큰 데이터셋은 고루틴과 채널로 병렬 처리 고려
3. break와 continue를 적절히 사용하여 불필요한 반복 방지
4. 중첩 루프에서는 레이블 사용으로 명확한 제어
5. 무한 루프 사용 시 반드시 종료 조건 확인
*/
