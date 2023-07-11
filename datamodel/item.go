package datamodel

import "gorm.io/gorm"

type Item struct {
	IDModelIntrinsic
	ItemKey string `gorm:"column:item_key;unique" json:"item_key"`
	Name    string `gorm:"column:name" json:"name"`

	TimestampModel
}

func (i *Item) TableName() string {
	return "item_tab"
}

// -------- Item List

type ItemList []*Item

func (l *ItemList) LoadAll(db *gorm.DB) error {
	err := db.Model(&Item{}).Find(&l).Error
	if err != nil {
		return err
	}

	return nil
}
