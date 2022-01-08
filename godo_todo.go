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

//Creating a struct of type TODO
type TODO struct {
	Text string
	Date string
}

//Creating a struct of type DB
type DB struct {
	Inprogress []TODO
	Completed  []TODO
	Forlater   []TODO
}

func ReadDB(filename string) (db DB) {
	//Reading
	byteData, err := ioutil.ReadFile(filename)
	// var db DB
	if err != nil {
		fmt.Printf("\nFailed to read the data from %s : %s", filename, err)
	}

	//Using unmarshal to convert byte to json
	json.Unmarshal(byteData, &db)
	return db
}

func CreateNewTask(filename string, db DB) {
	//Writing -> Modifying the data content
	var new_todo TODO
	//Creating a new todo
	fmt.Printf("New Task : ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	str_input := strings.TrimSpace(input)
	new_todo.Text = str_input
	// new_todo.Text = str_input
	presentTime := time.Now()
	new_todo.Date = presentTime.Format("01-02-2006 15:04:05")

	//appending the new todo to the list and encoding it to marshal
	db.Inprogress = append(db.Inprogress, new_todo)
	db_marshal, err := json.Marshal(db)
	if err != nil {
		fmt.Printf("\nFailed to convert to Marshal encoding : %s", err)
	}

	//Finally writing the file back to the db.json
	err = ioutil.WriteFile(filename, db_marshal, 0)
	if err != nil {
		fmt.Printf("\nFailed to Write data to %s : %s", filename, err)
	}
	fmt.Printf("\nNew Task Created Successfully\n -> %s", new_todo.Text)
}

func ShowTasks(filename string) {
	// status_idx := map[interface{}]interface{}{
	// 	1: "Inprogress",
	// 	2: "Completed",
	// 	3: "Forlater",
	// }
	// fmt.Printf("1.Inprogress  2. Completed  3.Forlater\nFor: ")
	// reader := bufio.NewReader(os.Stdin)
	// input, _ := reader.ReadString('\n')
	// index, _ := strconv.ParseInt(strings.TrimSpace(input), 0, 64)
	// index -= 1
	var io_db = ReadDB(filename)
	// fmt.Printf("%s Tasks", status_idx[index])
	inprogress := io_db.Inprogress

	for i := 0; i < len(inprogress); i++ {
		fmt.Printf("%d). %s\n", i+1, inprogress[i].Text)
	}
}

func MoveToCompleted(filename string, db DB) {
	//Writing -> Modifying the data content
	fmt.Printf("Task index which you want to mark as completed : ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	index, _ := strconv.ParseInt(strings.TrimSpace(input), 0, 64)
	index -= 1
	//Moving from inprogress to completed
	db.Completed = append(db.Completed, db.Inprogress[index])
	//Removing it from inprogress
	fmt.Printf("Moved to Completed,\n-> %s", db.Inprogress[index].Text)
	db.Inprogress = append(db.Inprogress[:index], db.Inprogress[index+1:]...)

	db_marshal, err := json.Marshal(db)
	if err != nil {
		fmt.Printf("\nFailed to convert to Marshal encoding : %s", err)
	}

	//Finally writing the file back to the db.json
	err = ioutil.WriteFile(filename, db_marshal, 0)
	if err != nil {
		fmt.Printf("\nFailed to Write data to %s : %s", filename, err)
	}
}

func MoveToForLater(filename string, db DB) {
	//Writing -> Modifying the data content
	fmt.Printf("Task index which you want to save for later : ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	index, _ := strconv.ParseInt(strings.TrimSpace(input), 0, 64)
	index -= 1
	//Moving from inprogress to completed
	db.Forlater = append(db.Forlater, db.Inprogress[index])
	//Removing it from inprogress
	fmt.Printf("Moved to ForLater,\n-> %s", db.Inprogress[index].Text)
	db.Inprogress = append(db.Inprogress[:index], db.Inprogress[index+1:]...)

	db_marshal, err := json.Marshal(db)
	if err != nil {
		fmt.Printf("\nFailed to convert to Marshal encoding : %s", err)
	}

	//Finally writing the file back to the db.json
	err = ioutil.WriteFile(filename, db_marshal, 0)
	if err != nil {
		fmt.Printf("\nFailed to Write data to %s : %s", filename, err)
	}
}

func main() {
	filename := "db.json"
	fmt.Println("Creating a todo app using golang")
	var io_db = ReadDB(filename)
	// fmt.Println(io_db)
	for {
		fmt.Println("\n1. Create")
		fmt.Println("2. Show Tasks")
		fmt.Println("3. Move to Completed")
		fmt.Println("4. Move to For Later")
		fmt.Println("5. Exit")
		fmt.Printf("\nEnter the Choice : ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		value, _ := strconv.ParseInt(strings.TrimSpace(input), 0, 64)

		switch value {
		case 1:
			CreateNewTask(filename, io_db)
		case 2:
			ShowTasks(filename)
		case 3:
			MoveToCompleted(filename, io_db)
		case 4:
			MoveToForLater(filename, io_db)
		case 5:
			fmt.Println("Bye✌")
			os.Exit(3)
		default:
			fmt.Println("Invalid")
		}
	}
}
