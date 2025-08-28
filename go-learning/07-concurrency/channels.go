// channels.go - Go의 채널 (Channels)
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("=== Go 채널 (Channels) ===")

	// 1. 기본 채널 사용법
	fmt.Println("\n1. 기본 채널:")

	// 채널 생성
	messages := make(chan string)

	// 고루틴에서 채널로 데이터 전송
	go func() {
		messages <- "안녕하세요!" // 채널에 데이터 전송
	}()

	// 채널에서 데이터 수신
	msg := <-messages
	fmt.Printf("받은 메시지: %s\n", msg)

	// 2. 버퍼드 채널 (Buffered Channel)
	fmt.Println("\n2. 버퍼드 채널:")

	// 버퍼 크기가 3인 채널 생성
	buffered := make(chan int, 3)

	// 고루틴 없이도 버퍼 크기만큼 전송 가능
	buffered <- 1
	buffered <- 2
	buffered <- 3

	// 데이터 수신
	fmt.Printf("버퍼드 채널에서 수신: %d\n", <-buffered)
	fmt.Printf("버퍼드 채널에서 수신: %d\n", <-buffered)
	fmt.Printf("버퍼드 채널에서 수신: %d\n", <-buffered)

	// 3. 채널 방향성 (단방향 채널)
	fmt.Println("\n3. 채널 방향성:")

	numbers := make(chan int)
	go sendNumbers(numbers) // 전송 전용 채널로 전달
	receiveNumbers(numbers) // 수신 전용 채널로 전달

	// 4. 채널 닫기
	fmt.Println("\n4. 채널 닫기:")

	jobs := make(chan int, 3)
	done := make(chan bool)

	go func() {
		for {
			job, ok := <-jobs
			if !ok {
				fmt.Println("모든 작업 완료")
				done <- true
				return
			}
			fmt.Printf("작업 처리: %d\n", job)
		}
	}()

	// 작업 전송
	for i := 1; i <= 3; i++ {
		jobs <- i
	}
	close(jobs) // 채널 닫기

	<-done // 완료 대기

	// 5. range를 사용한 채널 순회
	fmt.Println("\n5. range로 채널 순회:")

	queue := make(chan string, 3)
	queue <- "첫 번째"
	queue <- "두 번째"
	queue <- "세 번째"
	close(queue)

	// 채널이 닫힐 때까지 모든 값 수신
	for item := range queue {
		fmt.Printf("큐에서 수신: %s\n", item)
	}

	// 6. select 문 - 다중 채널 처리
	fmt.Println("\n6. select 문:")

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "채널1에서 메시지"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "채널2에서 메시지"
	}()

	// 여러 채널 중 준비된 채널에서 먼저 수신
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Printf("ch1 수신: %s\n", msg1)
		case msg2 := <-ch2:
			fmt.Printf("ch2 수신: %s\n", msg2)
		}
	}

	// 7. select의 default 케이스
	fmt.Println("\n7. select default 케이스:")

	nonBlockingCh := make(chan string, 1)

	select {
	case msg := <-nonBlockingCh:
		fmt.Printf("수신된 메시지: %s\n", msg)
	default:
		fmt.Println("수신할 메시지가 없음 (논블로킹)")
	}

	nonBlockingCh <- "이제 메시지가 있음"

	select {
	case msg := <-nonBlockingCh:
		fmt.Printf("수신된 메시지: %s\n", msg)
	default:
		fmt.Println("수신할 메시지가 없음")
	}

	// 8. 타임아웃과 함께 사용하는 select
	fmt.Println("\n8. 타임아웃이 있는 select:")

	slowCh := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		slowCh <- "느린 응답"
	}()

	select {
	case msg := <-slowCh:
		fmt.Printf("응답 수신: %s\n", msg)
	case <-time.After(2 * time.Second):
		fmt.Println("타임아웃! 응답이 너무 늦음")
	}

	// 9. 워커 풀 패턴
	fmt.Println("\n9. 워커 풀 패턴:")

	const numWorkers = 3
	const numJobs = 10

	jobQueue := make(chan Job, numJobs)
	resultQueue := make(chan Result, numJobs)

	// 워커들 시작
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobQueue, resultQueue)
	}

	// 작업들 생성
	for j := 1; j <= numJobs; j++ {
		job := Job{ID: j, Data: fmt.Sprintf("작업데이터-%d", j)}
		jobQueue <- job
	}
	close(jobQueue)

	// 결과 수집
	for r := 1; r <= numJobs; r++ {
		result := <-resultQueue
		fmt.Printf("결과: ID=%d, 처리결과=%s\n", result.JobID, result.Output)
	}

	// 10. 생산자-소비자 패턴
	fmt.Println("\n10. 생산자-소비자 패턴:")

	dataCh := make(chan int, 5)
	quit := make(chan bool)

	// 생산자
	go producer(dataCh)

	// 소비자
	go consumer(dataCh, quit)

	<-quit // 소비자 완료 대기

	// 11. 파이프라인 패턴
	fmt.Println("\n11. 파이프라인 패턴:")

	// 숫자 생성 -> 제곱 -> 필터링
	numbersCh := generateNumbers(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	squaresCh := squareNumbers(numbersCh)
	evensCh := filterEvenNumbers(squaresCh)

	fmt.Println("파이프라인 결과 (짝수 제곱수들):")
	for result := range evensCh {
		fmt.Printf("%d ", result)
	}
	fmt.Println()

	// 12. 채널을 활용한 실무 패턴들
	fmt.Println("\n12. 실무 패턴들:")

	// 팬아웃 패턴 (작업 분산)
	demonstrateFanOut()

	// 팬인 패턴 (결과 수집)
	demonstrateFanIn()

	// 취소 가능한 작업
	demonstrateCancellation()

	// 13. 채널 기반 이벤트 시스템
	fmt.Println("\n13. 이벤트 시스템:")

	eventBus := NewEventBus()

	// 이벤트 리스너 등록
	eventBus.Subscribe("user.login", func(data interface{}) {
		fmt.Printf("사용자 로그인 이벤트: %v\n", data)
	})

	eventBus.Subscribe("user.logout", func(data interface{}) {
		fmt.Printf("사용자 로그아웃 이벤트: %v\n", data)
	})

	// 이벤트 발생
	eventBus.Publish("user.login", "김사용자")
	eventBus.Publish("user.logout", "김사용자")

	time.Sleep(100 * time.Millisecond) // 이벤트 처리 대기
}

