package database

import (
	"github.com/St-Pavlov-Data-Department/backend/constants"
	"github.com/St-Pavlov-Data-Department/backend/datamodel"
	"github.com/St-Pavlov-Data-Department/backend/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

func checkSqliteDBExist(dbPath string) error {
	logger := log.CurrentModuleLogger()

	dbFileInfo, err := os.Stat(dbPath)
	switch {
	case os.IsNotExist(err):
		logger.WithField("sqlite_path", dbPath).
			Warnf("Sqlite db file not exist")
		if err := initSqlite(dbPath); err != nil {
			logger.WithError(err).
				Errorf("failed to initialize database, exitting")
			return err
		}
	case err != nil:
		logger.WithError(err).
			WithField("sqlite_path", dbPath).
			Errorf("Failed to open sqlite db file, exiting")
		return err
	case dbFileInfo.IsDir():
		logger.WithField("sqlite_path", dbPath).
			Errorf("Detected sqlite db file path, but the path is a directory instead of a file")
		return constants.ErrSqliteDBPathNotFile
	default:
		logger.WithField("sqlite_path", dbPath).
			Infof("Detected sqlite db file")
	}

	return nil
}

func ConnectSqlite(sqlitePath string) (*gorm.DB, error) {
	if err := checkSqliteDBExist(sqlitePath); err != nil {
		return nil, err
	}

	gormDB, err := gorm.Open(
		sqlite.Open(sqlitePath),
	)
	return gormDB, err
}

func initSqlite(dbPath string) error {
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
