package main

import (
	"fmt"

	"github.com/robertobouses/testing_ejercicio5/user"
)

func main() {
	var base, porcentaje float64
	base = 4
	porcentaje = 2

	fmt.Println("Iniciando Aplicación de cálculo de Iva")
	valor := user.Iva(base, porcentaje)
	fmt.Println(valor)
	numeroTelefono := 123456789
	res1, res2 := user.Telefono(numeroTelefono)
	fmt.Println(res1)
	fmt.Println(res2)

	numeroTelefono2 := 3957
	res3, res4 := user.Telefono(numeroTelefono2)
	fmt.Println(res3)
	fmt.Println(res4)

	slice := []int{1, 2, 3, 4, 5}

	user.PrintSlice(slice, 2)
	user.PrintSlice(slice, 89)

	user.PrintSlice2(slice, 8)
	user.PrintSlice2(slice, 2)

}
