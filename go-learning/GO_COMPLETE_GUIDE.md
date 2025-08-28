# Go 언어 완전 학습 가이드 🚀

이 디렉토리는 **Go 언어의 A부터 Z까지 모든 것**을 다루는 포괄적인 학습 자료입니다.

## 📚 학습 구조 (Study Structure)

### 🎯 기초 단계 (Fundamentals)
```
01-basics/              # 기본 문법과 구조
├── hello.go           # 첫 번째 Go 프로그램
└── syntax.go          # Go 문법 기초

02-data-types/          # 데이터 타입
├── basic-types.go     # 기본 데이터 타입들
└── variables.go       # 변수와 상수

03-control-flow/        # 제어 구조
├── conditionals.go    # 조건문 (if, switch)
└── loops.go           # 반복문 (for)

04-functions/           # 함수
├── basic-functions.go # 기본 함수 사용법
└── advanced-functions.go # 고급 함수 기법
```

### 🏗️ 중급 단계 (Intermediate)
```
05-structs-interfaces/  # 구조체와 인터페이스
├── structs.go         # 구조체 정의와 사용
└── interfaces.go      # 인터페이스와 다형성

06-error-handling/      # 에러 처리
└── basic-errors.go    # Go 에러 처리 패턴

07-concurrency/         # 동시성
├── goroutines.go      # 고루틴 사용법
└── channels.go        # 채널을 통한 통신

08-packages/           # 패키지 시스템
├── packages.go        # 패키지 개념과 구조
└── modules.go         # Go Modules 관리
```

### 🚀 고급 단계 (Advanced)
```
09-testing/            # 테스팅
├── basic-testing.go   # 테스트 대상 함수들
└── basic-testing_test.go # 테스트 코드

10-web-development/     # 웹 개발
└── http-server.go     # HTTP 서버 구축

11-advanced-topics/     # 고급 주제들
├── generics.go        # 제네릭 (Go 1.18+)
├── reflection.go      # 리플렉션
├── context.go         # Context 사용법
└── performance.go     # 성능 최적화

12-collections/        # 컬렉션과 자료구조
├── arrays-slices.go   # 배열과 슬라이스
├── maps.go           # 맵 사용법
└── custom-types.go   # 사용자 정의 타입

13-io-operations/      # 입출력 작업
├── file-operations.go # 파일 처리
├── json-xml.go       # JSON/XML 처리
└── networking.go     # 네트워크 통신

14-performance/        # 성능과 최적화
├── benchmarking.go   # 벤치마킹
├── profiling.go      # 프로파일링
└── memory-management.go # 메모리 관리

15-deployment/         # 배포와 운영
├── docker.go         # Docker 컨테이너화
├── logging.go        # 로깅 시스템
└── monitoring.go     # 모니터링
```

### 🛠️ 실습 프로젝트 (Projects)
```
projects/
├── calculator/        # 계산기 프로그램
├── todo-app/         # TODO 애플리케이션
├── web-server/       # 웹 서버
├── cli-tool/         # 명령줄 도구
└── microservice/     # 마이크로서비스
```

## 🎓 학습 순서 (Learning Path)

### 1단계: 기초 마스터 (1-2주)
1. **01-basics**: Go 설치, 기본 문법, Hello World
2. **02-data-types**: 모든 데이터 타입과 변수 사용법
3. **03-control-flow**: if, switch, for 문 완전 정복
4. **04-functions**: 함수 정의, 호출, 고급 기법

### 2단계: 핵심 개념 (2-3주)  
5. **05-structs-interfaces**: 객체지향적 설계
6. **06-error-handling**: Go만의 에러 처리 방식
7. **07-concurrency**: 고루틴과 채널 (Go의 핵심!)
8. **08-packages**: 모듈 시스템과 의존성 관리

### 3단계: 실무 역량 (3-4주)
9. **09-testing**: 테스트 주도 개발
10. **10-web-development**: HTTP 서버와 API 개발
11. **11-advanced-topics**: 제네릭, 리플렉션 등 고급 기능
12. **12-collections**: 자료구조 활용

