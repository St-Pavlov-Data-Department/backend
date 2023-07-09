package controller

import (
	"github.com/St-Pavlov-Data-Department/backend/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type PavlovController struct {
	Cfg    *config.Config
	router *gin.Engine
	db     *gorm.DB
}

func New(cfg *config.Config, db *gorm.DB) *PavlovController {
	controller := &PavlovController{
		Cfg: cfg,
		db:  db,
	}
	controller.init()

	return controller
}

func (r *PavlovController) init() {
	// init gin router
	gin.DisableConsoleColor()
	gin.SetMode(r.Cfg.GinMode)

	ginRouter := gin.New()

	r.router = ginRouter

	// register handlers
	r.router.GET("/", r.indexHandler)
}

func (r *PavlovController) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.router.ServeHTTP(w, req)
}
