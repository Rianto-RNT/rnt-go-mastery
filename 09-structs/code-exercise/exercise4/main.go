package main

import "fmt"

// 2. Add another field of type grades to person struct type (embedded struct).
type person struct{
	name string
	age int
	colors []string
	gr grades
}

// 1. Create a new struct type called grades with  2 fields: grade and course
type grades struct {
	grade  int
	course string
}

func main() {
	// 3. Change the initialization of the struct values called me and you to contain information about the grades.
	me := person{
		name : "Morrison",
		age: 29,
		colors: []string{"Black", "White"},
		gr: grades{grade: 88, course: "Java"},
	}

	you := person{
		name: "Jet",
		age: 20,
		colors: []string{"pink", "blue"},
		gr: grades{grade: 100, course: "Biography"},
	}

	// 4. Change the grade and the course of one struct value to "C#" and 99.
	me.gr.grade = 99
	me.gr.course = "C#"

	// 5. Print out the struct values.
	fmt.Printf("%v\n", me)
	fmt.Printf("%+v\n", you)
}