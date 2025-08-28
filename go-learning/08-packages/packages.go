// packages.go - Go의 패키지 시스템
package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println("=== Go 패키지 시스템 ===")

	// 1. 패키지 기본 개념
	fmt.Println("\n1. 패키지 기본 개념:")
	fmt.Println("- 패키지는 Go의 코드 구성 단위")
	fmt.Println("- 모든 Go 파일은 package 선언으로 시작")
	fmt.Println("- main 패키지는 실행 가능한 프로그램")
	fmt.Println("- 같은 디렉토리의 파일들은 같은 패키지")

	// 2. 표준 라이브러리 패키지들 소개
	fmt.Println("\n2. 주요 표준 라이브러리:")
	demonstrateStandardLibraries()

	// 3. 패키지 import 방법들
	fmt.Println("\n3. import 문 사용법:")
	demonstrateImportMethods()

	// 4. 패키지 가시성 규칙
	fmt.Println("\n4. 가시성 규칙 (Exported vs Unexported):")
	demonstrateVisibilityRules()

	// 5. 패키지 초기화
	fmt.Println("\n5. 패키지 초기화:")
	demonstratePackageInitialization()

	// 6. 사용자 정의 패키지 구조 예제
	fmt.Println("\n6. 사용자 정의 패키지 구조:")
	demonstrateCustomPackageStructure()

	// 7. Go Modules 사용법
	fmt.Println("\n7. Go Modules:")
	demonstrateGoModules()

	// 8. 패키지 문서화
	fmt.Println("\n8. 패키지 문서화:")
	demonstratePackageDocumentation()

	// 9. 내부 패키지 (Internal packages)
	fmt.Println("\n9. 내부 패키지:")
	demonstrateInternalPackages()

	// 10. 벤더링 (Vendoring)
	fmt.Println("\n10. 벤더링:")
	demonstrateVendoring()
}

// 2. 표준 라이브러리 데모
func demonstrateStandardLibraries() {
	fmt.Println("주요 표준 라이브러리 패키지들:")

	standardLibs := map[string]string{
		"fmt":           "형식화된 입출력",
		"os":            "운영체제 인터페이스",
		"io":            "I/O 기본 인터페이스",
		"net/http":      "HTTP 클라이언트/서버",
		"encoding/json": "JSON 인코딩/디코딩",
		"database/sql":  "SQL 데이터베이스 인터페이스",
		"context":       "취소, 데드라인, 요청 범위 값",
		"sync":          "동기화 프리미티브",
		"time":          "시간 관련 기능",
		"strings":       "문자열 조작",
		"strconv":       "문자열 변환",
		"math":          "수학 함수",
		"crypto":        "암호화 관련",
		"log":           "로깅 기능",
		"flag":          "명령행 플래그 파싱",
	}

	for pkg, desc := range standardLibs {
		fmt.Printf("  %-15s: %s\n", pkg, desc)
	}
}

// 3. import 방법 데모
func demonstrateImportMethods() {
	fmt.Println("import 문의 다양한 사용법:")

	examples := []string{
		`import "fmt"                    // 기본 import`,
		`import f "fmt"                  // 별칭 사용`,
		`import . "fmt"                  // dot import (권장하지 않음)`,
		`import _ "database/sql/driver"   // blank import (side effect만)`,
		`import (                        // 그룹 import
    "fmt"
    "os"
    "time"
)`,
	}

	for _, example := range examples {
		fmt.Printf("%s\n", example)
	}
}

// 4. 가시성 규칙 데모
func demonstrateVisibilityRules() {
	fmt.Println("Go의 가시성 규칙:")
	fmt.Println("- 대문자로 시작: Exported (다른 패키지에서 접근 가능)")
	fmt.Println("- 소문자로 시작: Unexported (같은 패키지 내에서만 접근)")
	fmt.Println()

	fmt.Println("예제:")
	fmt.Println("  type User struct {")
	fmt.Println("      Name    string  // Exported field")
	fmt.Println("      age     int     // unexported field")
	fmt.Println("  }")
	fmt.Println()
	fmt.Println("  func (u *User) GetAge() int {     // Exported method")
	fmt.Println("      return u.age")
	fmt.Println("  }")
	fmt.Println()
	fmt.Println("  func (u *User) validate() error { // unexported method")
	fmt.Println("      // validation logic")
	fmt.Println("      return nil")
	fmt.Println("  }")
}

// 5. 패키지 초기화 데모
func demonstratePackageInitialization() {
	fmt.Println("패키지 초기화 순서:")
	fmt.Println("1. import된 패키지들이 먼저 초기화")
	fmt.Println("2. 패키지 레벨 변수들이 선언 순서대로 초기화")
	fmt.Println("3. init() 함수들이 실행 (파일별로 여러 개 가능)")
	fmt.Println("4. main() 함수 실행 (main 패키지인 경우)")
	fmt.Println()

	fmt.Println("init() 함수 예제:")
	fmt.Println("  var config *Config")
	fmt.Println()
	fmt.Println("  func init() {")
	fmt.Println("      config = loadConfig()")
	fmt.Println("      fmt.Println('패키지 초기화 완료')")
	fmt.Println("  }")
}