// 3. 채널 방향성 함수들
func sendNumbers(ch chan<- int) { // 전송 전용 채널
	for i := 1; i <= 5; i++ {
		ch <- i
		fmt.Printf("전송: %d\n", i)
	}
	close(ch)
}

func receiveNumbers(ch <-chan int) { // 수신 전용 채널
	for num := range ch {
		fmt.Printf("수신: %d\n", num)
	}
}

// 9. 워커 풀을 위한 구조체들
type Job struct {
	ID   int
	Data string
}

type Result struct {
	JobID  int
	Output string
}

func worker(id int, jobs <-chan Job, results chan<- Result) {
	for job := range jobs {
		fmt.Printf("워커 %d가 작업 %d 처리 시작\n", id, job.ID)

		// 작업 시뮬레이션
		time.Sleep(time.Duration(rand.Intn(1000)+500) * time.Millisecond)

		result := Result{
			JobID:  job.ID,
			Output: fmt.Sprintf("처리됨-%s", job.Data),
		}

		results <- result
		fmt.Printf("워커 %d가 작업 %d 처리 완료\n", id, job.ID)
	}
}

// 10. 생산자-소비자 함수들
func producer(dataCh chan<- int) {
	defer close(dataCh)

	for i := 1; i <= 10; i++ {
		fmt.Printf("생산: %d\n", i)
		dataCh <- i
		time.Sleep(200 * time.Millisecond)
	}
	fmt.Println("생산자 종료")
}

func consumer(dataCh <-chan int, quit chan<- bool) {
	for {
		data, ok := <-dataCh
		if !ok {
			fmt.Println("소비자 종료")
			quit <- true
			return
		}

		fmt.Printf("소비: %d\n", data)
		time.Sleep(300 * time.Millisecond)
	}
}

// 11. 파이프라인 함수들
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

func filterEvenNumbers(in <-chan int) <-chan int {
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

// 12. 실무 패턴 데모 함수들
func demonstrateFanOut() {
	fmt.Println("팬아웃 패턴 (작업 분산):")

	work := make(chan int, 10)

	// 작업 생성
	go func() {
		defer close(work)
		for i := 1; i <= 10; i++ {
			work <- i
		}
	}()

	// 여러 워커가 동시에 처리
	const numWorkers = 3
	results := make(chan string, 10)

	for i := 0; i < numWorkers; i++ {
		go func(workerID int) {
			for job := range work {
				result := fmt.Sprintf("워커%d가 작업%d 완료", workerID, job)
				results <- result
			}
		}(i + 1)
	}

	// 결과 수집 (별도 고루틴)
	go func() {
		time.Sleep(2 * time.Second) // 모든 작업 완료 대기
		close(results)
	}()

	for result := range results {
		fmt.Printf("  %s\n", result)
	}
}

func demonstrateFanIn() {
	fmt.Println("팬인 패턴 (결과 수집):")

	// 여러 소스에서 데이터 생성
	source1 := make(chan string)
	source2 := make(chan string)
	source3 := make(chan string)

	go func() {
		defer close(source1)
		for i := 1; i <= 3; i++ {
			source1 <- fmt.Sprintf("소스1-데이터%d", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		defer close(source2)
		for i := 1; i <= 3; i++ {
			source2 <- fmt.Sprintf("소스2-데이터%d", i)
			time.Sleep(150 * time.Millisecond)
		}
	}()

	go func() {
		defer close(source3)
		for i := 1; i <= 3; i++ {
			source3 <- fmt.Sprintf("소스3-데이터%d", i)
			time.Sleep(200 * time.Millisecond)
		}
	}()

	// 모든 소스의 결과를 하나의 채널로 병합
	merged := fanIn(source1, source2, source3)

	for data := range merged {
		fmt.Printf("  병합된 데이터: %s\n", data)
	}
}

func fanIn(channels ...<-chan string) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)

		// 모든 채널이 닫힐 때까지 수신
		for len(channels) > 0 {
			// select로 준비된 채널에서 수신
			switch len(channels) {
			case 3:
				select {
				case data, ok := <-channels[0]:
					if ok {
						out <- data
					} else {
						channels = channels[1:] // 닫힌 채널 제거
					}
				case data, ok := <-channels[1]:
					if ok {
						out <- data
					} else {
						channels = append(channels[:1], channels[2:]...)
					}
				case data, ok := <-channels[2]:
					if ok {
						out <- data
					} else {
						channels = channels[:2]
					}
				}
			case 2:
				select {
				case data, ok := <-channels[0]:
					if ok {
						out <- data
					} else {
						channels = channels[1:]
					}
				case data, ok := <-channels[1]:
					if ok {
						out <- data
					} else {
						channels = channels[:1]
					}
				}
			case 1:
				if data, ok := <-channels[0]; ok {
					out <- data
				} else {
					channels = channels[:0]
				}
			}
		}
	}()

	return out
}

