package main

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Gaoey/covid-api/models"
	"github.com/Gaoey/covid-api/summary"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/covid/summary", func(c *gin.Context) {
		covidStatResp, err := fetchCovidStatApi()
		if err != nil {
			c.JSON(http.StatusInternalServerError,
				gin.H{
					"statusCode":   http.StatusInternalServerError,
					"errorMessage": err.Error(),
				})
			return
		}

		result, err := summary.GetSummary(covidStatResp.Data)
		if err != nil {
			c.JSON(http.StatusInternalServerError,
				gin.H{
					"statusCode":   http.StatusInternalServerError,
					"errorMessage": err.Error(),
				})
			return
		}

		c.JSON(http.StatusOK,
			gin.H{
				"statusCode": http.StatusOK,
				"data":       result,
			})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func fetchCovidStatApi() (*models.CovidStatResponse, error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	resp, err := http.Get("https://static.wongnai.com/devinterview/covid-cases.json")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var covidStatResponse models.CovidStatResponse
	if err := json.Unmarshal(bodyBytes, &covidStatResponse); err != nil {
		return nil, err
	}

	return &covidStatResponse, nil
}
