package standard

import (
	"github.com/loebfly/keruyun-sdk-go/internal/network"
	"github.com/loebfly/keruyun-sdk-go/kry_model"
	"github.com/loebfly/keruyun-sdk-go/kry_standard"
)

func (s standardApi) QueryKposLocal(req kry_standard.KposlocalReq) kry_model.Result[kry_standard.KposlocalAddResp] {
	return network.JsonToResult[kry_standard.KposlocalAddResp](s.newPostJsonOptions(UriKposLocalOrderDetail, req))
}

func (s standardApi) QueryKposLocalAdd(req kry_standard.KposlocalAddReq) kry_model.Result[kry_standard.KposlocalAddResp] {
	return network.JsonToResult[kry_standard.KposlocalAddResp](s.newPostJsonOptions(UriKposLocalOrderAdd, req))
}

func (s standardApi) QueryTableStatus(req kry_standard.KposTableStatusQueryReq) kry_model.Result[kry_standard.TableStatusResp] {
	return network.JsonToResult[kry_standard.TableStatusResp](s.newPostJsonOptions(UriKposTableStatusQuery, req))
}
