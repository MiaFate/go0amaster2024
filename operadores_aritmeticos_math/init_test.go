package operadores_aritmeticos_math

import (
	"fmt"
	"math"
	"testing"
)

func TestResto(t *testing.T) {
	got := Resto()
	want := 1
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestGetPi(t *testing.T) {
	got := GetPi()
	want := math.Pi
	if got != want {
		t.Errorf("got %f; want %f", got, want)
	}
}

func TestGetPower(t *testing.T) {
	got := GetPower(2, 4)
	want := float64(16)
	if got != want {
		t.Errorf("got %f; want %f", got, want)
	}
}

func ExampleResto() {
	fmt.Println(Resto())
	// Output: 1
}

func ExampleGetPower() {
	fmt.Println(GetPower(2, 4))
	// Output: 16
}
