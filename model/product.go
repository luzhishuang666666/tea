package model

type Product struct {
	Id     int    `json:"id" gorm:"column:id"`
	Name   string `json:"name" gorm:"column:name"`
	Number string `json:"number" gorm:"column:number"`
	Unit   string `json:"unit" gorm:"column:unit"`
	Remark string `json:"remark" gorm:"column:remark"`
}
