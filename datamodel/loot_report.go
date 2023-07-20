package datamodel

import (
	"github.com/St-Pavlov-Data-Department/backend/requests"
	"github.com/St-Pavlov-Data-Department/backend/utils"
	"gorm.io/gorm"
)

type LootReport struct {
	IDModelIntrinsic

	EpisodeID int64  `gorm:"column:episode_id;index:idx_server_episode,priority:2" json:"episode_id"`
	Server    string `gorm:"column:server;index:idx_server_episode,priority:1" json:"server"`
	Source    string `gorm:"column:source" json:"source"`
	Version   string `gorm:"column:version" json:"version"`

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

	if err := l.rawLoadByRequest(db, req); err != nil {
		return err
	}

	// filter items
	queryItems := utils.NewSet(req.Items...)
	for _, r := range *l {
		lootItemList := &LootItemList{}
		if err := lootItemList.LoadByReportUUID(db, r.UUID); err != nil {
			return err
		}

		if queryItems.IsEmpty() {
			// without filter
			r.Loot = *lootItemList
		} else {
			// only load required items
			for _, lootItem := range *lootItemList {
				if queryItems.Contains(lootItem.ItemID) {
					r.Loot = append(r.Loot, lootItem)
				}
			}
		}
	}

	return nil
}

func (l *LootReportList) rawLoadByRequest(db *gorm.DB, req *requests.QueryReportRequest) error {
	// filter server
	if req.Server != "" {
		db = db.Where("server = ?", req.Server)
	}

	// filter episodes
	if len(req.Episodes) > 0 {
		db = db.Where("episode_id in (?)", req.Episodes)
	}

	if err := db.Model(&LootReport{}).Find(l).Error; err != nil {
		return err
	}

	return nil
}
