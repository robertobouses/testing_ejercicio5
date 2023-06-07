package user

type Division struct {
	numerador   int
	denominador int
}

type Frase struct {
	Escritura string
}

var frase1 = Frase{Escritura: "no puede dividir entre cero"}

func (D Division) Dividendo() (Frase, int) {

	if D.denominador == 0 {
		return frase1, 0
	} else {
		return Frase{}, D.numerador / D.denominador
	}

}
