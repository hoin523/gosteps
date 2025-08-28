// modules.go - Go Modules 심화 가이드
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("=== Go Modules 심화 가이드 ===")

	// 1. Go Modules 개요
	fmt.Println("\n1. Go Modules 개요:")
	demonstrateModulesOverview()

	// 2. 모듈 생성과 관리
	fmt.Println("\n2. 모듈 생성과 관리:")
	demonstrateModuleCreation()

	// 3. 의존성 관리
	fmt.Println("\n3. 의존성 관리:")
	demonstrateDependencyManagement()

	// 4. 버전 관리
	fmt.Println("\n4. 버전 관리:")
	demonstrateVersionManagement()

	// 5. go.mod 파일 상세
	fmt.Println("\n5. go.mod 파일 상세:")
	demonstrateGoModFile()

	// 6. go.sum 파일과 보안
	fmt.Println("\n6. go.sum 파일과 보안:")
	demonstrateGoSumFile()

	// 7. 프록시와 체크섬 데이터베이스
	fmt.Println("\n7. 프록시와 체크섬 데이터베이스:")
	demonstrateProxyAndChecksum()

	// 8. 워크스페이스 모드
	fmt.Println("\n8. 워크스페이스 모드 (Go 1.18+):")
	demonstrateWorkspaceMode()

	// 9. 모듈 베스트 프랙티스
	fmt.Println("\n9. 모듈 베스트 프랙티스:")
	demonstrateBestPractices()

	// 10. 문제 해결
	fmt.Println("\n10. 문제 해결:")
	demonstrateTroubleshooting()
}

func demonstrateModulesOverview() {
	fmt.Println("Go Modules이란?")
	fmt.Println("- Go 1.11부터 도입된 의존성 관리 시스템")
	fmt.Println("- GOPATH에 의존하지 않는 패키지 관리")
	fmt.Println("- 버전 관리와 재현 가능한 빌드 지원")
	fmt.Println("- 분산된 버전 제어 시스템과 통합")
	fmt.Println()

	fmt.Println("주요 개념:")
	concepts := map[string]string{
		"Module":      "관련된 Go 패키지들의 컬렉션",
		"go.mod":      "모듈의 루트에 있는 모듈 정의 파일",
		"go.sum":      "의존성의 체크섬을 저장하는 파일",
		"Module Path": "모듈의 고유 식별자 (보통 VCS URL)",
		"Version":     "모듈의 특정 버전 (태그, 브랜치, 커밋)",
	}

	for term, desc := range concepts {
		fmt.Printf("  %-12s: %s\n", term, desc)
	}
}

func demonstrateModuleCreation() {
	fmt.Println("모듈 생성과 관리 명령어:")
	fmt.Println()

	fmt.Println("1. 새 모듈 초기화:")
	fmt.Println("   go mod init [module-path]")
	fmt.Println("   예: go mod init github.com/username/myproject")
	fmt.Println()

	fmt.Println("2. 기존 프로젝트를 모듈로 변환:")
	fmt.Println("   cd existing-project")
	fmt.Println("   go mod init [module-path]")
	fmt.Println("   go mod tidy")
	fmt.Println()

	fmt.Println("3. 모듈 정보 확인:")
	fmt.Println("   go list -m all          # 모든 의존성 나열")
	fmt.Println("   go mod graph            # 의존성 그래프 표시")
	fmt.Println("   go mod why [package]    # 패키지가 필요한 이유")
	fmt.Println()

	// 실제 go.mod 파일 예제 생성
	createExampleGoMod()
}

func demonstrateDependencyManagement() {
	fmt.Println("의존성 관리 명령어:")
	fmt.Println()

	fmt.Println("1. 의존성 추가:")
	fmt.Println("   go get [package]                    # 최신 버전")
	fmt.Println("   go get [package]@latest            # 최신 버전 명시")
	fmt.Println("   go get [package]@v1.2.3           # 특정 버전")
	fmt.Println("   go get [package]@master           # 특정 브랜치")
	fmt.Println()

	fmt.Println("2. 의존성 업데이트:")
	fmt.Println("   go get -u [package]                # 패키지 업데이트")
	fmt.Println("   go get -u ./...                    # 모든 의존성 업데이트")
	fmt.Println("   go get -u=patch ./...              # 패치 버전만 업데이트")
	fmt.Println()

	fmt.Println("3. 의존성 제거:")
	fmt.Println("   go mod tidy                        # 불필요한 의존성 제거")
	fmt.Println("   # 코드에서 사용을 제거한 후 go mod tidy 실행")
	fmt.Println()

	fmt.Println("4. 의존성 다운로드:")
	fmt.Println("   go mod download                    # 모든 의존성 다운로드")
	fmt.Println("   go mod download [package]          # 특정 패키지 다운로드")
	fmt.Println()

	fmt.Println("5. 로컬 개발:")
	fmt.Println("   go mod edit -replace=github.com/original/pkg=../local/pkg")
	fmt.Println("   # 로컬 경로로 패키지 대체")
}

