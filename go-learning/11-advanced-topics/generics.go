// generics.go - Go 제네릭 (Go 1.18+)
package main

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func main() {
	fmt.Println("=== Go 제네릭 (Generics) ===")

	// 1. 기본 제네릭 함수
	fmt.Println("\n1. 기본 제네릭 함수:")
	demonstrateBasicGenerics()

	// 2. 타입 제약 (Type Constraints)
	fmt.Println("\n2. 타입 제약:")
	demonstrateTypeConstraints()

	// 3. 제네릭 슬라이스 함수들
	fmt.Println("\n3. 제네릭 슬라이스 함수:")
	demonstrateGenericSliceFunctions()

	// 4. 제네릭 맵 함수들
	fmt.Println("\n4. 제네릭 맵 함수:")
	demonstrateGenericMapFunctions()

	// 5. 제네릭 구조체
	fmt.Println("\n5. 제네릭 구조체:")
	demonstrateGenericStructs()

	// 6. 제네릭 인터페이스
	fmt.Println("\n6. 제네릭 인터페이스:")
	demonstrateGenericInterfaces()

	// 7. 타입 추론
	fmt.Println("\n7. 타입 추론:")
	demonstrateTypeInference()

	// 8. 실용적인 제네릭 예제들
	fmt.Println("\n8. 실용적인 예제들:")
	demonstratePracticalExamples()
}

// 1. 기본 제네릭 함수들
func demonstrateBasicGenerics() {
	// 제네릭 함수로 다양한 타입에 대해 작업
	fmt.Printf("정수 최대값: %d\n", Max(5, 10))
	fmt.Printf("실수 최대값: %.2f\n", Max(3.14, 2.71))
	fmt.Printf("문자열 최대값: %s\n", Max("apple", "banana"))

	// 슬라이스의 첫 번째 요소
	intSlice := []int{1, 2, 3, 4, 5}
	stringSlice := []string{"Go", "is", "awesome"}

	fmt.Printf("정수 슬라이스 첫 요소: %d\n", First(intSlice))
	fmt.Printf("문자열 슬라이스 첫 요소: %s\n", First(stringSlice))

	// 두 값 교체
	a, b := 10, 20
	fmt.Printf("교체 전: a=%d, b=%d\n", a, b)
	a, b = Swap(a, b)
	fmt.Printf("교체 후: a=%d, b=%d\n", a, b)
}

// 비교 가능한 타입들에 대한 최대값
func Max[T comparable](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// 슬라이스의 첫 번째 요소 반환
func First[T any](slice []T) T {
	var zero T
	if len(slice) == 0 {
		return zero // 제로값 반환
	}
	return slice[0]
}

// 두 값 교체
func Swap[T any](a, b T) (T, T) {
	return b, a
}

// 2. 타입 제약 예제들
func demonstrateTypeConstraints() {
	numbers := []int{5, 2, 8, 1, 9, 3}
	fmt.Printf("정수 합계: %d\n", Sum(numbers))

	floats := []float64{1.1, 2.2, 3.3}
	fmt.Printf("실수 합계: %.2f\n", Sum(floats))

	// 정렬 가능한 슬라이스 정렬
	fruits := []string{"banana", "apple", "cherry"}
	fmt.Printf("정렬 전: %v\n", fruits)
	SortSlice(fruits)
	fmt.Printf("정렬 후: %v\n", fruits)
}

// 숫자 타입 제약
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

// 숫자 슬라이스의 합계
func Sum[T Number](numbers []T) T {
	var sum T
	for _, num := range numbers {
		sum += num
	}
	return sum
}

// 정렬 가능한 타입 제약
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 | ~string
}

// 제네릭 슬라이스 정렬
func SortSlice[T Ordered](slice []T) {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
}

// 3. 제네릭 슬라이스 함수들
func demonstrateGenericSliceFunctions() {
	numbers := []int{1, 2, 3, 4, 5}

	// Map 함수
	doubled := Map(numbers, func(x int) int { return x * 2 })
	fmt.Printf("원본: %v, 2배: %v\n", numbers, doubled)

	// 다른 타입으로 변환
	strings := Map(numbers, func(x int) string { return fmt.Sprintf("num-%d", x) })
	fmt.Printf("문자열 변환: %v\n", strings)

	// Filter 함수
	evens := Filter(numbers, func(x int) bool { return x%2 == 0 })
	fmt.Printf("짝수 필터: %v\n", evens)

	// Reduce 함수
	sum := Reduce(numbers, 0, func(acc, x int) int { return acc + x })
	fmt.Printf("합계: %d\n", sum)

	// Contains 함수
	fmt.Printf("3 포함: %t\n", Contains(numbers, 3))
	fmt.Printf("10 포함: %t\n", Contains(numbers, 10))

	// Unique 함수
	duplicates := []int{1, 2, 2, 3, 3, 3, 4}
	unique := Unique(duplicates)
	fmt.Printf("중복제거: %v -> %v\n", duplicates, unique)
}

