package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"time"
	"strconv"
	"encoding/json"
	"io/ioutil"
)

type Todo struct {
	ID		int       `json:"id"`
	Text 	string    `json:"text"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

// TodoApp manages our todo list
type TodoApp struct {
	Todos []Todo `json:"todos"`
	NextID int    `json:"next_id"`
	filename string 
}

// NewTodoApp creates a new TodoApp instance
func NewTodoApp(filename string) *TodoApp {
	app := &TodoApp{
		Todos: []Todo{},
		NextID: 1,
		filename: filename,
	}
	app.LoadFromFIle()
	return app
}

// AddTodo adds a new todo item
func (app *TodoApp) AddTodo(text string) {
	todo := Todo{
		ID: app.NextID,
		Text: strings.TrimSpace(text),
		Completed: false,
		CreatedAt: time.Now(),
	}

	app.Todos = append(app.Todos, todo)
	app.NextID++
	fmt.Printf("✅ Added todo #%d: %s\n", todo.ID, todo.Text)
}

// ListTodos displays all todos
func (app *TodoApp) ListTodos() {
	if len(app.Todos) == 0 {
		fmt.Println("No todos found. add one with 'add <description>'")
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

// completeTodo marks a todo as completed
func (app *TodoApp) CompletedTodo(id int) {
	for i, todo := range app.Todos {
		if todo.ID == id {
			if app.Todos[i].Completed {
				fmt.Printf("Todo #%d is already completed.\n", id)
				return
		}
		app.Todos[i].Completed = true
		app.SaveToFile()
		fmt.Printf("✅ Todo #%d marked as completed!\n", id)
		return
	}
}
  fmt.Printf("Todo with ID %d not found.\n", id)
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
	fmt.Printf("Todo with ID %d not found.\n", id)
	}

	// SaveToFile persists todos to JSON
	func (app *TodoApp) SaveToFile() {
		data, err := json.MarshalIndent(app, "", "  ")
		if err != nil {
			fmt.Println("Error saving todos:", err)
			return
		}
		err = ioutil.WriteFile(app.filename, data, 0644)
		if err != nil {
			fmt.Println("Error writing to file:", err)
		}
	}

	// LoadFromFile loads todos from JSON file
	func (app *TodoApp) LoadFromFIle() {
		data, err := ioutil.ReadFile(app.filename)
		if err != nil {
			return // File may not exist yet, ignore error
		}
		err = json.Unmarshal(data, app)
		if err !=nil {
			fmt.Println("Error loading todos:", err)
		}
	}		