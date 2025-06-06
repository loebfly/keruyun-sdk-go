package standard

import (
	"github.com/loebfly/keruyun-sdk-go/internal/network"
	"github.com/loebfly/keruyun-sdk-go/kry_model"
	"github.com/loebfly/keruyun-sdk-go/kry_standard"
)

func (s standardApi) BusinessIncomePromoList(req kry_standard.BusinessIncomeReq) kry_model.Result[kry_standard.BusinessIncomePromoResp] {
	return network.JsonToResult[kry_standard.BusinessIncomePromoResp](s.newPostJsonOptions(UriBusinessIncomePromoList, req))
}

func (s standardApi) BusinessIncomeList(req kry_standard.BusinessIncomeReq) kry_model.Result[kry_standard.BusinessIncomeResp] {
	return network.JsonToResult[kry_standard.BusinessIncomeResp](s.newPostJsonOptions(UriBusinessIncomeList, req))
}

func (s standardApi) BusinessIncomePromoStatistics(req kry_standard.BusinessIncomePromoStatisticsReq) kry_model.Result[kry_standard.BusinessIncomePromoStatisticsResp] {
	return network.JsonToResult[kry_standard.BusinessIncomePromoStatisticsResp](s.newPostJsonOptions(UriBusinessIncomePromoStatistics, req))
}
func (s standardApi) BusinessIncomeConstituteList(req kry_standard.BusinessIncomeReq) kry_model.Result[kry_standard.BusinessIncomeConstituteResp] {
	return network.JsonToResult[kry_standard.BusinessIncomeConstituteResp](s.newPostJsonOptions(UriBusinessIncomeConstituteList, req))
}

// func (s standardApi) PaidIncomeList(req kry_standard.BookConfirmReq) kry_model.Result[kry_standard.BookConfirmResp] {
// 	return network.JsonToResult[kry_standard.BookConfirmResp](s.newPostJsonOptions(UriBookConfirm, req))
// }

// func (s standardApi) PaymentReconciliationList(req kry_standard.BookCancelReq) kry_model.Result[kry_standard.BookCancelResp] {
// 	return network.JsonToResult[kry_standard.BookCancelResp](s.newPostJsonOptions(UriBookCancel, req))
// }
