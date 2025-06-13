//Código por luis henrique fernandes vieira
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

func sortDescending(array []int) []int {
	array = sortAscending(array)
	var size int = len(array)
	rev := make([]int, size)

	for i := 0; i < size; i++ {
		rev[size-1-i] = array[i]
	}
	return rev
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
	primes := []int{2, 2, 2, 5, 3, 3, 11, 5, 11, 13, 5, 2, 3, 11}
	executar_estrategia(&primes, removeDuplicates)
	fmt.Println(primes)
}