func demonstrateVersionManagement() {
	fmt.Println("버전 관리 및 Semantic Versioning:")
	fmt.Println()

	fmt.Println("버전 형식: vMAJOR.MINOR.PATCH")
	fmt.Println("  MAJOR: 호환되지 않는 API 변경")
	fmt.Println("  MINOR: 이전 버전과 호환되는 기능 추가")
	fmt.Println("  PATCH: 이전 버전과 호환되는 버그 수정")
	fmt.Println()

	fmt.Println("버전 선택 예제:")
	versionExamples := []string{
		"go get package@v1.2.3      # 정확한 버전",
		"go get package@v1.2        # v1.2.x 중 최신",
		"go get package@v1          # v1.x.x 중 최신",
		"go get package@latest      # 최신 버전",
		"go get package@upgrade     # 업그레이드 가능한 최신",
		"go get package@patch       # 패치 버전만",
		"go get package@master      # master 브랜치",
		"go get package@commit-hash # 특정 커밋",
	}

	for _, example := range versionExamples {
		fmt.Printf("  %s\n", example)
	}
	fmt.Println()

	fmt.Println("주 버전(Major Version) 관리:")
	fmt.Println("  v0: 개발 버전 (호환성 보장 없음)")
	fmt.Println("  v1: 안정 버전 (하위 호환성 보장)")
	fmt.Println("  v2+: /v2, /v3 등을 모듈 경로에 추가")
	fmt.Println("       예: github.com/user/repo/v2")
}

func demonstrateGoModFile() {
	fmt.Println("go.mod 파일 상세 구조:")
	fmt.Println()

	goModExample := `module github.com/username/myproject

go 1.21

require (
    github.com/gin-gonic/gin v1.9.1
    github.com/lib/pq v1.10.9
    github.com/stretchr/testify v1.8.4
)

require (
    // indirect dependencies (자동 관리)
    github.com/bytedance/sonic v1.9.1 // indirect
    github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311 // indirect
    github.com/gabriel-vasile/mimetype v1.4.2 // indirect
)

replace (
    // 로컬 개발이나 포크된 버전 사용
    github.com/original/package => ../local/package
    github.com/broken/package => github.com/fixed/package v1.2.3
)

exclude (
    // 특정 버전 제외 (보안상 문제가 있는 버전 등)
    github.com/problematic/package v1.0.0
    github.com/vulnerable/package v2.1.0
)

retract (
    // 자신의 모듈의 특정 버전을 철회
    v1.0.0 // 중대한 버그가 있는 버전
    [v1.1.0, v1.2.0] // 버전 범위 철회
)`

	fmt.Println(goModExample)
	fmt.Println()

	fmt.Println("go.mod 편집 명령어:")
	editCommands := []string{
		"go mod edit -module=new-name                    # 모듈명 변경",
		"go mod edit -require=package@version           # require 추가",
		"go mod edit -droprequire=package               # require 제거",
		"go mod edit -replace=old=new                   # replace 추가",
		"go mod edit -dropreplace=old                   # replace 제거",
		"go mod edit -exclude=package@version           # exclude 추가",
		"go mod edit -dropexclude=package@version       # exclude 제거",
		"go mod edit -retract=version                   # retract 추가",
		"go mod edit -dropretract=version               # retract 제거",
	}

	for _, cmd := range editCommands {
		fmt.Printf("  %s\n", cmd)
	}
}

