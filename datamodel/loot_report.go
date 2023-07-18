package datamodel

import (
	"github.com/St-Pavlov-Data-Department/backend/requests"
	"github.com/St-Pavlov-Data-Department/backend/utils"
	"gorm.io/gorm"
)

type LootReport struct {
	IDModelIntrinsic

	StageID string `gorm:"column:stage_id;index:idx_server_stage,priority:2" json:"stage_id"`
	Server  string `gorm:"column:server;index:idx_server_stage,priority:1" json:"server"`
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

// --------

type LootReportList []*LootReport

func (l *LootReportList) LoadByRequest(db *gorm.DB, req *requests.QueryReportRequest) error {

	if err := utils.WithTransaction(db,
		func(tx *gorm.DB) error {

			// filter server
			if req.Server != "" {
				db = db.Where("server = ?", req.Server)
			}

			// filter stages
			if len(req.Stages) > 0 {
				db = db.Where("stage_id in (?)", req.Stages)
			}

			if err := db.Model(&LootReport{}).Find(l).Error; err != nil {
				return err
			}

			for _, r := range *l {
				lootItemList := &LootItemList{}
				if err := lootItemList.LoadByReportUUID(tx, r.UUID); err != nil {
					return err
				}

				r.Loot = *lootItemList
			}

			return nil
		},
	); err != nil {
		return err
	}

	return nil
}
