package main

import (
	"fmt"
	"math"
)

func main() {
	//Operadores aritmeticos
	x := 10
	y := 50

	// SUMA
	result := x + y
	fmt.Println("SUMA ES: ", result)

	// RESTA
	result = y - x
	fmt.Println("RESTA ES: ", result)

	// MULTIPLICACIÓN
	result = x * y
	fmt.Println("MULTIPLICACIÓN ES: ", result)

	// DIVISIÓN
	result = y / x
	fmt.Println("DIVISIÓN ES: ", result)

	// MODULO o RESIDUO
	result = y % x
	fmt.Println("MODULO ES: ", result)

	// INCREMENTAR
	x++
	fmt.Println("INCREMENTAR ES: ", x)

	// DECREMENTAR
	x--
	fmt.Println("DECREMENTAR ES: ", x)

	// RETO DE LA CLASE
	// Encontrar el area de un rectangulo, trapecio y de un circulo

	// Rectangulo
	baseRec := 20
	altRec := 30
	result = baseRec * altRec / 2
	fmt.Println("El area de un rectangulo es: ", result)

	// Trapecio
	baseMenor := 5
	baseMayor := 10
	altTra := 6
	result = (baseMayor + baseMenor) * altTra / 2
	fmt.Println("El area de un Trapecio es de: ", result)

	// Circulo
	// Tuve que buscar ayuda para resolver este problema
	radio := 10
	results := math.Pi * math.Pow(float64(radio), 2)

	fmt.Println("El area del circulo es: ", results)
}
