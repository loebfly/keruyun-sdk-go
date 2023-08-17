package keruyun

import (
	internalConfig "github.com/loebfly/keruyun-sdk-go/internal/config"
	internalStandard "github.com/loebfly/keruyun-sdk-go/internal/standard"
	"github.com/loebfly/keruyun-sdk-go/kry_model"
	"github.com/loebfly/keruyun-sdk-go/kry_standard"
)

// RegisterSdk 注册SDK
func RegisterSdk(cfg kry_model.SdkConfig) {
	internalConfig.Global.Domain = cfg.Domain
	internalConfig.Global.AppKey = cfg.AppKey
	internalConfig.Global.SecretKey = cfg.SecretKey
	if cfg.Version != "" {
		internalConfig.Global.Version = cfg.Version
	}
	if cfg.SetTokenForShopIdHandle != nil {
		internalConfig.Global.SetTokenForShopIdHandle = cfg.SetTokenForShopIdHandle
	}
	if cfg.GetTokenForShopIdHandle != nil {
		internalConfig.Global.GetTokenForShopIdHandle = cfg.GetTokenForShopIdHandle
	}
	if cfg.PrintApiLogHandle != nil {
		internalConfig.Global.PrintApiLogHandle = cfg.PrintApiLogHandle
	}
}

// NewStandardAPI 创建智享版API
func NewStandardAPI(shopId int64) (kry_standard.Api, error) {
	return internalStandard.NewAPI(shopId)
}
