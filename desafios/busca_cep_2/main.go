package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

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

	// Caso a url seja diferente de /
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	cepParam := r.URL.Query().Get("cep")

	// Caso o cep não tenha sido informado
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Chama a função busca cep passando o cep como parametro
	cep, error := BuscaCep(cepParam)

	// Algo como o except
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Vai encodar a struct e jogar no nosso writer (response ao usuário)
	json.NewEncoder(w).Encode(cep)
}

func BuscaCep(cep string) (*ViaCEP, error) {

	// Faz a requisição com o cep passado como argumento
	resp, error := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
	if error != nil {
		return nil, error
	}

	defer resp.Body.Close()

	body, error := ioutil.ReadAll(resp.Body)
	if error != nil {
		return nil, error
	}

	var c ViaCEP
	error = json.Unmarshal(body, &c)

	return &c, nil
}
