package main

import "fmt"

type listFunction func([]int) []int

func executar_estrategia(lista *[]int, estrategia listFunction) {
	*lista = estrategia(*lista)
}

func removeDuplicates(array []int) []int {
	sorted := sortAscending(array)
	cpr := array[0]
	rem := make([]int, len(array))
	rem[0] = cpr
	k := 1

	for i := 1; i < len(array); i++ {
		if cpr != sorted[i] {
			rem[k] = sorted[i]
			cpr = sorted[i]
			k++
		}
	}
	return rem[0:k]
}

func main() {
	primes := []int{11, 2, 2, 2, 5, 3, 3, 11, 5, 11, 13, 5, 2, 3, 11}
	// executar_estrategia(&primes, removeDuplicates)
	fmt.Println(primes)

	listaSemDuplicatas := executar_estrategia(primes, removeDuplicates)
	fmt.Println("Sem Duplicatas: ", listaSemDuplicatas)
}
