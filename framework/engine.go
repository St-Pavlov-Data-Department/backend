package framework

import (
	"context"
	"github.com/St-Pavlov-Data-Department/backend/config"
	"github.com/St-Pavlov-Data-Department/backend/constants"
	"github.com/St-Pavlov-Data-Department/backend/controller"
	"github.com/St-Pavlov-Data-Department/backend/gameresource"
	"github.com/St-Pavlov-Data-Department/backend/log"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"os"
	"time"
)

type PavlovEngine struct {
	Cfg *config.Config

	Logger       *logrus.Logger
	Server       *http.Server
	Db           *gorm.DB
	GameResource *gameresource.Resource
}

func New() *PavlovEngine {
	return &PavlovEngine{Cfg: &config.Config{}}
}

func (e *PavlovEngine) Init() error {
	logger := log.CurrentModuleLogger()

	fileInfo, err := os.Stat(constants.PavlovConfigFilePath)
	switch {
	case os.IsNotExist(err):
		logger.WithField("config_filepath", constants.PavlovConfigFilePath).
			Warnf("No configuration file detected, generating one. You can ignore this warning if it is the first time running.")
		err = generateExampleConfig(constants.PavlovConfigFilePath)
		if err != nil {
			logger.WithError(err).
				Errorf("failed to generate example config, exitting")
			return err
		}
		return constants.ErrExampleConfigGenerated

	case err != nil:
		logger.WithError(err).
			WithField("config_filepath", constants.PavlovConfigFilePath).
			Errorf("Failed to check the configuration file, exiting")
		return err

	default:
		if fileInfo.IsDir() {
			logger.WithField("config_filepath", constants.PavlovConfigFilePath).
				Errorf("Detected configuration file path, but the path is a directory instead of a file")
			return constants.ErrConfigPathNotFile
		} else {
			logger.WithField("config_filepath", constants.PavlovConfigFilePath).
				Infof("Detected configuration file, starting with existing configuration file")
		}
	}

	confReloadCallback := func(config interface{}) {
		logger.Infof("config file reloaded: %v", config)
	}
	err = e.loadConfig(confReloadCallback)
	if err != nil {
		logger.WithError(err).
			Errorf("Failed to load configuration file! Please check if the configuration file format is correct")
		return err
	}

	// init logger
	logLevel, err := logrus.ParseLevel(e.Cfg.LogLevel)
	if err != nil {
		logrus.SetLevel(logrus.DebugLevel)
		logger.Warn("Unable to recognize logLevel, will use Debug level")
	} else {
		logrus.SetLevel(logLevel)
		logger.Infof("Set log level to %s", logLevel.String())
	}
	e.Logger = log.StandardLogger()

	// load game resources
	e.GameResource, err = gameresource.NewFromPath(e.Cfg.GameResourcePath)
	if err != nil {
		logger.WithError(err).
			Errorf("load game resource error")
		return err
	}

	// init database
	if err := e.connectDB(); err != nil {
		logger.WithError(err).
			Errorf("connect database error")
		return err
	}

	// init tasks

	// init router
	e.InitHTTPServer()

	return nil
}

func (e *PavlovEngine) InitHTTPServer() {

	apiController := controller.New(
		e.Cfg, e.Db, e.GameResource,
	)

	e.Server = &http.Server{
		Addr:    e.Cfg.ListenAddr,
		Handler: apiController,
	}
}

func (e *PavlovEngine) StartService() {

	go func() {
		if err := e.Server.ListenAndServe(); err != nil {
			e.Logger.WithError(err).
				Error("router error while listening")
		}
	}()

}

func (e *PavlovEngine) Stop() error {
	e.Logger.Info("St.Pavlov framework stopping ...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(e.Cfg.ServerShutdownMaxWaitSeconds)*time.Second)
	defer cancel()

	select {
	case <-ctx.Done():
		e.Logger.Error("server shutdown timeout")
	default:
		e.Logger.Info("server exiting")
	}
	e.Logger.Info("server exited")

	if err := e.Server.Shutdown(ctx); err != nil {
		e.Logger.WithError(err).
			Errorf("Shutdown http Server error")
		return err
	}

	dbConn, err := e.Db.DB()
	if err != nil {
		e.Logger.WithError(err).
			Errorf("get gorm db connection for close error")
	}
	if err := dbConn.Close(); err != nil {
		e.Logger.WithError(err).
			Errorf("Close Sqlite connection error")
		return err
	}

	return nil
}
