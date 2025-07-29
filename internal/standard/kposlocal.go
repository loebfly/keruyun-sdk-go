package standard

import (
	"github.com/loebfly/keruyun-sdk-go/internal/network"
	"github.com/loebfly/keruyun-sdk-go/kry_model"
	"github.com/loebfly/keruyun-sdk-go/kry_standard"
)

func (s standardApi) QueryKposLocal(req kry_standard.KposlocalReq) kry_model.Result[kry_standard.KposLocalResp] {
	return network.JsonToResult[kry_standard.KposLocalResp](s.newPostJsonOptions(UriKposLocalOrderDetail, req))
}
