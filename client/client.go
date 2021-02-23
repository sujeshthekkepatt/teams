package client

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"teams/types"
)

func createClient() http.Client {

	teamsClient := http.Client{}
	return teamsClient

}

func createRequest(url string, body []byte) (http.Request, error) {

	teamsRequest, err := http.NewRequestWithContext(context.Background(), "POST", url, bytes.NewBuffer(body))

	if err != nil {
		log.Fatal(err)
		return *teamsRequest, err
	}
	teamsRequest.Header.Set("Content-Type", "application/json")

	return *teamsRequest, nil
}

//TransportEvent Helps to create a client
func TransportEvent(teamsPayload []byte, teamsURL string) (*http.Response, error) {

	client := createClient()
	request, err := createRequest(teamsURL, teamsPayload)
	if err != nil {
		return nil, err
	}
	teamsResponse, err := client.Do(&request)
	if err != nil {
		return nil, err
	}

	return teamsResponse, nil
}

//MarshalMessage marshal message to json
func MarshalMessage(message types.Payload) (teamMessage []byte, err error) {

	teamMessage, err = json.Marshal(message)

	return teamMessage, err

}
