package goruntine

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/tour/tree"
)

type goruntine struct{}

var Handler = new(goruntine)

func (g goruntine) Run() {
	// Goroutines run in the same address space, so access to shared memory must be synchronized.
	go say("world")
	say("hello")

	// Test how the Senders and Receviers to cowork in the synchronize channel.
	syncReadWrite()

	// Throws `fatal error: all goroutines are asleep - deadlock!` when bufferSize < 2
	bufferedChannel(2)

	// Test the result about getting value from channel by for loop `range` but don't provide the channel close method in the sender.
	channelClose(5)

	// Observe the interaction between multipal communication channel in one goruntine.
	selectCase()

	// Use goruntine to operate the node of binary trees.
	compareBinaryTrees()

	// Make a scenario that multiple goruntines modify the same map concurrently.
	mutualExclusion()
}

func say(s string) {
	for i := 0; i < 2; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s, i)
	}
}

func syncReadWrite() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x := <-c
	fmt.Println("after activate reader")

	y := <-c
	fmt.Println("after another activate reader")

	fmt.Println(x, y, x+y)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}

	fmt.Println("activate sender")
	c <- sum
}

func bufferedChannel(bufferSize int) {
	// Sends to a buffered channel block only when the buffer is full.
	// Receives block when the buffer is empty.

	ch := make(chan int, bufferSize)
	ch <- 1
	ch <- 2
}

func fibonacci(n int, c chan int) {
	prev, next := 0, 1
	for i := 0; i < n; i++ {
		c <- prev
		prev, next = next, prev+next
	}

	// Note: Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.
	// Note2: Channels aren't like files; you don't usually need to close them. Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a `range` loop.
	close(c)
}

func channelClose(bufferSize int) {
	c := make(chan int, bufferSize)

	go fibonacci(cap(c), c)

	for i := range c {
		fmt.Println(i)
	}
}

func fib2(ch, quit chan int) {
	prev, next := 0, 1
	for {
		select {
		// ch sender
		case ch <- prev:
			prev, next = next, prev+next
		// quit recevier
		case <-quit:
			fmt.Println("trigger quit")
			// exit for-loop
			return
		}
	}
}

// The select statement lets a goroutine wait on multiple communication operations.
func selectCase() {
	fmt.Println("-selectCase-")

	ch := make(chan int)
	quit := make(chan int)

	go func() {
		// ch recevier
		for i := 0; i < 5; i++ {
			fmt.Println(<-ch)
		}

		// quit sender
		quit <- 0
	}()

	fib2(ch, quit)
}

func walkRecursive(t *tree.Tree, ch chan int) {
	if t != nil {
		walkRecursive(t.Left, ch)
		ch <- t.Value
		walkRecursive(t.Right, ch)
	}
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func walk(t *tree.Tree, ch chan int) {
	walkRecursive(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go walk(t1, ch1)
	go walk(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if ok1 != ok2 || v1 != v2 {
			return false
		}

		if !ok1 {
			break
		}
	}

	return true
}

func compareBinaryTrees() {
	fmt.Println("-compare two binary trees-")

	fmt.Println(same(tree.New(1), tree.New(1)))
	fmt.Println(same(tree.New(1), tree.New(2)))
	fmt.Println(same(tree.New(2), tree.New(2)))
}

type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

func (c *SafeCounter) inc(key string) {
	c.mu.Lock()
	c.v[key]++
	c.mu.Unlock()
}

func (c *SafeCounter) value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[key]
}

func mutualExclusion() {

	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.value("somekey"))
}
