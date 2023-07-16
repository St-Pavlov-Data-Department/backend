package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (r *PavlovController) _404Handler(c *gin.Context) {
	c.JSON(http.StatusNotFound,
		nil,
	)
}
