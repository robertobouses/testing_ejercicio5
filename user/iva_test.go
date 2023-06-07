package user_test

import (
	"testing"

	"github.com/robertobouses/testing_ejercicio5/user"
)

func TestIva(t *testing.T) {
	base := 10.00
	porcentaje := 10.00
	valor := user.Iva(base, porcentaje)
	expected := 100.00

	if valor != expected {
		t.Errorf("Se espera %f, y se obtuvo %f", expected, valor)
	}

}
func TestIva2(t *testing.T) {
	base := 20.00
	porcentaje := 20.00
	valor := user.Iva(base, porcentaje)
	expected := 400.00

	if valor != expected {
		t.Errorf("Se espera %f, y se obtuvo %f", expected, valor)
	}

}
