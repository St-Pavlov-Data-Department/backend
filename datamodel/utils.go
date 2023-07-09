package datamodel

import "gorm.io/gorm"

type IDModelIntrinsic struct {
	ID uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
}

func InitDataModel(dbConn *gorm.DB) {
	_ = dbConn.AutoMigrate(&Item{})
}
