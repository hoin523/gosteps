// http-server.go - Go로 웹 서버 개발하기
package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// 데이터 모델들
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type Server struct {
	users []User
}

// 메모리 저장소 (실제 프로젝트에서는 데이터베이스 사용)
var server = &Server{
	users: []User{
		{ID: 1, Name: "김철수", Email: "kim@example.com"},
		{ID: 2, Name: "이영희", Email: "lee@example.com"},
		{ID: 3, Name: "박민수", Email: "park@example.com"},
	},
}

func main() {
	fmt.Println("=== Go 웹 서버 개발 ===")

	// 1. 기본 HTTP 핸들러 등록
	setupBasicHandlers()

	// 2. RESTful API 엔드포인트 등록
	setupAPIHandlers()

	// 3. 미들웨어 사용 예제
	setupMiddlewareHandlers()

	// 4. 정적 파일 서빙
	setupStaticFileServing()

	// 5. 템플릿 핸들러
	setupTemplateHandlers()

	// 서버 시작
	port := ":8080"
	fmt.Printf("서버가 포트 %s에서 실행중입니다...\n", port)
	fmt.Println("\n사용 가능한 엔드포인트:")
	printEndpoints()

	log.Fatal(http.ListenAndServe(port, nil))
}

// 1. 기본 HTTP 핸들러들
func setupBasicHandlers() {
	// 루트 경로
	http.HandleFunc("/", homeHandler)

	// 단순한 텍스트 응답
	http.HandleFunc("/hello", helloHandler)

	// URL 파라미터 처리
	http.HandleFunc("/greet/", greetHandler)

	// 여러 HTTP 메서드 처리
	http.HandleFunc("/method-test", methodTestHandler)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	html := `
<!DOCTYPE html>
<html>
<head>
    <title>Go 웹 서버 예제</title>
    <meta charset="utf-8">
</head>
<body>
    <h1>Go 웹 서버에 오신 것을 환영합니다!</h1>
    <h2>사용 가능한 엔드포인트:</h2>
    <ul>
        <li><a href="/hello">Hello 페이지</a></li>
        <li><a href="/greet/김철수">Greet 페이지</a></li>
        <li><a href="/api/users">사용자 API</a></li>
        <li><a href="/template">템플릿 예제</a></li>
        <li><a href="/static/">정적 파일</a></li>
    </ul>
    <p>현재 시간: ` + time.Now().Format("2006-01-02 15:04:05") + `</p>
</body>
</html>`

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, html)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}

	fmt.Fprintf(w, "안녕하세요, %s님!", name)
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	// URL에서 이름 추출
	path := strings.TrimPrefix(r.URL.Path, "/greet/")
	name := strings.Split(path, "/")[0]

	if name == "" {
		http.Error(w, "이름을 입력해주세요", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `
<h1>안녕하세요, %s님!</h1>
<p>현재 시간: %s</p>
<p><a href="/">홈으로 돌아가기</a></p>
`, name, time.Now().Format("2006-01-02 15:04:05"))
}

func methodTestHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprint(w, "GET 요청을 받았습니다")
	case http.MethodPost:
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "요청 본문을 읽을 수 없습니다", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "POST 요청을 받았습니다. 본문: %s", string(body))
	case http.MethodPut:
		fmt.Fprint(w, "PUT 요청을 받았습니다")
	case http.MethodDelete:
		fmt.Fprint(w, "DELETE 요청을 받았습니다")
	default:
		http.Error(w, "지원하지 않는 메서드입니다", http.StatusMethodNotAllowed)
	}
}

