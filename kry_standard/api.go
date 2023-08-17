package kry_standard

import "github.com/loebfly/keruyun-sdk-go/kry_model"

type Api interface {
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
}
