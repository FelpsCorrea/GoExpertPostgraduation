package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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

	// Percorre todos argumentos passados na execução
	for _, cep := range os.Args[1:] {

		// Faz a requisição com o cep passado como argumento
		req, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer a requisição: %v\n", err)
		}

		// Para fechar o Body assim que executar todo o código
		defer req.Body.Close()

		// Para ler o Body da requisição
		res, err := io.ReadAll(req.Body)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler a resposta: %v\n", err)
		}

		// Convertendo o body recebido para Struct
		var data ViaCEP

		err = json.Unmarshal(res, &data)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao converter json para struct: %v\n", err)
		}

		// Criar arquivo
		file, err := os.Create("enderecoViaCEP.txt")

		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao criar o arquivo: %v\n", err)
		}

		// Fechar o arquivo após execução
		defer file.Close()

		_, err = file.WriteString(fmt.Sprintf("**ENDEREÇO COMPLETO**\nRua: %s\nEstado: %s\nCidade: %s\nCEP: %s\nBairro: %s", data.Logradouro, data.Uf, data.Localidade, data.Cep, data.Bairro))

		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao escrever no arquivo: %v\n", err)
		}

		fmt.Println("Arquivo criado com sucesso")
	}

}