// 6. 사용자 정의 패키지 구조 데모
func demonstrateCustomPackageStructure() {
	fmt.Println("권장하는 프로젝트 구조:")

	structure := []string{
		"myapp/",
		"├── cmd/",
		"│   └── myapp/",
		"│       └── main.go          # 애플리케이션 진입점",
		"├── pkg/",
		"│   ├── models/",
		"│   │   ├── user.go         # 사용자 모델",
		"│   │   └── product.go      # 상품 모델",
		"│   ├── services/",
		"│   │   ├── user_service.go # 사용자 서비스",
		"│   │   └── auth_service.go # 인증 서비스",
		"│   └── utils/",
		"│       ├── validator.go    # 유틸리티 함수",
		"│       └── helpers.go",
		"├── internal/",
		"│   ├── config/",
		"│   │   └── config.go       # 설정 관리",
		"│   └── database/",
		"│       └── db.go           # 데이터베이스 연결",
		"├── api/",
		"│   └── handlers/",
		"│       ├── user_handler.go # HTTP 핸들러",
		"│       └── auth_handler.go",
		"├── web/",
		"│   ├── static/             # 정적 파일",
		"│   └── templates/          # 템플릿",
		"├── scripts/",
		"│   └── deploy.sh           # 배포 스크립트",
		"├── docs/",
		"│   └── README.md           # 문서",
		"├── go.mod                  # Go 모듈 정의",
		"├── go.sum                  # 의존성 체크섬",
		"└── Makefile               # 빌드 스크립트",
	}

	for _, line := range structure {
		fmt.Println(line)
	}
}

// 7. Go Modules 데모
func demonstrateGoModules() {
	fmt.Println("Go Modules 사용법:")
	fmt.Println()

	fmt.Println("1. 모듈 초기화:")
	fmt.Println("   go mod init github.com/username/projectname")
	fmt.Println()

	fmt.Println("2. 의존성 추가:")
	fmt.Println("   go get github.com/gorilla/mux")
	fmt.Println("   go get github.com/gin-gonic/gin@v1.9.0")
	fmt.Println()

	fmt.Println("3. 의존성 관리:")
	fmt.Println("   go mod tidy      # 불필요한 의존성 제거")
	fmt.Println("   go mod download  # 의존성 다운로드")
	fmt.Println("   go mod verify    # 의존성 검증")
	fmt.Println("   go mod vendor    # vendor 디렉토리 생성")
	fmt.Println()

	fmt.Println("4. go.mod 파일 예제:")
	goModExample := `module github.com/username/myapp

go 1.21

require (
    github.com/gin-gonic/gin v1.9.1
    github.com/gorilla/mux v1.8.0
    github.com/lib/pq v1.10.9
)

require (
    github.com/bytedance/sonic v1.9.1 // indirect
    github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311 // indirect
    // ... other indirect dependencies
)`
	fmt.Println(goModExample)
}

// 8. 패키지 문서화 데모
func demonstratePackageDocumentation() {
	fmt.Println("Go 패키지 문서화 방법:")
	fmt.Println()

	fmt.Println("1. 패키지 문서화:")
	fmt.Println("   // Package math provides basic mathematical functions.")
	fmt.Println("   //")
	fmt.Println("   // This package implements elementary functions such as")
	fmt.Println("   // trigonometric functions, logarithms, etc.")
	fmt.Println("   package math")
	fmt.Println()

	fmt.Println("2. 함수 문서화:")
	fmt.Println("   // Add returns the sum of a and b.")
	fmt.Println("   //")
	fmt.Println("   // Example:")
	fmt.Println("   //   result := Add(5, 3) // returns 8")
	fmt.Println("   func Add(a, b int) int {")
	fmt.Println("       return a + b")
	fmt.Println("   }")
	fmt.Println()

	fmt.Println("3. 문서 확인:")
	fmt.Println("   go doc package_name")
	fmt.Println("   go doc package_name.FunctionName")
	fmt.Println("   godoc -http=:6060  # 로컬 문서 서버 실행")
}

// 9. 내부 패키지 데모
func demonstrateInternalPackages() {
	fmt.Println("내부 패키지 (Internal Packages):")
	fmt.Println("- 'internal' 디렉토리의 패키지들")
	fmt.Println("- 같은 서브트리의 패키지에서만 import 가능")
	fmt.Println("- 외부 패키지에서 접근 불가")
	fmt.Println()

	fmt.Println("예제 구조:")
	internalExample := []string{
		"myapp/",
		"├── pkg/",
		"│   └── public/",
		"│       └── api.go         # 외부에서 접근 가능",
		"└── internal/",
		"    ├── auth/",
		"    │   └── auth.go        # myapp 내부에서만 접근 가능",
		"    └── database/",
		"        └── db.go          # myapp 내부에서만 접근 가능",
	}

	for _, line := range internalExample {
		fmt.Println(line)
	}
}

