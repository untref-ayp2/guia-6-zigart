package ejercicios

import (
	"guia6/linkedlist"
	// agregar import faltante
)

func Interseccion(s1 []string, s2 []string) linkedlist.LinkedList[string] {
	list := linkedlist.NewLinkedList[string]()

	for _, valueOfFirstSlice := range s1 {
		for _, value := range s2 {
			if valueOfFirstSlice == value {
				list.InsertAt(value, 0)
			}
		}
	}

	return *list
}
