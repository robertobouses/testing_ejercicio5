package user

import "fmt"

func PrintSlice(slice []int, n int) {
	if n > len(slice) {
		n = len(slice)
	}
	fmt.Println(slice[:n])

}
