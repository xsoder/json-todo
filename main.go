package main

import (
	"encoding/json"
	"fmt"
	"testing"
    "os"
)

type task struct {
	date  string
	tasks string
}

func addTask() {

}

func deleteTask() {

}
package main

import (
	"encoding/json"
	"fmt"
	"testing"
    "os"
)

type task struct {
	date  string
	tasks string
}

func addTask() {

}

func deleteTask() {

}

func markTask() {

}

//func displayTask() {


func checkFile() (bool, error) {
	tasks := task{date: "17/6/2025", tasks: "First test"}
	json_file, err := json.Marshal(tasks.tasks)
	if err != nil {
		return false, err
	}
	return true, nil
}



//func initTask() {


// func markTask() {

// }
// func displayTask() {

// }
func testFile(t *testing.T) {
	state, err := checkFile()
	if state != true || err == nil {
		t.Errorf(Hello("") = %t, %v, want "", error, state, err)
	}
}

func main() {
    if len (os.Args) < 2 {
        checkFile()
    } else {
        switch os.Args[1] {
            case "add":
                addTask()

            case "delete":
                deleteTask()

            case "mark":
                markTask()

            case "display":
                displayTask()

            case "init":
                initTask()

        }
    }
}
func markTask() {

}

//func displayTask() {


func checkFile() (bool, error) {
	tasks := task{date: "17/6/2025", tasks: "First test"}
	json_file, err := json.Marshal(tasks.tasks)
	if err != nil {
		return false, err
	}
	return true, nil
}



//func initTask() {


// func markTask() {

// }
// func displayTask() {

// }
func testFile(t *testing.T) {
	state, err := checkFile()
	if state != true || err == nil {
		t.Errorf(Hello("") = %t, %v, want "", error, state, err)
	}
}

func main() {
    if len (os.Args) < 2 {
        checkFile()
    } else {
        switch os.Args[1] {
            case "add":
                addTask()

            case "delete":
                deleteTask()

            case "mark":
                markTask()

            case "display":
                displayTask()

            case "init":
                initTask()

        }
    }
}
