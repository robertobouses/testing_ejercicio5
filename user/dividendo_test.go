package user_test

import (
	"testing"

	"github.com/robertobouses/testing_ejercicio5/user"
)

func TestDividendo(t *testing.T) {
	numerador := 10
	denominador := 10
	expected := 1
	valor := user.Dividendo(numerador, denominador)
	if valor != expected {
		t.Errorf("Se espera %f, y se obtuvo %f", expected, valor)
	}

}

func (D Division) Dividendo() (Frase, int) {

	if D.denominador == 0 {
		return frase1, 0
	} else {
		return Frase{}, D.nominador / D.denominador
	}

}
