// 这个示例程序展示如何用无缓冲的通道来模拟 -> 接力比赛
// 4个goroutine间的接力比赛
// 无缓冲的例子
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	// 创建一个无缓冲通道
	baton := make(chan int)

	// 为最后一位跑步者计数加一
	wg.Add(1)

	// 第一位跑步者 持有接力棒
	go Runner(baton)

	// 比赛开始
	baton <- 1

	// 等待比赛结束
	wg.Wait()
}

func Runner(baton chan int) {

	var newRunner int
	// 等待接力棒
	runner := <-baton

	// 开始跑步
	fmt.Printf("Runner %d Running With Baton \n", runner)

	// 创建下一个跑步者
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d To The Line\n", runner)
		go Runner(baton)
	}

	// 围绕跑道跑
	time.Sleep(100 * time.Millisecond)

	// 比赛结束了么
	if runner == 4 {
		fmt.Printf("Runner %d Finished,Race Over\n", runner)
		wg.Done()
		return
	}

	// 将接力棒交给下一位跑步者
	fmt.Printf("Runner %d Exchange With Runner %d\n", runner, newRunner)

	baton <- newRunner
}
