package datamodel

import (
	"github.com/St-Pavlov-Data-Department/backend/requests"
	"gorm.io/gorm"
)

type Matrix struct {
	IDModelIntrinsic

	StageId        string `gorm:"column:stage_id;index:uk_stage_id_item_id,unique,priority:1" json:"stage_id"`
	ItemId         string `gorm:"column:item_id;index:uk_stage_id_item_id,unique,priority:2" json:"item_id"`
	StartTimeMilli int64  `gorm:"column:start_time_milli;index:idx_start_time_milli" json:"start_time_milli"`
	EndTimeMilli   int64  `gorm:"column:end_time_milli;index:idx_end_time_milli" json:"end_time_milli"`
	Quantity       int    `gorm:"column:quantity" json:"quantity"`
	ReplayCount    int    `gorm:"column:replay_count" json:"replay_count"`

	TimestampModel
}

func (m *Matrix) TableName() string {
	return "matrix_tab"
}

func (m *Matrix) Save(db *gorm.DB) error {
	return db.Save(m).Error
}

// --------

type MatrixList []*Matrix

func (m *MatrixList) LoadByRequest(db *gorm.DB, req *requests.MatrixRequest) error {
	// filter stages
	if req.Stages != "" {
		db = db.Where("stage_id in ? ", req.Stages)
	}

	// filter items
	if req.Items != "" {
		db = db.Where("item_id in ?", req.Items)
	}

	// filter server
	if req.Server != "" {
		db = db.Where("server = ?", req.Server)
	}

	// TODO: show_closed_stages and personal_data parameters are not supported yet.

	if err := db.Model(&Matrix{}).Find(m).Error; err != nil {
		return err
	}

	return nil
}
