package package_factory

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"net/http"
	"os"
)

type Player struct {
	SummonerName string `json:"summonerName"`
}

type GrandmasterLeague struct {
	Tier    string   `json:"tier"`
	Entries []Player `json:"entries"`
}

func GetMaster() ([]Player, error) {
	err := godotenv.Load("../.env")
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	region := os.Getenv("REGION")
	urlKey := os.Getenv("KEY")
	url := fmt.Sprintf("https://%s.api.riotgames.com/tft/league/v1/grandmaster?api_key=%s", region, urlKey)
	fmt.Println(url)

	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP request: %v", err)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making HTTP request: %v", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading HTTP response: %v", err)
	}
	// fmt.Println(string(body))

	var response GrandmasterLeague
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %v", err)
	}

	return response.Entries, nil
}
