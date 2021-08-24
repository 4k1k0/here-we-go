package main

import "fmt"

func hello() {
	fmt.Println("hello")
}

func world() {
	fmt.Println("world")
}

func main() {
	go hello()
	go world()
	fmt.Println("Main function")
}
