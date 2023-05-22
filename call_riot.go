package main

import (
	"Typering/riot_data/package_factory"
	"fmt"
)

func main() {
	players, err := package_factory.GetMaster()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	i := 0
	for _, player := range players {
		i += 1
		summonerName := fmt.Sprintf("%d. Summoner Name: %s", i, player.SummonerName)
		fmt.Println(summonerName)
	}
}
