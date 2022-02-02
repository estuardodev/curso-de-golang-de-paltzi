package main

import "fmt"

func main() {
	// Declaracion de variables
	helloMenssage := "Hello"
	worldMenssage := "World"

	// Println
	fmt.Println(helloMenssage, worldMenssage)
	fmt.Println(helloMenssage, worldMenssage)

	// Printf
	nombre := "Platzi"
	cursos := 700
	fmt.Printf("%s tiene mas de %d cursos\n", nombre, cursos)
	fmt.Printf("%v tiene mas de %v cursos\n", nombre, cursos)

	// Sprintf
	message := fmt.Sprintf("%s tiene mas de %d cursos", nombre, cursos)
	fmt.Println(message)

	//Tipos de datos
	fmt.Printf("helloMessage: %T\n", helloMenssage)
	fmt.Printf("cursos: %T\n", cursos)
}
