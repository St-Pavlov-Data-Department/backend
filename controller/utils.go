package controller

import (
	"github.com/St-Pavlov-Data-Department/backend/constants"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RespJSONWithError(c *gin.Context, data interface{}, err error) {
	if err != nil {
		ErrorJSON(c, err)
	} else {
		RespJSON(c, data,
			constants.NoErrorCode,
			"",
		)
	}
}

func ErrorJSON(c *gin.Context, err error) {
	RespJSON(c, nil,
		constants.GeneralErrorCode,
		err.Error(),
	)
}

func RespJSON(c *gin.Context, data interface{}, errorCode int64, errorMessage string) {
	c.JSON(http.StatusOK,
		gin.H{
			"error_code":    errorCode,
			"error_message": errorMessage,
			"data":          data,
		},
	)
}
