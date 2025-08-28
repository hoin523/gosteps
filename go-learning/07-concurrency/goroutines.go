// goroutines.go - Go의 고루틴 (Goroutines)
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== Go 고루틴 (Goroutines) ===")

	// 1. 기본 고루틴 사용법
	fmt.Println("\n1. 기본 고루틴:")

	// 메인 스레드에서 실행
	sayHello("메인 스레드")

	// 고루틴으로 실행 (go 키워드 사용)
	go sayHello("고루틴 1")
	go sayHello("고루틴 2")
	go sayHello("고루틴 3")

	// 고루틴이 실행될 시간을 주기 위해 잠시 대기
	time.Sleep(2 * time.Second)

	// 2. 익명 함수를 고루틴으로 실행
	fmt.Println("\n2. 익명 함수 고루틴:")

	for i := 1; i <= 3; i++ {
		go func(id int) {
			fmt.Printf("익명 고루틴 %d 실행\n", id)
		}(i) // 매개변수로 i 전달 (클로저 문제 방지)
	}

	time.Sleep(1 * time.Second)

	// 3. WaitGroup을 사용한 고루틴 동기화
	fmt.Println("\n3. WaitGroup 사용:")

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1) // 고루틴 카운터 증가

		go func(workerID int) {
			defer wg.Done() // 고루틴 완료 시 카운터 감소
			doWork(workerID)
		}(i)
	}

	wg.Wait() // 모든 고루틴이 완료될 때까지 대기
	fmt.Println("모든 워커 완료!")

	// 4. 고루틴과 데이터 레이스 문제
	fmt.Println("\n4. 데이터 레이스 문제:")

	// 문제가 있는 코드 (데이터 레이스)
	counter := 0

	for i := 0; i < 10; i++ {
		go func() {
			counter++ // 동시 접근 문제
		}()
	}

	time.Sleep(1 * time.Second)
	fmt.Printf("데이터 레이스가 있는 카운터: %d (예상: 10)\n", counter)

	// 5. Mutex를 사용한 동기화
	fmt.Println("\n5. Mutex로 동기화:")

	var mu sync.Mutex
	var safeCounter int

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock() // 락 획득
			safeCounter++
			mu.Unlock() // 락 해제
		}()
	}

	wg.Wait()
	fmt.Printf("Mutex로 보호된 카운터: %d\n", safeCounter)

	// 6. 고루틴 풀 패턴
	fmt.Println("\n6. 고루틴 풀:")

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 워커 고루틴들 시작
	const numWorkers = 3
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	// 작업들 전송
	const numJobs = 9
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// 결과 수집
	for r := 1; r <= numJobs; r++ {
		<-results
	}

	// 7. 고루틴 수명주기와 누수 방지
	fmt.Println("\n7. 고루틴 수명주기:")

	fmt.Printf("현재 고루틴 수: %d\n", runtime.NumGoroutine())

	// 컨텍스트를 사용한 고루틴 종료
	stopCh := make(chan struct{})

	go longRunningTask("작업1", stopCh)
	go longRunningTask("작업2", stopCh)

	time.Sleep(3 * time.Second)
	fmt.Println("모든 작업 중지 신호 전송")
	close(stopCh) // 모든 고루틴에 중지 신호

	time.Sleep(1 * time.Second)
	fmt.Printf("중지 후 고루틴 수: %d\n", runtime.NumGoroutine())

	// 8. 팬아웃/팬인 패턴
	fmt.Println("\n8. 팬아웃/팬인 패턴:")

	in := generateNumbers(1, 2, 3, 4, 5)

	// 팬아웃: 여러 고루틴으로 작업 분산
	c1 := square(in)
	c2 := square(in)

	// 팬인: 여러 채널의 결과를 하나로 병합
	for n := range merge(c1, c2) {
		fmt.Printf("제곱 결과: %d\n", n)
	}

	// 9. 파이프라인 패턴
	fmt.Println("\n9. 파이프라인 패턴:")

	numbers := generateSequence(10)
	squares := squareNumbers(numbers)
	filtered := filterEven(squares)

	for result := range filtered {
		fmt.Printf("파이프라인 결과: %d\n", result)
	}

	// 10. 실무 예제: 병렬 웹 크롤링 시뮬레이션
	fmt.Println("\n10. 병렬 웹 크롤링:")

	urls := []string{
		"https://example.com/1",
		"https://example.com/2",
		"https://example.com/3",
		"https://example.com/4",
		"https://example.com/5",
	}

	results := make(chan CrawlResult, len(urls))

	// 병렬로 웹페이지 크롤링
	for _, url := range urls {
		go crawlURL(url, results)
	}

	// 결과 수집
	for i := 0; i < len(urls); i++ {
		result := <-results
		if result.Error != nil {
			fmt.Printf("크롤링 실패 %s: %v\n", result.URL, result.Error)
		} else {
			fmt.Printf("크롤링 성공 %s: %d bytes\n", result.URL, result.Size)
		}
	}

	// 11. 고루틴 최적화 팁들
	fmt.Println("\n11. 고루틴 최적화:")

	demonstrateGoroutineOptimization()
}

// 1. 기본 함수들
func sayHello(name string) {
	for i := 1; i <= 3; i++ {
		fmt.Printf("%s: 안녕하세요! (%d/3)\n", name, i)
		time.Sleep(500 * time.Millisecond)
	}
}

