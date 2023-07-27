package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type BuscarCep struct {
	Cep        string `json:"cep"`
	Bairro     string `json:"bairro"`
	Cidade     string `json:"localidade"`
	Logradouro string `json:"logradouro"`
	UF         string `json:"uf"`
}

func main() {
	var cep string
	fmt.Print("Digite o CEP: ")
	fmt.Scanln(&cep)

	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	response, erro := http.Get(url)
	if erro != nil {
		fmt.Println("Erro ao buscar o CEP:", erro)
		return
	}
	defer response.Body.Close()

	var endereco BuscarCep
	erro = json.NewDecoder(response.Body).Decode(&endereco)
	if erro != nil {
		fmt.Println("Erro ao decodificar a resposta:", erro)
		return
	}

	fmt.Println("CEP:", endereco.Cep)
	fmt.Println("Bairro:", endereco.Bairro)
	fmt.Println("Cidade:", endereco.Cidade)
	fmt.Println("Logradouro:", endereco.Logradouro)
	fmt.Println("UF:", endereco.UF)
	
}
