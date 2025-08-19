package chapter6

import "fmt"

// 1. 배열 예제
func arrayExample() {
	fmt.Println("=== Array Example ===")
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array:", arr) // Array: [1 2 3 4 5]

	// 배열 요소 접근
	for i, v := range arr {
		fmt.Printf("arr[%d] = %d\n", i, v)
		// arr[0] = 1
		// arr[1] = 2
		// arr[2] = 3
		// arr[3] = 4
		// arr[4] = 5
	}
	fmt.Println()
}

// 2. 슬라이스 예제
func sliceExample() {
	fmt.Println("=== Slice Example ===")
	// 초기 슬라이스
	slice := []string{"Go", "Python", "Java"}
	fmt.Println("Initial slice:", slice) // Initial slice: [Go Python Java]

	// 슬라이스 잘라내기 (Slice view)
	subSlice := slice[1:3]                           // Python, Java
	fmt.Println("Sub-slice (slice[1:3]):", subSlice) // Sub-slice (slice[1:3]): [Python Java]

	// 슬라이스에 요소 추가 (append)
	slice = append(slice, "Rust")
	fmt.Println("After append:", slice) // After append: [Go Python Java Rust]

	// 원본 배열과 메모리 참조 차이 확인
	fmt.Printf("Original slice addr: %p\n", slice) // Original slice addr: 주소값 출력
	fmt.Printf("Sub-slice addr: %p\n", subSlice)   // Sub-slice addr: 주소값 출력
	fmt.Println()
}

// 3. 맵 예제
func mapExample() {
	fmt.Println("=== Map Example ===")
	m := map[string]int{
		"Alice": 30,
		"Bob":   25,
	}
	fmt.Println("Map:", m) // Map: map[Alice:30 Bob:25]

	// 맵 요소 접근
	for k, v := range m {
		fmt.Printf("%s => %d\n", k, v)
		// Alice => 30
		// Bob => 25
	}

	// 맵에 새로운 요소 추가
	m["Charlie"] = 35
	fmt.Println("After adding Charlie:", m) // After adding Charlie: map[Alice:30 Bob:25 Charlie:35]
	fmt.Println()
}
