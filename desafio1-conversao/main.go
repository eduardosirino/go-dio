package main

import "fmt"

func converteKelvinParaCelsius(temperaturaKelvin float64) float64 {
	return temperaturaKelvin - 273.15
}

func main() {
	temperaturaEbulicaoAguaKelvin := 373.15
	temperaturaEbulicaoAguaCelsius := converteKelvinParaCelsius(temperaturaEbulicaoAguaKelvin)
	fmt.Printf("Temperatura da ebulição da água em Kelvin é: %.2f\n", temperaturaEbulicaoAguaKelvin)
	fmt.Printf("Temperatura da ebulição da água em Celsius é: %.2f\n", temperaturaEbulicaoAguaCelsius)
}
