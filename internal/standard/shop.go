package standard

import (
	"github.com/loebfly/keruyun-sdk-go/internal/config"
	"github.com/loebfly/keruyun-sdk-go/internal/network"
	"github.com/loebfly/keruyun-sdk-go/kry_result"
	"github.com/loebfly/keruyun-sdk-go/standard"
)

func (s standardApi) ShopQueryStoreDetail() kry_result.Result[standard.ShopQueryStoreDetailResp] {
	return network.JsonToResult[standard.ShopQueryStoreDetailResp](network.JsonOptions{
		Host:   config.Global.Domain,
		Uri:    UriShopQueryStoreDetail,
		Method: network.POST,
		Header: nil,
		Query:  s.NewPostQuery(),
		JSON:   nil,
	})
}

func (s standardApi) ShopQueryTableList(req standard.ShopQueryTableListReq) kry_result.Result[standard.ShopQueryTableListResp] {
	return network.JsonToResult[standard.ShopQueryTableListResp](network.JsonOptions{
		Host:   config.Global.Domain,
		Uri:    UriShopQueryTableList,
		Method: network.POST,
		Header: nil,
		Query:  s.NewPostQuery(),
		JSON:   req,
	})
}

func (s standardApi) ShopQueryServiceFeeList(req standard.ShopQueryServiceFeeListReq) kry_result.Result[standard.ShopQueryServiceFeeListResp] {
	return network.JsonToResult[standard.ShopQueryServiceFeeListResp](network.JsonOptions{
		Host:   config.Global.Domain,
		Uri:    UriShopQueryServiceFeeList,
		Method: network.POST,
		Header: nil,
		Query:  s.NewPostQuery(),
		JSON:   req,
	})
}
