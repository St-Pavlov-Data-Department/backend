package database

import (
	"fmt"
	"github.com/St-Pavlov-Data-Department/backend/config"
	"github.com/St-Pavlov-Data-Department/backend/constants"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"strings"
)

func ConnectMySQL(connInfo *config.MySQLConnInfo) (*gorm.DB, error) {

	// validate connection info
	if strings.TrimSpace(connInfo.Username) == "" {
		return nil, constants.ErrConfMySQLUsernameUnset
	}
	if strings.TrimSpace(connInfo.Password) == "" {
		return nil, constants.ErrConfMySQLPasswordUnset
	}
	if strings.TrimSpace(connInfo.Host) == "" {
		return nil, constants.ErrConfMySQLHostUnset
	}
	if strings.TrimSpace(connInfo.Port) == "" {
		return nil, constants.ErrConfMySQLPortUnset
	}
	if strings.TrimSpace(connInfo.DBName) == "" {
		return nil, constants.ErrConfMySQLDBNameUnset
	}

	dsn := fmt.Sprintf(constants.DSNFormat,
		connInfo.Username, connInfo.Password,
		connInfo.Host, connInfo.Port,
		connInfo.DBName,
		connInfo.ConnectTimeout,
	)

	gormDB, err := gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{
			Logger: logger.Default.LogMode(connInfo.GormLogLevel),
		},
	)
	if err != nil {
		return nil, err
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(connInfo.MaxOpenConns)
	sqlDB.SetMaxIdleConns(connInfo.MaxIdleConns)

	return gormDB, nil
}
