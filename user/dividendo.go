package user

type Division struct {
	Numerador   int
	Denominador int
}

type Frase struct {
	Escritura string
}

var frase1 = Frase{Escritura: "no puede dividir entre cero"}

func (d Division) Dividendo() (Frase, int) {

	if d.Denominador == 0 {
		return frase1, 0
	} else {
		return Frase{}, d.Numerador / d.Denominador
	}

}
