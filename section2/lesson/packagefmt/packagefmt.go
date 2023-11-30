package main

import "fmt"

var a int
var b string = "Programa"
var c bool
var d bool = true

func main() {
	e := 42
	f := "dice hola mundo"
	g := `Esto es un string`
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println("Mi", b, f)
	fmt.Println(g)

}
