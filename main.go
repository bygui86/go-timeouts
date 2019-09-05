package main

import (
	"context"
	"fmt"
	"time"

	"github.com/bygui86/go-timeouts/netcall"
)

func main() {

	noTimeout()

	withTimeout(true)

	withTimeout(false)
}

func noTimeout() {

	res, err := netcall.GetHttpResponse(context.Background(), false)
	if err != nil {
		fmt.Printf("no-timeout - err %v\n", err)
	} else {
		fmt.Printf("no-timeout - res %v\n", res)
	}
}

func withTimeout(acceptableTimeout bool) {

	var ctx context.Context
	var cancelFunc context.CancelFunc
	if acceptableTimeout {
		ctx, cancelFunc = context.WithTimeout(context.Background(), 1*time.Second)
	} else {
		ctx, cancelFunc = context.WithTimeout(context.Background(), 1*time.Millisecond)
	}
	defer cancelFunc()

	res, err := netcall.GetHttpResponse(ctx, true)
	if err != nil {
		fmt.Printf("with-timeout - err %v\n", err)
	} else {
		fmt.Printf("with-timeout - res %v\n", res)
	}
}
