package requests

type UploadReportRequest struct {
	EpisodeID int64  `json:"episode_id"`
	Server    string `json:"server"`
	Source    string `json:"source"`
	Version   string `json:"version"`

	Loot        []*LootItem `json:"loot"`
	ReplayLevel int64       `json:"replay_level"`

	ClientIP string `json:"-"`
}

type LootItem struct {
	ItemID   int64  `json:"item_id"`
	LootType string `json:"loot_type"`
	Quantity int64  `json:"quantity"`
}

// --------

// QueryReportRequest request for querying reports
type QueryReportRequest struct {
	Server   string
	Episodes []int64
	Items    []int64
}
