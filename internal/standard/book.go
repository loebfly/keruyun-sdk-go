package standard

import (
	"github.com/loebfly/keruyun-sdk-go/internal/network"
	"github.com/loebfly/keruyun-sdk-go/kry_model"
	"github.com/loebfly/keruyun-sdk-go/kry_standard"
)

func (s standardApi) BookSave(req kry_standard.BookSaveReq) kry_model.Result[kry_standard.BookSaveResp] {
	return network.JsonToResult[kry_standard.BookSaveResp](s.newPostJsonOptions(UriBookSave, req))
}

func (s standardApi) BookQueryShopBookPeriodTime(req kry_standard.BookQueryPeriodTimeReq) kry_model.Result[kry_standard.BookQueryPeriodTimeResp] {
	return network.JsonToResult[kry_standard.BookQueryPeriodTimeResp](s.newPostJsonOptions(UriBookQueryShopBookPeriodTime, req))
}

func (s standardApi) BookQueryOrderInfo(req kry_standard.BookQueryOrderReq) kry_model.Result[kry_standard.BookQueryOrderResp] {
	return network.JsonToResult[kry_standard.BookQueryOrderResp](s.newPostJsonOptions(UriBookQuery, req))
}

func (s standardApi) BookQueryTableInfo(req kry_standard.BookQueryTableInfoReq) kry_model.Result[kry_standard.BookTbaleInfoResp] {
	return network.JsonToResult[kry_standard.BookTbaleInfoResp](s.newPostJsonOptions(UriBookQueryBookTableInfo, req))
}

func (s standardApi) BookConfirm(req kry_standard.BookConfirmReq) kry_model.Result[kry_standard.BookConfirmResp] {
	return network.JsonToResult[kry_standard.BookConfirmResp](s.newPostJsonOptions(UriBookConfirm, req))
}

func (s standardApi) BookCancel(req kry_standard.BookCancelReq) kry_model.Result[kry_standard.BookCancelResp] {
	return network.JsonToResult[kry_standard.BookCancelResp](s.newPostJsonOptions(UriBookCancel, req))
}
