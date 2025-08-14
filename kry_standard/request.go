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
	Gender   int    `json:"gender"`             // 性别,0女1男2其它
	Mobile   string `json:"mobile"`             // 手机号,11位
	Name     string `json:"name"`               // 姓名
	ShopId   string `json:"shopId"`             // 入会门店id
	Birthday string `json:"birthday,omitempty"` // 生日日期,yyyy-MM-dd格式
}

type CrmUpdateReq struct {
	Birthday     int64  `json:"birthday"`
	Address      string `json:"address"`
	Gender       int    `json:"gender"`
	Remark       string `json:"remark"`
	CustomerType int    `json:"customerType"` // 顾客类型
	CustomerId   string `json:"customerId"`
	Name         string `json:"name"`
	Invoice      string `json:"invoice"`
	Email        string `json:"email"`
}

type CrmQueryByMobileReq struct {
	Mobile string `json:"mobile"`
}
type CrmQueryByIdsReq struct {
	CustomerIds []int `json:"customerIds"`
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

type CrmTemplateListReq struct {
	TypeList   []string `json:"typeList"`
	StatusList []string `json:"statusList"`
	PageNo     string   `json:"pageNo"`
	PageSize   string   `json:"pageSize"`
}

var TypeList = map[int]string{
	1: "CASH",               // 代金券
	2: "DISCOUNT",           // 折扣券
	3: "GIFT",               // 礼品券
	4: "CONDITION_DISCOUNT", // 第N份折扣券
	5: "BOUGHT_GIFT",        // 买A赠B券
}

var StatusList = map[int]string{
	1: "UNUSED",       // 未使用
	2: "USED",         // 使用中
	3: "NO_INVENTORY", // 无库存
	4: "INVALID",      // 已失效
	5: "AUDITING",     // 待审核
	6: "REJECTED",     //  驳回
}

var MemberCouponStatusList = map[int]string{
	1: "NORMAL",  // 未使用
	2: "FREEZED", // 冻结中
	3: "INVALID", // 已失效
	4: "EXPIRED", // 已过期
	5: "ISUSED",  // 已使用
}

type CrmTemplateInfoReq struct {
	TemplateId  string `json:"templateId"`
	NeedDeleted bool   `json:"needDeleted"`
}

type CrmTemplateShopReq struct {
	PageNo     string `json:"pageNo"`
	PageSize   string `json:"pageSize"`
	TemplateId string `json:"templateId"`
}

type CrmTemplateDishReq struct {
	PageNo     string `json:"pageNo"`
	PageSize   string `json:"pageSize"`
	TemplateId string `json:"templateId"`
}

type CrmCouponSendReq struct {
	ActivityId       string `json:"activityId"`
	OutPreInstanceId string `json:"outPreInstanceId"`
	CustomerId       string `json:"customerId"`
	TemplateId       string `json:"templateId"`
}

type CrmCouponInvalidReq struct {
	ResourceCodes []string `json:"resourceCodes"`
}

type CrmCouponQueryReq struct {
	PageNo     string   `json:"pageNo"`
	CustomerId string   `json:"customerId"`
	PageSize   string   `json:"pageSize"`
	StatusList []string `json:"statusList"`
}

type CrmCouponInfoReq struct {
	ResourceCode string `json:"resourceCode"`
	CustomerId   string `json:"customerId"`
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
	OrgId          string                      `json:"orgId"`          // 库存组织id
	DepotId        string                      `json:"depotId"`        // 入库仓库id
	BussDate       string                      `json:"bussDate"`       // 业务日期
	BusinessType   string                      `json:"businessType"`   // 业务类型(固定传入)：0001
	Remarks        string                      `json:"remarks"`        // 备注
	CreateUserId   string                      `json:"createUserId"`   // 创建人id(如果没有，默认传system)
	CreateUserName string                      `json:"createUserName"` // 创建人名称
	DetailList     []StockOtherInReceiveDetail `json:"detailList"`     // 入库物品明细
}

type StockOtherInReceiveDetail struct {
	GoodsCode    string `json:"goodsCode"`    // 物品编码
	GoodsBatchNo string `json:"goodsBatchNo"` // 实物批号
	ProDate      string `json:"proDate"`      // 生产日期
	ExpDate      string `json:"expDate"`      // 有效期至
	GoodsQty     int    `json:"goodsQty"`     // 标准单位入库数量
	DualGoodsQty int    `json:"dualGoodsQty"` // 辅助单位入库数量
	UnitPrice    int    `json:"unitPrice"`    // 成本单价
	GoodsAmt     int    `json:"goodsAmt"`     // 成本金额
	Remarks      string `json:"remarks"`      // 备注
}

type StockOtherOutReceiveReq struct {
	OrgId          string                       `json:"orgId"`          // 库存组织id
	DepotId        string                       `json:"depotId"`        // 入库仓库id
	BussDate       string                       `json:"bussDate"`       // 业务日期
	BusinessType   string                       `json:"businessType"`   // 业务类型(固定传入)：0002
	Remarks        string                       `json:"remarks"`        // 备注
	CreateUserId   string                       `json:"createUserId"`   // 创建人id(如果没有，默认传system)
	CreateUserName string                       `json:"createUserName"` // 创建人名称
	DetailList     []StockOtherOutReceiveDetail `json:"detailList"`     // 入库物品明细
}

type StockOtherOutReceiveDetail struct {
	GoodsCode    string `json:"goodsCode"`    // 物品编码
	GoodsBatchNo string `json:"goodsBatchNo"` // 实物批号
	ProDate      string `json:"proDate"`      // 生产日期
	ExpDate      string `json:"expDate"`      // 有效期至
	GoodsQty     int    `json:"goodsQty"`     // 标准单位入库数量
	DualGoodsQty int    `json:"dualGoodsQty"` // 辅助单位入库数量
	Remarks      string `json:"remarks"`      // 备注
}

type StockPurchaseInReceiveReq struct {
	PurchaseDate string                         `json:"purchaseDate"` // 采购日期
	ArrivalDate  string                         `json:"arrivalDate"`  // 到货日期
	ThirdBillNo  string                         `json:"thirdBillNo"`  // 第三方单号
	Remarks      string                         `json:"remarks"`      // 备注
	OrgId        string                         `json:"orgId"`        // 门店机构id
	DepotId      string                         `json:"depotId"`      // 仓库id
	SupplyId     string                         `json:"supplyId"`     // 供应商id
	Details      []StockPurchaseInReceiveDetail `json:"details"`      // 入库物品明细
	UserName     string                         `json:"userName"`     // 操作人
}

type StockPurchaseInReceiveDetail struct {
	GoodsId  string  `json:"goodsId"`  // 物品id
	OrdNum   float64 `json:"ordNum"`   // 采购单位数量
	GoodsAmt int     `json:"goodsAmt"` // 物品金额（含税）
	DualNum  int     `json:"dualNum"`  // 辅助单位数量
	TaxRatio string  `json:"taxRatio"` // 税率
	Remarks  string  `json:"remarks"`  // 物品备注
}

type StockPurchaseBackReceiveReq struct {
	OrgId         string                              `json:"orgId"`         // 门店机构id
	SupplyId      string                              `json:"supplyId"`      // 供应商id
	DepotId       string                              `json:"depotId"`       // 仓库id
	BussDate      string                              `json:"bussDate"`      // 业务日期
	Remarks       string                              `json:"remarks"`       // 备注
	UserName      string                              `json:"userName"`      // 操作人
	DetailDTOList []StockPurchaseBackReceiveDetailDTO `json:"detailDTOList"` // 待退货明细
}

type StockPurchaseBackReceiveDetailDTO struct {
	GoodsId  string `json:"goodsId"`  // 物品id
	UnitNum  string `json:"unitNum"`  // 标准单位数量
	DualNum  string `json:"dualNum"`  // 辅助单位数量
	TaxRatio string `json:"taxRatio"` // 税率
	GoodsAmt string `json:"goodsAmt"` // 物品金额（含税）
	Remarks  string `json:"remarks"`  // 备注
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

type BookQueryReq struct {
	BookOrderNo string `json:"bookOrderNo"`
}

type BookSaveReq struct {
	OrderNo                string `json:"orderNo"`
	DinnerTime             string `json:"dinnerTime"`
	TableNum               int    `json:"tableNum"`
	DinersNum              int    `json:"dinersNum"`
	CustomerName           string `json:"customerName"`
	CustomerSex            string `json:"customerSex"`
	CustomerPhone          string `json:"customerPhone"`
	BusinessType           string `json:"businessType"`
	DeviceType             string `json:"deviceType"`
	DeviceId               string `json:"deviceId"`
	SkipConflictFlag       bool   `json:"skipConflictFlag"`
	TableBookRecordDtoList []struct {
		AreaId    string `json:"areaId"`
		AreaName  string `json:"areaName"`
		TableId   string `json:"tableId"`
		TableName string `json:"tableName"`
	} `json:"tableBookRecordDtoList"`
	Note                  string       `json:"note"`
	OutBizNo              string       `json:"outBizNo"`
	SendMessageFlag       bool         `json:"sendMessageFlag"`
	BookingTimeType       string       `json:"bookingTimeType"`
	PeriodTimeList        []PeriodTime `json:"periodTimeList"`
	OperatorUserId        string       `json:"operatorUserId"`
	OperatorUserName      string       `json:"operatorUserName"`
	BookingTypeName       string       `json:"bookingTypeName"`
	BookingTitle          string       `json:"bookingTitle"`
	ArrivalTime           string       `json:"arrivalTime"`
	ColdTableCount        int          `json:"coldTableCount"`
	EmptyTableCount       int          `json:"emptyTableCount"`
	FullTableCount        int          `json:"fullTableCount"`
	DiningStandard        *string      `json:"diningStandard"`
	CommissionStaffId     string       `json:"commissionStaffId"`
	CommissionStaffName   string       `json:"commissionStaffName"`
	CommissionStaffSource string       `json:"commissionStaffSource"`
	CommissionStaffPhone  string       `json:"commissionStaffPhone"`
	AddItemList           []struct {
		DishType            string            `json:"dishType"`
		SkuId               string            `json:"skuId"`
		SellPrice           float64           `json:"sellPrice"`
		DishNum             float64           `json:"dishNum"`
		UnitId              string            `json:"unitId"`
		UnitName            string            `json:"unitName"`
		WeighFlag           string            `json:"weighFlag"`
		CategoryId          string            `json:"categoryId"`
		CookingWayGroupList []CookingWayGroup `json:"cookingWayGroupList"`
		ComboGroupList      []struct {
			ComboGroupId         string `json:"comboGroupId"`
			ComboGroupDetailList []struct {
				CookingWayGroupList []CookingWayGroup `json:"cookingWayGroupList"`
				SideDishGroupList   []struct {
					SideDishGroupId         string `json:"sideDishGroupId"`
					SideDishGroupDetailList []struct {
						SideDishId  int `json:"sideDishId"`
						SideDishNum int `json:"sideDishNum"`
					} `json:"sideDishGroupDetailList"`
				} `json:"sideDishGroupList"`
				SingleDishId   string  `json:"singleDishId"`
				DishNum        int     `json:"dishNum"`
				WeighFlag      string  `json:"weighFlag"`
				DishName       string  `json:"dishName"`
				SellPrice      float64 `json:"sellPrice"`
				UnitId         string  `json:"unitId"`
				UnitName       string  `json:"unitName"`
				ItemCategoryId string  `json:"itemCategoryId"`
			} `json:"comboGroupDetailList"`
		} `json:"comboGroupList"`
		SideDishGroupList []struct {
			SideDishGroupId         string `json:"sideDishGroupId"`
			SideDishGroupDetailList []struct {
				SideDishId  int `json:"sideDishId"`
				SideDishNum int `json:"sideDishNum"`
			} `json:"sideDishGroupDetailList"`
		} `json:"sideDishGroupList"`
		CartItemNoteList []CartItemNote `json:"cartItemNoteList"`
	} `json:"addItemList"`
}

// PeriodTime 结构体
type PeriodTime struct {
	PeriodId   string `json:"periodId"`
	PeriodName string `json:"periodName"`
	StartTime  string `json:"startTime"`
	EndTime    string `json:"endTime"`
}

// CookingWayGroup 结构体
type CookingWayGroup struct {
	CookingWayGroupId string `json:"cookingWayGroupId"`
	CookingWayList    []struct {
		CookingWayId string `json:"cookingWayId"`
	} `json:"cookingWayList"`
}

// CartItemNote 结构体
type CartItemNote struct {
}

type BookQueryPeriodTimeReq struct {
	BusinessDate string `json:"businessDate"` // 格式  2025-12-12
}

type BookQueryOrderReq struct {
	BookOrderNo      string `json:"bookOrderNo"`
	OperatorUserId   string `json:"operatorUserId"`
	OperatorUserName string `json:"operatorUserName"`
}

type BookQueryTableInfoReq struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
}

type BookConfirmReq struct {
	OrderNo          string `json:"OrderNo"`
	OperatorUserId   string `json:"operatorUserId"`
	OperatorUserName string `json:"operatorUserName"`
}

type BookCancelReq struct {
	OrderNo                      string `json:"orderNo"`
	ShopBookOrderCancelReasonDto struct {
		ReasonContent string `json:"reasonContent"`
	} `json:"shopBookOrderCancelReasonDto"`
	SendMessageFlag  bool   `json:"sendMessageFlag"`
	OperatorUserId   string `json:"operatorUserId"`
	OperatorUserName string `json:"operatorUserName"`
}

type BusinessIncomeReq struct {
	StatisticsByShop      bool     `json:"statisticsByShop"`
	OrgStatisticsType     string   `json:"orgStatisticsType"`
	CouponStatisticalType string   `json:"couponStatisticalType"`
	StoreStatisticalType  string   `json:"storeStatisticalType"`
	PeriodType            string   `json:"periodType"`
	OrderSourceList       []string `json:"orderSourceList"`
	OrderTypeList         []string `json:"orderTypeList"`
	DateRange             struct {
		DateType  string `json:"dateType"`
		StartDate int64  `json:"startDate"`
		EndDate   int64  `json:"endDate"`
	} `json:"dateRange"`
	PageBean struct {
		PageNum  int `json:"pageNum"`
		PageSize int `json:"pageSize"`
	} `json:"pageBean"`
	ShopIds   []int  `json:"shopIds"`
	RequestId string `json:"requestId"`
}

type BusinessIncomePromoStatisticsReq struct {
	DateRange struct {
		DateType  string `json:"dateType"`
		StartDate int64  `json:"startDate"`
		EndDate   int64  `json:"endDate"`
	} `json:"dateRange"`
	OrderSources []string `json:"orderSources"`
	OrderTypes   []string `json:"orderTypes"`
	PromoTypes   []string `json:"promoTypes"`
	PromoName    string   `json:"promoName"`
	PageBean     struct {
		PageNum  int `json:"pageNum"`
		PageSize int `json:"pageSize"`
	} `json:"pageBean"`
	ShopIds []int `json:"shopIds"`
}

type KposlocalReq struct {
	RequestId     string         `json:"requestId"`
	RequestHeader map[string]any `json:"requestHeader"`
	RequestBody   struct {
		OrderId        string `json:"orderId"`
		BusinessFormat string `json:"businessFormat"`
		SceneCode      string `json:"sceneCode"`
		CartId         string `json:"cartId"`
	} `json:"requestBody"`
}

type KposlocalAddReq struct {
	RequestId     string         `json:"requestId"`
	RequestHeader map[string]any `json:"requestHeader"`
	RequestBody   struct {
		OrderId               string           `json:"orderId"`
		BusinessFormat        string           `json:"businessFormat"`
		SceneCode             string           `json:"sceneCode"`
		GoodList              []AddItemRequest `json:"goodList"`
		OperateItemSourceType string           `json:"operateItemSourceType,omitempty"`
	} `json:"requestBody"`
}

type AddItemRequest struct {
	SceneCode               string                `json:"sceneCode,omitempty"`
	SkuId                   string                `json:"skuId"`
	SpuId                   string                `json:"SpuId"`
	TemporaryDishFlag       string                `json:"temporaryDishFlag,omitempty"`
	MustOrderFlag           string                `json:"mustOrderFlag,omitempty"`
	ItemName                string                `json:"itemName,omitempty"`
	ItemType                string                `json:"itemType,omitempty"`
	ItemExtType             string                `json:"itemExtType,omitempty"`
	TicketExitIdList        []string              `json:"ticketExitIdList,omitempty"`
	ItemPriceStr            string                `json:"itemPriceStr,omitempty"`
	ItemNumStr              float64               `json:"itemNumStr,omitempty"`
	WeighDishFlag           string                `json:"weighDishFlag,omitempty"`
	UnitId                  string                `json:"unitId,omitempty"`
	UnitName                string                `json:"unitName,omitempty"`
	SecondUnitId            string                `json:"secondUnitId,omitempty"`
	SecondUnitName          string                `json:"secondUnitName,omitempty"`
	Weigh                   float64               `json:"weigh,omitempty"`
	RequisiteCookMethodFlag string                `json:"requisiteCookMethodFlag,omitempty"`
	RequisiteSideDishFlag   string                `json:"requisiteSideDishFlag,omitempty"`
	CookAttachGroupList     []CookAttachGroup     `json:"cookAttachGroupList,omitempty"`
	AddSpiceAttachGroupList []AddSpiceAttachGroup `json:"addSpiceAttachGroupList,omitempty"`
	ComboGroupList          []ComboGroup          `json:"comboGroupList,omitempty"`
	OrderUnitId             string                `json:"orderUnitId,omitempty"`
	OrderUnitName           string                `json:"orderUnitName,omitempty"`
	OrderUnitNum            float64               `json:"orderUnitNum,omitempty"`
	DoubleUnitWeighDishFlag string                `json:"doubleUnitWeighDishFlag,omitempty"`
	CartItemNoteList        []string              `json:"cartItemNoteList,omitempty"`
	CartItemId              string                `json:"cartItemId,omitempty"`
	ChangePriceFlag         string                `json:"changePriceFlag,omitempty"`
	WaitCallFlag            string                `json:"waitCallFlag,omitempty"`
	PriceType               string                `json:"priceType,omitempty"`
	CookbookId              string                `json:"cookbookId,omitempty"`
	CookbookType            string                `json:"cookbookType,omitempty"`
	MealNum                 float64               `json:"mealNum,omitempty"`
	SinglePackFlag          string                `json:"singlePackFlag,omitempty"`
	ItemNoCard              string                `json:"itemNoCard,omitempty"`
}

type CookAttachGroup struct {
	CookAttachGroupId string       `json:"cookAttachGroupId"`
	CookAttachIdList  []CookAttach `json:"cookAttachIdList"`
}

type CookAttach struct {
	CookAttachId  string `json:"cookAttachId"`
	TemporaryFlag string `json:"temporaryFlag"`
}

type AddSpiceAttachGroup struct {
	AddSpiceGroupId    string           `json:"addSpiceGroupId"`
	AddSpiceAttachList []AddSpiceAttach `json:"addSpiceAttachList"`
}
type AddSpiceAttach struct {
	ChildSkuId  string `json:"childSkuId"`
	ChildNumStr string `json:"childNumStr"`
}

type ComboGroup struct {
	ComboGroupId string       `json:"comboGroupId"`
	ChildSkuList []ComboChild `json:"childSkuList"`
}

type ComboChild struct {
	CartItemId                   string                `json:"cartItemId"`
	ChildWeighFlag               string                `json:"childWeighFlag"`
	ChildWeighItemPrice          string                `json:"childWeighItemPrice,omitempty"`
	RequisiteCookMethodFlag      string                `json:"requisiteCookMethodFlag"`
	ChildSkuId                   string                `json:"childSkuId"`
	ChildNumStr                  float64               `json:"childNumStr"`
	CartItemNoteList             []string              `json:"cartItemNoteList"`
	CookAttachGroupList          []CookAttachGroup     `json:"cookAttachGroupList"`
	AddSpiceAttachGroups         []AddSpiceAttachGroup `json:"addSpiceAttachGroups"`
	AlternativeSkuId             string                `josn:"alternativeSkuId,omitempty"`
	ChildOrderUnitNumStr         float64               `json:"childOrderUnitNumStr,omitempty"`
	ChildDoubleUnitWeighDishFlag string                `json:"childDoubleUnitWeighDishFlag"`
	RequisiteSideDishFlag        string                `json:"requisiteSideDishFlag"`
	SinglePackFlag               string                `json:"singlePackFlag"`
	WaitCallFlag                 string                `json:"waitCallFlag"`
	TemporaryDishFlag            string                `json:"temporaryDishFlag,omitempty"`
	ItemName                     string                `json:"itemName,omitempty"`
	TicketExitIdList             []string              `json:"ticketExitIdList"`
	ItemPriceStr                 string                `json:"itemPriceStr,omitempty"`
	ItemNumStr                   float64               `json:"itemNumStr,omitempty"`
	WeighDishFlag                string                `json:"weighDishFlag,omitempty"`
	UnitId                       string                `json:"unitId,omitempty"`
	UnitName                     string                `json:"unitName,omitempty"`
	SecondUnitId                 string                `json:"secondUnitId,omitempty"`
	SecondUnitName               string                `json:"secondUnitName,omitempty"`
	Weigh                        float64               `json:"weigh,omitempty"`
}

type KposTableStatusQueryReq struct {
	PullPosFlag bool   `json:"pullPosFlag"`
	TableId     string `json:"tableId"`
}
