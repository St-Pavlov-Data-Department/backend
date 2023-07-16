package controller

import (
	"github.com/St-Pavlov-Data-Department/backend/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (r *PavlovController) documentHandler(c *gin.Context) {
	logger := log.CurrentModuleLogger()
	logger.Info("documentHandler")

	c.HTML(http.StatusOK,
		"redoc.html",
		nil,
	)
}
