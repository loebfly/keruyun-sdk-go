package kry_standard

import "github.com/loebfly/keruyun-sdk-go/kry_model"

type Api interface {
	/*
		门店相关接口
	*/
	// ShopQueryStoreDetail 根据门店id获取门店详情
	ShopQueryStoreDetail() kry_model.Result[ShopStoreDetailResp]
	// ShopQueryTableList 获取门店桌台列表
	ShopQueryTableList(req ShopQueryTableListReq) kry_model.Result[ShopTableListResp]
	// ShopQueryServiceFeeList 查询门店服务费列表
	ShopQueryServiceFeeList(req ShopQueryServiceFeeListReq) kry_model.Result[ShopServiceFeeListResp]
	// ShopQueryTableInfo 根据桌台ID获取桌台信息
	ShopQueryTableInfo(req ShopQueryTableInfoReq) kry_model.Result[ShopTableInfoResp]
	// ShopQueryServiceFeeInfo 根据服务费ID获取服务费详情
	ShopQueryServiceFeeInfo(req ShopQueryServiceFeeInfoReq) kry_model.Result[ShopServiceFeeListResp]

	/*
		菜品相关接口
	*/
	// DishQueryPage 门店分页查询菜品基本信息
	DishQueryPage(req DishQueryPageReq) kry_model.Result[DishPageResp]
	// DishQueryDetail 门店查询菜品详情
	DishQueryDetail(req DishQueryDetailReq) kry_model.Result[DishDetailResp]
	// DishQueryCategory 门店查询分类信息
	DishQueryCategory(req DishQueryCategoryReq) kry_model.Result[DishCategoryResp]

	/*
		会员相关
	*/
	// CrmCreate 会员创建
	CrmCreate(req CrmCreateReq) kry_model.Result[CrmCreateResp]
	// CrmQueryByMobile 根据手机号查询会员
	CrmQueryByMobile(req CrmQueryByMobileReq) kry_model.Result[CrmCustomerInfoResp]
	// CrmQueryProperty 会员资产查询
	CrmQueryProperty(req CrmQueryPropertyReq) kry_model.Result[CrmCustomerPropertyResp]
	// CrmDirectCharge 储值充值接口
	CrmDirectCharge(req CrmDirectChargeReq) kry_model.Result[CrmDirectChargeResp]

	/*
		订单相关
	*/
	// OrderQueryDetail 订单详情查询接口
	OrderQueryDetail(req OrderQueryDetailReq) kry_model.Result[OrderDetailResp]
	// OrderQueryList 订单列表查询接口
	OrderQueryList(req OrderQueryListReq) kry_model.Result[OrderListResp]

	/*
		供应链2.0
	*/
	// StockLossOutQueryList 报损出库单列表查询
	StockLossOutQueryList(req StockInOutQueryListReq) kry_model.Result[StockInOutListResp]
	// StockLossOutQueryDetail 报损出库单详情
	StockLossOutQueryDetail(req StockLossOutQueryDetailReq) kry_model.Result[StockInOutDetailResp]
	// StockOtherInQueryDetail 其他入库单详情
	StockOtherInQueryDetail(req StockOtherInOutQueryDetailReq) kry_model.Result[StockInOutDetailResp]
	// StockOtherInQueryList 其他入库单列表查询
	StockOtherInQueryList(req StockInOutQueryListReq) kry_model.Result[StockInOutListResp]
	// StockOtherOutQueryList 其他出库单列表查询
	StockOtherOutQueryList(req StockInOutQueryListReq) kry_model.Result[StockInOutListResp]
	// StockOtherOutQueryDetail 其他出库单详情
	StockOtherOutQueryDetail(req StockOtherInOutQueryDetailReq) kry_model.Result[StockInOutDetailResp]
	// StockTransferInternalQueryList 组织内调拨订单列表查询
	StockTransferInternalQueryList(req StockTransferInternalQueryListReq) kry_model.Result[StockTransferInternalListResp]
	// StockTransferInternalQueryDetail 组织内调拨订单详情
	StockTransferInternalQueryDetail(req StockTransferInternalQueryDetailReq) kry_model.Result[StockTransferInternalDetailResp]
	// StockCheckQueryList 盘点单列表查询
	StockCheckQueryList(req StockCheckQueryListReq) kry_model.Result[StockCheckListResp]
}
