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

//
// // O Tipo é a frente do nome!
// func add(x int, y int) int {
// 	return x + y
// }
//
// // Pode omitir tipo repetido
// func add2(x, y int) int {
// 	return x + y
// }
//
// // multiple return values
// func swap(x, y string) (string, string) {
// 	return y, x
// }
//
// // variable declaration
// func printbool() {
// 	// inicializado com zero value(false)
// 	var x bool
// 	y := true
// 	var z = true
// 	fmt.Print(x, y, z) // faltou o z!
// }
//
// // loops
// func loops() {
// 	sum := 0
// 	for i := 0; i < 10; i++ {
// 		sum += i
// 	}
// 	fmt.Println(sum)
// }
//
// // for == while
// func main() {
// 	sum := 1
// 	for sum < 1000 {
// 		sum += sum
// 	}
// 	fmt.Println(sum)
// }
//
// // for range!
// var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
//
// func main() {
// 	for i, v := range pow {
// 		fmt.Printf("2**%d = %d\n", i, v)
// 	}
// }

// // if scope!
// func pow(x, n, lim float64) float64 {
// 	if v := math.Pow(x, n); v < lim {
// 		return v
// 	} else {
// 		fmt.Printf("%g >= %g\n", v, lim)
// 	}
// 	// can't use v here, though
// 	return lim
// }
//
// func main() {
// 	fmt.Println(
// 		pow(3, 2, 10),
// 		pow(3, 3, 20),
// 	)
// }
//
// // switch
// func main() {
// 	fmt.Print("Go runs on ")
// 	switch os := runtime.GOOS; os {
// 	case "darwin":
// 		fmt.Println("OS X.")
// 	case "linux":
// 		fmt.Println("Linux.")
// 	default:
// 		// freebsd, openbsd,
// 		// plan9, windows...
// 		fmt.Printf("%s.", os)
// 	}
// }
//
// Ponteiros
// func main() {
// 	i, j := 42, 2701
//
// 	p := &i         // point to i
// 	fmt.Println(*p) // read i through the pointer
// 	*p = 21         // set i through the pointer
// 	fmt.Println(i)  // see the new value of i
//
// 	p = &j         // point to j
// 	*p = *p / 37   // divide j through the pointer
// 	fmt.Println(j) // see the new value of j
// }

// // Structs tructs tructs
// type Vertex struct {
// 	X int
// 	Y int
// }
//
// func main() {
// 	v := Vertex{1, 2}
// 	fmt.Println()
// 	v.X = 4
// 	fmt.Println(v.X)
// 	// v2 = Vertex{X: 1}  // Y:0 is implicit
// 	// v3 = Vertex{}      // X:0 and Y:0
// 	// p  = &Vertex{1, 2} // has type *Vertex
// }

// // Arrays/slices
// func main() {
// 	var a [2]string
// 	a[0] = "Hello"
// 	a[1] = "World"
// 	fmt.Println(a[0], a[1])
// 	fmt.Println(a)
//
// 	primes := [6]int{2, 3, 5, 7, 11, 13}
// 	fmt.Println(primes)
// }
//
// func main() {
// 	primes := [6]int{2, 3, 5, 7, 11, 13}
//
// 	var s []int = primes[1:4]
// 	fmt.Println(s)
// }
//
// // Appending to slices
// func main() {
//   // Cuidado: nil slice não tem um array!
// 	var s []int
// 	printSlice(s)
//
// 	// append works on nil slices.
// 	s = append(s, 0)
// 	printSlice(s)
//
// 	// The slice grows as needed.
// 	s = append(s, 1)
// 	printSlice(s)
//
// 	// We can add more than one element at a time.
// 	s = append(s, 2, 3, 4)
// 	printSlice(s)
// }
//
// func printSlice(s []int) {
// 	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
// }

// // maps
// // maps == dict(mais ou menos)
// type Vertex struct {
// 	Lat, Long float64
// }
//
// // CUIDADO: Nil map!
// var m map[string]Vertex
//
// func main() {
// 	m = make(map[string]Vertex)
// 	m["Bell Labs"] = Vertex{
// 		40.68433, -74.39967,
// 	}
// 	fmt.Println(m["Bell Labs"])
// }

// // map mutations
// func main() {
// 	m := make(map[string]int)
//
// 	m["Answer"] = 42
// 	fmt.Println("The value:", m["Answer"])
//
// 	m["Answer"] = 48
// 	fmt.Println("The value:", m["Answer"])
//
// 	delete(m, "Answer")
// 	fmt.Println("The value:", m["Answer"])
//
// 	v, ok := m["Answer"]
// 	fmt.Println("The value:", v, "Present?", ok)
// }
//

// // Interfaces
// type Abser interface {
// 	Abs() float64
// }
//
// func main() {
// 	var a Abser
// 	f := MyFloat(-math.Sqrt2)
// 	v := Vertex{3, 4}
//
// 	a = f  // a MyFloat implements Abser
// 	a = &v // a *Vertex implements Abser
//
// 	// In the following line, v is a Vertex (not *Vertex)
// 	// and does NOT implement Abser.
// 	a = v
//
// 	fmt.Println(a.Abs())
// }
//
// type MyFloat float64
//
// func (f MyFloat) Abs() float64 {
// 	if f < 0 {
// 		return float64(-f)
// 	}
// 	return float64(f)
// }
//
// type Vertex struct {
// 	X, Y float64
// }
//
// func (v *Vertex) Abs() float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }
//
// // Type Assertions
// func main() {
// 	var i interface{} = "hello"
//
// 	s := i.(string)
// 	fmt.Println(s)
//
// 	s, ok := i.(string)
// 	fmt.Println(s, ok)
//
// 	f, ok := i.(float64)
// 	fmt.Println(f, ok)
//
// 	f = i.(float64) // panic
// 	fmt.Println(f)
// }

// // Goroutines
// func fibonacci(n int, c chan int) {
// 	x, y := 0, 1
// 	for i := 0; i < n; i++ {
// 		c <- x
// 		x, y = y, x+y
// 	}
// 	close(c)
// }
//
// func main() {
// 	c := make(chan int, 10)
// 	go fibonacci(cap(c), c)
// 	for i := range c {
// 		fmt.Println(i)
// 	}
// }
