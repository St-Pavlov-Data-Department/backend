package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (r *PavlovController) indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK,
		"index.html",
		nil,
	)
}
