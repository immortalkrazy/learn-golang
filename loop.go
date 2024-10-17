package main

import (
	"fmt"
)

func main() {
	var taskOne = "This is task One"
	var taskTwo = "This is task Three"
	var taskThr = "This is task three"

	fmt.Println(taskOne, taskTwo, taskThr)

	var taskList = []string{taskOne, taskTwo, taskThr}

	fmt.Println("Items: ")
	var items = []string{"itemOne", "itemTwo", "itemThree"}

	fmt.Println(items)

	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	fmt.Println("--- Now print items and List ---")
	for index, item := range items {
		fmt.Printf("%d: %s\n", index+1, item)
	}

	fmt.Println("--- Now Pring task List ---")
	for index, task := range taskList {
		fmt.Printf("%d: %s\n", index+1, task)
	}
}
