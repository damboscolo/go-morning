package usecase

import (
	"go-morning/infrastructure/http"

	"github.com/gin-gonic/gin"
)

func SendMorningMessage(c *gin.Context) {
	message := "Bom dia família ❤️"
	http.SendTelegramMessage(message)
}
