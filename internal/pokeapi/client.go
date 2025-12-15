package pokeapi

import (
	"encoding/json"
	"net/http"
	"time"
	"github.com/richardw55555/pokedexcli/internal/pokecache"
)

type Client struct {
    cache      *pokecache.Cache
    httpClient http.Client
}

func NewClient(timeout, cacheInterval time.Duration) Client {
    return Client{
        cache: pokecache.NewCache(cacheInterval),
        httpClient: http.Client{
            Timeout: timeout,
        },
    }
}

func unmarshalData(jsonData []byte) (LocationAreaResponse, error) {
	var resp LocationAreaResponse

	if err := json.Unmarshal(jsonData, &resp); err != nil {
		return LocationAreaResponse{}, err
	}

	return resp, nil
}