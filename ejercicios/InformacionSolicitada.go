package ejercicios

import (
	"guia6/dictionary"
)

func InformacionSolicitada(entrada dictionary.Dictionary[string, []string]) *dictionary.Dictionary[string, []string] {
	d := dictionary.NewDictionary[string, []string]()

	for _, dia := range entrada.GetKeys() {
		for _, alumno := range entrada.Get(dia) {
			arr := []string{}
			if d.Contains(alumno) {
				arr = d.Get(alumno)
			}
			arr = append(arr, dia)
			d.Put(alumno, arr)
		}
	}
	return &d
}
