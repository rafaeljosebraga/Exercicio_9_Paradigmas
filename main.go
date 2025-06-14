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
  
func ordenarDecrescente(lista []int) []int {
	tamanho := len(lista)
	copia := make([]int, tamanho)
	copy(copia, lista)

	for i := 0; i < tamanho-1; i++ {
		for j := 0; j < tamanho-1-i; j++ {
			if copia[j] < copia[j+1] {
				copia[j], copia[j+1] = copia[j+1], copia[j]
			}
		}
	}
	return copia

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
	fmt.Println("Remoção de duplicatas: ",primes)
	
	primes2 := []int{11, 2, 2, 2, 5, 3, 3, 11, 5, 11, 13, 5, 2, 3, 11}
	executar_estrategia(&primes2, sortAscending)	
	fmt.Println("Lista Crescente: ",primes2)

  lista1 := []int{11, 2, 2, 2, 5, 3, 3, 11, 5, 11, 13, 5, 2, 3, 11}
	executar_estrategia(&lista1, filterEven)
	fmt.Println("Numeros pares ", lista1)
  
	lista2 := []int{11, 2, 2, 2, 5, 3, 3, 11, 5, 11, 13, 5, 2, 3, 11}
	executar_estrategia(&lista2, ordenarDecrescente)
	fmt.Println("Lista Decrescente: ", lista2)
}