func demonstrateGoSumFile() {
	fmt.Println("go.sum 파일과 보안:")
	fmt.Println()

	fmt.Println("go.sum의 역할:")
	fmt.Println("- 의존성의 암호화 체크섬 저장")
	fmt.Println("- 의존성 변조 탐지")
	fmt.Println("- 재현 가능한 빌드 보장")
	fmt.Println("- 버전 관리에 포함되어야 함")
	fmt.Println()

	fmt.Println("go.sum 파일 형식:")
	goSumExample := `github.com/gin-gonic/gin v1.9.1 h1:4idEAncQnU5cB7BeOkPtxjfCSye0AAm1R0RVIqJ+Jmg=
github.com/gin-gonic/gin v1.9.1/go.mod h1:hPrL7YrpYKXt5YId3A/Tnip5kqbEAP+KLuI3SUcPTeU=
github.com/lib/pq v1.10.9 h1:YXG7RB+JIjhP29X+OtkiDnYaXQwpS4JEWq7dtCCRUEw=
github.com/lib/pq v1.10.9/go.mod h1:AlVN5x4E4T544tWzH6hKfbfQvm3HdbOxrmggDNAPY9o=`

	fmt.Println(goSumExample)
	fmt.Println()

	fmt.Println("체크섬 검증:")
	fmt.Println("  go mod verify              # 모든 의존성 검증")
	fmt.Println("  go clean -modcache         # 모듈 캐시 정리")
	fmt.Println()

	fmt.Println("보안 모범 사례:")
	securityPractices := []string{
		"go.sum 파일을 반드시 버전 관리에 포함",
		"의존성 업데이트 시 go.sum 변경사항 검토",
		"GOSUMDB를 통한 자동 체크섬 검증 활용",
		"사설 모듈의 경우 GOPRIVATE 환경변수 설정",
		"의존성 취약점 스캔 도구 활용",
	}

	for _, practice := range securityPractices {
		fmt.Printf("  - %s\n", practice)
	}
}

func demonstrateProxyAndChecksum() {
	fmt.Println("Go 모듈 프록시와 체크섬 데이터베이스:")
	fmt.Println()

	fmt.Println("GOPROXY (모듈 프록시):")
	fmt.Println("- 기본값: https://proxy.golang.org,direct")
	fmt.Println("- 모듈 다운로드 속도 향상")
	fmt.Println("- 가용성 보장")
	fmt.Println("- 캐싱 및 미러링")
	fmt.Println()

	fmt.Println("GOSUMDB (체크섬 데이터베이스):")
	fmt.Println("- 기본값: sum.golang.org")
	fmt.Println("- 모듈 무결성 검증")
	fmt.Println("- 변조 방지")
	fmt.Println()

	fmt.Println("환경변수 설정:")
	envExamples := []string{
		"export GOPROXY=https://goproxy.cn,direct         # 중국 프록시 사용",
		"export GOPROXY=https://proxy.golang.org,direct   # 기본 프록시",
		"export GOPROXY=direct                            # 프록시 사용 안 함",
		"export GOPRIVATE=*.corp.com                      # 사설 모듈 설정",
		"export GONOPROXY=*.internal.com                  # 프록시 제외",
		"export GONOSUMDB=*.private.com                   # 체크섬 DB 제외",
		"export GOSUMDB=off                               # 체크섬 검증 비활성화",
	}

	for _, env := range envExamples {
		fmt.Printf("  %s\n", env)
	}
}

func demonstrateWorkspaceMode() {
	fmt.Println("워크스페이스 모드 (Go 1.18+):")
	fmt.Println()

	fmt.Println("워크스페이스란?")
	fmt.Println("- 여러 모듈을 함께 개발할 때 사용")
	fmt.Println("- 로컬 의존성 관리 간소화")
	fmt.Println("- replace 지시문 대신 사용")
	fmt.Println()

	fmt.Println("워크스페이스 생성:")
	fmt.Println("  go work init ./module1 ./module2")
	fmt.Println("  go work use ./module3")
	fmt.Println("  go work sync")
	fmt.Println()

	fmt.Println("go.work 파일 예제:")
	goWorkExample := `go 1.21

use (
    ./api
    ./frontend  
    ./backend
    ./shared
)

replace (
    github.com/external/package => ./local/package
)`

	fmt.Println(goWorkExample)
	fmt.Println()

	fmt.Println("워크스페이스 명령어:")
	workspaceCommands := []string{
		"go work init [modules...]    # 워크스페이스 초기화",
		"go work use [modules...]     # 모듈 추가",
		"go work edit -use=./module   # 모듈 추가 (편집)",
		"go work edit -dropuse=./mod  # 모듈 제거",
		"go work sync                 # 워크스페이스 동기화",
	}

	for _, cmd := range workspaceCommands {
		fmt.Printf("  %s\n", cmd)
	}
}

