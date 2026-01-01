package pokeapi

import (
    "encoding/json"
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

type Location struct {
    Name              string `json:"name"`
    PokemonEncounters []struct {
        Pokemon struct {
            Name string `json:"name"`
            URL  string `json:"url"`
        } `json:"pokemon"`
    } `json:"pokemon_encounters"`
}

const baseURL = "https://pokeapi.co/api/v2"

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

func (c *Client) GetLocation(locationName string) (Location, error) {
    fullURL := baseURL + "/location-area/" + locationName

    if val, ok := c.cache.Get(fullURL); ok {
        loc := Location{}
	    if err := json.Unmarshal(val, &loc); err != nil {
		    return Location{}, err
	    }

        return loc, nil
    }

    req, err := http.NewRequest("GET", fullURL, nil)
    if err != nil {
        return Location{}, err
    }

    res, err := c.httpClient.Do(req)
    if err != nil {
		return Location{}, err
	}
	defer res.Body.Close()
    
    dat, err := io.ReadAll(res.Body)
	if err != nil {
		return Location{}, err
	}

    loc := Location{}
    if err := json.Unmarshal(dat, &loc); err != nil {
        return Location{}, err
    }

    c.cache.Add(fullURL, dat)

    return loc, nil
}