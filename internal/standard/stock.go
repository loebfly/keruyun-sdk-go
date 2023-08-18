package standard

import (
	"github.com/loebfly/keruyun-sdk-go/internal/network"
	"github.com/loebfly/keruyun-sdk-go/kry_model"
	"github.com/loebfly/keruyun-sdk-go/kry_standard"
)

func (s standardApi) StockLossOutQueryList(req kry_standard.StockInOutQueryListReq) kry_model.Result[kry_standard.StockInOutListResp] {
	return network.JsonToResult[kry_standard.StockInOutListResp](s.newPostJsonOptions(UriStockLossOutQueryList, req))
}

func (s standardApi) StockLossOutQueryDetail(req kry_standard.StockLossOutQueryDetailReq) kry_model.Result[kry_standard.StockInOutDetailResp] {
	return network.JsonToResult[kry_standard.StockInOutDetailResp](s.newPostJsonOptions(UriStockLossOutQueryDetail, req))
}

func (s standardApi) StockOtherInQueryDetail(req kry_standard.StockOtherInOutQueryDetailReq) kry_model.Result[kry_standard.StockInOutDetailResp] {
	return network.JsonToResult[kry_standard.StockInOutDetailResp](s.newPostJsonOptions(UriStockOtherInQueryDetail, req))
}

func (s standardApi) StockOtherInQueryList(req kry_standard.StockInOutQueryListReq) kry_model.Result[kry_standard.StockInOutListResp] {
	return network.JsonToResult[kry_standard.StockInOutListResp](s.newPostJsonOptions(UriStockOtherInQueryList, req))
}

func (s standardApi) StockOtherOutQueryList(req kry_standard.StockInOutQueryListReq) kry_model.Result[kry_standard.StockInOutListResp] {
	return network.JsonToResult[kry_standard.StockInOutListResp](s.newPostJsonOptions(UriStockOtherOutQueryList, req))
}

func (s standardApi) StockOtherOutQueryDetail(req kry_standard.StockOtherInOutQueryDetailReq) kry_model.Result[kry_standard.StockInOutDetailResp] {
	return network.JsonToResult[kry_standard.StockInOutDetailResp](s.newPostJsonOptions(UriStockOtherOutQueryDetail, req))
}

func (s standardApi) StockTransferInternalQueryList(req kry_standard.StockTransferInternalQueryListReq) kry_model.Result[kry_standard.StockTransferInternalListResp] {
	return network.JsonToResult[kry_standard.StockTransferInternalListResp](s.newPostJsonOptions(UriStockTransferInternalQueryList, req))
}

func (s standardApi) StockTransferInternalQueryDetail(req kry_standard.StockTransferInternalQueryDetailReq) kry_model.Result[kry_standard.StockTransferInOutDetailResp] {
	return network.JsonToResult[kry_standard.StockTransferInOutDetailResp](s.newPostJsonOptions(UriStockTransferInternalQueryDetail, req))
}

func (s standardApi) StockCheckQueryList(req kry_standard.StockCheckQueryListReq) kry_model.Result[kry_standard.StockCheckListResp] {
	return network.JsonToResult[kry_standard.StockCheckListResp](s.newPostJsonOptions(UriStockCheckQueryList, req))
}

func (s standardApi) StockCheckQueryDetail(req kry_standard.StockCheckQueryDetailReq) kry_model.Result[kry_standard.StockCheckDetailResp] {
	return network.JsonToResult[kry_standard.StockCheckDetailResp](s.newPostJsonOptions(UriStockCheckQueryDetail, req))
}

func (s standardApi) StockTransferOutQueryDetail(req kry_standard.StockTransferInOutQueryDetailReq) kry_model.Result[kry_standard.StockTransferInOutDetailResp] {
	return network.JsonToResult[kry_standard.StockTransferInOutDetailResp](s.newPostJsonOptions(UriStockTransferOutQueryDetail, req))
}

func (s standardApi) StockTransferOutStoreToStoreQueryList(req kry_standard.StockTransferInOutStoreToStoreQueryListReq) kry_model.Result[kry_standard.StockInOutListResp] {
	return network.JsonToResult[kry_standard.StockInOutListResp](s.newPostJsonOptions(UriStockTransferOutStoreToStoreQueryList, req))
}

func (s standardApi) StockTransferInStoreToStoreQueryList(req kry_standard.StockTransferInOutStoreToStoreQueryListReq) kry_model.Result[kry_standard.StockInOutListResp] {
	return network.JsonToResult[kry_standard.StockInOutListResp](s.newPostJsonOptions(UriStockTransferInStoreToStoreQueryList, req))
}

func (s standardApi) StockTransferInQueryDetail(req kry_standard.StockTransferInOutQueryDetailReq) kry_model.Result[kry_standard.StockTransferInOutDetailResp] {
	return network.JsonToResult[kry_standard.StockTransferInOutDetailResp](s.newPostJsonOptions(UriStockTransferInQueryDetail, req))
}

func (s standardApi) StockSaleOutQueryDetail(req kry_standard.StockSaleOutQueryDetailReq) kry_model.Result[kry_standard.StockSaleInOutDetailResp] {
	return network.JsonToResult[kry_standard.StockSaleInOutDetailResp](s.newPostJsonOptions(UriStockSaleOutQueryDetail, req))
}

func (s standardApi) StockSaleOutQueryList(req kry_standard.StockSaleInOutQueryListReq) kry_model.Result[kry_standard.StockSaleOutInOutListResp] {
	return network.JsonToResult[kry_standard.StockSaleOutInOutListResp](s.newPostJsonOptions(UriStockSaleOutQueryList, req))
}

func (s standardApi) StockOtherInReceive(req kry_standard.StockOtherInReceiveReq) kry_model.Result[kry_standard.StockOtherInOutReceiveResp] {
	return network.JsonToResult[kry_standard.StockOtherInOutReceiveResp](s.newPostJsonOptions(UriStockOtherInReceive, req))
}

func (s standardApi) StockOtherOutReceive(req kry_standard.StockOtherOutReceiveReq) kry_model.Result[kry_standard.StockOtherInOutReceiveResp] {
	return network.JsonToResult[kry_standard.StockOtherInOutReceiveResp](s.newPostJsonOptions(UriStockOtherOutReceive, req))
}

func (s standardApi) StockPurchaseInReceive(req kry_standard.StockPurchaseInReceiveReq) kry_model.Result[kry_standard.StockPurchaseInOutReceiveResp] {
	return network.JsonToResult[kry_standard.StockPurchaseInOutReceiveResp](s.newPostJsonOptions(UriStockPurchaseInReceive, req))
}

func (s standardApi) StockPurchaseBackReceive(req kry_standard.StockPurchaseBackReceiveReq) kry_model.Result[kry_standard.StockPurchaseInOutReceiveResp] {
	return network.JsonToResult[kry_standard.StockPurchaseInOutReceiveResp](s.newPostJsonOptions(UriStockPurchaseBackReceive, req))
}

func (s standardApi) StockQuery(req kry_standard.StockQueryReq) kry_model.Result[kry_standard.StockQueryResp] {
	return network.JsonToResult[kry_standard.StockQueryResp](s.newPostJsonOptions(UriStockQuery, req))
}

func (s standardApi) OrgQuery(req kry_standard.OrgQueryReq) kry_model.Result[kry_standard.OrgQueryResp] {
	return network.JsonToResult[kry_standard.OrgQueryResp](s.newPostJsonOptions(UriOrgQuery, req))
}
