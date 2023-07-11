package datamodel

import "gorm.io/gorm"

type LootItem struct {
	IDModelIntrinsic
	ReportUUID string `gorm:"column:report_uuid" json:"report_uuid"`

	ItemID   string `gorm:"column:item_id" json:"item_id"`
	LootType string `gorm:"column:loot_type" json:"loot_type"`
	Quantity int64  `gorm:"column:quantity" json:"quantity"`

	UUIDModel
	TimestampModel
}

func (i *LootItem) TableName() string {
	return "loot_item_tab"
}

func (i *LootItem) BeforeCreate(db *gorm.DB) error {
	if pavlovUUID, err := i.GetUUID(db); err == nil {
		i.UUID = pavlovUUID.UUID
		return nil
	} else {
		return err
	}
}

func (i *LootItem) Save(db *gorm.DB) error {
	return db.Save(i).Error
}
