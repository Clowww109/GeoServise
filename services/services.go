package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ekomobile/dadata/v2"
	"github.com/ekomobile/dadata/v2/api/clean"
	"github.com/ekomobile/dadata/v2/api/model"
	"github.com/ekomobile/dadata/v2/client"
	"io"
	"log"
	"net/http"
	"os"
)

type RequestAddressSearch struct {
	Query string `json:"query"`
}

type ResponseAddress struct {
	Addresses []model.Address `json:"addresses"`
}

type RequestAddressGeocode struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

func AddCredential(credentials client.Credentials) client.Option {
	return client.WithCredentialProvider(&credentials)
}

func NewClientCredentials() client.Credentials {
	return client.Credentials{
		ApiKeyValue:    "86ba621e7a8aa0efa0fd8057cd4e889d734cbbd8",
		SecretKeyValue: "9a72b95eb027c3120069c6da0755a7fcfd3299d9",
	}
}

func NewWorkApi() *clean.Api {

	return dadata.NewCleanApi(AddCredential(NewClientCredentials()))
}

// @Summary getAddress
// @Tags address
// @Description return address from geo
// @Accept json
// @Produce json
// @Param input body RequestAddressSearch true "Нужно указать query(улица + дом в строке)"
// @Success 200
// @Failure 400 "Неверный формат запроса"
// @Failure 500 "Cервис https://dadata.ru не доступен"
// @Router /api/address/search [post]
func GetAddress(w http.ResponseWriter, r *http.Request) {
	var ra RequestAddressSearch
	data, err := io.ReadAll(r.Body)

	err = json.Unmarshal(data, &ra)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	api := NewWorkApi()
	result, err := api.Address(context.Background(), ra.Query)
	if err != nil {
		fmt.Println(err)
		return
	}
	var res ResponseAddress
	for _, r := range result {
		res.Addresses = append(res.Addresses, *r)
	}

	data, err = json.Marshal(res.Addresses)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	log.Println("Данные отправлены клиенту")
	w.Write(data)
}

// @Summary getGeocode
// @Tags geocode
// @Description return closest address to point
// @Accept json
// @Produce json
// @Param input body RequestAddressGeocode true "Нужно указать query(улица + дом в строке)"
// @Success 200
// @Failure 400 "Неверный формат запроса"
// @Failure 500 "Cервис https://dadata.ru не доступен"
// @Router /api/address/geocode [post]
func GetGeocode(w http.ResponseWriter, r *http.Request) {

	data, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	url := "https://suggestions.dadata.ru/suggestions/api/4_1/rs/geolocate/address"
	token := "86ba621e7a8aa0efa0fd8057cd4e889d734cbbd8"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Token "+token)

	cli := http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		if resp.StatusCode == http.StatusBadRequest {
			w.WriteHeader(400)
		} else if resp.StatusCode == http.StatusInternalServerError {
			w.WriteHeader(500)
		}
		fmt.Println(err)
		os.Exit(1)
	}
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	log.Println("Данные отправлены клиенту")
	w.Write(data)
}
