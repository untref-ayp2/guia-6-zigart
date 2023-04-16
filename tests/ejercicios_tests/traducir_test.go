package tests

import (
	"guia6/dictionary"
	"guia6/ejercicios"
	"testing"
)

func TestTraducir(t *testing.T) {
	dic := dictionary.NewDictionary[string, string]()
	dic.Put("Dungeons", "Calabozos")
	dic.Put("Dragons", "Dragones")
	salida := ejercicios.Traducir("Dungeons & Dragons", dic)
	if salida != "Calabozos error Dragones" {
		t.Errorf("La traducción es %v, pero deberían ser %v", salida, "Calabozos error Dragones")
	}
}
