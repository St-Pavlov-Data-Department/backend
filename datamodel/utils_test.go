package datamodel

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm/logger"
	"testing"
)

func Test_NewTestDB(t *testing.T) {
	dbName := "Test_NewTestDB"
	dbPath := testDBPath(dbName)
	assert.NotNil(t, dbPath)

	testDB, remove := NewTestDB(dbName, logger.Silent)
	assert.NotNil(t, testDB)
	assert.NotNil(t, remove)

	remove()
}
