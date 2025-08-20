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

