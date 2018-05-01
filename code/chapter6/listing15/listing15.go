// This sample program demonstrates how to use the atomic
// package functions Store and Load to provide safe access
// to numeric types.
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	// shutdown is a flag to alert running goroutines to shutdown.
	shutdown int64

	// wg is used to wait for the program to finish.
	wg sync.WaitGroup
)

// main is the entry point for all Go programs.
func main() {
	// Add a count of two, one for each goroutine.
	wg.Add(2)

	// Create two goroutines.
	go doWork("A")
	go doWork("B")

	// Give the goroutines time to run.
	time.Sleep(1 * time.Second)

	// Safely flag it is time to shutdown.
	fmt.Println("Shutdown Now")
	atomic.StoreInt64(&shutdown, 1)

	// Wait for the goroutines to finish.
	wg.Wait()
}

// 在这个例子中,启动了两个 goroutine,并完成一些工作。
// 在各自循环的每次迭代之后,在第 52 行中 goroutine 会使用 LoadInt64 来检查 shutdown 变量的值。
// 这个函数会安全地返回 shutdown 变量的一个副本。如果这个副本的值为 1,goroutine 就会跳出循环并终止。
// 在第 35 行中,main 函数使用 StoreInt64 函数来安全地修改 shutdown 变量的值。
// 如果 哪个 doWork goroutine 试图在 main 函数调用 StoreInt64 的同时调用 LoadInt64 函数,那 么原子函数会将这些调用互相同步,保证这些操作都是安全的,不会进入竞争状态。

// doWork simulates a goroutine performing work and
// checking the Shutdown flag to terminate early.
func doWork(name string) {
	// Schedule the call to Done to tell main we are done.
	defer wg.Done()

	for {
		fmt.Printf("Doing %s Work\n", name)
		time.Sleep(250 * time.Millisecond)

		// Do we need to shutdown.
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutting %s Down\n", name)
			break
		}
	}
}
