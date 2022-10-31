package api

import (
	"go-morning/entity/request"
	infra_http "go-morning/infrastructure/http"
	infra_repository "go-morning/infrastructure/repository"
	"go-morning/mapper"

	"net/http"

	"github.com/gin-gonic/gin"
)

func SendMorningMessage(c *gin.Context) {
	message := "Bom dia família ❤️"
	infra_http.SendTelegramMessage(message)
}

func SaveDani(c *gin.Context) {
	var body request.UpsertDani
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusAccepted, &body)
	dbmodel := mapper.Request_to_dbmodel(body)

	infra_repository.SaveDani(dbmodel)
}
