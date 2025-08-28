// main.go - TODO 애플리케이션
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Todo 아이템 구조체
type Todo struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Completed   bool       `json:"completed"`
	Priority    Priority   `json:"priority"`
	Category    string     `json:"category"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DueDate     *time.Time `json:"due_date,omitempty"`
}

// 우선순위 타입
type Priority int

const (
	Low Priority = iota + 1
	Medium
	High
	Critical
)

// Priority의 문자열 표현
func (p Priority) String() string {
	switch p {
	case Low:
		return "낮음"
	case Medium:
		return "보통"
	case High:
		return "높음"
	case Critical:
		return "긴급"
	default:
		return "알 수 없음"
	}
}

// TodoManager 구조체
type TodoManager struct {
	todos    []Todo
	nextID   int
	filename string
}

// NewTodoManager 생성자
func NewTodoManager(filename string) *TodoManager {
	tm := &TodoManager{
		todos:    make([]Todo, 0),
		nextID:   1,
		filename: filename,
	}

	// 파일에서 데이터 로드
	tm.Load()
	return tm
}

// AddTodo 새로운 TODO 추가
func (tm *TodoManager) AddTodo(title, description, category string, priority Priority, dueDate *time.Time) {
	todo := Todo{
		ID:          tm.nextID,
		Title:       title,
		Description: description,
		Completed:   false,
		Priority:    priority,
		Category:    category,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DueDate:     dueDate,
	}

	tm.todos = append(tm.todos, todo)
	tm.nextID++
	tm.Save()

	fmt.Printf("TODO가 추가되었습니다! (ID: %d)\n", todo.ID)
}

// GetAllTodos 모든 TODO 조회
func (tm *TodoManager) GetAllTodos() []Todo {
	return tm.todos[:]
}

// GetTodoByID ID로 TODO 조회
func (tm *TodoManager) GetTodoByID(id int) (*Todo, error) {
	for i := range tm.todos {
		if tm.todos[i].ID == id {
			return &tm.todos[i], nil
		}
	}
	return nil, fmt.Errorf("ID %d인 TODO를 찾을 수 없습니다", id)
}

// UpdateTodo TODO 업데이트
func (tm *TodoManager) UpdateTodo(id int, title, description, category string, priority Priority, dueDate *time.Time) error {
	todo, err := tm.GetTodoByID(id)
	if err != nil {
		return err
	}

	if title != "" {
		todo.Title = title
	}
	if description != "" {
		todo.Description = description
	}
	if category != "" {
		todo.Category = category
	}
	if priority != 0 {
		todo.Priority = priority
	}
	todo.DueDate = dueDate
	todo.UpdatedAt = time.Now()

	tm.Save()
	fmt.Printf("TODO (ID: %d)가 업데이트되었습니다.\n", id)
	return nil
}

// CompleteTodo TODO 완료 처리
func (tm *TodoManager) CompleteTodo(id int) error {
	todo, err := tm.GetTodoByID(id)
	if err != nil {
		return err
	}

	todo.Completed = true
	todo.UpdatedAt = time.Now()
	tm.Save()

	fmt.Printf("TODO (ID: %d) '%s'가 완료되었습니다! 🎉\n", id, todo.Title)
	return nil
}

// DeleteTodo TODO 삭제
func (tm *TodoManager) DeleteTodo(id int) error {
	for i, todo := range tm.todos {
		if todo.ID == id {
			tm.todos = append(tm.todos[:i], tm.todos[i+1:]...)
			tm.Save()
			fmt.Printf("TODO (ID: %d)가 삭제되었습니다.\n", id)
			return nil
		}
	}
	return fmt.Errorf("ID %d인 TODO를 찾을 수 없습니다", id)
}

// ListTodos TODO 목록 출력
func (tm *TodoManager) ListTodos(filter TodoFilter) {
	todos := tm.FilterTodos(filter)

	if len(todos) == 0 {
		fmt.Println("표시할 TODO가 없습니다.")
		return
	}

	// 정렬
	tm.SortTodos(todos, filter.SortBy)

	fmt.Printf("\n=== TODO 목록 (%d개) ===\n", len(todos))

	for _, todo := range todos {
		status := "⭕"
		if todo.Completed {
			status = "✅"
		}

		// 우선순위 아이콘
		priorityIcon := ""
		switch todo.Priority {
		case Critical:
			priorityIcon = "🔥"
		case High:
			priorityIcon = "❗"
		case Medium:
			priorityIcon = "➖"
		case Low:
			priorityIcon = "⬇️"
		}

		fmt.Printf("%s [%d] %s %s - %s\n", status, todo.ID, priorityIcon, todo.Title, todo.Priority)

		if todo.Description != "" {
			fmt.Printf("    설명: %s\n", todo.Description)
		}

		if todo.Category != "" {
			fmt.Printf("    카테고리: %s\n", todo.Category)
		}

		if todo.DueDate != nil {
			dueStr := todo.DueDate.Format("2006-01-02 15:04")
			if time.Now().After(*todo.DueDate) && !todo.Completed {
				fmt.Printf("    ⚠️ 마감일: %s (지났음)\n", dueStr)
			} else {
				fmt.Printf("    📅 마감일: %s\n", dueStr)
			}
		}

		fmt.Printf("    생성: %s\n", todo.CreatedAt.Format("2006-01-02 15:04"))

		if !todo.UpdatedAt.Equal(todo.CreatedAt) {
			fmt.Printf("    수정: %s\n", todo.UpdatedAt.Format("2006-01-02 15:04"))
		}

		fmt.Println()
	}
}

// TodoFilter 필터링 옵션
type TodoFilter struct {
	ShowCompleted bool
	Category      string
	Priority      Priority
	SortBy        string
}

// FilterTodos TODO 필터링
func (tm *TodoManager) FilterTodos(filter TodoFilter) []Todo {
	var filtered []Todo

	for _, todo := range tm.todos {
		// 완료된 항목 필터
		if !filter.ShowCompleted && todo.Completed {
			continue
		}

		// 카테고리 필터
		if filter.Category != "" && todo.Category != filter.Category {
			continue
		}

		// 우선순위 필터
		if filter.Priority != 0 && todo.Priority != filter.Priority {
			continue
		}

		filtered = append(filtered, todo)
	}

	return filtered
}

// SortTodos TODO 정렬
func (tm *TodoManager) SortTodos(todos []Todo, sortBy string) {
	switch sortBy {
	case "priority":
		sort.Slice(todos, func(i, j int) bool {
			return todos[i].Priority > todos[j].Priority // 높은 우선순위부터
		})
	case "created":
		sort.Slice(todos, func(i, j int) bool {
			return todos[i].CreatedAt.After(todos[j].CreatedAt) // 최근 생성부터
		})
	case "due":
		sort.Slice(todos, func(i, j int) bool {
			if todos[i].DueDate == nil && todos[j].DueDate == nil {
				return false
			}
			if todos[i].DueDate == nil {
				return false
			}
			if todos[j].DueDate == nil {
				return true
			}
			return todos[i].DueDate.Before(*todos[j].DueDate)
		})
	default: // "id"
		sort.Slice(todos, func(i, j int) bool {
			return todos[i].ID < todos[j].ID
		})
	}
}

// GetStatistics 통계 정보
func (tm *TodoManager) GetStatistics() {
	total := len(tm.todos)
	completed := 0
	overdue := 0

	categories := make(map[string]int)
	priorities := make(map[Priority]int)

	now := time.Now()

	for _, todo := range tm.todos {
		if todo.Completed {
			completed++
		}

		if todo.DueDate != nil && now.After(*todo.DueDate) && !todo.Completed {
			overdue++
		}

		categories[todo.Category]++
		priorities[todo.Priority]++
	}

	fmt.Println("\n=== TODO 통계 ===")
	fmt.Printf("전체 TODO: %d개\n", total)
	fmt.Printf("완료된 TODO: %d개 (%.1f%%)\n", completed, float64(completed)/float64(total)*100)
	fmt.Printf("미완료 TODO: %d개\n", total-completed)
	fmt.Printf("기한 초과: %d개\n", overdue)

	if len(categories) > 0 {
		fmt.Println("\n카테고리별:")
		for category, count := range categories {
			if category == "" {
				category = "(없음)"
			}
			fmt.Printf("  %s: %d개\n", category, count)
		}
	}

	fmt.Println("\n우선순위별:")
	for priority, count := range priorities {
		fmt.Printf("  %s: %d개\n", priority, count)
	}
	fmt.Println()
}

// Save TODO 목록을 파일에 저장
func (tm *TodoManager) Save() error {
	data, err := json.MarshalIndent(tm.todos, "", "  ")
	if err != nil {
		return fmt.Errorf("JSON 인코딩 오류: %v", err)
	}

	err = os.WriteFile(tm.filename, data, 0644)
	if err != nil {
		return fmt.Errorf("파일 저장 오류: %v", err)
	}

	return nil
}

// Load 파일에서 TODO 목록 로드
func (tm *TodoManager) Load() error {
	file, err := os.Open(tm.filename)
	if err != nil {
		// 파일이 없으면 새로 시작
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("파일 열기 오류: %v", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("파일 읽기 오류: %v", err)
	}

	if len(data) == 0 {
		return nil
	}

	err = json.Unmarshal(data, &tm.todos)
	if err != nil {
		return fmt.Errorf("JSON 디코딩 오류: %v", err)
	}

	// 다음 ID 설정
	maxID := 0
	for _, todo := range tm.todos {
		if todo.ID > maxID {
			maxID = todo.ID
		}
	}
	tm.nextID = maxID + 1

	return nil
}

// CLI 관련 함수들
func main() {
	if len(os.Args) < 2 {
		showUsage()
		return
	}

	filename := "todos.json"
	tm := NewTodoManager(filename)

	command := os.Args[1]

	switch command {
	case "add":
		handleAddCommand(tm, os.Args[2:])
	case "list":
		handleListCommand(tm, os.Args[2:])
	case "complete":
		handleCompleteCommand(tm, os.Args[2:])
	case "delete":
		handleDeleteCommand(tm, os.Args[2:])
	case "update":
		handleUpdateCommand(tm, os.Args[2:])
	case "stats":
		tm.GetStatistics()
	case "help":
		showHelp()
	default:
		fmt.Printf("알 수 없는 명령어: %s\n", command)
		showUsage()
	}
}

func handleAddCommand(tm *TodoManager, args []string) {
	if len(args) == 0 {
		fmt.Println("TODO 제목을 입력하세요.")
		return
	}

	title := args[0]
	description := ""
	category := ""
	priority := Medium
	var dueDate *time.Time

	// 플래그 파싱
	for i := 1; i < len(args); i += 2 {
		if i+1 >= len(args) {
			break
		}

		flag := args[i]
		value := args[i+1]

		switch flag {
		case "-desc", "-d":
			description = value
		case "-category", "-c":
			category = value
		case "-priority", "-p":
			switch strings.ToLower(value) {
			case "low", "낮음", "1":
				priority = Low
			case "medium", "보통", "2":
				priority = Medium
			case "high", "높음", "3":
				priority = High
			case "critical", "긴급", "4":
				priority = Critical
			}
		case "-due":
			if parsedTime, err := time.Parse("2006-01-02", value); err == nil {
				dueDate = &parsedTime
			} else if parsedTime, err := time.Parse("2006-01-02 15:04", value); err == nil {
				dueDate = &parsedTime
			} else {
				fmt.Printf("잘못된 날짜 형식: %s (올바른 형식: 2006-01-02 또는 2006-01-02 15:04)\n", value)
				return
			}
		}
	}

	tm.AddTodo(title, description, category, priority, dueDate)
}

func handleListCommand(tm *TodoManager, args []string) {
	filter := TodoFilter{
		ShowCompleted: false,
		SortBy:        "id",
	}

	// 플래그 파싱
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-all", "-a":
			filter.ShowCompleted = true
		case "-category", "-c":
			if i+1 < len(args) {
				filter.Category = args[i+1]
				i++
			}
		case "-priority", "-p":
			if i+1 < len(args) {
				switch strings.ToLower(args[i+1]) {
				case "low", "낮음", "1":
					filter.Priority = Low
				case "medium", "보통", "2":
					filter.Priority = Medium
				case "high", "높음", "3":
					filter.Priority = High
				case "critical", "긴급", "4":
					filter.Priority = Critical
				}
				i++
			}
		case "-sort", "-s":
			if i+1 < len(args) {
				filter.SortBy = args[i+1]
				i++
			}
		}
	}

	tm.ListTodos(filter)
}

func handleCompleteCommand(tm *TodoManager, args []string) {
	if len(args) == 0 {
		fmt.Println("완료할 TODO의 ID를 입력하세요.")
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("잘못된 ID 형식: %s\n", args[0])
		return
	}

	if err := tm.CompleteTodo(id); err != nil {
		fmt.Printf("오류: %v\n", err)
	}
}

func handleDeleteCommand(tm *TodoManager, args []string) {
	if len(args) == 0 {
		fmt.Println("삭제할 TODO의 ID를 입력하세요.")
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("잘못된 ID 형식: %s\n", args[0])
		return
	}

	if err := tm.DeleteTodo(id); err != nil {
		fmt.Printf("오류: %v\n", err)
	}
}

func handleUpdateCommand(tm *TodoManager, args []string) {
	if len(args) == 0 {
		fmt.Println("업데이트할 TODO의 ID를 입력하세요.")
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("잘못된 ID 형식: %s\n", args[0])
		return
	}

	var title, description, category string
	var priority Priority
	var dueDate *time.Time

	// 플래그 파싱
	for i := 1; i < len(args); i += 2 {
		if i+1 >= len(args) {
			break
		}

		flag := args[i]
		value := args[i+1]

		switch flag {
		case "-title", "-t":
			title = value
		case "-desc", "-d":
			description = value
		case "-category", "-c":
			category = value
		case "-priority", "-p":
			switch strings.ToLower(value) {
			case "low", "낮음", "1":
				priority = Low
			case "medium", "보통", "2":
				priority = Medium
			case "high", "높음", "3":
				priority = High
			case "critical", "긴급", "4":
				priority = Critical
			}
		case "-due":
			if parsedTime, err := time.Parse("2006-01-02", value); err == nil {
				dueDate = &parsedTime
			} else if parsedTime, err := time.Parse("2006-01-02 15:04", value); err == nil {
				dueDate = &parsedTime
			}
		}
	}

	if err := tm.UpdateTodo(id, title, description, category, priority, dueDate); err != nil {
		fmt.Printf("오류: %v\n", err)
	}
}

func showUsage() {
	fmt.Println("사용법: go run main.go <명령어> [옵션]")
	fmt.Println("명령어 목록: add, list, complete, delete, update, stats, help")
	fmt.Println("자세한 도움말: go run main.go help")
}

func showHelp() {
	fmt.Println(`=== TODO 앱 도움말 ===

