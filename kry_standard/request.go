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
