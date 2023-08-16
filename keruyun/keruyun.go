package keruyun

import (
	internalConfig "github.com/loebfly/keruyun-sdk-go/internal/config"
	internalStandard "github.com/loebfly/keruyun-sdk-go/internal/standard"
	"github.com/loebfly/keruyun-sdk-go/kry_config"
	"github.com/loebfly/keruyun-sdk-go/standard"
)

// RegisterSdk 注册SDK
func RegisterSdk(cfg kry_config.Config) {
	internalConfig.Global.Domain = cfg.Domain
	internalConfig.Global.AppKey = cfg.AppKey
	internalConfig.Global.SecretKey = cfg.SecretKey
	if cfg.Version != "" {
		internalConfig.Global.Version = cfg.Version
	}
}

// NewStandardAPI 创建智享版API
func NewStandardAPI(shopId int64) (standard.Api, error) {
	return internalStandard.NewAPI(shopId)
}
