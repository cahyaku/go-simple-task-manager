package main

import (
	"fmt"
)

type Task struct {
	Title string
	Done  bool
}

func main() {
	var tasks []Task

	for {
		showMenuList()
		input, err := ReadInt("Enter your choice: ")
		if err != nil {
			fmt.Println("Please input number between 1 and 5 👺")
			continue
		}

		switch input {
		case 1:
			fmt.Println("====== Show tasks =====")
			showTasks(tasks)
			fmt.Println()
		case 2:
			addTask(&tasks)
			fmt.Println()

		case 3:
			markTaskAsDone(&tasks)
			fmt.Println()

		case 4:
			deleteTask(&tasks)
			fmt.Println()

		case 5:
			fmt.Println("Exit")
			fmt.Println("Thank you for using this app bye byeee.... 👺")
			return
		default:
			fmt.Println("Please input number between 1 and 5 👺👺👺👺👺")
			fmt.Println()
			continue
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
}

/**
 * Function to add tasks
 */
func addTask(tasks *[]Task) {
	fmt.Println("===== Add task =====")
	taskName := ReadLine("Enter task title: ")

	*tasks = append(*tasks, Task{
		Title: taskName,
		Done:  false,
	})
	fmt.Println("Task added successfully ✔")
}

func showTasks(tasks []Task) {

	if isEmptyTasks(tasks) {
		fmt.Println("Empty task list, nothing to show")
	}

	fmt.Println("Task list:")
	for i, v := range tasks {
		taskStatus := "[]"
		if v.Done {
			taskStatus = "[X]"
		}

		fmt.Printf("%d. %s %s\n", i+1, v.Title, taskStatus)
	}
}

func markTaskAsDone(tasks *[]Task) {

	if isEmptyTasks(*tasks) {
		fmt.Println("Empty task list, nothing to mark as done!")
	}

	fmt.Println("===== Mark task as done =====")
	showTasks(*tasks)

	taskId, err := ReadInt("Enter task id: ")
	if err != nil {
		fmt.Println("Task id must be a number.")
		return
	}

	if taskId > len(*tasks) || taskId <= 0 {
		fmt.Println("Invalid task id, please try again.")
		return
	}

	(*tasks)[taskId-1].Done = true

	fmt.Println((*tasks)[taskId-1].Title, (*tasks)[taskId-1].Done)
	fmt.Println("Task marked as done successfully ✔")
}

func deleteTask(tasks *[]Task) {
	if isEmptyTasks(*tasks) {
		fmt.Println("Empty task list, nothing to delete!")
	}

	fmt.Println("===== Delete task =====")
	showTasks(*tasks)

	taskId, err := ReadInt("Enter task id: ")
	if err != nil {
		fmt.Println("Task id must be a number.")
		return
	}

	if taskId > len(*tasks) || taskId <= 0 {
		fmt.Println("Invalid task id, please try again.")
		return
	}

	// Harus pakai append karena slice tidak bisa dihapus langsung
	*tasks = append((*tasks)[:taskId-1], (*tasks)[taskId:]...)
	fmt.Println("Successfully deleted task", taskId)
}

/**
 * Function to check whether there is data
 */
func isEmptyTasks(tasks []Task) bool {
	return len(tasks) == 0
}
