package datamodel

type Item struct {
	IDModelIntrinsic
	ItemKey string `gorm:"column:item_key;unique" json:"item_key"`
	Name    string `gorm:"column:name" json:"name"`
}
