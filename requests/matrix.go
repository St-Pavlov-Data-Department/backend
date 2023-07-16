package requests

type MatrixRequest struct {
	Stages           string `json:"stages"`
	Items            string `json:"items"`
	Server           string `json:"server"`
	ShowClosedStages bool   `json:"show_closed_stages"`
	PersonalData     string `json:"personal_data"`
}
