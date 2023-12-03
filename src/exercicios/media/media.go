package main

import "fmt"

func main() {

	var a float32
	var b float32
	var c float32

	fmt.Println("Digite as notas abaixo: ")
	fmt.Print("1º: ")
	fmt.Scan(&a)

	fmt.Print("2º: ")
	fmt.Scan(&b)

	fmt.Print("3º: ")
	fmt.Scan(&c)

	msg, media := calculamedia(a, b, c)
	fmt.Println(msg, media)

}
func calculamedia(valorA float32, valorB float32, valorC float32) (string, float32) {

	return "Sua média é:", (valorA + valorB + valorC) / 3
}
