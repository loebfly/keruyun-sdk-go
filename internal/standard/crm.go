package standard

import (
	"github.com/loebfly/keruyun-sdk-go/internal/network"
	"github.com/loebfly/keruyun-sdk-go/kry_model"
	"github.com/loebfly/keruyun-sdk-go/kry_standard"
)

func (s standardApi) CrmCreate(req kry_standard.CrmCreateReq) kry_model.Result[kry_standard.CrmCreateResp] {
	return network.JsonToResult[kry_standard.CrmCreateResp](s.newPostJsonOptions(UriCrmCreate, req))
}

func (s standardApi) CrmUpdate(req kry_standard.CrmUpdateReq) kry_model.Result[kry_standard.CrmUpdateResp] {
	return network.JsonToResult[kry_standard.CrmUpdateResp](s.newPostJsonOptions(UriCrmUpdate, req))
}

func (s standardApi) CrmQueryByMobile(req kry_standard.CrmQueryByMobileReq) kry_model.Result[kry_standard.CrmCustomerInfoResp] {
	return network.JsonToResult[kry_standard.CrmCustomerInfoResp](s.newPostJsonOptions(UriCrmQueryByMobile, req))
}

func (s standardApi) CrmQueryByIds(req kry_standard.CrmQueryByIdsReq) kry_model.Result[kry_standard.CrmCustomerIdsResp] {
	return network.JsonToResult[kry_standard.CrmCustomerIdsResp](s.newPostJsonOptions(UriCrmQueryByIds, req))
}

func (s standardApi) CrmQueryProperty(req kry_standard.CrmQueryPropertyReq) kry_model.Result[kry_standard.CrmCustomerPropertyResp] {
	return network.JsonToResult[kry_standard.CrmCustomerPropertyResp](s.newPostJsonOptions(UriCrmQueryProperty, req))
}

func (s standardApi) CrmDirectCharge(req kry_standard.CrmDirectChargeReq) kry_model.Result[kry_standard.CrmDirectChargeResp] {
	return network.JsonToResult[kry_standard.CrmDirectChargeResp](s.newPostJsonOptions(UriCrmDirectCharge, req))
}

func (s standardApi) CrmTemplateList(req kry_standard.CrmTemplateListReq) kry_model.Result[kry_standard.CrmTemplateListResp] {
	return network.JsonToResult[kry_standard.CrmTemplateListResp](s.newPostJsonOptions(UriCrmTemplateList, req))
}

func (s standardApi) CrmTemplateInfo(req kry_standard.CrmTemplateInfoReq) kry_model.Result[kry_standard.CrmTemplateInfoResp] {
	return network.JsonToResult[kry_standard.CrmTemplateInfoResp](s.newPostJsonOptions(UriCrmTemplateInfo, req))
}

func (s standardApi) CrmTemplateShopList(req kry_standard.CrmTemplateShopReq) kry_model.Result[kry_standard.CrmTemplateShopResp] {
	return network.JsonToResult[kry_standard.CrmTemplateShopResp](s.newPostJsonOptions(UriCrmTemplateShopList, req))
}

func (s standardApi) CrmTemplateDishList(req kry_standard.CrmTemplateDishReq) kry_model.Result[kry_standard.CrmTemplateDishResp] {
	return network.JsonToResult[kry_standard.CrmTemplateDishResp](s.newPostJsonOptions(UriCrmTemplateDishList, req))
}

func (s standardApi) CrmCouponSend(req kry_standard.CrmCouponSendReq) kry_model.Result[kry_standard.CrmCouponSendResp] {
	return network.JsonToResult[kry_standard.CrmCouponSendResp](s.newPostJsonOptions(UriCrmCouponSend, req))
}

func (s standardApi) CrmCouponInvalid(req kry_standard.CrmCouponInvalidReq) kry_model.Result[kry_standard.CrmCouponInvalidResp] {
	return network.JsonToResult[kry_standard.CrmCouponInvalidResp](s.newPostJsonOptions(UriCrmCouponInvalid, req))
}

func (s standardApi) CrmCouponQueryList(req kry_standard.CrmCouponQueryReq) kry_model.Result[kry_standard.CrmCouponQueryResp] {
	return network.JsonToResult[kry_standard.CrmCouponQueryResp](s.newPostJsonOptions(UriCrmCouponQueryList, req))
}

func (s standardApi) CrmCouponQueryInfo(req kry_standard.CrmCouponInfoReq) kry_model.Result[kry_standard.CrmCouponInfoResp] {
	return network.JsonToResult[kry_standard.CrmCouponInfoResp](s.newPostJsonOptions(UriCrmCouponQueryInfo, req))
}