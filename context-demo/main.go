package main

import (
	"context"
	"fmt"
	"reflect"
	"time"
)

const shortDuration = 1 * time.Millisecond

func main() {
	tooslow := fmt.Errorf("too slow")
	fmt.Println(reflect.TypeOf(shortDuration))
	// ctx, cancel := context.WithTimeout(context.Background(), shortDuration)
	ctx, cancel := context.WithDeadlineCause(context.Background(), shortDuration, tooslow)
	time.Sleep(2 * time.Second)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
