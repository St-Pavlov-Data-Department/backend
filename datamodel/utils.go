package datamodel

import "gorm.io/gorm"

type IDModelIntrinsic struct {
	ID uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
}

func InitDataModel(db *gorm.DB) {
	_ = db.AutoMigrate(&UUID{})
	_ = db.AutoMigrate(&Item{})

	_ = db.AutoMigrate(&LootReport{})
	_ = db.AutoMigrate(&LootItem{})

}
