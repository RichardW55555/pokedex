package pokeapi

type LocationAreaResponse struct {
    Next     string `json:"next"`
    Previous string `json:"previous"`
    Results  []struct {
        Name string `json:"name"`
        URL  string `json:"url"`
    } `json:"results"`
}

func GetLocationAreas(url string) (LocationAreaResponse, error) {
    return getHTTP(url)
}