// Map 함수 (슬라이스의 각 요소를 변환)
func Map[T, U any](slice []T, fn func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

// Filter 함수 (조건에 맞는 요소만 선택)
func Filter[T any](slice []T, predicate func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// Reduce 함수 (슬라이스를 하나의 값으로 축약)
func Reduce[T, U any](slice []T, initial U, fn func(U, T) U) U {
	result := initial
	for _, v := range slice {
		result = fn(result, v)
	}
	return result
}

// Contains 함수 (요소 포함 여부 확인)
func Contains[T comparable](slice []T, target T) bool {
	for _, v := range slice {
		if v == target {
			return true
		}
	}
	return false
}

// Unique 함수 (중복 요소 제거)
func Unique[T comparable](slice []T) []T {
	seen := make(map[T]bool)
	var result []T

	for _, v := range slice {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

// 4. 제네릭 맵 함수들
func demonstrateGenericMapFunctions() {
	// 맵 키들 가져오기
	intMap := map[string]int{"a": 1, "b": 2, "c": 3}
	keys := Keys(intMap)
	fmt.Printf("맵 키들: %v\n", keys)

	// 맵 값들 가져오기
	values := Values(intMap)
	fmt.Printf("맵 값들: %v\n", values)

	// 맵 변환
	doubled := MapValues(intMap, func(v int) int { return v * 2 })
	fmt.Printf("값 2배: %v\n", doubled)

	// 맵 필터링
	filtered := FilterMap(intMap, func(k string, v int) bool { return v > 1 })
	fmt.Printf("1보다 큰 값들: %v\n", filtered)
}

// 맵의 모든 키 반환
func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// 맵의 모든 값 반환
func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

// 맵의 값들을 변환
func MapValues[K comparable, V, U any](m map[K]V, fn func(V) U) map[K]U {
	result := make(map[K]U, len(m))
	for k, v := range m {
		result[k] = fn(v)
	}
	return result
}

// 맵 필터링
func FilterMap[K comparable, V any](m map[K]V, predicate func(K, V) bool) map[K]V {
	result := make(map[K]V)
	for k, v := range m {
		if predicate(k, v) {
			result[k] = v
		}
	}
	return result
}

// 5. 제네릭 구조체들
func demonstrateGenericStructs() {
	// 제네릭 스택
	intStack := NewStack[int]()
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)

	fmt.Printf("스택 크기: %d\n", intStack.Size())
	fmt.Printf("팝: %d\n", intStack.Pop())
	fmt.Printf("픽: %d\n", intStack.Peek())

	// 문자열 스택
	stringStack := NewStack[string]()
	stringStack.Push("Go")
	stringStack.Push("Generics")
	fmt.Printf("문자열 스택 팝: %s\n", stringStack.Pop())

	// 제네릭 페어
	intPair := NewPair(10, 20)
	fmt.Printf("정수 페어: %v\n", intPair)

	stringPair := NewPair("Hello", "World")
	fmt.Printf("문자열 페어: %v\n", stringPair)

	// Optional 타입
	some := Some(42)
	none := None[int]()

	fmt.Printf("Some 값: %v\n", some.Get())
	fmt.Printf("None 존재: %t\n", none.IsPresent())
}

// 제네릭 스택
type Stack[T any] struct {
	items []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{items: make([]T, 0)}
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() T {
	var zero T
	if len(s.items) == 0 {
		return zero
	}

	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item
}

func (s *Stack[T]) Peek() T {
	var zero T
	if len(s.items) == 0 {
		return zero
	}
	return s.items[len(s.items)-1]
}

func (s *Stack[T]) Size() int {
	return len(s.items)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// 제네릭 페어
type Pair[T, U any] struct {
	First  T
	Second U
}

func NewPair[T, U any](first T, second U) Pair[T, U] {
	return Pair[T, U]{First: first, Second: second}
}

func (p Pair[T, U]) String() string {
	return fmt.Sprintf("(%v, %v)", p.First, p.Second)
}

// Optional 타입 (Go에는 없는 널 안전성)
type Optional[T any] struct {
	value   T
	present bool
}

func Some[T any](value T) Optional[T] {
	return Optional[T]{value: value, present: true}
}

func None[T any]() Optional[T] {
	return Optional[T]{present: false}
}

func (o Optional[T]) IsPresent() bool {
	return o.present
}

func (o Optional[T]) Get() T {
	return o.value
}

func (o Optional[T]) GetOrElse(defaultValue T) T {
	if o.present {
		return o.value
	}
	return defaultValue
}

// 6. 제네릭 인터페이스
func demonstrateGenericInterfaces() {
	// 제네릭 컨테이너 인터페이스 사용
	var intContainer Container[int] = NewListContainer[int]()
	intContainer.Add(1)
	intContainer.Add(2)
	intContainer.Add(3)

	fmt.Printf("컨테이너 크기: %d\n", intContainer.Size())
	fmt.Printf("컨테이너 요소들: %v\n", intContainer.GetAll())

	// 제네릭 비교자 사용
	var intComparator Comparator[int] = IntComparator{}
	fmt.Printf("5와 3 비교: %d\n", intComparator.Compare(5, 3))

	var stringComparator Comparator[string] = StringComparator{}
	fmt.Printf("\"apple\"과 \"banana\" 비교: %d\n", stringComparator.Compare("apple", "banana"))
}

// 제네릭 컨테이너 인터페이스
type Container[T any] interface {
	Add(item T)
	Remove(item T) bool
	Contains(item T) bool
	Size() int
	GetAll() []T
}

// 리스트 기반 컨테이너 구현
type ListContainer[T comparable] struct {
	items []T
}

func NewListContainer[T comparable]() *ListContainer[T] {
	return &ListContainer[T]{items: make([]T, 0)}
}

func (lc *ListContainer[T]) Add(item T) {
	lc.items = append(lc.items, item)
}

func (lc *ListContainer[T]) Remove(item T) bool {
	for i, v := range lc.items {
		if v == item {
			lc.items = append(lc.items[:i], lc.items[i+1:]...)
			return true
		}
	}
	return false
}

func (lc *ListContainer[T]) Contains(item T) bool {
	for _, v := range lc.items {
		if v == item {
			return true
		}
	}
	return false
}

func (lc *ListContainer[T]) Size() int {
	return len(lc.items)
}

func (lc *ListContainer[T]) GetAll() []T {
	return lc.items[:]
}

// 제네릭 비교자 인터페이스
type Comparator[T any] interface {
	Compare(a, b T) int // -1: a < b, 0: a == b, 1: a > b
}

type IntComparator struct{}

func (IntComparator) Compare(a, b int) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

type StringComparator struct{}

func (StringComparator) Compare(a, b string) int {
	return strings.Compare(a, b)
}

// 7. 타입 추론 예제
func demonstrateTypeInference() {
	// 타입을 명시적으로 지정하지 않아도 추론됨
	result1 := Max(10, 20)        // int로 추론
	result2 := Max(3.14, 2.71)    // float64로 추론
	result3 := Max("hello", "hi") // string으로 추론

	fmt.Printf("타입 추론 결과: %T=%v, %T=%.2f, %T=%s\n",
		result1, result1, result2, result2, result3, result3)

	// 함수 매개변수에서 타입 추론
	numbers := []int{1, 2, 3, 4, 5}
	processSlice(numbers, func(x int) int { return x * x }) // 타입 추론

	strings := []string{"Go", "generics", "are", "awesome"}
	processSlice(strings, func(s string) string { return strings.ToUpper(s) })
}

func processSlice[T any](slice []T, fn func(T) T) {
	fmt.Printf("처리 전: %v\n", slice)
	result := Map(slice, fn)
	fmt.Printf("처리 후: %v\n", result)
}

// 8. 실용적인 제네릭 예제들
func demonstratePracticalExamples() {
	// 결과 타입 (성공/실패)
	successResult := Ok[string, error]("성공!")
	errorResult := Err[string, error](fmt.Errorf("실패!"))

	fmt.Printf("성공 결과: %s\n", handleResult(successResult))
	fmt.Printf("실패 결과: %s\n", handleResult(errorResult))

	// 제네릭 캐시
	cache := NewLRUCache[string, int](3)
	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3)
	cache.Put("d", 4) // "a"가 제거됨 (LRU)

	if value, found := cache.Get("a"); found {
		fmt.Printf("캐시에서 'a' 발견: %d\n", value)
	} else {
		fmt.Println("캐시에서 'a' 없음 (LRU로 제거됨)")
	}

	if value, found := cache.Get("b"); found {
		fmt.Printf("캐시에서 'b' 발견: %d\n", value)
	}
}

// Result 타입 (Rust의 Result와 유사)
type Result[T, E any] struct {
	value T
	err   E
	isOk  bool
}

func Ok[T, E any](value T) Result[T, E] {
	return Result[T, E]{value: value, isOk: true}
}

func Err[T, E any](err E) Result[T, E] {
	return Result[T, E]{err: err, isOk: false}
}

func (r Result[T, E]) IsOk() bool {
	return r.isOk
}

func (r Result[T, E]) IsErr() bool {
	return !r.isOk
}

func (r Result[T, E]) Unwrap() T {
	return r.value
}

func (r Result[T, E]) UnwrapErr() E {
	return r.err
}

func handleResult[T, E any](result Result[T, E]) string {
	if result.IsOk() {
		return fmt.Sprintf("성공: %v", result.Unwrap())
	}
	return fmt.Sprintf("실패: %v", result.UnwrapErr())
}

// 간단한 LRU 캐시
type LRUCache[K comparable, V any] struct {
	capacity int
	cache    map[K]*Node[K, V]
	head     *Node[K, V]
	tail     *Node[K, V]
}

type Node[K comparable, V any] struct {
	key   K
	value V
	prev  *Node[K, V]
	next  *Node[K, V]
}

func NewLRUCache[K comparable, V any](capacity int) *LRUCache[K, V] {
	cache := &LRUCache[K, V]{
		capacity: capacity,
		cache:    make(map[K]*Node[K, V]),
	}

	// 더미 헤드와 테일 노드
	cache.head = &Node[K, V]{}
	cache.tail = &Node[K, V]{}
	cache.head.next = cache.tail
	cache.tail.prev = cache.head

	return cache
}

func (lru *LRUCache[K, V]) Get(key K) (V, bool) {
	var zero V
	if node, exists := lru.cache[key]; exists {
		lru.moveToHead(node)
		return node.value, true
	}
	return zero, false
}

func (lru *LRUCache[K, V]) Put(key K, value V) {
	if node, exists := lru.cache[key]; exists {
		node.value = value
		lru.moveToHead(node)
	} else {
		newNode := &Node[K, V]{key: key, value: value}

		if len(lru.cache) >= lru.capacity {
			lru.removeTail()
		}

		lru.cache[key] = newNode
		lru.addToHead(newNode)
	}
}

func (lru *LRUCache[K, V]) addToHead(node *Node[K, V]) {
	node.prev = lru.head
	node.next = lru.head.next
	lru.head.next.prev = node
	lru.head.next = node
}

func (lru *LRUCache[K, V]) removeNode(node *Node[K, V]) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (lru *LRUCache[K, V]) moveToHead(node *Node[K, V]) {
	lru.removeNode(node)
	lru.addToHead(node)
}

func (lru *LRUCache[K, V]) removeTail() {
	lastNode := lru.tail.prev
	lru.removeNode(lastNode)
	delete(lru.cache, lastNode.key)
}

/*
Go 제네릭 (Generics) 완전 가이드:

1. 제네릭이란?:
   - 타입을 매개변수화하여 재사용 가능한 코드 작성
   - Go 1.18부터 도입
   - 컴파일 타임에 타입 안전성 보장
   - 런타임 성능 영향 최소화

2. 문법:
   - func FuncName[T TypeConstraint](params) ReturnType
   - type TypeName[T TypeConstraint] struct { ... }
   - interface InterfaceName[T TypeConstraint] { ... }

3. 타입 제약 (Type Constraints):
   - any: 모든 타입 허용 (interface{}와 동일)
   - comparable: 비교 연산자(==, !=) 사용 가능한 타입
   - 사용자 정의 제약: 타입 유니온 사용

4. 타입 유니온:
   - ~int | ~string | ~float64 형태
   - ~ 기호는 underlying type 포함
   - 여러 타입을 하나의 제약으로 그룹화

5. 타입 추론:
   - 대부분의 경우 타입 명시 불필요
   - 컴파일러가 문맥에서 타입 추론
   - 복잡한 경우만 명시적 타입 지정

6. 실용적 활용:
   - 컬렉션 조작 함수 (Map, Filter, Reduce)
   - 자료구조 (Stack, Queue, Cache)
   - 옵셔널 타입, 결과 타입
   - 컨테이너 및 래퍼 타입

7. 성능 고려사항:
   - 제네릭은 컴파일 타임에 특화됨
   - interface{} 사용보다 빠름
   - 타입 assertion 비용 제거
   - 메모리 할당 최적화

8. 모범 사례:
   - 간단한 타입 제약 사용
   - 제네릭 함수는 순수 함수로 작성
   - 복잡한 제약보다는 인터페이스 활용
   - 기존 interface{} 코드의 점진적 마이그레이션

9. 주의사항:
   - 과도한 제네릭 사용 지양
   - 가독성과 유지보수성 고려
   - 타입 제약 복잡성 최소화
   - 컴파일 시간 증가 가능성

10. 제한사항:
    - 메서드에는 타입 매개변수 추가 불가
    - 타입 스위치에서 타입 매개변수 사용 제한
    - 익명 구조체에서 제네릭 사용 불가
*/
