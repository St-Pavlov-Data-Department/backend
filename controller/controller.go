package controller

import (
	"github.com/St-Pavlov-Data-Department/backend/config"
	"github.com/St-Pavlov-Data-Department/backend/gameresource"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type PavlovController struct {
	Cfg          *config.Config
	router       *gin.Engine
	db           *gorm.DB
	gameResource *gameresource.Resource
}

func New(cfg *config.Config, db *gorm.DB, gameResource *gameresource.Resource) *PavlovController {
	controller := &PavlovController{
		Cfg:          cfg,
		db:           db,
		gameResource: gameResource,
	}
	controller.init()

	return controller
}

func (r *PavlovController) init() {
	// init gin router
	gin.DisableConsoleColor()
	gin.SetMode(r.Cfg.GinMode)

	ginRouter := gin.New()
	ginRouter.LoadHTMLGlob("template/**/*")
	ginRouter.Static("./static", "static")

	r.router = ginRouter

	r.router.NoRoute(r._404Handler)
	// register handlers
	r.router.GET("/", r.indexHandler)

	// register documentation
	r.router.GET("/api/doc", r.documentHandler)

	// register APIs
	r.router.GET("/api/version", r.versionHandler)
	r.router.POST("/api/report", r.reportHandler)
	r.router.GET("/api/matrix", r.matrixHandler)
}

func (r *PavlovController) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.router.ServeHTTP(w, req)
}
