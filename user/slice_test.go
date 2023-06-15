package user_test

import (
	"fmt"
	"testing"

	"github.com/robertobouses/testing_ejercicio5/user"
)

func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestPrintSlice(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}

	n := 3
	expected := []int{1, 2, 3}
	user.PrintSlice(slice, n)
	fmt.Println("_______slice[:n]", slice[:n])
	if !equalSlices(expected, slice[:n]) {
		t.Errorf("El resultado %v es diferente de lo esperado %v", slice[:n], expected)
	}
}
