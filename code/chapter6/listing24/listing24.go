package main

// 这个实例程序展示如何使用
// 有缓冲的通道和固定数目的gorutgoroutine来处理一堆工作

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4
	taskLoad         = 10
)

// wg用来等待程序完成
var wg sync.WaitGroup

// init 初始化包, Go语言运行时会在其他代码执行之前
func init() {
	// 初始化随机数种子
	rand.Seed(time.Now().Unix())
}

// main是所有Go程序的入口
func main() {
	// 创建一个有缓冲的通道来管理工作
	tasks := make(chan string, taskLoad)

	// 启动goroutine来处理工作
	wg.Add(numberGoroutines)
	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}

	// 增加一组要完成的工作
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task : %d", post)
	}

	// 当所有工作都处理完时关闭通道
	// 以便所有的goroutine退出
	close(tasks)

	// 等待所有工作完成
	wg.Wait()
}

// workder 作为goroutine启动来处理
// 从有缓冲的通道传的工作
func worker(tasks chan string, worker int) {
	// 通知函数已经返回
	defer wg.Done()

	for {
		// 等待分配工作 会阻塞
		task, ok := <-tasks
		if !ok {
			fmt.Printf("Worker :%d : Shutting Down\n", worker)
			return
		}

		// 显示我们开始工作了
		fmt.Printf("Worker: %d : Started %s\n", worker, task)

		// 随机等待一段时间来模拟工作
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Microsecond)

		// 显示我们完成了工作
		fmt.Printf("Worker: %d : Complated %s\n", worker, task)
	}
}
