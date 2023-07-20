package gameresource

import (
	"encoding/json"
	"github.com/St-Pavlov-Data-Department/backend/utils"
	"os"
)

type JumpInfoRaw struct {
	JumpInfo
	ParamRaw string `json:"param"`
}

type JumpInfo struct {
	Id     int64      `json:"id"`
	Name   string     `json:"name"`
	OpenId int        `json:"openId"`
	Param  *Reference `json:"-"`
}

type Jump struct {
	jumps map[int64]*JumpInfo
}

func (j *Jump) LoadFromJson(jsonPath string) error {
	jsonContent, err := os.ReadFile(jsonPath)
	if err != nil {
		return err
	}

	jumpListRaw := new([]*JumpInfoRaw)
	if err := json.Unmarshal(jsonContent, jumpListRaw); err != nil {
		return err
	}

	if j.jumps == nil {
		j.jumps = map[int64]*JumpInfo{}
	}
	for _, jump := range *jumpListRaw {
		jump.Param = j.parseReference(jump.ParamRaw)

		j.jumps[jump.Id] = &jump.JumpInfo
	}

	return nil
}

func (j *Jump) Contains(key int64) bool {
	_, ok := j.jumps[key]
	return ok
}

func (j *Jump) Get(key int64) (val *JumpInfo, ok bool) {
	val, ok = j.jumps[key]
	return val, ok
}

func (j *Jump) parseReference(refStr string) *Reference {
	tableMap := map[int64]Table{
		1:   STORE_ENTRANCE,
		3:   CHAPTER,
		4:   EPISODE,
		100: ACTIVITY,
	}

	chunk := utils.RefToInt64Arr(refStr)
	switch len(chunk) {
	case 1:
		// pure ID, local jummp
		return &Reference{
			Table:   JUMP,
			ID:      chunk[0],
			Special: 0,
		}
	case 2:
		// TABLE # ID
		return &Reference{
			Table:   tableMap[chunk[0]],
			ID:      chunk[1],
			Special: 0,
		}
	case 3:
		// TABLE # ID # Special
		return &Reference{
			Table:   tableMap[chunk[0]],
			ID:      chunk[1],
			Special: chunk[2],
		}

	default:
		return nil
	}

}
