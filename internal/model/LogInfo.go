package model

type Log struct {
	ID   uint `gorm:primaryKey`
	IP   string
	Date string
	Verb string
}
