package standard

import (
	"github.com/loebfly/keruyun-sdk-go/internal/config"
	"github.com/loebfly/keruyun-sdk-go/internal/network"
	"github.com/loebfly/keruyun-sdk-go/kry_model"
	"github.com/loebfly/keruyun-sdk-go/kry_standard"
)

func (s standardApi) ShopQueryStoreDetail() kry_model.Result[kry_standard.ShopQueryStoreDetailResp] {
	return network.JsonToResult[kry_standard.ShopQueryStoreDetailResp](network.JsonOptions{
		Host:   config.Global.Domain,
		Uri:    UriShopQueryStoreDetail,
		Method: network.POST,
		Header: nil,
		Query:  s.NewPostQuery(),
		JSON:   nil,
	})
}

func (s standardApi) ShopQueryTableList(req kry_standard.ShopQueryTableListReq) kry_model.Result[kry_standard.ShopQueryTableListResp] {
	return network.JsonToResult[kry_standard.ShopQueryTableListResp](network.JsonOptions{
		Host:   config.Global.Domain,
		Uri:    UriShopQueryTableList,
		Method: network.POST,
		Header: nil,
		Query:  s.NewPostQuery(),
		JSON:   req,
	})
}

func (s standardApi) ShopQueryServiceFeeList(req kry_standard.ShopQueryServiceFeeListReq) kry_model.Result[kry_standard.ShopQueryServiceFeeListResp] {
	return network.JsonToResult[kry_standard.ShopQueryServiceFeeListResp](network.JsonOptions{
		Host:   config.Global.Domain,
		Uri:    UriShopQueryServiceFeeList,
		Method: network.POST,
		Header: nil,
		Query:  s.NewPostQuery(),
		JSON:   req,
	})
}
