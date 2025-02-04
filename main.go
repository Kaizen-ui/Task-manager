package main

import "fmt"

func main() {

	// Concept Focus: Array declaration, initialization, and fixed-size constraints. and Accessing/modifying elements by index.
	// Array for Daily Task max 5 values
	dailyTasks := [5]string{"Code", "Readbooks", "Workout"}

	// add daily task
	dailyTasks[3] = "Wash plates"
	dailyTasks[4] = "Feed the Fish"

	fmt.Printf("Daily Tasks:\n[0] %v \n[1] %v \n[2] %v \n[3] %v \n[4] %v\n\n\n", dailyTasks[0], dailyTasks[1], dailyTasks[2], dailyTasks[3], dailyTasks[4])

	// Call the weekly function
	weeklyTask()
}

func weeklyTask() {

	//Concept Focus: Slice declaration, dynamic resizing with append(), and index-based operations.
	// Slice for Weekly task
	weekly := []string{"Plan Week", "Grocery Shopping", "Call family", "Cook dinner", "Write journal"}

	// Delete the index 2 of the slice "Call Family"
	delete := 2
	weekly = append(weekly[:delete], weekly[delete+1:]...)
	// Insert "Relax between Cook dinner and Write Journal"
	insertTask := "Relax"
	toInsert := 3
	weekly = append(weekly[:toInsert], append([]string{insertTask}, weekly[toInsert:]...)...)

	fmt.Printf("Weekly Tasks (First 3):\n[0] %v\n[1] %v\n[3] %v", weekly[0], weekly[1], weekly[3])
}
