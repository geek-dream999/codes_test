package ctx

import (
	"context"
	"fmt"
	"time"
)

/**   0.超时信号  context.WithTimeout(context.Background(), 1*time.Second)

 */

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}

/** 2.取消信号  context.WithCancel
一旦我们执行返回的取消函数，当前上下文以及它的子上下文都会被取消，所有的 Goroutine 都会同步收到这一取消信号。
*/

func cancel() {
	// 创建根上下文
	ctx := context.Background()

	// 创建带有取消能力的上下文
	ctx, cancel := context.WithCancel(ctx)

	// 启动一个 goroutine，在一段时间后取消上下文
	go func() {
		time.Sleep(3 * time.Second)
		cancel() // 取消上下文
	}()

	// 执行某些操作，检查上下文是否已被取消
	for {
		select {
		case <-ctx.Done():
			fmt.Println("上下文已取消")
			return
		default:
			fmt.Println("继续执行操作")
			time.Sleep(1 * time.Second)
		}
	}
}

/**  1.默认上下文  context.Background、context.TODO
这两个私有变量都是通过 new(emptyCtx) 语句初始化的，它们是指向私有结构体 context.emptyCtx 的指针，这是最简单、最常用的上下文类型：
源码如下，可以发现通过空方法实现了context的所有方法
*/

type emptyCtx int

func (*emptyCtx) Deadline() (deadline time.Time, ok bool) {
	return
}

func (*emptyCtx) Done() <-chan struct{} {
	return nil
}

func (*emptyCtx) Err() error {
	return nil
}

func (*emptyCtx) Value(key interface{}) interface{} {
	return nil
}
