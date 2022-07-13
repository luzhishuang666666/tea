package model

type Warehouse struct {
	Id          int    `json:"id" gorm:"column:id"`
	Name        string `json:"name" gorm:"column:name"`
	Expectation string `json:"expectation,omitempty" gorm:"column:expectation"`
	Actual      string `json:"actual,omitempty" gorm:"column:actual"`
	Createtime  string `json:"createtime" gorm:"column:createtime"`
	Updatetime  string `json:"updatetime,omitempty" gorm:"column:updatetime"`
	Remark      string `json:"remark" gorm:"column:remark"`
}