명령어:
  add <제목> [옵션]     - TODO 추가
  list [옵션]          - TODO 목록 보기
  complete <ID>        - TODO 완료 처리
  delete <ID>          - TODO 삭제
  update <ID> [옵션]   - TODO 수정
  stats               - 통계 보기
  help                - 이 도움말

add 옵션:
  -desc <설명>        - 설명 추가
  -category <카테고리> - 카테고리 설정
  -priority <우선순위> - 우선순위 (low/medium/high/critical)
  -due <날짜>         - 마감일 (2006-01-02 또는 2006-01-02 15:04)

list 옵션:
  -all               - 완료된 항목도 표시
  -category <카테고리> - 카테고리로 필터링
  -priority <우선순위> - 우선순위로 필터링
  -sort <정렬기준>    - 정렬 (id/priority/created/due)

update 옵션:
  -title <제목>       - 제목 변경
  -desc <설명>        - 설명 변경
  -category <카테고리> - 카테고리 변경
  -priority <우선순위> - 우선순위 변경
  -due <날짜>         - 마감일 변경

예시:
  go run main.go add "Go 공부하기" -desc "Go 언어 기초 학습" -priority high -due 2024-12-31
  go run main.go list -all -sort priority
  go run main.go complete 1
  go run main.go update 1 -title "Go 고급 학습" -priority critical`)
}
