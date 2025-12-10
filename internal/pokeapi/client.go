package pokeapi

import (
	"io"
	"net/http"
	"encoding/json"
)

func getHTTP(url string) (LocationAreaResponse, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	resp, err := unmarshalData(body)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	return resp, nil
}

func unmarshalData(jsonData []byte) (LocationAreaResponse, error) {
	var resp LocationAreaResponse

	if err := json.Unmarshal(jsonData, &resp); err != nil {
		return LocationAreaResponse{}, err
	}

	return resp, nil
}