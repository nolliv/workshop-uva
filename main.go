// Comentários iguais aos do C
/* assim também */

// Organização do código por pacotes
package main

import "fmt"

func main() {
	// exported em maiúsculo
	// UNICODE ALL THINGS!!!!!!1!!
	fmt.Println("Hello, Modafukás!")
}

// O Tipo é a frente do nome!
func add(x int, y int) int {
	return x + y
}

// Pode omitir tipo repetido
func add2(x, y int) int {
	return x + y
}

// multiple return values
func swap(x, y string) (string, string) {
	return y, x
}

// variable declaration
func printbool() {
	// inicializado com zero value(false)
	var x bool
	y := true
	var z = true
	fmt.Print(x, y, z) // faltou o z!
}

// loops
func loops() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
