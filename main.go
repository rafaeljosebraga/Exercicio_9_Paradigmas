package main

import "fmt"

type listFunction func([]int) []int

func executar_estrategia(lista *[]int, estrategia listFunction) {
	*lista = estrategia(*lista)
}

func filterEven(array []int) []int {
	rem := make([]int, len(array))
	k := 0

	for i := 0; i < len(array); i++ {
		if array[i]%2 == 0 {
			rem[k] = array[i]
			k++
		}
	}
	return rem[0:k]
}

func main() {
	primes := []int{11, 2, 2, 2, 5, 3, 3, 11, 5, 11, 13, 5, 2, 3, 11}
	// executar_estrategia(&primes, removeDuplicates)
	fmt.Println(primes)
	
	primes2 := []int{11, 2, 2, 2, 5, 3, 3, 11, 5, 11, 13, 5, 2, 3, 11}
	executar_estrategia(&primes2, sortAscending)	
	fmt.Println(primes2)
}
