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

func (s standardApi) StockTransferInternalQueryDetail(req kry_standard.StockTransferInternalQueryDetailReq) kry_model.Result[kry_standard.StockTransferInternalDetailResp] {
	return network.JsonToResult[kry_standard.StockTransferInternalDetailResp](s.newPostJsonOptions(UriStockTransferInternalQueryDetail, req))
}

func (s standardApi) StockCheckQueryList(req kry_standard.StockCheckQueryListReq) kry_model.Result[kry_standard.StockCheckListResp] {
	return network.JsonToResult[kry_standard.StockCheckListResp](s.newPostJsonOptions(UriStockCheckQueryList, req))
}
