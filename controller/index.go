package controller

import (
	"github.com/St-Pavlov-Data-Department/backend/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (r *PavlovController) indexHandler(c *gin.Context) {
	logger := log.CurrentModuleLogger()
	logger.Info("indexHandler")

	c.HTML(http.StatusOK,
		"index.html",
		gin.H{
			"Items": []interface{}{
				map[string]string{
					"name": "API Document",
					"url":  "/api/doc",
				},
			},
		},
	)
}
