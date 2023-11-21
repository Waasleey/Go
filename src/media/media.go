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

	_, _, media := calculamedia(a, b, c)
	fmt.Println(media)

}
func calculamedia(valorA float32, valorB float32, valorC float32) (float32, float32, float32) {

	soma := (valorA + valorB + valorC)
	divisao := soma / 3
	valortotal := divisao

	return soma, divisao, valortotal
}
