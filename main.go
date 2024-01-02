package main

import "fmt"

func main() {
	// declare variable

	// 1. zero value
	// var i int
	// var s string
	// var ok bool

	// 2. assign value
	// var i int = 13
	// var s string = "hello"
	// var ok bool = true

	// 3. go will ausume data type
	// var i = 13
	// var s = "hello"
	// var ok = true

	// 4. use colon (only using in function)
	// i := 13
	// s := "hello"
	// ok := true

	// constant
	// const defaultValue int = 1
	// const defaultTitle = "GO"

	// use iota. every number after first value will plus one
	// const (
	// 	sunday = iota
	// 	monday
	// 	tuesday
	// 	wednesday
	// 	thursday
	// 	friday
	// 	saturday
	// )

	fmt.Println("Hello world, Again :)")
}

	// function

	// no return
	func greeting(firstName, lastName string) {
		fmt.Println("Hello", firstName, lastName)
	}

	// return 1 value
	func add(a, b int) int {
		return a + b
	}

	// return 2 value
	func swap(a, b int) (int, int) {
		return b, a
	}
