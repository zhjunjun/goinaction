// 这个示例程序展示如何使用互斥锁来
// 定义一段需要同步访问的代码临界区
// 资源的同步访问
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int

	wg sync.WaitGroup

	mutex sync.Mutex
)

func main() {
	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg.Wait()
	fmt.Printf("Final Counter: %d \n", counter)
}

func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// 加锁
		mutex.Lock()
		{
			value := counter
			runtime.Gosched()

			value++

			counter = value
		}
		// 释放锁
		mutex.Unlock()
	}
}
