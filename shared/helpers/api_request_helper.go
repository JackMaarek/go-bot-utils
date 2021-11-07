package helpers

import (
	"bytes"
	"encoding/json"
	"github.com/JackMaarek/go-bot-utils/shared/enum"
	"github.com/JackMaarek/go-bot-utils/shared/env"
	"net/http"

)

func PerformRequest(route enum.Routes, method string, data interface{}) (*http.Response, error) {
	client := &http.Client{}
	baseUrl := env.GetVariable("DOMAIN_API_URL")
	requestBody, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	url := baseUrl + string(route)
	r, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))

	response, err := client.Do(r)
	if err != nil {
		return nil, err
	}

	return response, err
}
