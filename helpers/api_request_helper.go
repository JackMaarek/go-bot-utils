package helpers

import (
	"bytes"
	"encoding/json"
	"github.com/JackMaarek/go-bot-utils/env"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
)

func PerformRequest(route string, method string, data interface{}) (*http.Response, error) {
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

func DeserializeResponseFromObject(data io.ReadCloser, object interface{}) error {
	deserializedData, err := ioutil.ReadAll(data)
	if err != nil {
		log.Error("could ot read response for deserialization")
		return err
	}

	err = json.Unmarshal(deserializedData, &object)
	if err != nil {
		log.Error("could not unmarshal response data to object")
		return err
	}
	return nil
}
