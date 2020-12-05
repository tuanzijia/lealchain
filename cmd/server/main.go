package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTicker(3 * time.Second)
	i := 1
	for {
		select {
		case <-t.C:
			fmt.Println("区块高度:", i, "有效交易:", 0, "无效交易:", 0)
			fmt.Println("区块高度:", i, "确认交易:", 0)
			i++
		}
	}
	// for i := 1; ; i++ {
	//	<-time.After(3 * time.Second)
	//	fmt.Println("区块高度:", i, "有效交易:", 0, "无效交易:", 0)
	//	fmt.Println("区块高度:", i, "确认交易:", 0)
	// }
}
