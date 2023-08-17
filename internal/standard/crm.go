package standard

import (
	"github.com/loebfly/keruyun-sdk-go/internal/network"
	"github.com/loebfly/keruyun-sdk-go/kry_model"
	"github.com/loebfly/keruyun-sdk-go/kry_standard"
)

func (s standardApi) CrmCreate(req kry_standard.CrmCreateReq) kry_model.Result[kry_standard.CrmCreateResp] {
	return network.JsonToResult[kry_standard.CrmCreateResp](s.newPostJsonOptions(UriCrmCreate, req))
}

func (s standardApi) CrmQueryByMobile(req kry_standard.CrmQueryByMobileReq) kry_model.Result[kry_standard.CrmCustomerInfoResp] {
	return network.JsonToResult[kry_standard.CrmCustomerInfoResp](s.newPostJsonOptions(UriCrmQueryByMobile, req))
}

func (s standardApi) CrmQueryProperty(req kry_standard.CrmQueryPropertyReq) kry_model.Result[kry_standard.CrmCustomerPropertyResp] {
	return network.JsonToResult[kry_standard.CrmCustomerPropertyResp](s.newPostJsonOptions(UriCrmQueryProperty, req))
}

func (s standardApi) CrmDirectCharge(req kry_standard.CrmDirectChargeReq) kry_model.Result[kry_standard.CrmDirectChargeResp] {
	return network.JsonToResult[kry_standard.CrmDirectChargeResp](s.newPostJsonOptions(UriCrmDirectCharge, req))
}
