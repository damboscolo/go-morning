package api

import (
	"fmt"
	"go-morning/adapter/repository"
	"go-morning/adapter/rest"
	"go-morning/mapper"
	"go-morning/model/request"
	"math/rand"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
)

func SendMorningMessage(c *gin.Context) {
	var mornings, err = repository.GetMorningMessages()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	rand.Seed(time.Now().UnixNano())
	randIdx := rand.Intn(len(mornings))
	rest.SendTelegramMessage(mornings[randIdx].Message)
	fmt.Println("result=", mornings[randIdx])
}

func SaveDani(c *gin.Context) {
	var body request.UpsertDani
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	dbmodel := mapper.Request_to_dbmodel(body)

	repository.SaveDani(dbmodel)
}
