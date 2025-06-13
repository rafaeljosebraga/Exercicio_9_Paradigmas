package main

import "fmt"

type listFunction func([]int) []int

func executar_estrategia(lista *[]int, estrategia listFunction) {
	*lista = estrategia(*lista)
}

func swap(pos1, pos2 *int) {
	var aux int = *pos1
	*pos1 = *pos2
	*pos2 = aux
}

func sortAscending(array []int) []int {
	var swapped bool = true
	var start int = 0
	var end int = len(array) - 1

	for swapped {
		swapped = false

		for i := start; i < end; i++ {
			if array[i] > array[i+1] {
				swap(&array[i], &array[i+1])
				swapped = true
			}
		}
		if !swapped {
			break
		}
		swapped = false
		end--

		for i := end - 1; i >= start; i-- {
			if array[i] > array[i+1] {
				swap(&array[i], &array[i+1])
				swapped = true
			}
		}
		start++
	}
	return array
}

func main() {
	primes := []int{11, 2, 2, 2, 5, 3, 3, 11, 5, 11, 13, 5, 2, 3, 11}
	// executar_estrategia(&primes, removeDuplicates)
	fmt.Println(primes)
}
