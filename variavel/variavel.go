package main

import "fmt"

var x = "Hello World!" //Escopo global

func main() { //Escopo da funcao
	//var x = "Hello World!" //var x string = "Hello World!"
	fmt.Println(x)
}
