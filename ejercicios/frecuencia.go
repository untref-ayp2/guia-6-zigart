package ejercicios

import (
	"guia6/dictionary"

	"github.com/agrison/go-commons-lang/stringUtils"
)

func Frecuencia(s string) *dictionary.Dictionary[string, int] {
	frec := dictionary.NewDictionary[string, int]()
	for _, letter := range s {
		letra := string(letter)

		if stringUtils.IsAlpha(letra) {
			if frec.Contains(letra) {
				frec.Put(letra, frec.Get(letra)+1)
			} else {
				frec.Put(letra, 1)
			}
		}
	}
	return &frec
}
