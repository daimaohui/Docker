package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("开始计算...")
	start := time.Now()
	n := 1024 * 1024 * 1024 * 4 * 100
	for i := 0; i < n; i++ {
		_ = i + 1
	}
	fmt.Println("计算已完成，耗时：", time.Since(start).Seconds(), "秒")
}
