package framework

import (
	"github.com/St-Pavlov-Data-Department/backend/constants"
	"github.com/St-Pavlov-Data-Department/backend/datamodel"
	"github.com/St-Pavlov-Data-Department/backend/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

func (e *PavlovEngine) connectDB() error {
	logger := log.CurrentModuleLogger()

	dbFileInfo, err := os.Stat(e.Cfg.SqlitePath)
	switch {
	case os.IsNotExist(err):
		logger.WithField("sqlite_path", e.Cfg.SqlitePath).
			Warnf("Sqlite db file not exist")
		if err := e.initDB(e.Cfg.SqlitePath); err != nil {
			logger.WithError(err).
				Errorf("failed to initialize database, exitting")
			return err
		}
	case err != nil:
		logger.WithError(err).
			WithField("sqlite_path", e.Cfg.SqlitePath).
			Errorf("Failed to open sqlite db file, exiting")
		return err
	case dbFileInfo.IsDir():
		logger.WithField("sqlite_path", e.Cfg.SqlitePath).
			Errorf("Detected sqlite db file path, but the path is a directory instead of a file")
		return constants.SqliteDBPathNotFileErr
	default:
		logger.WithField("sqlite_path", e.Cfg.SqlitePath).
			Infof("Detected sqlite db file")
	}

	gormDB, err := gorm.Open(
		sqlite.Open(e.Cfg.SqlitePath),
	)
	if err != nil {
		logger.WithError(err).
			Errorf("Failed to open sqlite!")
		return err
	}
	e.Db = gormDB

	return nil
}

func (e *PavlovEngine) initDB(dbPath string) error {
	logger := log.CurrentModuleLogger()

	db, err := gorm.Open(
		sqlite.Open(dbPath),
	)
	dbConn, _ := db.DB()

	if err != nil {
		logger.WithError(err).
			Errorf("initialize db error")
		return err
	}

	datamodel.InitDataModel(db)

	if err := dbConn.Close(); err != nil {
		logger.WithError(err).
			Errorf("close db connection error")
		return err
	}

	return nil
}