### 4단계: 전문가 레벨 (4-6주)
13. **13-io-operations**: 파일, 네트워크, 데이터 처리
14. **14-performance**: 성능 측정과 최적화
15. **15-deployment**: 실제 배포와 운영
16. **projects**: 실무 프로젝트 구현

## 🛠️ 개발 환경 설정

### Go 설치
```bash
# 공식 사이트에서 다운로드
https://golang.org/dl/

# 설치 확인
go version

# 모듈 초기화
go mod init go-learning
```

### 권장 에디터
- **VS Code** + Go Extension
- **GoLand** (JetBrains)
- **Vim/Neovim** + vim-go

### 유용한 도구들
```bash
# 코드 포맷팅
go fmt ./...

# 코드 정리
go mod tidy

# 테스트 실행
go test ./...

# 벤치마크
go test -bench=.

# 레이스 컨디션 검사
go test -race

# 빌드
go build

# 실행
go run main.go
```

## 📖 Go 언어 핵심 특징

### ✨ Go의 장점
- **단순성**: 25개의 키워드만 사용
- **빠른 컴파일**: 초고속 빌드 시간
- **동시성**: 고루틴과 채널로 쉬운 병렬 처리
- **메모리 안전**: 가비지 컬렉터 내장
- **정적 바이너리**: 의존성 없는 단일 실행 파일
- **크로스 플랫폼**: 다양한 OS/아키텍처 지원

### 🎯 주요 사용 영역
- **클라우드 서비스**: Docker, Kubernetes, Terraform
- **웹 서버**: Gin, Echo, Fiber 프레임워크
- **마이크로서비스**: gRPC, REST API
- **CLI 도구**: Cobra, urfave/cli
- **데이터베이스**: InfluxDB, CockroachDB
- **블록체인**: Ethereum, Hyperledger

### 🏢 주요 사용 기업
- **Google**: Kubernetes, Go 언어 개발
- **Docker**: 컨테이너 기술
- **Uber**: 마이크로서비스 아키텍처
- **Netflix**: 클라우드 인프라
- **Dropbox**: 파일 저장 시스템
- **Twitch**: 실시간 스트리밍

## 📚 각 디렉토리별 상세 설명

### 01-basics (기초)
Go 언어의 첫걸음! 설치부터 첫 프로그램 작성까지.
- 패키지 시스템 이해
- 기본 문법과 규칙
- Hello World 프로그램

### 02-data-types (데이터 타입)
Go의 모든 데이터 타입을 완벽 이해.
- 기본 타입: int, float, string, bool
- 복합 타입: array, slice, map, struct
- 포인터와 메모리 관리

### 03-control-flow (제어 구조)
프로그램의 흐름을 제어하는 모든 방법.
- 조건문: if, switch (fallthrough 없음!)
- 반복문: for (유일한 반복문)
- break, continue, goto

### 04-functions (함수)
Go 함수의 모든 것 - 기본부터 고급까지.
- 다중 반환값
- 명명된 반환값
- 가변 매개변수
- 클로저와 고차함수

### 05-structs-interfaces (구조체와 인터페이스)
Go의 타입 시스템과 객체지향 개념.
- 구조체 정의와 임베딩
- 메서드와 리시버
- 인터페이스와 덕 타이핑
- 다형성 구현

### 06-error-handling (에러 처리)
Go만의 독특한 에러 처리 방식.
- error 인터페이스
- 커스텀 에러 타입
- 에러 래핑과 언래핑
- panic과 recover

### 07-concurrency (동시성)
Go의 핵심 기능! 고루틴과 채널.
- 고루틴 생성과 관리
- 채널을 통한 통신
- select 문
- sync 패키지 활용

### 08-packages (패키지)
코드 구성과 모듈 관리.
- 패키지 시스템
- Go Modules
- 의존성 관리
- 버전 관리

