package main

import (
	"fmt"
)

func main() {
	// wybierzKomende("aslkdjkas")
	// wybierzKomende("lista ikonek")
	wybierzKomende("druga ikonka")
	ikonka := []string{"oo"}
	for index := range ikonka {
		fmt.Println("index numer", index, " dlugosc tablicy", len(ikonka))
	}
	fmt.Println(ikonka[len(ikonka)-1])
}

func listaIcon() []string {
	return []string{"ikonka1", "ikonk2", "ikonka3"}
}

func wybierzKomende(cmd string) {
	ikonki := listaIcon()
	switch cmd {
	case "lista ikonek":
		wyswietlIkonki(ikonki)
	case "druga ikonka":
		wyswieltIkonkeNumer(3, ikonki)
	default:
		fmt.Println("Wpisales nieprawidlowa komende")
	}
}

func wyswietlIkonki(ikonki []string) {
	fmt.Println(ikonki)
}

func wyswieltIkonkeNumer(index int, ikonki []string) {
	fmt.Println(len(ikonki), index)
	if len(ikonki) < index+1 {
		fmt.Println("Nie ma takiej ikonki")
		return
	}
	fmt.Println(ikonki[index])
}
