package main

import (
	"fmt"
	"runtime"
)

func stringAllocation() string {
	var s string
	var stats runtime.MemStats
	for i := 0; i < 100; i++ {
		s = fmt.Sprint("string")
		runtime.ReadMemStats(&stats)
		fmt.Println("HeapSys:   ", stats.HeapSys)
		fmt.Println("HeapAlloc: ", stats.HeapAlloc)
	}
	return s
}

func main() {
	var stats runtime.MemStats
	//runtime.ReadMemStats(&stats)
	//fmt.Println(stats.HeapSys)
	//for i := 0; i < 370500; i++ {
	//	_ = fmt.Sprint("string")
	//}
	runtime.ReadMemStats(&stats)
	fmt.Println("TotalAlloc:  ", stats.TotalAlloc)
	fmt.Println("Sys:         ", stats.Sys)
	fmt.Println("HeapIdle:    ", stats.HeapIdle)
	fmt.Println("HeapInuse:   ", stats.HeapInuse)
	fmt.Println("HeapReleased:", stats.HeapReleased)
	fmt.Println("HeapSys:     ", stats.HeapSys)
	fmt.Println("HeapAlloc:   ", stats.HeapAlloc)
	//stringAllocation()
}
