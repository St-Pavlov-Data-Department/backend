package framework

import (
	"github.com/St-Pavlov-Data-Department/backend/database"
	"github.com/St-Pavlov-Data-Department/backend/log"
)

func (e *PavlovEngine) connectDB() error {
	logger := log.CurrentModuleLogger()

	// gormDB, err := database.ConnectSqlite(e.Cfg.SqlitePath)
	gormDB, err := database.ConnectMySQL(e.Cfg.DBConnInfo)
	if err != nil {
		logger.WithError(err).
			Error("connect database error")
		return err
	}
	e.Db = gormDB

	return nil
}
