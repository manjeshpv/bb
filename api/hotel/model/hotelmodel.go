package hotelmodel

type Hotel struct {
	Id        int64  `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	ChainId  int64 `db:"chain_id" json:"chain_id"`
}

type Hotels []Hotel
