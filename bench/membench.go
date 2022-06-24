package bench

import "fmt"

func stringAllocation() string {
	var s string
	for i := 0; i < 100; i++ {
		s = fmt.Sprint("string")
	}
	return s
}
