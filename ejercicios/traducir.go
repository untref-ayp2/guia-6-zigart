package ejercicios

import (
	"fmt"
	"guia6/dictionary"
)

func Traducir(texto string, dic dictionary.Dictionary[string, string]) string {
	stringFiltered := Split(texto, " ")
	result := ""
	for i := 0; i < len(stringFiltered); i++ {
		if dic.Contains(stringFiltered[i]) {

			if i == len(stringFiltered)-1 {
				result += dic.Get(stringFiltered[i])
			} else {
				result += dic.Get(stringFiltered[i]) + " "
			}
		} else {
			result += "error "
		}
	}
	return result
}

func Split(str, sep string) []string {
	strings := make([]string, 0)
	aux := 0
	for i := 0; i < len(str); i++ {

		if str[i] == sep[0] {
			strings = append(strings, string(str[aux:i]))
			aux = i + 1
		}

		if i == len(str)-1 {
			strings = append(strings, string(str[aux:]))
		}
	}
	fmt.Print(strings)
	return strings
}
