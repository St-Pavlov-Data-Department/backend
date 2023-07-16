package controller

import (
	"github.com/St-Pavlov-Data-Department/backend/buildinfo"
	"github.com/gin-gonic/gin"
)

func (r *PavlovController) versionHandler(c *gin.Context) {
	RespJSON(c,
		buildinfo.New(),
		0,
		"",
	)
}