// 2. RESTful API 핸들러들
func setupAPIHandlers() {
	http.HandleFunc("/api/users", usersHandler)
	http.HandleFunc("/api/users/", userHandler)
	http.HandleFunc("/api/health", healthHandler)
	http.HandleFunc("/api/echo", echoHandler)
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		// 모든 사용자 조회
		response := Response{
			Status:  "success",
			Message: "사용자 목록 조회 성공",
			Data:    server.users,
		}
		json.NewEncoder(w).Encode(response)

	case http.MethodPost:
		// 새 사용자 생성
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "잘못된 JSON 형식", http.StatusBadRequest)
			return
		}

		// ID 자동 생성
		user.ID = len(server.users) + 1
		server.users = append(server.users, user)

		response := Response{
			Status:  "success",
			Message: "사용자 생성 성공",
			Data:    user,
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)

	default:
		http.Error(w, "지원하지 않는 메서드", http.StatusMethodNotAllowed)
	}
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	// URL에서 사용자 ID 추출
	path := strings.TrimPrefix(r.URL.Path, "/api/users/")
	idStr := strings.Split(path, "/")[0]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "유효하지 않은 사용자 ID", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		// 특정 사용자 조회
		user := findUserByID(id)
		if user == nil {
			response := Response{
				Status:  "error",
				Message: "사용자를 찾을 수 없습니다",
			}
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(response)
			return
		}

		response := Response{
			Status:  "success",
			Message: "사용자 조회 성공",
			Data:    user,
		}
		json.NewEncoder(w).Encode(response)

	case http.MethodPut:
		// 사용자 정보 수정
		userIndex := findUserIndexByID(id)
		if userIndex == -1 {
			response := Response{
				Status:  "error",
				Message: "사용자를 찾을 수 없습니다",
			}
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(response)
			return
		}

		var updatedUser User
		err := json.NewDecoder(r.Body).Decode(&updatedUser)
		if err != nil {
			http.Error(w, "잘못된 JSON 형식", http.StatusBadRequest)
			return
		}

		updatedUser.ID = id
		server.users[userIndex] = updatedUser

		response := Response{
			Status:  "success",
			Message: "사용자 수정 성공",
			Data:    updatedUser,
		}
		json.NewEncoder(w).Encode(response)

	case http.MethodDelete:
		// 사용자 삭제
		userIndex := findUserIndexByID(id)
		if userIndex == -1 {
			response := Response{
				Status:  "error",
				Message: "사용자를 찾을 수 없습니다",
			}
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(response)
			return
		}

		// 슬라이스에서 사용자 제거
		server.users = append(server.users[:userIndex], server.users[userIndex+1:]...)

		response := Response{
			Status:  "success",
			Message: "사용자 삭제 성공",
		}
		json.NewEncoder(w).Encode(response)

	default:
		http.Error(w, "지원하지 않는 메서드", http.StatusMethodNotAllowed)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	health := map[string]interface{}{
		"status":    "healthy",
		"timestamp": time.Now().Format(time.RFC3339),
		"uptime":    "1h 30m",
		"version":   "1.0.0",
	}

	json.NewEncoder(w).Encode(health)
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	echo := map[string]interface{}{
		"method":  r.Method,
		"url":     r.URL.String(),
		"headers": r.Header,
		"query":   r.URL.Query(),
	}

	if r.Method == http.MethodPost || r.Method == http.MethodPut {
		body, err := io.ReadAll(r.Body)
		if err == nil {
			echo["body"] = string(body)
		}
	}

	json.NewEncoder(w).Encode(echo)
}

// 3. 미들웨어 예제
func setupMiddlewareHandlers() {
	// 미들웨어를 적용한 핸들러
	http.Handle("/api/protected", loggingMiddleware(authMiddleware(http.HandlerFunc(protectedHandler))))
}

// 로깅 미들웨어
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// 요청 로깅
		log.Printf("시작 %s %s", r.Method, r.URL.Path)

		// 다음 핸들러 실행
		next.ServeHTTP(w, r)

		// 응답 시간 로깅
		log.Printf("완료 %s %s - %v", r.Method, r.URL.Path, time.Since(start))
	})
}

// 인증 미들웨어 (간단한 API 키 체크)
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-Key")

		if apiKey != "secret-api-key" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)

			response := Response{
				Status:  "error",
				Message: "유효하지 않은 API 키입니다",
			}
			json.NewEncoder(w).Encode(response)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func protectedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := Response{
		Status:  "success",
		Message: "보호된 리소스에 접근했습니다",
		Data: map[string]string{
			"secret": "이것은 비밀 데이터입니다",
		},
	}

	json.NewEncoder(w).Encode(response)
}

// 4. 정적 파일 서빙
func setupStaticFileServing() {
	// /static/ 경로로 정적 파일 서빙
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
}

// 5. 템플릿 핸들러
func setupTemplateHandlers() {
	http.HandleFunc("/template", templateHandler)
	http.HandleFunc("/user-list", userListTemplateHandler)
}

func templateHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := `
<!DOCTYPE html>
<html>
<head>
    <title>Go 템플릿 예제</title>
    <meta charset="utf-8">
</head>
<body>
    <h1>{{.Title}}</h1>
    <p>안녕하세요, {{.Name}}님!</p>
    <p>현재 시간: {{.Time}}</p>
    <p>좋아하는 색깔:</p>
    <ul>
    {{range .Colors}}
        <li>{{.}}</li>
    {{end}}
    </ul>
</body>
</html>`

	t, err := template.New("page").Parse(tmpl)
	if err != nil {
		http.Error(w, "템플릿 파싱 오류", http.StatusInternalServerError)
		return
	}

	data := struct {
		Title  string
		Name   string
		Time   string
		Colors []string
	}{
		Title:  "Go 템플릿 데모",
		Name:   "개발자",
		Time:   time.Now().Format("2006-01-02 15:04:05"),
		Colors: []string{"빨강", "파랑", "초록", "노랑"},
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t.Execute(w, data)
}

func userListTemplateHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := `
<!DOCTYPE html>
<html>
<head>
    <title>사용자 목록</title>
    <meta charset="utf-8">
    <style>
        table { border-collapse: collapse; width: 100%; }
        th, td { border: 1px solid #ddd; padding: 8px; }
        th { background-color: #f2f2f2; }
    </style>
</head>
<body>
    <h1>사용자 목록</h1>
    <table>
        <tr>
            <th>ID</th>
            <th>이름</th>
            <th>이메일</th>
        </tr>
        {{range .Users}}
        <tr>
            <td>{{.ID}}</td>
            <td>{{.Name}}</td>
            <td>{{.Email}}</td>
        </tr>
        {{end}}
    </table>
    <p>총 {{len .Users}}명의 사용자가 있습니다.</p>
    <p><a href="/">홈으로</a></p>
</body>
</html>`

	t, err := template.New("users").Parse(tmpl)
	if err != nil {
		http.Error(w, "템플릿 파싱 오류", http.StatusInternalServerError)
		return
	}

	data := struct {
		Users []User
	}{
		Users: server.users,
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t.Execute(w, data)
}

// 유틸리티 함수들
func findUserByID(id int) *User {
	for _, user := range server.users {
		if user.ID == id {
			return &user
		}
	}
	return nil
}

func findUserIndexByID(id int) int {
	for i, user := range server.users {
		if user.ID == id {
			return i
		}
	}
	return -1
}

func printEndpoints() {
	endpoints := []string{
		"GET  /                      - 홈페이지",
		"GET  /hello?name=이름        - Hello 페이지",
		"GET  /greet/{name}          - 인사 페이지",
		"*    /method-test           - HTTP 메서드 테스트",
		"GET  /api/users             - 모든 사용자 조회",
		"POST /api/users             - 새 사용자 생성",
		"GET  /api/users/{id}        - 특정 사용자 조회",
		"PUT  /api/users/{id}        - 사용자 정보 수정",
		"DEL  /api/users/{id}        - 사용자 삭제",
		"GET  /api/health            - 서버 상태 확인",
		"*    /api/echo              - 요청 에코",
		"GET  /api/protected         - 보호된 리소스 (X-API-Key: secret-api-key)",
		"GET  /template              - 템플릿 예제",
		"GET  /user-list             - 사용자 목록 템플릿",
		"GET  /static/{file}         - 정적 파일",
	}

	for _, endpoint := range endpoints {
		fmt.Printf("  %s\n", endpoint)
	}

	fmt.Println("\nAPI 사용 예제:")
	fmt.Println("curl http://localhost:8080/api/users")
	fmt.Println("curl -X POST -H 'Content-Type: application/json' -d '{\"name\":\"홍길동\",\"email\":\"hong@example.com\"}' http://localhost:8080/api/users")
	fmt.Println("curl -H 'X-API-Key: secret-api-key' http://localhost:8080/api/protected")
}

/*
Go HTTP 웹 서버 개발 가이드:

1. 기본 HTTP 서버:
   - net/http 패키지 사용
   - http.HandleFunc()로 라우트 등록
   - http.ListenAndServe()로 서버 시작

2. HTTP 핸들러:
   - func(http.ResponseWriter, *http.Request) 시그니처
   - ResponseWriter로 응답 작성
   - Request에서 요청 정보 읽기

3. URL 라우팅:
   - 기본 ServeMux 사용
   - 패턴 매칭 (/path/, /exact/path)
   - 써드파티 라우터 사용 가능 (gorilla/mux 등)

4. HTTP 메서드 처리:
   - r.Method로 메서드 구분
   - RESTful API 설계 원칙 적용
   - 적절한 HTTP 상태 코드 반환

5. 요청 처리:
   - URL 파라미터: r.URL.Query()
   - 경로 파라미터: URL 파싱 필요
   - 요청 본문: r.Body 읽기
   - 헤더: r.Header

6. 응답 처리:
   - 헤더 설정: w.Header().Set()
   - 상태 코드: w.WriteHeader()
   - 본문 작성: w.Write() 또는 fmt.Fprint()

7. JSON API:
   - json.NewEncoder().Encode() (응답)
   - json.NewDecoder().Decode() (요청)
   - 적절한 Content-Type 헤더

8. 미들웨어:
   - http.Handler 인터페이스 활용
   - 체이닝 패턴 구현
   - 로깅, 인증, CORS 등

9. 템플릿:
   - html/template 패키지
   - 안전한 HTML 생성
   - 데이터 바인딩

10. 정적 파일:
    - http.FileServer() 사용
    - http.StripPrefix()로 경로 조정
    - 캐싱 헤더 설정 가능

11. 보안 고려사항:
    - HTTPS 사용
    - CSRF 보호
    - 입력값 검증
    - SQL 인젝션 방지

12. 성능 최적화:
    - Keep-Alive 연결
    - 압축 사용
    - 캐싱 전략
    - 연결 풀링
*/
