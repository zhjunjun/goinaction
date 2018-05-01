// 这个示例程序展示如何用无缓冲的通道来模拟
// 网球的例子
// 无缓冲的例子
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	// 创建一个无缓冲通道
	court := make(chan int)

	wg.Add(2)

	// 启动两个球员
	go player("Nada1", court)
	go player("Dajo1", court)

	// 发球
	court <- 1

	// 等待游戏结束
	wg.Wait()
}

func player(name string, court chan int) {

	// 通知main函数已经完成工作
	defer wg.Done()

	for {
		// 等待球被打过来
		ball, ok := <-court
		// 如果通道关闭 我们就赢了
		if !ok {
			fmt.Printf("player %s Won", name)
			return
		}
		// 选随机数,然后用这个数来判断我们是否
		n := rand.Intn(100)

		if n%13 == 0 {
			fmt.Printf("palyer %s Missed", name)
			close(court)
			return
		}
		// 显示击球数 并将击球数加1
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++
		court <- ball
	}
}
