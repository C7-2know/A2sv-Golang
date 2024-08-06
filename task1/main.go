package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to grade calulator\n")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Println("How many courses do you take?")
	var num int
	fmt.Scanf("%d\n", &num)
	subject_grade := make(map[string]float64)
	var grades []float64
	for n := 0; n < num; n++ {
		fmt.Println(n)
		fmt.Print("Please enter the course name:")
		course, _ := reader.ReadString('\n')
		course = strings.TrimSpace(course)

		fmt.Println("Please enter your " + course + " grade: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		grade, err:= strconv.ParseFloat(input, 64)
		if grade < 0 || grade > 100 || err != nil {
			fmt.Println("Please enter a valid grade")
			input, _ = reader.ReadString('\n')
			input = strings.TrimSpace(input)
			grade, err= strconv.ParseFloat(input, 64)
			if grade < 0 || grade > 100 || err != nil  {
				fmt.Println("Please enter a valid grade")
				n--
				continue
			}
		}

		fmt.Println(grade,input)
		subject_grade[course] = grade
		grades = append(grades, grade)
	}

	avrage := calculateAverage(grades)

	fmt.Println("Hello " + name)
	for k, v := range subject_grade {
		fmt.Println("your grade in " + k + " is " + fmt.Sprintf("%.2f", v))
	}
	fmt.Println("your avrage grade is " + strconv.Itoa(avrage))
}

func calculateAverage(grades []float64) int {
	total := 0.0
	for i := range grades {
		total += grades[i]
	}
	return int(total / float64(len(grades)))
}
