package main

import (
	"fmt"
	"sync"
	"time"
)

// func main() {
// 	tasksNum := 10

// 	dataCh := make(chan interface{})
// 	resp := make([]interface{}, 0, tasksNum)

// 	// 1.0版本存在slice写入慢, 导致主goroutine退出,有时数据不全的情况.
// 	// 2.0下面引入了一个stopCh, 让主goroutine等待所有数据写入完成后再退出.
// 	stopCh := make(chan struct{}, 1)

// 	go func() {
// 		for data := range dataCh {
// 			resp = append(resp, data)
// 		}
// 		stopCh <- struct{}{}
// 	}()

// 	var wg sync.WaitGroup
// 	// wg.Add(tasksNum)
// 	for i := 0; i < tasksNum; i++ {
// 		wg.Add(1)
// 		go func(ch chan<- interface{}) {
// 			defer wg.Done()
// 			ch <- time.Now().UnixNano()
// 		}(dataCh)
// 	}
// 	wg.Wait()
// 	close(dataCh)

// 	<-stopCh

// 	fmt.Printf("resp total: %v\n", len(resp))
// 	fmt.Printf("resp: %v", resp)
// }

// 3.0版本
func main() {
	tasksNum := 10
	dataCh := make(chan interface{})
	resp := make([]interface{}, 0, tasksNum)

	// 启动写 goroutine，推进并发获取数据进程，将获取到的数据聚合到 channel 中
	go func() {
		// 保证获取到所有数据后，通过 channel 传递到读协程手中
		var wg sync.WaitGroup
		for i := 0; i < tasksNum; i++ {
			wg.Add(1)
			go func(ch chan<- interface{}) {
				defer wg.Done()
				ch <- time.Now().UnixNano()
			}(dataCh)
		}
		// 确保所有数据都写入 channel 后，关闭 channel
		wg.Wait()
		close(dataCh)
	}()

	for data := range dataCh {
		resp = append(resp, data)
	}
	fmt.Printf("resp total: %v\n", len(resp))
	fmt.Printf("resp: %v\n", resp)
}
