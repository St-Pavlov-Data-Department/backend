package datamodel

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm/logger"
	"testing"
)

func Test_UUID_GetUUID(t *testing.T) {
	testDB := NewTestDB("pavlov", logger.Info)
	InitDataModel(testDB)
	defer FreeTestDB("pavlov", testDB)

	uuidModel := &UUIDModel{}
	uuid, err := uuidModel.GetUUID(testDB)
	assert.NotNil(t, uuid)
	assert.NoError(t, err)

	var count int64
	err = testDB.Model(uuid).Count(&count).Error
	assert.NoError(t, err)
	assert.Equal(t, int64(1), count)

	uuidLoader := &UUID{}
	err = testDB.Model(uuidLoader).First(uuidLoader).Error
	assert.NoError(t, err)
	assert.Equal(t, uuid.UUID, uuidLoader.UUID)
}
