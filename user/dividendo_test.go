package user_test

import (
	"testing"

	"github.com/robertobouses/testing_ejercicio5/user"
)

func TestDividendo(t *testing.T) {
	numerador := 10
	denominador := 10
	expected := 1
	_, valor := user.Division{Numerador: numerador, Denominador: denominador}.Dividendo()
	if valor != expected {
		t.Errorf("Se espera %d, y se obtuvo %d", expected, valor)
	}
}

func TestDividendo2(t *testing.T) {
	numerador := 10
	denominador := 0
	expected := user.Frase{Escritura: "no puede dividir entre cero"}
	frase, _ := user.Division{Numerador: numerador, Denominador: denominador}.Dividendo()
	if frase != expected {
		t.Errorf("Se espera '%s', y se obtuvo '%s'", expected.Escritura, frase.Escritura)
	}
}
