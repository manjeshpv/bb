package chainmodel

type Chain struct {
	Id        int64  `db:"id, primarykey, autoincrement" json:"id"`
	Name string `db:"name" json:"name"`
}

type Chains []Chain
