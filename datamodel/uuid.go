package datamodel

import (
	"fmt"
	"gorm.io/gorm"
	"math/rand"
)

const GetUUIDRetryLimit = 5

type UUID struct {
	IDModelIntrinsic

	UUIDModel
	TimestampModel
}

func (u *UUID) TableName() string {
	return "uuid_tab"
}

type UUIDModel struct {
	UUID string `gorm:"column:uuid; not null;index:uk_uuid,unique" json:"uuid"`
}

func (u *UUIDModel) GetUUID(db *gorm.DB) (*UUID, error) {
	uuid := &UUID{}
	for count := 0; count < GetUUIDRetryLimit; count++ {
		uuidStr := randUUID()
		uuid.UUID = uuidStr
		if dbErr := db.Model(uuid).Save(uuid).Error; dbErr == nil {
			return uuid, nil
		}
	}

	return nil, fmt.Errorf("could not generate valid UUID in %d retries", GetUUIDRetryLimit)
}

func randUUID() string {

	// generate random number by timestamp
	// [0 ... 27](28 bit) current timestamp, [28...63](36 bits) random number
	// (a full timestamp takes 32 bits, we emit the first 4 bits because they hardly change)
	uuidUint64 := rand.Uint64() //nolint:gosec, gomnd

	// UUID is a 16-character string in HEX format
	return fmt.Sprintf("%016x", uuidUint64)
}
