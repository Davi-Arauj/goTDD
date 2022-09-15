package estruturasmetodosinterfaces

import "math"

type Forma interface {
	Area() float64
}

type Retangulo struct {
	Largura float64
	Altura  float64
}

type Circulo struct {
	Raio float64
}

type Triangulo struct{
	Base float64
	Altura float64
}

func (t Triangulo) Area() float64{
	return t.Altura * (t.Base/2)
}

func (c Circulo) Area() float64 {
	return math.Pi * c.Raio * c.Raio
}

func (r Retangulo) Area() float64 {
	return r.Largura * r.Altura
}

func Perimetro(retangulo Retangulo) (perimetro float64) {
	perimetro = 2 * (retangulo.Altura + retangulo.Largura)
	return
}
