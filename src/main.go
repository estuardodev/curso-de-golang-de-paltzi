package main

import "fmt"

func helloFunction(message string) {

	fmt.Println(message)

}

func tripeArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

// RETO DE LA CLASE
func areaRectangulo(base, altura int) int {
	var resultado int = (base * altura) / 2

	return resultado
}

func main() {
	//Funciones
	helloFunction("Hola Mundo")
	tripeArgument(1, 2, "Hola")

	value := returnValue(2)
	fmt.Println("Value: ", value)

	value1, _ := doubleReturn(2)
	fmt.Println("value1: ", value1)

	// RETO DE LA CLASE
	respuesta := areaRectangulo(10, 10)
	fmt.Println("El area del rectangulo es: ", respuesta)

}