func demonstrateCancellation() {
	fmt.Println("취소 가능한 작업:")

	cancel := make(chan struct{})
	done := make(chan bool)

	// 장시간 실행되는 작업
	go func() {
		defer func() { done <- true }()

		ticker := time.NewTicker(200 * time.Millisecond)
		defer ticker.Stop()

		for i := 1; i <= 10; i++ {
			select {
			case <-cancel:
				fmt.Printf("  작업이 %d단계에서 취소됨\n", i)
				return
			case <-ticker.C:
				fmt.Printf("  작업 진행: %d/10\n", i)
			}
		}
		fmt.Println("  작업 완료!")
	}()

	// 3초 후 취소
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("  작업 취소 신호 전송")
		close(cancel)
	}()

	<-done
}

// 13. 이벤트 시스템
type EventBus struct {
	subscribers map[string][]chan interface{}
}

func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[string][]chan interface{}),
	}
}

func (eb *EventBus) Subscribe(eventType string, handler func(interface{})) {
	ch := make(chan interface{}, 1)
	eb.subscribers[eventType] = append(eb.subscribers[eventType], ch)

	// 이벤트 처리 고루틴
	go func() {
		for data := range ch {
			handler(data)
		}
	}()
}

func (eb *EventBus) Publish(eventType string, data interface{}) {
	if channels, exists := eb.subscribers[eventType]; exists {
		for _, ch := range channels {
			select {
			case ch <- data:
			default:
				fmt.Printf("이벤트 드롭됨: %s\n", eventType)
			}
		}
	}
}

/*
Go 채널 핵심 개념:

1. 채널이란?
   - 고루틴 간 통신을 위한 파이프라인
   - "공유 메모리로 통신하지 말고, 통신으로 메모리를 공유하라"
   - 타입 안전성 보장
   - 동기화 메커니즘 제공

2. 채널 종류:
   - 언버퍼드 채널: make(chan Type)
     * 동기적 통신 (송신자는 수신자가 받을 때까지 블록)
   - 버퍼드 채널: make(chan Type, size)
     * 비동기적 통신 (버퍼가 가득 찰 때까지 블록 안 됨)

3. 채널 방향성:
   - 양방향: chan Type
   - 송신 전용: chan<- Type
   - 수신 전용: <-chan Type

4. 채널 연산:
   - 전송: ch <- value
   - 수신: value := <-ch
   - 수신 (ok 패턴): value, ok := <-ch
   - 닫기: close(ch)
   - 방향 제한으로 컴파일 타임 안전성 확보

5. select 문:
   - 여러 채널 작업 중 하나 선택
   - default 케이스로 논블로킹 처리
   - time.After()로 타임아웃 구현
   - 랜덤하게 준비된 케이스 선택

6. 채널 패턴:
   - 워커 풀: 작업 분산 처리
   - 파이프라인: 순차적 데이터 변환
   - 팬아웃: 작업을 여러 워커에 분산
   - 팬인: 여러 소스의 결과를 하나로 수집
   - 생산자-소비자: 데이터 생성과 소비 분리

7. 채널 사용 원칙:
   - 송신자가 채널을 닫는다
   - 닫힌 채널에서는 수신만 가능
   - 닫힌 채널에 송신하면 패닉
   - nil 채널은 항상 블록

8. 성능 고려사항:
   - 버퍼 크기 최적화
   - 불필요한 고루틴 생성 방지
   - 채널 누수 방지 (적절한 닫기)
   - select문에서 default 케이스 활용

9. 디버깅 팁:
   - 데드락 방지 (순환 대기 상황 주의)
   - 고루틴 누수 탐지
   - 채널 버퍼 크기 모니터링
   - race detector 활용

10. 실무 활용:
    - HTTP 요청 처리 풀
    - 이벤트 기반 아키텍처
    - 백그라운드 작업 처리
    - 실시간 데이터 스트리밍
    - 마이크로서비스 간 통신
*/
