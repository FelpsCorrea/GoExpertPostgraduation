package main

import (
	"context"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	// Cria um contexto vazio
	ctx := context.Background()

	// Aplica uma regra de timeout
	ctx, cancel := context.WithTimeout(ctx, time.Second)

	// Deixa a função cancel em background para rodar depois de todo sistema terminar a execução
	defer cancel()

	// Prepara a requisição
	req, err := http.NewRequestWithContext(ctx, "GET", "http://google.com", nil)

	if err != nil {
		panic(err)
	}

	// Executa a requisição
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	println(string(body))

}
