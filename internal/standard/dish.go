package standard

import (
	"github.com/loebfly/keruyun-sdk-go/internal/network"
	"github.com/loebfly/keruyun-sdk-go/kry_model"
	"github.com/loebfly/keruyun-sdk-go/kry_standard"
)

func (s standardApi) DishQueryPage(req kry_standard.DishQueryPageReq) kry_model.Result[kry_standard.DishPageResp] {
	return network.JsonToResult[kry_standard.DishPageResp](s.newPostJsonOptions(UriDishQueryPage, req))
}

func (s standardApi) DishQueryDetail(req kry_standard.DishQueryDetailReq) kry_model.Result[kry_standard.DishDetailResp] {
	return network.JsonToResult[kry_standard.DishDetailResp](s.newPostJsonOptions(UriDishQueryDetail, req))
}

func (s standardApi) DishQueryCategory(req kry_standard.DishQueryCategoryReq) kry_model.Result[kry_standard.DishCategoryResp] {
	return network.JsonToResult[kry_standard.DishCategoryResp](s.newPostJsonOptions(UriDishQueryCategory, req))
}
