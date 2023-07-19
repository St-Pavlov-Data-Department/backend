package requests

type MatrixRequest struct {
	Stages           []int64 `json:"stages"`
	Items            []int64 `json:"items"`
	Server           string  `json:"server"`
	ShowClosedStages bool    `json:"show_closed_stages"`
	PersonalData     string  `json:"personal_data"`
}
