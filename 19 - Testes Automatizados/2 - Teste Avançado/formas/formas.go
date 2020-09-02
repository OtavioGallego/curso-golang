package formas

import (
	"math"
)

// Forma representa uma forma geométrica
type Forma interface {
	area() float64
}

// Retangulo é uma estrutura que contém altura e largura
type Retangulo struct {
	Altura  float64
	Largura float64
}

// Area calcula a área do retângulo
func (r Retangulo) Area() float64 {
	return r.Altura * r.Largura
}

// Circulo representa uma estrutura que contém raio
type Circulo struct {
	raio float64
}

// Area calcula a área do círculo
func (c Circulo) Area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}
