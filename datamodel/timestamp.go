package datamodel

type TimestampModel struct {
	CreatedTSInt int64 `json:"created_ts_int" gorm:"column:created_ts_int;autoCreateTime:milli"`
	UpdatedTSInt int64 `json:"updated_ts_int" gorm:"column:updated_ts_int;autoUpdateTime:milli"`
}
