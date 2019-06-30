package main

import("fmt")

func calc(taskChan, resChan chan int, exitChan chan bool) {
    defer func() {
        err := recover()
        if err != nil {
            fmt.Println("error...")
            return 
        }
    }()
    
    for v := range taskChan {
        // 任务处理逻辑
        flag := true
        for i := 2; i < v; i++ {
            if v%i == 0 {
                flag = false
                break
            }
        }
        if flag {
            //结果进chan
            resChan <- v
        }
    }
    //处理完进退出chan
    exitChan <- true
}

func main() {
    //任务chan
    intChan := make(chan int, 1000)
    //结果chan
    resChan := make(chan int, 1000)
    //退出chan
    exitChan := make(chan bool, 8)

    go func() {
        for i := 0; i < 1000; i++ {
            intChan <- i
        }
        close(intChan)
    }()

    //启动8个goroutine做任务
    for i := 0; i < 8; i++ {
        go calc(intChan, resChan, exitChan)
    }

    go func() {
        //等所有goroutine结束
        for i := 0; i < 8; i++ {
            <-exitChan
        }
        close(resChan)
        close(exitChan)
    }()

    for v := range resChan {
        fmt.Println(v)
    }
}
