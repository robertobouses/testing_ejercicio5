package user

import "fmt"

var frase2 = Frase{Escritura: "el número es incorrecto"}

func Telefono(numero int) (string, Frase) {
	cifras := len(fmt.Sprint(numero))
	if cifras == 9 {
		return "el número existe y es correcto", Frase{}
	} else {
		return "", frase2
	}

}