### 09-testing (테스팅)
테스트 주도 개발 (TDD) 실습.
- 기본 테스트 작성
- 테이블 기반 테스트
- 벤치마크 테스트
- 테스트 커버리지

### 10-web-development (웹 개발)
Go로 웹 애플리케이션 개발.
- HTTP 서버 구축
- REST API 개발
- 미들웨어 작성
- 템플릿 엔진

## 🚀 실습 프로젝트 안내

각 프로젝트는 학습한 내용을 실제로 적용해보는 종합 실습입니다.

### calculator/ - 계산기
- 기본 문법과 함수 활용
- 명령줄 인터페이스
- 에러 처리

### todo-app/ - TODO 애플리케이션  
- 구조체와 메서드
- 파일 I/O
- JSON 처리

### web-server/ - 웹 서버
- HTTP 서버 구현
- REST API 설계
- 동시성 처리

### cli-tool/ - 명령줄 도구
- 플래그 처리
- 파일 시스템 조작
- 외부 명령 실행

### microservice/ - 마이크로서비스
- gRPC 통신
- 데이터베이스 연동
- Docker 컨테이너화

## 📝 학습 팁

### 🎯 효과적인 학습 방법
1. **코드를 직접 타이핑**: 복사-붙여넣기 금지!
2. **작은 변경 실험**: 코드를 수정해보며 결과 관찰
3. **주석 읽기**: 각 예제의 상세한 설명 숙지
4. **오류 경험**: 의도적으로 오류를 만들어 보기
5. **문서 참고**: 공식 문서와 함께 학습

### 📚 추천 학습 리소스
- **공식 문서**: https://golang.org/doc/
- **Go Tour**: https://tour.golang.org/
- **Effective Go**: https://golang.org/doc/effective_go.html
- **Go Blog**: https://blog.golang.org/
- **Go Playground**: https://play.golang.org/

### 🔧 디버깅 및 문제 해결
```bash
# 문법 검사
go vet ./...

# 의존성 그래프
go mod graph

# 빌드 태그 확인
go list -tags

# 환경 변수 확인  
go env

# 도움말
go help [command]
```

## 🎖️ 학습 완료 체크리스트

### 기초 레벨 ✅
- [ ] Go 설치 및 환경 설정
- [ ] Hello World 프로그램 작성
- [ ] 모든 데이터 타입 사용해보기
- [ ] 함수 정의하고 호출하기
- [ ] 조건문과 반복문 활용

### 중급 레벨 ✅  
- [ ] 구조체와 메서드 정의
- [ ] 인터페이스 구현
- [ ] 에러 처리 패턴 적용
- [ ] 고루틴과 채널 사용
- [ ] 패키지 만들고 관리하기

### 고급 레벨 ✅
- [ ] 테스트 코드 작성
- [ ] HTTP 서버 구축
- [ ] 제네릭 활용
- [ ] 성능 최적화
- [ ] 실제 프로젝트 배포

## 🤝 커뮤니티와 도움

### 한국 Go 커뮤니티
- **Go Korea**: Facebook 그룹
- **GDG Golang Korea**: 구글 개발자 그룹
- **Discord/Slack**: 실시간 질의응답

### 국제 커뮤니티
- **Go Forum**: https://forum.golangbridge.org/
- **Reddit**: r/golang
- **Stack Overflow**: golang 태그

---

## 🎉 시작하기

지금 당장 `01-basics/hello.go`부터 시작하세요!

```bash
cd 01-basics
go run hello.go
```

**Go 언어와 함께하는 멋진 여정을 시작합니다! 🚀**

---

*이 가이드는 Go 언어를 처음 배우는 초보자부터 실무에서 활용하고자 하는 개발자까지 모든 레벨을 대상으로 합니다. 각자의 속도에 맞춰 천천히 학습하되, 반드시 코드를 직접 작성하며 실습하시기 바랍니다.*