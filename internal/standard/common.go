package standard

import (
	"github.com/loebfly/keruyun-sdk-go/internal/network"
	"github.com/loebfly/keruyun-sdk-go/kry_standard"
)

func (s standardApi) KryCrm(req kry_standard.KryCommon) (any, error) {

	opt := s.newPostJsonOptions(req.Uri, req.Req)
	return network.JsonToAny(opt)
}
