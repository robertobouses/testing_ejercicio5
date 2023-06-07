package main

import (
	"fmt"

	"github.com/robertobouses/testing_ejercicio5/user"
)

func main() {
	var base, porcentaje float64

	fmt.Println("Iniciando Aplicación de cálculo de Iva")
	valor := user.Iva(base, porcentaje)
	fmt.Println(valor)

}
