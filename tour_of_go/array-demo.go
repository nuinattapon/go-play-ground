package main

import "fmt"

func main() {
	// grades := [...]int{97, 85, 93}
	// fmt.Printf("Grades: %v\n", grades)
	var students [5]string
	students[0] = "Lisa"
	students[1] = "Ahmed"
	students[2] = "Arnold"
	fmt.Printf("students: %v, %T\n", students, students)
	fmt.Printf("students[1]: %v, %T\n", students[1], students[1])
	fmt.Printf("Number of students: len %v, cap %v\n", len(students), cap(students))

}
