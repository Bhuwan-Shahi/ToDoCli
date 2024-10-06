package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

// SLICE TOD
type Todos []Todo

// FUNCTIONS
func (todos *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
	*todos = append(*todos, todo)

}

func (todos *Todos) validateINdex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("Invalid Index")
		fmt.Println(err)
		return err
	}
	return nil //no error if valid indext
}

// DELETING THE TODOS
func (todos *Todos) delete(index int) error {
	t := todos

	if err := t.validateINdex(index); err != nil {
		return err
	}
	*todos = append((*todos)[:index], (*todos)[index+1:]...)
	return nil
}

// TOGGOLE BETWEEN ❌,✅
func (todos *Todos) toggle(index int) error {
	t := todos

	if err := t.validateINdex(index); err != nil {
		return err
	}

	isCompleted := (*todos)[index].Completed
	if !isCompleted {
		completionTime := time.Now()
		(*todos)[index].CompletedAt = &completionTime
	}

	(*todos)[index].Completed = !isCompleted

	return nil
}

// EDITING THE TODOS
func (todos *Todos) edit(index int, title string) error {
	t := *todos
	if err := t.validateINdex(index); err != nil {
		return err
	}
	t[index].Title = title
	return nil
}

// PRINTING ALL THE TODOS
func (todos *Todos) print() {
	table := table.New(os.Stdout)

	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "Created AT", "Completed At")
	for index, t := range *todos {
		completed := "❌"
		completedAt := ""

		if t.Completed {
			completed = "✅"
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format(time.RFC1123)
			}
		}

		table.AddRow(strconv.Itoa(index), t.Title, completed, t.CreatedAt.Format(time.RFC1123), completedAt)
	}
	table.Render()
}
