package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

// Runner在给定的超时时间内执行一组任务
// 并且在操作系统发送中断信号时结束这些任务
type Runner struct {
	// interrupt 通道报告从操作系统
	// 发送信号
	interrupt chan os.Signal

	// complate通道报告处理任务已经完成
	complete chan error

	// timeout报告处理任务已经超时
	timeout <-chan time.Time

	// tasks 持有一组以索引顺序依次执行的
	tasks []func(int)
}

// test
