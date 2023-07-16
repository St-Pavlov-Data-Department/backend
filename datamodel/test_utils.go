package datamodel

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"path/filepath"
)

func NewTestDB(dbName string, logLevel logger.LogLevel) (
	*gorm.DB, func(),
) {
	dbPath := testDBPath(dbName)
	fmt.Println("db path: ", dbPath)

	testDB, err := gorm.Open(
		sqlite.Open(dbPath),
		&gorm.Config{
			Logger: logger.Default.LogMode(logLevel),
		},
	)
	if err != nil {
		fmt.Println("db err: ", err)
		return nil, nil
	}

	InitDataModel(testDB)

	remove := func() {
		FreeTestDB(dbName, testDB)
	}

	return testDB, remove
}

func testDBPath(dbName string) string {
	return filepath.Join("/tmp/pavlov/",
		fmt.Sprintf("%s.sqlite", dbName),
	)
}

func FreeTestDB(dbName string, db *gorm.DB) {
	dbPath := testDBPath(dbName)

	dbConn, _ := db.DB()
	if err := dbConn.Close(); err != nil {
		fmt.Println(err)
	}

	if err := os.Remove(dbPath); err != nil {
		fmt.Println(err)
	}
}
