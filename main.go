package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type BrasilApiResponse struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
}

type ViaCepResponse struct {
	Cep          string `json:"cep"`
	State        string `json:"uf"`
	City         string `json:"localidade"`
	Neighborhood string `json:"bairro"`
	Street       string `json:"logradouro"`
}

func callBrasilApi(ch chan BrasilApiResponse) {
	req, err := http.NewRequest("GET", "https://brasilapi.com.br/api/cep/v1/97015121", nil)
	if err != nil {
		panic(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	var result BrasilApiResponse
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		panic(err)
	}
	ch <- result
}

func callViaCepApi(ch chan ViaCepResponse) {
	req, err := http.NewRequest("GET", "http://viacep.com.br/ws/97015121/json/", nil)
	if err != nil {
		panic(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	var result ViaCepResponse
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		panic(err)
	}

	ch <- result
}

func main() {
	brasilApiChannel := make(chan BrasilApiResponse)
	viaCepChannel := make(chan ViaCepResponse)

	go callBrasilApi(brasilApiChannel)
	go callViaCepApi(viaCepChannel)

	select {
	case brasilApiResult := <-brasilApiChannel:
		println("Brasil API")
		println(brasilApiResult.Cep)
		println(brasilApiResult.State)
		println(brasilApiResult.City)
		println(brasilApiResult.Neighborhood)
		println(brasilApiResult.Street)
	case viaCepResult := <-viaCepChannel:
		println("Via Cep")
		println(viaCepResult.Cep)
		println(viaCepResult.State)
		println(viaCepResult.City)
		println(viaCepResult.Neighborhood)
		println(viaCepResult.Street)
	case <-time.After(1 * time.Second):
		println("Timeout")
	}

}
