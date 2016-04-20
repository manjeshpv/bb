package departmentmodel

import "time"

type Department struct {
	Id        int64 `gorm:"primary_key" db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	CreatedAt time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `db:"updatedAt" json:"updatedAt"`
}

type Users []Department
