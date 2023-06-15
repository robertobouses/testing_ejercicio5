package user

import "fmt"

func PrintSlice2(slice []int, n int) {
	if n > len(slice) {
		for i := len(slice); i < n; i++ {
			slice = append(slice, i+11)
		}
	}
	fmt.Println(slice[:n])
}