func demonstrateBestPractices() {
	fmt.Println("Go 모듈 베스트 프랙티스:")
	fmt.Println()

	bestPractices := map[string][]string{
		"모듈 구조": {
			"모듈당 하나의 저장소 권장",
			"루트에 go.mod 파일 위치",
			"명확하고 간결한 모듈명 사용",
			"모듈 경로는 실제 저장소 경로와 일치",
		},
		"버전 관리": {
			"Semantic Versioning 준수",
			"v0은 개발 버전으로만 사용",
			"호환성 깨는 변경은 메이저 버전 증가",
			"정기적인 버전 태그 생성",
		},
		"의존성 관리": {
			"최소한의 의존성만 추가",
			"정기적인 의존성 업데이트",
			"go mod tidy 정기 실행",
			"취약점 스캔 도구 활용",
		},
		"개발 워크플로": {
			"go.sum 파일을 버전 관리에 포함",
			"CI/CD에서 go mod verify 실행",
			"의존성 라이센스 확인",
			"워크스페이스 모드로 로컬 개발",
		},
	}

	for category, practices := range bestPractices {
		fmt.Printf("%s:\n", category)
		for _, practice := range practices {
			fmt.Printf("  - %s\n", practice)
		}
		fmt.Println()
	}
}

func demonstrateTroubleshooting() {
	fmt.Println("일반적인 문제와 해결책:")
	fmt.Println()

	troubles := map[string]string{
		"의존성을 찾을 수 없음": "go clean -modcache && go mod download",
		"체크섬 불일치":      "go clean -modcache && go mod tidy",
		"순환 의존성":       "아키텍처 재설계 필요 (인터페이스 활용)",
		"프록시 접속 문제":    "GOPROXY 환경변수 확인",
		"사설 저장소 접근":    "GOPRIVATE 환경변수 설정",
		"버전 충돌":        "go mod graph로 의존성 분석 후 버전 조정",
		"빌드 실패":        "go mod tidy && go mod verify",
		"느린 다운로드":      "GOPROXY를 지역 프록시로 설정",
	}

	for problem, solution := range troubles {
		fmt.Printf("문제: %s\n", problem)
		fmt.Printf("해결: %s\n", solution)
		fmt.Println()
	}

	fmt.Println("디버깅 명령어:")
	debugCommands := []string{
		"go mod graph                 # 의존성 그래프 출력",
		"go mod why package           # 패키지가 필요한 이유",
		"go list -m -versions package # 사용 가능한 버전 목록",
		"go mod download -json        # 다운로드 정보 JSON 출력",
		"go env GOPROXY GOSUMDB      # 관련 환경변수 확인",
		"go clean -modcache          # 모듈 캐시 완전 정리",
	}

	for _, cmd := range debugCommands {
		fmt.Printf("  %s\n", cmd)
	}
}

func createExampleGoMod() {
	fmt.Println("예제 go.mod 파일 구조:")

	// 현재 디렉토리에 예제 go.mod가 있는지 확인
	goModPath := "go.mod"
	if _, err := os.Stat(goModPath); os.IsNotExist(err) {
		fmt.Println("  (현재 디렉토리에 go.mod가 없어서 예제만 표시)")
	} else {
		fmt.Printf("  현재 디렉토리의 go.mod 파일: %s\n", goModPath)

		// go.mod 파일 읽기
		content, err := os.ReadFile(goModPath)
		if err == nil {
			fmt.Println("  내용:")
			lines := string(content)
			for _, line := range []string{lines} {
				fmt.Printf("    %s", line)
			}
		}
	}
}

/*
Go Modules 완전 가이드:

1. 핵심 개념:
   - Module: 버전 관리되는 Go 패키지들의 그룹
   - Version: Semantic Versioning (v1.2.3)
   - Dependency: 모듈이 의존하는 다른 모듈
   - Proxy: 모듈 다운로드를 위한 캐싱 서버

2. 주요 파일들:
   - go.mod: 모듈 정의 (이름, Go 버전, 의존성)
   - go.sum: 체크섬 (보안 및 재현성)
   - go.work: 워크스페이스 정의 (Go 1.18+)

3. 버전 선택 규칙:
   - 최소 버전 선택 (Minimal Version Selection)
   - 명시된 최소 버전 중 가장 높은 버전 선택
   - 하위 호환성 보장 (같은 메이저 버전 내)

4. 프록시 시스템:
   - proxy.golang.org: 공식 프록시
   - sum.golang.org: 체크섬 검증
   - 캐싱과 가용성 향상

5. 보안 기능:
   - 체크섬 검증으로 무결성 보장
   - 투명한 로그로 변조 방지
   - 취약점 데이터베이스 연동

6. 개발 워크플로:
   - 로컬 개발: replace 지시문
   - 멀티 모듈: 워크스페이스 모드
   - CI/CD: 검증과 재현성 확보

7. 마이그레이션:
   - GOPATH → Modules
   - 기존 프로젝트 변환
   - 의존성 관리 개선

8. 고급 기능:
   - 조건부 빌드
   - 플랫폼별 의존성
   - 내부 모듈 관리
   - 기업 환경 설정
*/
