package main

import "fmt"

func init() {
	initConfiguration()
}

func main() {
	fmt.Println("Hello world")
	fmt.Println(config.dict)

}
