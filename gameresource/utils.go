package gameresource

type Table uint64

const (
	JUMP Table = iota
	ITEM
	EPISODE
	CHAPTER
	ACTIVITY
	STORE_ENTRANCE
)

type Reference struct {
	Table   Table
	ID      int64
	Special int64
}
