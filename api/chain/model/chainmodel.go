package chainmodel

type Chain struct {
	Id        int64  `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	HotelId int64 `db:"hotel_id" json:"hotel_id"`
}

type Chains []Chain
