//context中的接口

// type Context interface {
//     // 当 context 被取消或者到了 deadline，返回一个被关闭的 channel
       // 这个channel就是被用来监测这个context对应的携程是否结束的
//     Done() <-chan struct{}

//     // 在 channel Done 关闭后，返回 context 取消原因
//     Err() error

//     // 返回 context 是否会被取消以及自动取消时间（即 deadline）
//     Deadline() (deadline time.Time, ok bool)

//     // 获取 key 对应的 value
//     Value(key interface{}) interface{}
// }

// type canceler interface {
//     cancel(removeFromParent bool, err error)
//     Done() <-chan struct{}
// }

// type cancelCtx struct {
//     Context

//     // 保护之后的字段
//     mu       sync.Mutex
//     done     chan struct{}
//     children map[canceler]struct{}
//     err      error
// }

// 可以这么理解 context就是对一个channel做了封装 还加入了一些传参的功能
// 可以在携程中一直传递下去
// 当你删除一个context时，可以在外侧通过他封装的这个channel检测到
// 然后得到这个携程结束了的消息

//基础的context实例

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go handle(ctx, 500*time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}