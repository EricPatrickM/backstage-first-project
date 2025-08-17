package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handlers
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", helloHandler)

	// Iniciar o servidor na porta 8080
	log.Println("Iniciando servidor na porta 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Erro ao iniciar o servidor: ", err)
	}
}
