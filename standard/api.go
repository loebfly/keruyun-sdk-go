package standard

import "github.com/loebfly/keruyun-sdk-go/kry_result"

type Api interface {
	// ShopQueryStoreDetail 根据门店id获取门店详情
	ShopQueryStoreDetail() kry_result.Result[ShopQueryStoreDetailResp]
	// ShopQueryTableList 获取门店桌台列表
	ShopQueryTableList(req ShopQueryTableListReq) kry_result.Result[ShopQueryTableListResp]
	// ShopQueryServiceFeeList 查询门店服务费列表
	ShopQueryServiceFeeList(req ShopQueryServiceFeeListReq) kry_result.Result[ShopQueryServiceFeeListResp]
}
