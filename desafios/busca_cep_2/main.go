package main

import "net/http"

func main() {
	// Definição de rota
	http.HandleFunc("/", BuscaCepHandler)
	// Abrindo servidor na porta 8080
	http.ListenAndServe(":8080", nil)
}

// w é o retorno ao usuário e r é como o "Body"
// func BuscaCEP(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Hello, World!"))
// }
func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World!"))
}
