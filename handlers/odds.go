package handlers

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
)

var oddsApiKey = os.Getenv("ODDS_API_KEY")

const (
	NFLUri     = "https://api.the-odds-api.com/v4/sports/americanfootball_nfl/odds"
	collegeUri = "https://api.the-odds-api.com/v4/sports/americanfootball_ncaaf/odds"
	oddsRoute  = "/odds/"
	bookMaker  = "draftkings"
	region     = "us"
	oddsFormat = "american"
)

type Spread struct {
	FavoriteTeam string `json:"favorite_team"`
	DogTeam      string `json:"dog_team"`
}

func GetSports(c *gin.Context) {
	client := http.Client{}

	request, err := http.NewRequest("GET", NFLUri, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "test"})
	}
	query := request.URL.Query()
	query.Add("apiKey", oddsApiKey)
	query.Add("regions", region)
	query.Add("oddsFormat", oddsFormat)
	query.Add("bookmaker", bookMaker)
	request.URL.RawQuery = query.Encode()

	response, err := client.Do(request)
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "test"})
	}

	c.JSON(200, string(body))
}

func GetNFLSpreads(c *gin.Context) {
	client := http.Client{}

	request, err := http.NewRequest("GET", NFLUri, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "test"})
	}
	query := request.URL.Query()
	query.Add("apiKey", oddsApiKey)
	query.Add("regions", region)
	query.Add("oddsFormat", oddsFormat)
	query.Add("bookmaker", bookMaker)

	request.URL.RawQuery = query.Encode()

	response, err := client.Do(request)
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "test"})
	}

	c.JSON(200, string(body))

}
