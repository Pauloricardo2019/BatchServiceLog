package model

type Log struct {
	ID   uint   `gorm:"primaryKey"`
	IP   string `gorm:"column:ip; type:varchar(100)"`
	Date string `gorm:"column:date; type:varchar(30)"`
	Verb string `gorm:"verb; type:varchar(10)"`
}
