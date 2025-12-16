package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	Title string
	Done  bool
}

func main() {
	var tasks []Task

	for {
		showMenuList()
		var input int
		fmt.Scanln(&input)

		switch input {
		case 1:
			fmt.Println("====== Show tasks =====")
			showTasks(tasks)
			fmt.Println()
		case 2:
			addTask(&tasks)
			fmt.Println()

		case 3:
			markTaskAsDone(tasks)
			fmt.Println()

		case 4:
			deleteTask(tasks)
			fmt.Println()

		case 5:
			fmt.Println("Exit")
			fmt.Println("Thank you for using this app bye byeee.... 👺")
			return
		default:
			fmt.Println("Please input number between 1 and 5 👺👺👺👺👺")
			fmt.Println()
		}
	}
}

/**
 * Function to display the list menu
 */
func showMenuList() {
	fmt.Println("*********************************************")
	fmt.Println("𐔌՞꜆.  ̫.꜀՞𐦯 SIMPLE TASK MANAGER 𐔌՞꜆.  ̫.꜀՞𐦯 ‪‪❤︎‬ ♡")
	fmt.Println("*********************************************")
	fmt.Println("Menu:")
	fmt.Println("1. Show tasks")
	fmt.Println("2. Add task")
	fmt.Println("3. Mark task as done")
	fmt.Println("4. Delete task")
	fmt.Println("5. Exit")
	fmt.Print("Enter your choice: ")
}

/**
 * Function to add tasks
 */
func addTask(tasks *[]Task) {
	fmt.Println("===== Add task =====")

	//var taskName string
	taskName := readInput("Enter task title: ")
	// kalau pakai scan hanya 1 kata awal yang tersimpan
	//fmt.Scanln(&taskName)

	//tasks = append(tasks, Task{Title: taskName, Done: false})

	*tasks = append(*tasks, Task{
		Title: taskName,
		Done:  false,
	})
	fmt.Println("Task added successfully ✔")
}

func showTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("Empty task list, nothing to show!")
	}

	fmt.Println("Task list:")
	for i, v := range tasks {
		taskStatus := "[]"
		if v.Done {
			taskStatus = "[X]"
		}
		//fmt.Printf("%d. %s (done: %t)\n", i+1, v.Title, v.Done)
		fmt.Printf("%d. %s %s\n", i+1, v.Title, taskStatus)
	}
}

func markTaskAsDone(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("Empty task list, nothing to mark as done!")
		return
	}

	fmt.Println("===== Mark task as done =====")
	showTasks(tasks)
	//var taskId int
	//fmt.Println("Enter task id: ")
	//fmt.Scanln(&taskId)

	input := readInput("Enter task id: ")
	taskId, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("Task id must be a number.")
		return
	}

	if taskId > len(tasks) || taskId <= 0 {
		fmt.Println("Invalid task id, please try again.")
		return
	} else {
		tasks[taskId-1].Done = true
	}
	fmt.Println(tasks[taskId-1].Title, tasks[taskId-1].Done)
	fmt.Println("Task marked as done successfully ✔")
}

func deleteTask(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("Empty task list, nothing to delete!")
		return
	}

	fmt.Println("===== Delete task =====")
	showTasks(tasks)
	//fmt.Println("Enter task id: ")
	//var taskId int
	//fmt.Scanln(&taskId)

	input := readInput("Enter task id: ")
	taskId, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("Task id must be a number.")
		return
	}

	if taskId > len(tasks) || taskId <= 0 {
		fmt.Println("Invalid task id, please try again.")
		return
	}
	// Harus pakai append karena slice tidak bisa dihapus langsung
	tasks = append(tasks[:taskId-1], tasks[taskId:]...)
	fmt.Println("Successfully deleted task", taskId)
}

/**
 * Ini untuk mendapatkan jumlah inputan lebih dari 1 kata.
 */
var reader = bufio.NewReader(os.Stdin)

func readInput(prompt string) string {
	fmt.Print(prompt)

	// ini membaca input satu baris penuh
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input) // menghapus spasi di awal dan akhir
}
