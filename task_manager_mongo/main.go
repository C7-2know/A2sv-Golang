package main

import (
	"fmt"
	"log"
	"task_manager/data"
	"task_manager/router"
)

func init() {
	fmt.Println("init runs")
	err := data.Dbconnect()
	if err != nil {
		fmt.Println("connection error")
		// panic(err)
		log.Fatal(err)
	}
}

func main() {
	router.Route()
}
