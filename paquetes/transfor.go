package main

//go-paquetes.git
import (
	"fmt"
	"paquetes/numeros"
	"paquetes/saludo"
)

func main() {
	fmt.Println(numeros.Convertir(16))
	//fmt.Println(numeros.IsPrime(19))
	//fmt.Println(greeting.WelcomeText)
	//fmt.Println(strings.Reverse("adicasa"))
	//fmt.Println(str.Count("Go es Amigable y agradable. Me gusta trabajar en Go", "Go"))

	fmt.Println(saludo.Holamundo() + " de adicasa")
}
