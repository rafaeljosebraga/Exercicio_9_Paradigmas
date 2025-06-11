package main

import "fmt"

type function func(string, string) bool

func validate(lista int[...], fun function) string {
	if fun(str, validator) {
		return str
	}
	return ""
}

func larger_than(str1, str2 string) bool {
	return len(str1) > len(str2)
}

func main() {
	a := "Luciana"
	b := "Linda"
	fmt.Println(validate(a, b, larger_than))
}
