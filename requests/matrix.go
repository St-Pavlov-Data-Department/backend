package requests

type MatrixRequest struct {
	Episodes           []int64 `json:"episodes"`
	Items              []int64 `json:"items"`
	Server             string  `json:"server"`
	ShowClosedEpisodes bool    `json:"show_closed_episodes"`
	PersonalData       string  `json:"personal_data"`
}
