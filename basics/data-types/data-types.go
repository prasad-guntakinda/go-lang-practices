package main

import (
	"fmt"
	"strings"
)

func main() {
	var name string = "John"
	var age uint = 25
	var avgScore float32 = 95.5
	var isPassed bool = true

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Average Score:", avgScore)
	fmt.Println("Passed:", isPassed)

	name2 := "Sam"
	age2 := 30
	avgScore2 := 85.5
	isPassed2 := true

	name3, age3, avgScore3, isPassed3 := "Ram", 18, 34.5, false

	fmt.Println("Student Details")
	fmt.Println(strings.Repeat("-", 15))
	fmt.Println(name, age, avgScore, isPassed)
	fmt.Println(name2, age2, avgScore2, isPassed2)
	fmt.Println(name3, age3, avgScore3, isPassed3)

	fmt.Println(strings.Repeat("=", 15))
	fmt.Println(strings.Repeat("\n", 3))
	fmt.Println("====== Pointers============")
	a := 10
	var b *int
	b = &a
	c := b
	fmt.Println("a=", a, "b=", *b, "c=", *c)
	a = 20
	fmt.Println("a=", a, "b=", *b, "c=", *c)

	var d int

	//d = b //cannot use b (variable of type *int) as int value in assignment
	d = *b // this works
	fmt.Println("d=", d)

	e := new(int)
	//e = 10 // cannot use 10 (type int) as type *int in assignment it's expected to be a pointer not a constant
	e = &a
	fmt.Println("e=", *e)

}
