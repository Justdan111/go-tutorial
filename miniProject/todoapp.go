package main

import (
	"fmt"
	"os"
	"strconv"
	"bufio"
	"strings"
	"time"
	"io/ioutil"
	"encoding/json"
)


// TodoItem represents a single todo item
type Todo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

// TodoApp manages todo items
type TodoApp struct {
	Todos []Todo `json:"todos"`
	NextID int    `json:"next_id"`
	filename string
}

func main() {

}