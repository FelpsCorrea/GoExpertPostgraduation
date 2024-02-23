package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// O contexto será o próprio da requisição
	ctx := r.Context()

	// Indica no Stdout que iniciou a requisição
	log.Println("Request iniciada")

	// Depois de executar tudo, indica no Stdout que a request foi encerrada
	defer log.Println("Request Finalizada")

	select {

	case <-time.After(5 * time.Second):
		log.Println("Request processada com sucesso")

		// Devolve resposta ao cliente
		w.Write([]byte("Request processada com sucesso"))

	// Caso haja um cancel do contexto do request (parada forçada pelo cliente)
	// Encerra os processamentos vinculados a esse request
	case <-ctx.Done():
		log.Println("Request cancelada pelo cliente")
	}
}
