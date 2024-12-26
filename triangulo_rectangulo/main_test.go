package main

import "testing"

// Calcular hipotenusa
func TestCalcHip(t *testing.T) {
	inicializar(2, 3)
	got := calcHip()
	want := 3.605551275463989
	if got != want {
		t.Errorf("calcHip() = %.*f, want %.*f", Presition, got, Presition, want)
	}
}

// Calcular area
func TestCalcArea(t *testing.T) {
	inicializar(2, 3)
	got := calcArea()
	want := 3.0
	if got != want {
		t.Errorf("calcArea() = %.*f, want %.*f", Presition, got, Presition, want)
	}
}

// Calcular perimetro
func TestCalcPerimeter(t *testing.T) {
	inicializar(2, 3)
	got := calcPerimeter()
	want := 8.60555127546399
	if got != want {
		t.Errorf("calcPerimeter() = %.*f, want %.*f", Presition, got, Presition, want)
	}
}
