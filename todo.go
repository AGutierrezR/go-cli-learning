package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type item struct {
	Task string
	Done bool
	CreateAt time.Time
	CompletedAt time.Time
}

type List []item

func (l *List) Add(task string) {
	t := item{
		Task: task,
		Done: false,
		CreateAt: time.Now(),
		CompletedAt: time.Time{},
	}

	*l = append(*l, t)
}

func (l *List) Complete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("item %d does not exist", i)
	}

	// Adjusting index for 0 based index
	ls[i-1].Done = true
	ls[i-1].CompletedAt = time.Now()

	return nil
}

// Delete method deletes a ToDo item from the list
func (l *List) Delete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt. Errorf("item %d does not exist", i)
	}

	//I Adjusting index for 0 based index
	*l = append (ls[:i-1], ls[i:]...)
	
	return nil
}

// Save method encodes the List as JSON and saves it
// using the provided file name
func (l *List) Save(filename string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, js, 0644)
}

// Get method opens the provided file name, decodes // the JSON data and parses it into a List
func (l *List) Get(filename string) error {
	// If the file does not exist will return and err
	file, err := os.ReadFile(filename)

	if err != nil {
		// Check if the file does not exists
		if errors.Is(err, os.ErrNotExist) { 
			return nil
		}
		return err 
	}

	// If files does not have bytes returns nill
	if len(file) == 0 { 
		return nil
	}

	// Store the content of the json in the pointer l, this returns a nill
	return json.Unmarshal(file, l) 
}