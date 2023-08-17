package kry_standard

type ShopQueryTableListReq struct {
	ShopAreaIds []string `json:"shopAreaIds"` // 区域ID列表，为空表示查询全部区域
	PageSize    int      `json:"pageSize"`    // 每页条目数，最大100条
	PageNum     int      `json:"pageNum"`     // 当前分页页数，从1开始
}

type ShopQueryServiceFeeListReq struct {
	TableAreaId string `json:"tableAreaId"` // 桌台所属区域ID
}

type ShopQueryTableInfoReq struct {
	TableId string `json:"tableId"` // 桌台ID
}

type ShopQueryServiceFeeInfoReq struct {
	ExtraChargeIds []string `json:"extraChargeIds"` // 服务费ID列表，最多100条
}

type DishQueryPageReq struct {
	CategoryIds []string `json:"categoryIds"` // 分类ID
	DishTypes   []string `json:"dishTypes"`   // 菜品类型， SINGLE：单菜 ，COMBO：套餐， SIDE：配料
	DishIds     []string `json:"dishIds"`     // 菜品ID
	PageIndex   int64    `json:"pageIndex"`   // 页码
	PageSize    int64    `json:"pageSize"`    // 一页数据数，最大100
}

type DishQueryDetailReq struct {
	DishIds []int64 `json:"dishIds"` // 菜品ID，上限5个
}

type DishQueryCategoryReq struct {
	CategoryTypes []string `json:"categoryTypes"` // 分类类型，DISH：菜品分类 SIDE_DISH_GROUP:加料分组
}
