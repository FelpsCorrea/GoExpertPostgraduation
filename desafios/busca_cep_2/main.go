package main

import "net/http"

func main() {
	// Definição de rota
	http.HandleFunc("/", BuscaCEP)
	// Abrindo servidor na porta 8080
	http.ListenAndServe(":8080", nil)
}

// w é o retorno ao usuário e r é como o "Body"
func BuscaCEP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
