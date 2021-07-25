package main

import (
	"fmt"
	"math"
	"runtime"
	"strings"
	"time"

	"example.com/go-demo-1/hello"
)

// Constants can't be declared using the := syntax
const Pi = 3.14
const (
	Big   = 1 << 100
	Small = Big >> 99
)

// The var statement declares a list of variables, and also includes initializers.
var c, python, java bool = true, false, false

// The above variable declarations also equal to the below which are factored into brackets.
// var (
// 	c      bool = true
// 	python bool = false
// 	java   bool = false
// )

type Vertex struct {
	x, y int
}

// When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
func add(x, y int) int {
	return x + y
}

// A function can return any number of results.
func swap(x, y string) (string, string) {
	return y, x
}

// Go's return values may be named. If so, they are treated as variables defined at the top of the function.
// This is known as a "naked" return. Naked return statements should be used only in short functions.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func forSyntax(n int) int {
	var sum int = 0
	// The init and post statements can be omitted. e.g. for ; sum <= n; {}
	// If drop the semicolons and only keep statement and condition expression stay, it present as 'while' syntax. e.g. for sum < 100 {}
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func ifSyntax(n int) bool {
	if n > 0 {
		return true
	} else if n2 := 10; n > n2 {
		return true
	} else {
		return false
	}
}

func switchSyntax() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("darwin")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Println("other: ", os)
	}

	// Switch with no condition can be a clean way to write long if-then-else chains.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

// Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.
func deferSytax() {

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

func Sqrt(x float64) float64 {
	z := float64(1)

	for math.Abs(z*z-x) > 1e-10 {
		z -= (z*z - x) / (2 * z)
	}

	return z
}

func pointerSyntax() {
	var p *int
	var i int = 10
	p = &i
	fmt.Println("pointers:", p, *p)
}

func printSlice(s []int) {
	// A slice has both a length and a capacity.
	// The length of a slice is the number of elements it contains.
	// The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func arraySliceSyntax() {
	fmt.Println("Arrays & Slices:")

	// An array's length is part of its type, so arrays cannot be resized
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a, a[1])

	primes := [6]int{2, 3, 5, 7, 11, 13}
	// A slice is formed by specifying two indices, a low and high bound, separated by a colon,
	// and the selects wii includes the low index element, but excludes the high one.
	var sliceArr []int = primes[1:3]
	fmt.Println(primes, sliceArr)

	// Remember, Slices are just like the reference to Arrays.
	sliceArr[0] = 30
	// The array primes now contains {2, 30, 5, 7, 11, 13}

	// A slice literal is like an array literal without the length.
	arrayLiteral := [3]bool{true, true, false}
	sliceLiteral := []bool{true, true, false}
	fmt.Println(arrayLiteral, sliceLiteral)

	s := []int{2, 3, 5, 7, 11, 13}
	// Slice the slice to give it zero length. (Capacity 6)
	s = s[:0]
	printSlice(s)
	// Extend its length. (Capacity 6)
	s = s[:4]
	printSlice(s)
	// Drop its first two values. (Capacity 6 => 4)
	s = s[2:]
	printSlice(s)
}

func makeSyntax() {
	// Slices can be created with the built-in `make` function; this is how you create dynamically-sized arrays.
	// make(type, length, capacity)
	a := make([]int, 5)
	printSlice(a)
	b := make([]int, 0, 5)
	printSlice(b)
}

func ticTacToe() {
	// Slices can contain any type, including other slices.

	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func main() {
	fmt.Println(hello.BaseHello())

	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

	fmt.Println(add(42, 13))

	// The := syntax can be used to replace a `var` declaration with implicit type, but only be allowed to use inside the function.
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))

	// If an initializer is present, the type can be ommited.
	var i = 1
	var complexValue complex64 = 1.2 + 12i
	fmt.Println(i, c, python, java, complexValue)

	// The len() will count number by the bytes, and range syntax count number by the rune.
	aa := "Hi,世界"
	n := 0
	for range aa {
		n++
	}
	fmt.Println(len(aa))
	fmt.Println(n)

	fmt.Println(forSyntax(10))
	fmt.Println(ifSyntax(10))

	fmt.Println("牛頓法: ", Sqrt(10))
	fmt.Println("Math庫: ", math.Sqrt(10))

	switchSyntax()
	deferSytax()

	pointerSyntax()

	v := Vertex{1, 2}
	p := &v
	// The explicit statement is `(*p).x`, but the language permits us instead to write just `p.x`.
	p.x = 1e9
	fmt.Println("Init Struct: ", v, v.x)

	arraySliceSyntax()
	ticTacToe()
}
