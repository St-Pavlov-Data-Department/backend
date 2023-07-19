package gameresource

import (
	"encoding/json"
	"os"
)

type ItemInfo struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	UseDesc     string `json:"useDesc"`
	Desc        string `json:"desc"`
	SubType     int    `json:"subType"`
	Icon        string `json:"icon"`
	Rare        int    `json:"rare"`
	HighQuality int    `json:"highQuality"`
	IsStackable int    `json:"isStackable"`
	IsShow      int    `json:"isShow"`
	IsTimeShow  int    `json:"isTimeShow"`
	Effect      string `json:"effect"`
	Cd          int    `json:"cd"`
	ExpireTime  string `json:"expireTime"`
	Price       string `json:"price"`
	Sources     string `json:"sources"`
	BoxOpen     string `json:"boxOpen"`
}

type Item struct {
	items map[int64]*ItemInfo
}

func (i *Item) LoadFromJson(jsonPath string) error {
	jsonContent, err := os.ReadFile(jsonPath)
	if err != nil {
		return err
	}

	jsonItemList := new([]*ItemInfo)
	if err := json.Unmarshal(jsonContent, jsonItemList); err != nil {
		return err
	}

	if i.items == nil {
		i.items = map[int64]*ItemInfo{}
	}
	for _, item := range *jsonItemList {
		i.items[item.Id] = item
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
