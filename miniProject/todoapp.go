package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

// Todo represents a single todo item
type Todo struct {
	ID        int       `json:"id"`
	Text      string    `json:"text"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

// TodoApp manages our todo list
type TodoApp struct {
	Todos    []Todo `json:"todos"`
	NextID   int    `json:"next_id"`
	filename string
}

// NewTodoApp creates a new TodoApp instance
func NewTodoApp(filename string) *TodoApp {
	app := &TodoApp{
		Todos:    []Todo{},
		NextID:   1,
		filename: filename,
	}
	app.LoadFromFile() // load existing todos from file if available
	return app
}

// AddTodo adds a new todo item
func (app *TodoApp) AddTodo(text string) {
	todo := Todo{
		ID:        app.NextID,
		Text:      strings.TrimSpace(text),
		Completed: false,
		CreatedAt: time.Now(),
	}

	app.Todos = append(app.Todos, todo)
	app.NextID++
	app.SaveToFile()

	fmt.Printf("✅ Added todo #%d: %s\n", todo.ID, todo.Text)
}

// ListTodos displays all todos
func (app *TodoApp) ListTodos() {
	if len(app.Todos) == 0 {
		fmt.Println("📝 No todos yet! Add one with 'add <description>'")
		return
	}

	fmt.Println("\n📋 Your Todos:")
	fmt.Println("=" + strings.Repeat("=", 50))

	for _, todo := range app.Todos {
		status := "⬜"
		if todo.Completed {
			status = "✅"
		}

		fmt.Printf("%s #%d: %s\n", status, todo.ID, todo.Text)
		fmt.Printf("   📅 Created: %s\n", todo.CreatedAt.Format("2006-01-02 15:04:05"))
	}
	fmt.Println("=" + strings.Repeat("=", 50))
}

// CompleteTodo marks a todo as completed
func (app *TodoApp) CompleteTodo(id int) {
	for i, todo := range app.Todos {
		if todo.ID == id {
			if app.Todos[i].Completed {
				fmt.Printf("⚠️ Todo #%d is already completed.\n", id)
				return
			}
			app.Todos[i].Completed = true
			app.SaveToFile()
			fmt.Printf("🎉 Todo #%d marked as completed!\n", id)
			return
		}
	}
	fmt.Printf("❌ Todo with ID %d not found.\n", id)
}

// DeleteTodo removes a todo by ID
func (app *TodoApp) DeleteTodo(id int) {
	for i, todo := range app.Todos {
		if todo.ID == id {
			app.Todos = append(app.Todos[:i], app.Todos[i+1:]...)
			app.SaveToFile()
			fmt.Printf("🗑️ Deleted todo #%d: %s\n", id, todo.Text)
			return
		}
	}
	fmt.Printf("❌ Todo with ID %d not found.\n", id)
}

// SaveToFile persists todos to JSON
func (app *TodoApp) SaveToFile() {
	data, err := json.MarshalIndent(app, "", "  ")
	if err != nil {
		fmt.Println("❌ Error saving todos:", err)
		return
	}
	err = ioutil.WriteFile(app.filename, data, 0644)
	if err != nil {
		fmt.Println("❌ Error writing file:", err)
	}
}

// LoadFromFile loads todos from JSON file
func (app *TodoApp) LoadFromFile() {
	data, err := ioutil.ReadFile(app.filename)
	if err != nil {
		return // file may not exist yet, that's fine
	}
	err = json.Unmarshal(data, app)
	if err != nil {
		fmt.Println("❌ Error loading todos:", err)
	}
}

// Command loop for CLI
func main() {
	app := NewTodoApp("todos.json")
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("🚀 Simple Go Todo App")
	fmt.Println("Commands: add <task>, list, complete <id>, delete <id>, quit")

	for {
		fmt.Print("\n👉 Enter command: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			continue
		}

		parts := strings.SplitN(input, " ", 2)
		command := parts[0]

		switch command {
		case "add":
			if len(parts) < 2 {
				fmt.Println("⚠️ Please provide a todo description.")
				continue
			}
			app.AddTodo(parts[1])

		case "list":
			app.ListTodos()

		case "complete":
			if len(parts) < 2 {
				fmt.Println("⚠️ Please provide the todo ID to complete.")
				continue
			}
			id, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("⚠️ Invalid ID. Please enter a number.")
				continue
			}
			app.CompleteTodo(id)

		case "delete":
			if len(parts) < 2 {
				fmt.Println("⚠️ Please provide the todo ID to delete.")
				continue
			}
			id, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("⚠️ Invalid ID. Please enter a number.")
				continue
			}
			app.DeleteTodo(id)

		case "quit", "exit":
			fmt.Println("👋 Goodbye!")
			return

		default:
			fmt.Println("❓ Unknown command. Try: add, list, complete, delete, quit")
		}
	}
}
