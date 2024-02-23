package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {

	// Abrindo o client HTTP
	c := http.Client{}

	// Slice de bytes para enviar no body da requisição
	jsonVar := bytes.NewBuffer([]byte(`{"name": "felipe"}`))

	// Requisição post
	resp, err := c.Post("http://google.com", "application/json", jsonVar)

	// Caso dê erro
	if err != nil {
		panic(err)
	}

	// Fechar automaticamente o body no final de tudo
	defer resp.Body.Close()

	// Copiar os dados de body e jogar no Stdout
	io.CopyBuffer(os.Stdout, resp.Body, nil)

}
