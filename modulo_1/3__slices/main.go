package main

import "fmt"

// Crie um slice de strings com 3 nomes e adicione um quarto nome usando append()
// e imprima todos os nomes.

// EXTRA: Imprimir na tela o cap() e len() do slice antes de adicionar o quarto nome
// e depois de adiciona-lo, observar a diferença

func main(){

	sl := []string{"Gabriel", "Lucas", "João"}

	fmt.Println(cap(sl))
	
	sl = append(sl, "Malu")

	fmt.Println(cap(sl))

	fmt.Println(sl)

}