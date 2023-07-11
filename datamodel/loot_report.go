package datamodel

import "gorm.io/gorm"

type LootReport struct {
	IDModelIntrinsic

	StageID string `gorm:"column:stage_id" json:"stage_id"`
	Server  string `gorm:"column:server" json:"server"`
	Source  string `gorm:"column:source" json:"source"`
	Version string `gorm:"column:version" json:"version"`

	Loot []*LootItem `gorm:"-"`

	ReplayLevel int64 `gorm:"column:replay_level" json:"replay_level"`

	ClientIP string `gorm:"column:client_ip" json:"client_ip"`

	UUIDModel
	TimestampModel
}

func (r *LootReport) TableName() string {
	return "loot_report_tab"
}

func (r *LootReport) BeforeCreate(db *gorm.DB) error {
	if pavlovUUID, err := r.GetUUID(db); err == nil {
		r.UUID = pavlovUUID.UUID
		return nil
	} else {
		return err
	}
}

func (r *LootReport) Save(db *gorm.DB) error {
	return db.Save(r).Error
}
