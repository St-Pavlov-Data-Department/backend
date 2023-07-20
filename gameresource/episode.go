package gameresource

import (
	"encoding/json"
	"github.com/St-Pavlov-Data-Department/backend/utils"
	"os"
	"strings"
)

// TODO: refactor: load EpisodeInfoRaw first
//	and parse into EpisodeInfo

type EpisodeInfoRaw struct {
	EpisodeInfo
	RewardListRaw string `json:"rewardList"`
}

type EpisodeInfo struct {
	Id                int64        `json:"id"`
	ChapterId         int          `json:"chapterId"`
	Type              int          `json:"type"`
	Name              string       `json:"name"`
	NameEn            string       `json:"name_En"`
	Desc              string       `json:"desc"`
	BattleDesc        string       `json:"battleDesc"`
	BeforeStory       int          `json:"beforeStory"`
	Story             string       `json:"story"`
	AfterStory        int          `json:"afterStory"`
	AutoSkipStory     int          `json:"autoSkipStory"`
	DecryptId         int          `json:"decryptId"`
	MapId             int          `json:"mapId"`
	Pic               string       `json:"pic"`
	Icon              string       `json:"icon"`
	PreEpisode        int          `json:"preEpisode"`
	UnlockEpisode     int          `json:"unlockEpisode"`
	ElementList       string       `json:"elementList"`
	Cost              string       `json:"cost"`
	FailCost          string       `json:"failCost"`
	FirstBattleId     int          `json:"firstBattleId"`
	BattleId          int          `json:"battleId"`
	DayChangeBonus    string       `json:"dayChangeBonus"`
	FirstBonus        int          `json:"firstBonus"`
	Bonus             int          `json:"bonus"`
	AdvancedBonus     int          `json:"advancedBonus"`
	RewardDisplayList string       `json:"rewardDisplayList"`
	RewardList        []*Reference `json:"-"`
	Bgmevent          int          `json:"bgmevent"`
	Navigationpic     int          `json:"navigationpic"`
	Year              int          `json:"year"`
	CanUseRecord      int          `json:"canUseRecord"`
	Time              string       `json:"time"`
	FreeBonus         int          `json:"freeBonus"`
	FreeDisplayList   string       `json:"freeDisplayList"`
	DayNum            int          `json:"dayNum"`
	SaveDayNum        int          `json:"saveDayNum"`
}

type Episode struct {
	episodes map[int64]*EpisodeInfo
}

func (e *Episode) LoadFromJson(jsonPath string) error {
	jsonContent, err := os.ReadFile(jsonPath)
	if err != nil {
		return err
	}

	jsonEpisodeList := new([]*EpisodeInfoRaw)
	if err := json.Unmarshal(jsonContent, jsonEpisodeList); err != nil {
		return err
	}

	if e.episodes == nil {
		e.episodes = map[int64]*EpisodeInfo{}
	}
	for _, episode := range *jsonEpisodeList {
		episode.RewardList = utils.Apply(
			strings.Split(episode.RewardListRaw, "|"),
			e.parseReference,
		)
		/*
			for _, refStr := range strings.Split(episode.RewardListRaw, "|") {
				episode.RewardList = append(episode.RewardList, e.parseReference(refStr))
			}
		*/

		e.episodes[episode.Id] = &episode.EpisodeInfo
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

func (e *Episode) parseReference(refStr string) *Reference {
	tableMap := map[int64]Table{
		1: ITEM,
		2: CURRENCY,
		9: EQUIP,
	}

	chunk := utils.RefToInt64Arr(refStr)
	switch len(chunk) {
	case 3:
		// if relativeTableID not known, return nil
		if _, ok := tableMap[chunk[0]]; !ok {
			return nil
		}
		// TABLE # ID # Special
		return &Reference{
			Table:   tableMap[chunk[0]],
			ID:      chunk[1],
			Special: chunk[2],
			Count:   1,
		}
	case 4:
		// if relativeTableID not known, return nil
		if _, ok := tableMap[chunk[0]]; !ok {
			return nil
		}
		// TABLE # ID # Special # Count
		return &Reference{
			Table:   tableMap[chunk[0]],
			ID:      chunk[1],
			Special: chunk[2],
			Count:   chunk[3],
		}

	default:
		return nil
	}

}
