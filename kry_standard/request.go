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

type CrmCreateReq struct {
	Gender   int    `json:"gender"`   // 性别,0女1男2其它
	Mobile   string `json:"mobile"`   // 手机号,11位
	Name     string `json:"name"`     // 姓名
	ShopId   string `json:"shopId"`   // 入会门店id
	Birthday string `json:"birthday"` // 生日日期,yyyy-MM-dd格式
}

type CrmQueryByMobileReq struct {
	Mobile string `json:"mobile"`
}

type CrmQueryPropertyReq struct {
	CustomerId string `json:"customerId"` // 用户id
	ShopId     string `json:"shopId"`     // 客如云门店id
}

type CrmDirectChargeReq struct {
	BizDate          string `json:"bizDate"`          // 充值时间，格式：yyyy-MM-dd HH:mm:ss
	ShopId           string `json:"shopId"`           // 门店id
	OperatorName     string `json:"operatorName"`     // 操作人名称（不含特殊字符，长度不超过64位）
	SecondBizChannel string `json:"secondBizChannel"` // 用户自定义二级充值渠道
	RealValue        int    `json:"realValue"`        // 充值金额，单位：分
	CustomerId       string `json:"customerId"`       // 用户id
	BizId            string `json:"bizId"`            // 充值业务id，订单号（不含特殊字符，长度不超过64位）
	OperatorId       string `json:"operatorId"`       // 操作人id（数字字符串，长度不超过64位）
}
