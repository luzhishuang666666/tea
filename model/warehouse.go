package model

import (
	"encoding/json"
)

type Warehouse struct {
	Id          int    `json:"id" gorm:"column:id"`
	Name        string `json:"name" gorm:"column:name"`
	Expectation string `json:"expectation,omitempty" gorm:"column:expectation"`
	Actual      string `json:"actual,omitempty" gorm:"column:actual"`
	Createtime  string `json:"createtime" gorm:"column:createtime"`
	Updatetime  string `json:"updatetime,omitempty" gorm:"column:updatetime"`
	Remark      string `json:"remark" gorm:"column:remark"`
}

type WarehouseExpAndAct []struct {
	MaterialID     int    `json:"materialId"`
	MaterialName   string `json:"materialName"`
	MaterialNumber string `json:"materialNumber"`
	Unit           string `json:"unit"`
	Amount         int    `json:"amount"`
}

type WarehouseDetail struct {
	ID          int                `json:"id"`
	Name        string             `json:"name"`
	CreateTime  string             `json:"createTime"`
	UpdateTime  string             `json:"updateTime"`
	Remark      string             `json:"remark"`
	Expectation WarehouseExpAndAct `json:"expectation"`
	Actual      WarehouseExpAndAct `json:"actual"`
}

func Warehouse2WarehouseDetail(warehouse Warehouse) (WarehouseDetail, error) {
	var warehouseDetail WarehouseDetail

	var ExpModel = WarehouseExpAndAct{}
	err := json.Unmarshal([]byte(warehouse.Expectation), &ExpModel)
	if err != nil {
		return warehouseDetail, err
	}

	var ActModel = WarehouseExpAndAct{}
	err = json.Unmarshal([]byte(warehouse.Actual), &ActModel)
	if err != nil {
		return warehouseDetail, err
	}

	warehouseDetail.ID = warehouse.Id
	warehouseDetail.Name = warehouse.Name
	warehouseDetail.CreateTime = warehouse.Createtime
	warehouseDetail.UpdateTime = warehouse.Updatetime
	warehouseDetail.Remark = warehouse.Remark
	warehouseDetail.Expectation = ExpModel
	warehouseDetail.Actual = ActModel

	return warehouseDetail, nil
}
