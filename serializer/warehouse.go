package serializer

type Warehouse struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Createtime string `json:"createtime"`
	Remark     string `json:"remark"`
}
