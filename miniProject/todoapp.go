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