package summoner_info_tab

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"net/http"
	"os"
)

type Summoner struct {
	ID        string `json:"id"`
	AccountID string `json:"accountId"`
	PUUID     string `json:"puuid"`
}

func Summoner_checker(summonerName string) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error loading .env file: %s\n", err.Error())
		return
	}

	region := os.Getenv("REGION")
	APIKey := os.Getenv("KEY")

	url := fmt.Sprintf("https://%s.api.riotgames.com/lol/summoner/v4/summoners/by-name/%s?api_key=%s", region, summonerName, APIKey)
	fmt.Println(url)
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error sending request: %s\n", err.Error())
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %s\n", err.Error())
		return
	}

	if response.StatusCode != http.StatusOK {
		fmt.Printf("Request failed with status code: %d\n", response.StatusCode)
		fmt.Println(string(body))
		return
	}

	var summoner Summoner
	err = json.Unmarshal(body, &summoner)
	if err != nil {
		fmt.Printf("Error parsing JSON response: %s\n", err.Error())
		return
	}

	fmt.Println("ID:", summoner.ID)
	fmt.Println("Account ID:", summoner.AccountID)
	fmt.Println("PUUID:", summoner.PUUID)
}
