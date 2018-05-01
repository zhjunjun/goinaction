package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 因此,调用 GOMAXPROCS 函数就为每个可用的物理处理器创建一个逻辑处理 器。
// 需要强调的是,使用多个逻辑处理器并不意味着性能更好。在修改任何语言运行时配置参数 的时候,都需要配合基准测试来评估程序的运行效果。
// 如果给调度器分配多个逻辑处理器,我们会看到之前的示例程序的输出行为会有些不同。
// 让 我们把逻辑处理器的数量改为 2,并再次运行第一个打印英文字母表的示例程序

func main() {

	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	// Declare an anonymous function and create a goroutine.
	go func() {
		// Schedule the call to Done to tell main we are done.
		defer wg.Done()

		// Display the alphabet three times.
		for count := 0; count < 30; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	go func() {
		// Schedule the call to Done to tell main we are done.
		defer wg.Done()

		// Display the alphabet three times.
		for count := 0; count < 30; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	// Wait for the goroutines to finish.
	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}
