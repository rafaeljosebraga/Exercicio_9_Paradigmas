package main

import "fmt"

type listFunction func([]int) []int

func executar_estrategia(lista *[]int, estrategia listFunction) {
	*lista = estrategia(*lista)
}

func removeDuplicates(lista []int) []int {
	chavesVistas := make(map[int]bool)
	resultado := []int{}

	for _, item := range lista {
		if _, visto := chavesVistas[item]; !visto {
			chavesVistas[item] = true
			resultado = append(resultado, item)
		}
	}
	return resultado
}

func main() {
	primes := []int{11, 2, 2, 2, 5, 3, 3, 11, 5, 11, 13, 5, 2, 3, 11}
	executar_estrategia(&primes, removeDuplicates)
	fmt.Println(primes)
}
