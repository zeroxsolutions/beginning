package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct {}

// Health Health check
// @Tags Health
// @Produce plain
// @Success 200 {string} OK
// @Router /health [get]
func (healthController *HealthController) Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "OK")
}

func NewHealthController() *HealthController {
	return &HealthController{}
}
