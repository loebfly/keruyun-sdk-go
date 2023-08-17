package standard

import (
	"github.com/loebfly/keruyun-sdk-go/internal/network"
	"github.com/loebfly/keruyun-sdk-go/kry_model"
	"github.com/loebfly/keruyun-sdk-go/kry_standard"
)

func (s standardApi) OrderQueryDetail(req kry_standard.OrderQueryDetailReq) kry_model.Result[kry_standard.OrderDetailResp] {
	return network.JsonToResult[kry_standard.OrderDetailResp](s.newPostJsonOptions(UriOrderQueryDetail, req))
}

func (s standardApi) OrderQueryList(req kry_standard.OrderQueryListReq) kry_model.Result[kry_standard.OrderListResp] {
	return network.JsonToResult[kry_standard.OrderListResp](s.newPostJsonOptions(UriOrderQueryList, req))
}
