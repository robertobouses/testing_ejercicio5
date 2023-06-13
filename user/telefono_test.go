package user_test

import (
	"testing"

	"github.com/robertobouses/testing_ejercicio5/user"
)

func TestTelefono(t *testing.T) {
	numeroValido := 123456789
	resultado, _ := user.Telefono(numeroValido)
	if resultado != "el número existe y es correcto" {
		t.Errorf("Se esperaba 'el número existe y es correcto', se obtuvo: %s", resultado)
	}

	numeroInvalido := 12345
	_, resultado2 := user.Telefono(numeroInvalido)
	if resultado2.Escritura != "el número es incorrecto" {
		t.Errorf("Se esperaba 'el número es incorrecto', se obtuvo: %s", resultado2.Escritura)
	}
}
