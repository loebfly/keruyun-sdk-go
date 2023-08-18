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

type OrderQueryDetailReq struct {
	OrderId string `json:"orderId"` // 	订单ID
}

type OrderQueryListReq struct {
	DateType        string   `json:"dateType"`        // 查询时间类型{1:营业日期(FINISH_BUSI_DATE) 2:下单时间(OPEN_TIME) 3:完结时间(FINISH_TIME) 4:结账时间(SETTLE_TIME)}
	StartDate       string   `json:"startDate"`       // 查询开始时间（限制查询时间范围不超过31天） 若查询时间类型为营业日期，则精确到天，示例：2021-06-02 其它查询时间类型精确到秒，示例：2021-06-02 04:30:00
	EndDate         string   `json:"endDate"`         // 查询结束时间（限制查询时间范围不超过31天） 若查询时间类型为营业日期，则精确到天，示例：2021-06-02 其它查询时间类型精确到秒，示例：2021-06-02 04:30:00
	OrderTypeList   []string `json:"orderTypeList"`   // 订单类型{1:堂食(FOR_HERE) 2:平台外卖(PLATFORM_TAKE_OUT) 3:自营外卖(SELF_TAKE_OUT)4:会员充值(MEMBER_STORE) 5:无单收银(NO_ORDER_CASHIER)6:销账订单(REPAYMENT_ORDER)}
	OrderStatusList []string `json:"orderStatusList"` // 订单状态{1:待处理(WAIT_PROCESSED) 2:待结账(WAIT_SETTLED) 3:已结账(SETTLED) 4:已退单(REFUND)5:已作废(INVALID)6:已取消(CANCELLED) 7:已拒绝(REJECTED) 8:已完成(SUCCESS)}
	PageBean        struct {
		PageNum  int `json:"pageNum"`  // 页码
		PageSize int `json:"pageSize"` // 页大小
	} `json:"pageBean"` // 分页
}

type StockInOutQueryListReq struct {
	StartDate       string `json:"startDate"`       // 业务日期-开始（不能与更新时间同时为空）
	EndDate         string `json:"endDate"`         // 业务日期-结束
	UpdateTimeStart string `json:"updateTimeStart"` // 更新时间-开始（不能与业务日期同时为空）
	UpdateTimeEnd   string `json:"updateTimeEnd"`   // 更新时间-结束
	OrgId           string `json:"orgId"`           // 机构id
	DepotId         string `json:"depotId"`         // 仓库id
	Status          int    `json:"status"`          // 状态：961：暂存，963：提交，962：已审核，964：已驳回，965：已反审，966：已作废
	Type            int    `json:"type"`            // 类型 1：配送差异报损出库 2：手动添加的报损出库
	OrgType         int    `json:"orgType"`         // 机构类型 1：门店 2：公司
	PageNum         int    `json:"pageNum"`         // 页码（默认1）
	PageSize        int    `json:"pageSize"`        // 页大小（默认20，最大为100）
}

type StockLossOutQueryDetailReq struct {
	Type int      `json:"type"` // 类型 1：配送差异报损出库 2：手动添加报损出库
	Ids  []string `json:"ids"`  // 单据id列表
}

type StockOtherInOutQueryDetailReq struct {
	Ids []string `json:"ids"` // 单据id列表
}

type StockTransferInternalQueryListReq struct {
	StartDate       string `json:"startDate"`       // 业务日期-开始（不能与更新时间同时为空）
	EndDate         string `json:"endDate"`         // 业务日期-结束
	UpdateTimeStart string `json:"updateTimeStart"` // 更新时间-开始（不能与业务日期同时为空）
	UpdateTimeEnd   string `json:"updateTimeEnd"`   // 更新时间-结束
	StoreId         string `json:"storeId"`         // 机构id
	OutDepotId      string `json:"outDepotId"`      // 出库仓库id
	InDepotId       string `json:"inDepotId"`       // 入库仓库id
	Status          int    `json:"status"`          // 状态961：暂存，963：提交，962：已审核，964：已驳回，965：已反审，966：已作废
	OrgType         int    `json:"orgType"`         // 机构类型 1：门店 2：公司
	PageNum         int    `json:"pageNum"`         // 页码(默认1)
	PageSize        int    `json:"pageSize"`        // 页大小（默认20，最大为100）
}

