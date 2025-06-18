package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

type task struct {
	date  string
	tasks string
}

func checkFile() (bool, error) {
	tasks := task{date: "17/6/2025", tasks: "First test"}
	json_file, err := json.Marshal(tasks.tasks)
	if err != nil {
		return false, err
	}
	_ = json_file
	fmt.Printf("json_file: %s", json_file)
	fmt.Println(tasks.date)
	fmt.Println(tasks.tasks)
	return true, nil
}

// func markTask() {
//
// }
// func displayTask() {
//
// }
func testFile(t *testing.T) {
	state, err := checkFile()
	if state != true || err == nil {
		t.Errorf(`Hello("") = %t, %v, want "", error`, state, err)
	}
}
func main() {
	checkFile()
	msg := "Hello world"
	log := 12
	fmt.Printf("%s", msg)
	fmt.Printf("%d", log)
}
