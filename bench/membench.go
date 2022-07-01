package main

import (
	"fmt"
	"runtime"
	"time"
)

//func stringAllocation() string {
//	var s string
//	var stats runtime.MemStats
//	for i := 0; i < 100; i++ {
//		s = fmt.Sprint("string")
//		runtime.ReadMemStats(&stats)
//		fmt.Println("HeapSys:   ", stats.HeapSys)
//		fmt.Println("HeapAlloc: ", stats.HeapAlloc)
//	}
//	return s
//}
//
//func main() {
//	var stats runtime.MemStats
//	//runtime.ReadMemStats(&stats)
//	//fmt.Println(stats.HeapSys)
//	//for i := 0; i < 370500; i++ {
//	//	_ = fmt.Sprint("string")
//	//}
//	runtime.ReadMemStats(&stats)
//	fmt.Println("TotalAlloc:  ", stats.TotalAlloc)
//	fmt.Println("Sys:         ", stats.Sys)
//	fmt.Println("HeapIdle:    ", stats.HeapIdle)
//	fmt.Println("HeapInuse:   ", stats.HeapInuse)
//	fmt.Println("HeapReleased:", stats.HeapReleased)
//	fmt.Println("HeapSys:     ", stats.HeapSys)
//	fmt.Println("HeapAlloc:   ", stats.HeapAlloc)
//	//stringAllocation()
//}

// go:noinline
func allocate(size int) []byte {
	return make([]byte, size)
}

func main() {
	const dur = 11 * 60 * 1e9
	const base int = 1 << 10
	start := time.Now().UnixNano()
	var buf []byte
	for {
		now := time.Now().UnixNano()
		if now-start > dur {
			break
		}
		buf = append(buf, allocate(base)...)
		time.Sleep(10 * time.Millisecond)
	}
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	println(ms.HeapSys)
	fmt.Println("end")
}
