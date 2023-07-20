package gameresource

import (
	"encoding/json"
	"github.com/St-Pavlov-Data-Department/backend/utils"
	"os"
	"strings"
)

type ItemInfoRaw struct {
	ItemInfo
	SourcesRaw string `json:"sources"`
}

type ItemInfo struct {
	Id          int64        `json:"id"`
	Name        string       `json:"name"`
	UseDesc     string       `json:"useDesc"`
	Desc        string       `json:"desc"`
	SubType     int          `json:"subType"`
	Icon        string       `json:"icon"`
	Rare        int          `json:"rare"`
	HighQuality int          `json:"highQuality"`
	IsStackable int          `json:"isStackable"`
	IsShow      int          `json:"isShow"`
	IsTimeShow  int          `json:"isTimeShow"`
	Effect      string       `json:"effect"`
	Cd          int          `json:"cd"`
	ExpireTime  string       `json:"expireTime"`
	Price       string       `json:"price"`
	Sources     []*Reference `json:"-"`
	BoxOpen     string       `json:"boxOpen"`
}

type Item struct {
	items map[int64]*ItemInfo
}

func (i *Item) LoadFromJson(jsonPath string) error {
	jsonContent, err := os.ReadFile(jsonPath)
	if err != nil {
		return err
	}

	itemListRaw := new([]*ItemInfoRaw)
	if err := json.Unmarshal(jsonContent, itemListRaw); err != nil {
		return err
	}

	if i.items == nil {
		i.items = map[int64]*ItemInfo{}
	}
	for _, item := range *itemListRaw {
		for _, refStr := range strings.Split(item.SourcesRaw, "|") {
			item.Sources = append(item.Sources, i.parseReference(refStr))
		}

		i.items[item.Id] = &item.ItemInfo
	}

	return nil
}

func (i *Item) Contains(key int64) bool {
	_, ok := i.items[key]
	return ok
}

func (i *Item) Get(key int64) (val *ItemInfo, ok bool) {
	val, ok = i.items[key]
	return val, ok
}

func (i *Item) parseReference(refStr string) *Reference {
	chunk := utils.RefToInt64Arr(refStr)
	switch len(chunk) {
	case 1:
		// JUMP_ID
		return &Reference{
			Table:   JUMP,
			ID:      chunk[0],
			Special: 0,
		}
	case 2:
		// JUMP_ID # SPECIAL
		return &Reference{
			Table:   JUMP,
			ID:      chunk[0],
			Special: chunk[1],
		}

	default:
		return nil
	}
}
