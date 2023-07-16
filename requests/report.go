package requests

type ReportRequest struct {
	StageID string `json:"stage_id"`
	Server  string `json:"server"`
	Source  string `json:"source"`
	Version string `json:"version"`

	Loot        []*LootItem `json:"loot"`
	ReplayLevel int64       `json:"replay_level"`

	ClientIP string `json:"-"`
}

type LootItem struct {
	ItemID   string `json:"item_id"`
	LootType string `json:"loot_type"`
	Quantity int64  `json:"quantity"`
}
