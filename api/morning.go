package api

import (
	"go-morning/model/request"
	"go-morning/adapter/rest"
	"go-morning/adapter/repository"
	"go-morning/mapper"

	"net/http"

	"github.com/gin-gonic/gin"
)

func SendMorningMessage(c *gin.Context) {
	message := "Bom dia família ❤️"
	rest.SendTelegramMessage(message)
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