type StockTransferInternalQueryDetailReq struct {
	Ids []string `json:"ids"`
}

type StockCheckQueryListReq struct {
	StartDate       string `json:"startDate"`       // 盘点业务日期-开始（与最后更新时间不能同时为空）
	EndDate         string `json:"endDate"`         // 盘点业务日期-开始
	UpdateTimeStart string `json:"updateTimeStart"` // 最后更新时间-开始（与盘点业务日期不能同时为空）
	UpdateTimeEnd   string `json:"updateTimeEnd"`   // 最后更新时间-结束
	StoreId         string `json:"storeId"`         // 机构id
	DepotId         string `json:"depotId"`         // 仓库id
	Status          string `json:"status"`          // 状态 961：暂存，963：提交，962：已审核，964：已驳回，965：已反审
	OrgType         int    `json:"orgType"`         // 机构类型 1：门店 2：公司
	PageNum         int    `json:"pageNum"`         // 页码（默认1）
	PageSize        int    `json:"pageSize"`        // 页大小（默认20，最大不能超过100）
}

type StockCheckQueryDetailReq struct {
	Id string `json:"id"` // 单据id
}

type StockTransferInOutQueryDetailReq struct {
	Ids []string `json:"ids"` // 单据id（一次性最多100个）
}

type StockTransferInOutStoreToStoreQueryListReq struct {
	StoreIds        []string `json:"storeIds"`        // 出库组织id
	StartDate       string   `json:"startDate"`       // 业务日期-开始（不能与开始更新时间同时为空）
	EndDate         string   `json:"endDate"`         // 业务日期-结束
	UpdateTimeStart string   `json:"updateTimeStart"` // 更新时间-开始（不能与开始业务日期同时为空）
	UpdateTimeEnd   string   `json:"updateTimeEnd"`   // 更新时间-结束
	StatusList      int      `json:"statusList"`      // 状态：962：已审核
	PageNum         int      `json:"pageNum"`         // 页码（默认1）
	PageSize        int      `json:"pageSize"`        // 页大小（默认20，最大100）
}

type StockSaleOutQueryDetailReq struct {
	Ids []string `json:"ids"` // 单据id列表
}

type StockSaleInOutQueryListReq struct {
	UpdateTimeStart string   `json:"updateTimeStart"` // 更新时间-开始（最多查询最近90天数据，最大间隔为7天）
	UpdateTimeEnd   string   `json:"updateTimeEnd"`   // 更新时间-结束
	OrgIds          []string `json:"orgIds"`          // 门店id列表
	DepotIds        []string `json:"depotIds"`        // 仓库id列表
	PageNum         int      `json:"pageNum"`         // 页码（默认1）
	PageSize        int      `json:"pageSize"`        // 页大小（默认20，最大支持100）
}

type StockOtherInReceiveReq struct {
	OrgId          string `json:"orgId"`          // 库存组织id
	DepotId        string `json:"depotId"`        // 入库仓库id
	BussDate       string `json:"bussDate"`       // 业务日期
	BusinessType   string `json:"businessType"`   // 业务类型(固定传入)：0001
	Remarks        string `json:"remarks"`        // 备注
	CreateUserId   string `json:"createUserId"`   // 创建人id(如果没有，默认传system)
	CreateUserName string `json:"createUserName"` // 创建人名称
	DetailList     []struct {
		GoodsCode    string `json:"goodsCode"`    // 物品编码
		GoodsBatchNo string `json:"goodsBatchNo"` // 实物批号
		ProDate      string `json:"proDate"`      // 生产日期
		ExpDate      string `json:"expDate"`      // 有效期至
		GoodsQty     int    `json:"goodsQty"`     // 标准单位入库数量
		DualGoodsQty int    `json:"dualGoodsQty"` // 辅助单位入库数量
		UnitPrice    int    `json:"unitPrice"`    // 成本单价
		GoodsAmt     int    `json:"goodsAmt"`     // 成本金额
		Remarks      string `json:"remarks"`      // 备注
	} `json:"detailList"` // 入库物品明细
}

