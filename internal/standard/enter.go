package standard

import (
	"errors"
	"github.com/loebfly/keruyun-sdk-go/internal/config"
	"github.com/loebfly/keruyun-sdk-go/internal/network"
	"github.com/loebfly/keruyun-sdk-go/kry_standard"
	"strconv"
	"time"
)

type standardApi struct {
	ShopId int64
	token  string
}

func (s standardApi) NewPostQuery() map[string]string {
	return map[string]string{
		network.SignPtrAppKey:    config.Global.AppKey,
		network.SingPtrShopId:    strconv.FormatInt(s.ShopId, 10),
		network.SignPtrTimestamp: strconv.FormatInt(time.Now().Unix(), 10),
		network.SignPtrVersion:   config.Global.Version,
		network.SignPtrToken:     s.token,
	}
}

func NewAPI(shopId int64) (kry_standard.Api, error) {
	// 获取Token
	if config.Global.GetTokenForShopIdHandle != nil {
		token := config.Global.GetTokenForShopIdHandle(shopId)
		if token != "" {
			return &standardApi{ShopId: shopId, token: token}, nil
		}
	}

	res := network.JsonToResult[map[string]string](network.JsonOptions{
		Host:   config.Global.Domain,
		Uri:    "/open/v1/token/get",
		Method: network.GET,
		Query: map[string]string{
			network.SignPtrAppKey:    config.Global.AppKey,
			network.SignPtrSecretKey: config.Global.SecretKey,
			network.SingPtrShopId:    strconv.FormatInt(shopId, 10),
			network.SignPtrTimestamp: strconv.FormatInt(time.Now().Unix(), 10),
			network.SignPtrVersion:   config.Global.Version,
		},
	})
	if !res.IsSuccess() {
		return nil, errors.New(res.ErrorMsg())
	}

	if config.Global.SetTokenForShopIdHandle != nil {
		config.Global.SetTokenForShopIdHandle(shopId, res.Result["token"])
	}
	return &standardApi{ShopId: shopId, token: res.Result["token"]}, nil
}
