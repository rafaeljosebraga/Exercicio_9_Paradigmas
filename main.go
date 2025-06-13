package main

import "fmt"

type listFunction func([]int) []int

func executar_estrategia(lista *[]int, estrategia listFunction) {
	*lista = estrategia(*lista)
}

func sortDescending(array []int) []int {
	array = sortAscending(array)
	var size int = len(array)
	rev := make([]int, size)

	for i := 0; i < size; i++ {
		rev[size-1-i] = array[i]
	}
	return rev
}

func main() {
	primes := []int{11, 2, 2, 2, 5, 3, 3, 11, 5, 11, 13, 5, 2, 3, 11}
	// executar_estrategia(&primes, removeDuplicates)
	fmt.Println(primes)
	
	SortDescendingList := executar_estrategia(primes, sortDescending)
	fmt.Println("Sorted descending list: ", listSortDescending)
}