type StockOtherOutReceiveReq struct {
	OrgId          string `json:"orgId"`          // 库存组织id
	DepotId        string `json:"depotId"`        // 入库仓库id
	BussDate       string `json:"bussDate"`       // 业务日期
	BusinessType   string `json:"businessType"`   // 业务类型(固定传入)：0002
	Remarks        string `json:"remarks"`        // 备注
	CreateUserId   string `json:"createUserId"`   // 创建人id(如果没有，默认传system)
	CreateUserName string `json:"createUserName"` // 创建人名称
	DetailList     []struct {
		GoodsCode    string `json:"goodsCode"`    // 物品编码
		GoodsBatchNo string `json:"goodsBatchNo"` // 实物批号
		ProDate      string `json:"proDate"`      // 生产日期
		ExpDate      string `json:"expDate"`      // 有效期至
		GoodsQty     int    `json:"goodsQty"`     // 标准单位入库数量
		DualGoodsQty int    `json:"dualGoodsQty"` // 辅助单位入库数量
		Remarks      string `json:"remarks"`      // 备注
	} `json:"detailList"` // 入库物品明细
}

type StockPurchaseInReceiveReq struct {
	PurchaseDate string `json:"purchaseDate"` // 采购日期
	ArrivalDate  string `json:"arrivalDate"`  // 到货日期
	ThirdBillNo  string `json:"thirdBillNo"`  // 第三方单号
	Remarks      string `json:"remarks"`      // 备注
	OrgId        string `json:"orgId"`        // 门店机构id
	DepotId      string `json:"depotId"`      // 仓库id
	SupplyId     string `json:"supplyId"`     // 供应商id
	Details      []struct {
		GoodsId  string  `json:"goodsId"`  // 物品id
		OrdNum   float64 `json:"ordNum"`   // 采购单位数量
		GoodsAmt int     `json:"goodsAmt"` // 物品金额（含税）
		DualNum  int     `json:"dualNum"`  // 辅助单位数量
		TaxRatio string  `json:"taxRatio"` // 税率
		Remarks  string  `json:"remarks"`  // 物品备注
	} `json:"details"` // 入库物品明细
	UserName string `json:"userName"` // 操作人
}

type StockPurchaseBackReceiveReq struct {
	OrgId         string `json:"orgId"`    // 门店机构id
	SupplyId      string `json:"supplyId"` // 供应商id
	DepotId       string `json:"depotId"`  // 仓库id
	BussDate      string `json:"bussDate"` // 业务日期
	Remarks       string `json:"remarks"`  // 备注
	UserName      string `json:"userName"` // 操作人
	DetailDTOList []struct {
		GoodsId  string `json:"goodsId"`  // 物品id
		UnitNum  string `json:"unitNum"`  // 标准单位数量
		DualNum  string `json:"dualNum"`  // 辅助单位数量
		TaxRatio string `json:"taxRatio"` // 税率
		GoodsAmt string `json:"goodsAmt"` // 物品金额（含税）
		Remarks  string `json:"remarks"`  // 备注
	} `json:"detailDTOList"` // 待退货明细
}

type StockQueryReq struct {
	OrgId      string   `json:"orgId"`      //组织ID
	DepotId    string   `json:"depotId"`    //仓库ID
	GoodsCodes []string `json:"goodsCodes"` //物品编码
	PageNum    int      `json:"pageNum"`    //页码（默认第一页）
	PageSize   int      `json:"pageSize"`   //每页条数（默认20,最大值100）
}

type OrgQueryReq struct {
	OrgType int `json:"orgType"` // 3门店，7 配送中心，10仓库
}
