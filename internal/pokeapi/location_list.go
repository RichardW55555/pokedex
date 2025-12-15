package pokeapi

import (
    "io"
    "net/http"
)

type LocationAreaResponse struct {
    Next     string `json:"next"`
    Previous string `json:"previous"`
    Results  []struct {
        Name string `json:"name"`
        URL  string `json:"url"`
    } `json:"results"`
}

func (c *Client) GetLocationAreas(url string) (LocationAreaResponse, error) {
    // 1. Try cache
    if val, ok := c.cache.Get(url); ok {
        resp, err := unmarshalData(val)
        if err != nil {
            return LocationAreaResponse{}, err
        }
        return resp, nil
    }
    
    // 2. HTTP request using c.httpClient
    req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}

    // 3. Store in cache
    c.cache.Add(url, body)

    // 4. Unmarshal and return
	resp, err := unmarshalData(body)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	return resp, nil
}