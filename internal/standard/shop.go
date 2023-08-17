package standard

import (
	"github.com/loebfly/keruyun-sdk-go/internal/network"
	"github.com/loebfly/keruyun-sdk-go/kry_model"
	"github.com/loebfly/keruyun-sdk-go/kry_standard"
)

func (s standardApi) ShopQueryStoreDetail() kry_model.Result[kry_standard.ShopStoreDetailResp] {
	return network.JsonToResult[kry_standard.ShopStoreDetailResp](s.newPostJsonOptions(UriShopQueryStoreDetail, nil))
}

func (s standardApi) ShopQueryTableList(req kry_standard.ShopQueryTableListReq) kry_model.Result[kry_standard.ShopTableListResp] {
	return network.JsonToResult[kry_standard.ShopTableListResp](s.newPostJsonOptions(UriShopQueryTableList, req))
}

func (s standardApi) ShopQueryServiceFeeList(req kry_standard.ShopQueryServiceFeeListReq) kry_model.Result[kry_standard.ShopServiceFeeListResp] {
	return network.JsonToResult[kry_standard.ShopServiceFeeListResp](s.newPostJsonOptions(UriShopQueryServiceFeeList, req))
}

func (s standardApi) ShopQueryTableInfo(req kry_standard.ShopQueryTableInfoReq) kry_model.Result[kry_standard.ShopTableInfoResp] {
	return network.JsonToResult[kry_standard.ShopTableInfoResp](s.newPostJsonOptions(UriShopQueryTableInfo, req))
}

func (s standardApi) ShopQueryServiceFeeInfo(req kry_standard.ShopQueryServiceFeeInfoReq) kry_model.Result[kry_standard.ShopServiceFeeListResp] {
	return network.JsonToResult[kry_standard.ShopServiceFeeListResp](s.newPostJsonOptions(UriShopQueryServiceFeeInfo, req))
}
