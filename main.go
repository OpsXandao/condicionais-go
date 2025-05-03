package main

import "fmt"

func saudacao(turno string) string {
	switch turno {
	case "manhã":
		return "Bom dia!"
	case "tarde":
		return "Boa tarde"
	case "noite":
		return "Boa noite"
	default:
		return "Turno inválido"
	}
}

func main() {
	fmt.Println(saudacao("manhã"))
}