func doWork(workerID int) {
	fmt.Printf("워커 %d 작업 시작\n", workerID)

	// 실제 작업 시뮬레이션
	duration := time.Duration(rand.Intn(3)+1) * time.Second
	time.Sleep(duration)

	fmt.Printf("워커 %d 작업 완료 (소요시간: %v)\n", workerID, duration)
}

// 6. 워커 함수
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("워커 %d 작업 %d 시작\n", id, job)
		time.Sleep(time.Second) // 작업 시뮬레이션

		result := job * 2 // 간단한 처리
		results <- result

		fmt.Printf("워커 %d 작업 %d 완료 (결과: %d)\n", id, job, result)
	}
}

// 7. 장시간 실행되는 작업
func longRunningTask(taskName string, stopCh <-chan struct{}) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-stopCh:
			fmt.Printf("%s 중지됨\n", taskName)
			return
		case <-ticker.C:
			fmt.Printf("%s 실행 중...\n", taskName)
		}
	}
}

// 8. 팬아웃/팬인 패턴 함수들
func generateNumbers(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)

	var wg sync.WaitGroup
	wg.Add(len(cs))

	for _, c := range cs {
		go func(ch <-chan int) {
			defer wg.Done()
			for n := range ch {
				out <- n
			}
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// 9. 파이프라인 패턴 함수들
func generateSequence(max int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 1; i <= max; i++ {
			out <- i
		}
	}()
	return out
}

func squareNumbers(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()
	return out
}

func filterEven(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			if n%2 == 0 {
				out <- n
			}
		}
	}()
	return out
}

// 10. 웹 크롤링 시뮬레이션
type CrawlResult struct {
	URL   string
	Size  int
	Error error
}

func crawlURL(url string, results chan<- CrawlResult) {
	// 실제 HTTP 요청 대신 시뮬레이션
	time.Sleep(time.Duration(rand.Intn(2)+1) * time.Second)

	// 랜덤하게 성공/실패 결정
	if rand.Intn(10) < 2 { // 20% 확률로 실패
		results <- CrawlResult{
			URL:   url,
			Error: fmt.Errorf("연결 시간 초과"),
		}
	} else {
		results <- CrawlResult{
			URL:  url,
			Size: rand.Intn(10000) + 1000, // 1000-11000 bytes
		}
	}
}

// 11. 고루틴 최적화 데모
func demonstrateGoroutineOptimization() {
	fmt.Println("고루틴 최적화 기법들:")

	// 1. 적절한 버퍼 크기
	buffered := make(chan int, 10) // 버퍼드 채널
	go func() {
		for i := 0; i < 5; i++ {
			buffered <- i
		}
		close(buffered)
	}()

	for val := range buffered {
		fmt.Printf("버퍼드 채널에서 수신: %d\n", val)
	}

	// 2. 고루틴 수 제한
	const maxGoroutines = 3
	semaphore := make(chan struct{}, maxGoroutines)

	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			semaphore <- struct{}{}        // 세마포어 획득
			defer func() { <-semaphore }() // 세마포어 해제

			fmt.Printf("제한된 고루틴 %d 실행\n", id)
			time.Sleep(100 * time.Millisecond)
		}(i)
	}
	wg.Wait()

	// 3. 고루틴 풀 재사용
	pool := &sync.Pool{
		New: func() interface{} {
			return make([]byte, 1024) // 1KB 버퍼
		},
	}

	// 풀에서 버퍼 가져오기
	buffer := pool.Get().([]byte)
	fmt.Printf("풀에서 가져온 버퍼 크기: %d bytes\n", len(buffer))

	// 풀에 버퍼 반납
	pool.Put(buffer)

	fmt.Printf("최종 고루틴 수: %d\n", runtime.NumGoroutine())
}

/*
Go 고루틴 핵심 개념:

1. 고루틴이란?
   - 경량 스레드, OS 스레드보다 적은 메모리 사용 (2KB 스택)
   - Go 런타임이 관리하는 사용자 레벨 스레드
   - 수백만 개의 고루틴을 동시에 실행 가능

2. 고루틴 vs 스레드:
   - 고루틴: 2KB 초기 스택, 동적 증가
   - OS 스레드: 1-8MB 고정 스택
   - 컨텍스트 스위칭 비용이 매우 낮음

3. 고루틴 생성:
   - go 키워드로 간단하게 생성
   - 함수, 메서드, 익명 함수 모두 가능
   - main 함수도 고루틴으로 실행됨

4. 동기화 방법:
   - WaitGroup: 여러 고루틴의 완료 대기
   - Mutex: 공유 자원의 동시 접근 제어
   - 채널: 고루틴 간 통신과 동기화
   - Select: 여러 채널 작업 중 하나 선택

5. 고루틴 패턴:
   - 워커 풀: 제한된 수의 워커로 작업 처리
   - 팬아웃/팬인: 작업 분산 후 결과 수집
   - 파이프라인: 단계적 데이터 처리
   - 생산자/소비자: 데이터 생성과 소비 분리

6. 주의사항:
   - 고루틴 누수 방지 (적절한 종료 처리)
   - 데이터 레이스 방지 (동기화 필요)
   - 무한정 고루틴 생성 방지 (세마포어 활용)
   - 메인 함수 종료 시 모든 고루틴 종료

7. 최적화 기법:
   - 버퍼드 채널로 블로킹 최소화
   - sync.Pool로 객체 재사용
   - 적절한 고루틴 수 제한
   - 컨텍스트로 취소 및 타임아웃 처리

8. 디버깅:
   - runtime.NumGoroutine()으로 고루틴 수 확인
   - go tool trace로 고루틴 분석
   - race detector로 데이터 레이스 탐지 (go run -race)
*/
