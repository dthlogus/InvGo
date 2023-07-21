package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dthlogus/InvGo/models"
	"github.com/dthlogus/InvGo/responses"
	"github.com/gin-gonic/gin"
)

func GetStocks() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		simbolos := c.Param("symbols")
		defer cancel()

		url := "https://yahoo-finance15.p.rapidapi.com/api/yahoo/qu/quote/" + simbolos

		req, _ := http.NewRequest("GET", url, nil)

		req.Header.Add("X-RapidAPI-Key", "dddcf7a2b2mshbbd05839cc265e0p1c7087jsn76cc414c5a76")
		req.Header.Add("X-RapidAPI-Host", "yahoo-finance15.p.rapidapi.com")

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		defer res.Body.Close()

		var stockData []models.StockData
		err = json.NewDecoder(res.Body).Decode(&stockData)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		var stockDtos []models.StockDTO
		stockDtos = models.ToDTOs(stockData)

		c.JSON(http.StatusOK,
			responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": stockDtos}},
		)
		fmt.Println(stockDtos)
	}
}
