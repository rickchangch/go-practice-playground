package gotour

import (
	"fmt"
	"math"
	"runtime"
	"strings"
	"time"

	"golang.org/x/tour/pic"
	"golang.org/x/tour/wc"
)

type gotour struct{}

var Handler = new(gotour)

func (g *gotour) Run() {

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

	// Arrays & Slices
	arraySliceSyntax()
	makeSyntax()
	ticTacToe()
	appendSyntax()

	rangeSyntax()

	pic.Show(Pic)

	// Maps
	mapSyntax()
	wc.Test(WordCount)

	// Functions
	functionAsParam()
	functionClosures()

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

	// Interface
	interfaceSyntax()
	typeAssertions()
	typeSwitch()
}

// Constants can't be declared using the := syntax
const Pi = 3.14
const (
	Big   = 1 << 100
	Small = Big >> 99
)

// The var statement declares a list of variables, and also includes initializers.
var c, python, java bool = true, false, false

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

/**
 * @see https://blog.golang.org/slices-intro
 */
func arraySliceSyntax() {
	fmt.Println("Arrays & Slices:")

	// An array's length is part of its type, so arrays cannot be resized
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a, a[1])

	// The below expression is equaled to [...]int{2,3,5,7,11,13}
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

func appendSyntax() {
	// built-in function: func append(s []T, vs ...T) []T

	// If the backing array of s is too small to fit all the given values, a bigger array will be allocated.
	// The returned slice will point to the newly allocated array.

	var s []int
	printSlice(s)

	// `append` works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)
}

func ticTacToe() {
	// Slices can contain any type, including other slices.

	// Create a tic-tac-toe board.
	board := [][]string{
		// []string{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
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

func rangeSyntax() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	// When ranging over a slice, two values are returned for each iteration.
	// The first is the index, and the second is a copy of the element at that index.
	for idx, val := range pow {
		fmt.Printf("2**%d = %d\n", idx, val)
	}

	// You can skip the index or value by assigning to _.
	for _, value := range pow {
		fmt.Println(value)
	}
}

func Pic(dx, dy int) (res [][]uint8) {
	res = make([][]uint8, dy)

	for y := range res {
		res[y] = make([]uint8, dx)
		for x := range res[y] {
			res[y][x] = uint8((x + y) / 2)
		}
	}

	return res
}

type Vertex2 struct {
	Lat, Long float64
}

func mapSyntax() {

	// A map maps keys to values.
	var m map[string]Vertex2 = make(map[string]Vertex2)

	m["Bell Labs"] = Vertex2{
		40.68433, -74.39967,
	}

	fmt.Println(m["Bell Labs"])

	// The usage of map literals are like struct literals.
	m = map[string]Vertex2{
		"Bell Labs": {40.68433, -74.39967},
		// You can also ommit the type of value, cuz it will follow the top-level type.
		"Google": {37.42202, -122.08408},
	}
	fmt.Println(m)

	m2 := make(map[string]int)

	m2["Answer"] = 42
	fmt.Println("The value:", m2["Answer"], m2)

	delete(m, "Answer")
	fmt.Println("The value:", m2["Answer"], m2)

	// The ok presents that if "Answer" is in the map.
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

func WordCount(s string) map[string]int {
	res := make(map[string]int)

	for _, word := range strings.Fields(s) {
		res[word] += 1
	}

	return res
}

func functionAsParam() {
	// Functions are values too. They can be passed around just like other values.

	compute := func(function func(float64, float64) float64) float64 {
		return function(3, 4)
	}

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(compute(hypot))
}

// adder() is a function that returns a func(int) that returns int
func adder() func(int) int {
	sum := 0
	// A closure is a function value that references variables from outside its body.
	// Each closure is bound to its own `sum` variable.
	return func(x int) int {
		sum += x
		return sum
	}
}
func functionClosures() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func fibonacci() func() int {
	previous, next := 0, 1
	return func() int {
		now := previous
		previous, next = next, next+previous
		return now
	}
}

/**
 * Methods
 */
// Go does not have classes. Instead, a `method` is a function with a recevier argument.
// The below statement presents Struct v has the member function Abs, e.g. v.Abs()
func (v Vertex2) Abs() float64 {
	return math.Sqrt(v.Lat*v.Lat + v.Long*v.Long)
}

// You can declare methods with pointer receivers. The method will modify the value to which the recevier points.
func (v *Vertex2) Scale(f float64) {
	v.Lat = v.Lat * f
	v.Long = v.Long * f
}

type MyFloat float64

// You can declare a method on non-struct types, too.
// But only allow to declare a method with a receiver whose type is defined in the same package.
// If change the recevier type `Myfloat` to `float64`, Go will build failed.
func (f MyFloat) Abs() float64 {
	return float64(f)
}

/**
 * Interfaces
 */
func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

type Abser interface {
	Abs() float64
}

type MyFloat3 float64

func (f MyFloat3) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex3 struct {
	X, Y float64
}

func (v *Vertex3) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func interfaceSyntax() {
	// The below statement can be implicitly simplfied as `var f Abser = MyFloat3(-math.Sqrt2)`
	var a Abser
	f := MyFloat3(-math.Sqrt2)
	v := Vertex3{3, 4}

	// Compile OK.
	a = f
	a = &v
	// Compile failed. Below v is a Vertex3, not *Vertex3
	// a = v

	fmt.Println(a.Abs())

	// The interface type that specifies zero methods is known as the empty interface:
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func typeAssertions() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	// If i doesn't hold (T), a type assertion can return two values. The ok value will report whether the assertion succeeded.
	f, ok := i.(float64)
	fmt.Println(f, ok)
	// If you don't catch assertion status, and i also doesn't hold (T), a panic will be throwed.
	// f = i.(float64)
}

func typeSwitch() {
	// When type assertion specifies type as switch case target, we can use it to build a function which permits any type of parameters.
	do := func(i interface{}) {
		switch v := i.(type) {
		case int:
			fmt.Printf("Twice %v is %v\n", v, v*2)
		case string:
			fmt.Printf("%q is %v bytes long\n", v, len(v))
		default:
			fmt.Printf("I don't know about type %T!\n", v)
		}
	}

	do(21)
	do("hello")
	do(true)
}
