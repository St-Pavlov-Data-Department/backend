package gameresource

import (
	"encoding/json"
	"os"
)

type EpisodeInfo struct {
	Id                int64  `json:"id"`
	ChapterId         int    `json:"chapterId"`
	Type              int    `json:"type"`
	Name              string `json:"name"`
	NameEn            string `json:"name_En"`
	Desc              string `json:"desc"`
	BattleDesc        string `json:"battleDesc"`
	BeforeStory       int    `json:"beforeStory"`
	Story             string `json:"story"`
	AfterStory        int    `json:"afterStory"`
	AutoSkipStory     int    `json:"autoSkipStory"`
	DecryptId         int    `json:"decryptId"`
	MapId             int    `json:"mapId"`
	Pic               string `json:"pic"`
	Icon              string `json:"icon"`
	PreEpisode        int    `json:"preEpisode"`
	UnlockEpisode     int    `json:"unlockEpisode"`
	ElementList       string `json:"elementList"`
	Cost              string `json:"cost"`
	FailCost          string `json:"failCost"`
	FirstBattleId     int    `json:"firstBattleId"`
	BattleId          int    `json:"battleId"`
	DayChangeBonus    string `json:"dayChangeBonus"`
	FirstBonus        int    `json:"firstBonus"`
	Bonus             int    `json:"bonus"`
	AdvancedBonus     int    `json:"advancedBonus"`
	RewardDisplayList string `json:"rewardDisplayList"`
	RewardList        string `json:"rewardList"`
	Bgmevent          int    `json:"bgmevent"`
	Navigationpic     int    `json:"navigationpic"`
	Year              int    `json:"year"`
	CanUseRecord      int    `json:"canUseRecord"`
	Time              string `json:"time"`
	FreeBonus         int    `json:"freeBonus"`
	FreeDisplayList   string `json:"freeDisplayList"`
	DayNum            int    `json:"dayNum"`
	SaveDayNum        int    `json:"saveDayNum"`
}

type Episode struct {
	episodes map[int64]*EpisodeInfo
}

func (e *Episode) LoadFromJson(jsonPath string) error {
	jsonContent, err := os.ReadFile(jsonPath)
	if err != nil {
		return err
	}

	jsonEpisodeList := new([]*EpisodeInfo)
	if err := json.Unmarshal(jsonContent, jsonEpisodeList); err != nil {
		return err
	}

	if e.episodes == nil {
		e.episodes = map[int64]*EpisodeInfo{}
	}
	for _, episode := range *jsonEpisodeList {
		e.episodes[episode.Id] = episode
	}

	return nil
}

func (e *Episode) Contains(key int64) bool {
	_, ok := e.episodes[key]
	return ok
}

func (e *Episode) Get(key int64) (val *EpisodeInfo, ok bool) {
	val, ok = e.episodes[key]
	return val, ok
}
