package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	// counter是所有goroutine都要增加其值的变量
	counter int64
	// wg用来等待程序结束
	wg sync.WaitGroup
)

// main是所有Go程序的入口
func main() {
	wg.Add(2)

	// 创建两个goroutine
	go incCounter(1)
	go incCounter(2)

	wg.Wait()

	fmt.Println("Final Counter:", counter)
}

// 现在,程序的第 43 行使用了 atmoic 包的 AddInt64 函数。这个函数会同步整型值的加法,
// 方法是强制同一时刻只能有一个 goroutine 运行并完成这个加法操作。
// 当 goroutine 试图去调用任 何原子函数时,这些 goroutine 都会自动根据所引用的变量做同步处理。现在我们得到了正确的 值 4。

// incCounter increments the package level counter variable.
func incCounter(id int) {
	// Schedule the call to Done to tell main we are done.
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// Safely Add One To Counter.
		atomic.AddInt64(&counter, 1)

		// Yield the thread and be placed back in queue.
		// 当前goroutine从线程退出,并放回到队列
		runtime.Gosched()
	}
}
