package main

import (
	"fmt"
)

func main() {
	var taskOne = "This is task One"
	var taskTwo = "This is task Three"
	var taskThr = "This is task three"

	var taskList = []string{taskOne, taskTwo, taskThr}

	fmt.Println("--- Printing task List ---")
	printTasks(taskList)

	taskList = addTask(taskList, "This is task Four")

	fmt.Println("Now Printing updated task List")
	printTasks(taskList)
}

func printTasks(listOfTasks []string) {
	for index, task := range listOfTasks {
		fmt.Printf("%d: %s\n", index+1, task)
	}
}

func addTask(taskList []string, newTask string) []string {
	var updatedTaskList = append(taskList, newTask)
	return updatedTaskList
}
