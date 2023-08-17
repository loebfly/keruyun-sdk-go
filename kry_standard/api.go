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
}
