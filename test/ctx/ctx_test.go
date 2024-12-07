package ctx

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// 测试Context 同步信号
func TestCtx(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go handle(ctx, 500*time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
}