// 10. 벤더링 데모
func demonstrateVendoring() {
	fmt.Println("벤더링 (Vendoring):")
	fmt.Println("- 의존성을 프로젝트 내부에 복사")
	fmt.Println("- 외부 의존성 변경에 대한 보호")
	fmt.Println("- 빌드 재현성 보장")
	fmt.Println()

	fmt.Println("벤더링 명령어:")
	fmt.Println("  go mod vendor    # vendor 디렉토리 생성")
	fmt.Println("  go build -mod=vendor  # vendor 사용하여 빌드")
	fmt.Println()

	fmt.Println("벤더링 후 구조:")
	vendorExample := []string{
		"myapp/",
		"├── vendor/",
		"│   ├── github.com/",
		"│   │   └── gorilla/",
		"│   │       └── mux/",
		"│   └── modules.txt         # 벤더된 모듈 정보",
		"├── go.mod",
		"└── go.sum",
	}

	for _, line := range vendorExample {
		fmt.Println(line)
	}

	// 11. 실제 패키지 분석 예제
	fmt.Println("\n11. 현재 디렉토리의 Go 파일 분석:")
	analyzeGoFiles()
}

// 실제 Go 파일 분석 함수
func analyzeGoFiles() {
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !strings.HasSuffix(path, ".go") {
			return nil
		}

		// Go 파일 파싱
		fset := token.NewFileSet()
		node, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if err != nil {
			return err
		}

		fmt.Printf("파일: %s\n", path)
		fmt.Printf("  패키지: %s\n", node.Name.Name)

		// import 문 분석
		if len(node.Imports) > 0 {
			fmt.Println("  imports:")
			for _, imp := range node.Imports {
				importPath := strings.Trim(imp.Path.Value, `"`)
				if imp.Name != nil {
					fmt.Printf("    %s %s\n", imp.Name.Name, importPath)
				} else {
					fmt.Printf("    %s\n", importPath)
				}
			}
		}

		// 함수 분석
		functions := []string{}
		for _, decl := range node.Decls {
			if fn, ok := decl.(*ast.FuncDecl); ok {
				if fn.Name.IsExported() {
					functions = append(functions, fn.Name.Name)
				}
			}
		}

		if len(functions) > 0 {
			fmt.Printf("  exported functions: %s\n", strings.Join(functions, ", "))
		}

		fmt.Println()
		return nil
	})

	if err != nil {
		fmt.Printf("파일 분석 중 에러: %v\n", err)
	}
}

/*
Go 패키지 시스템 핵심 개념:

1. 패키지 기본 원리:
   - 코드 구성의 기본 단위
   - 디렉토리 = 패키지 (같은 디렉토리의 모든 .go 파일은 같은 패키지)
   - main 패키지는 실행 가능한 프로그램
   - 다른 패키지들은 라이브러리

2. 가시성 규칙:
   - 대문자 시작: Exported (public)
   - 소문자 시작: Unexported (private)
   - 패키지 경계에서만 적용

3. Import 경로:
   - 표준 라이브러리: "fmt", "os", "net/http"
   - 로컬 패키지: "./utils", "../models"
   - 외부 패키지: "github.com/user/repo"

4. 패키지 초기화 순서:
   1. Import된 패키지들 초기화
   2. 패키지 레벨 변수 초기화
   3. init() 함수 실행
   4. main() 함수 실행

5. Go Modules (Go 1.11+):
   - 의존성 관리 시스템
   - go.mod: 모듈 정의
   - go.sum: 의존성 체크섬
   - 버전 관리 (Semantic Versioning)

6. 모범 사례:
   - 패키지명은 간결하고 명확하게
   - 순환 의존성 방지
   - internal 패키지로 내부 API 보호
   - 인터페이스는 사용하는 쪽에서 정의
   - 하나의 개념당 하나의 패키지

7. 프로젝트 구조:
   - cmd/: 실행 파일들
   - pkg/: 외부에서 사용 가능한 라이브러리
   - internal/: 내부 전용 패키지
   - api/: API 관련 코드
   - web/: 웹 관련 파일들

8. 문서화:
   - 패키지 문서: 패키지 선언 위의 주석
   - 함수 문서: 함수 선언 위의 주석
   - godoc 도구로 문서 생성
   - 예제 코드 포함 가능

9. 테스팅:
   - _test.go 파일들
   - 같은 패키지 또는 _test 패키지
   - go test 명령어로 실행

10. 배포와 버전 관리:
    - Git 태그로 버전 관리
    - go.mod의 require 절로 의존성 명시
    - go get으로 패키지 설치
    - 호환성 보장 (v1.x.x 내에서)
*/
