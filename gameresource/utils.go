package gameresource

type Table uint64

const (
	JUMP Table = iota
	ITEM
	CURRENCY
	EPISODE
	CHAPTER
	ACTIVITY
	EQUIP
	STORE_ENTRANCE
)

type Reference struct {
	Table   Table
	ID      int64
	Special int64
	Count   int64
}
