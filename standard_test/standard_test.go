package standard_test

import (
	"github.com/loebfly/keruyun-sdk-go/keruyun"
	"github.com/loebfly/keruyun-sdk-go/kry_config"
	"testing"
)

var shopTokenCache = map[int64]string{}

func Test_standard(t *testing.T) {
	t.Log("standard_test.go")
	keruyun.RegisterSdk(kry_config.Config{
		Version: "2.0",
		SetTokenForShopIdHandle: func(shopId int64, token string) {
			shopTokenCache[shopId] = token
		},
		GetTokenForShopIdHandle: func(shopId int64) string {
			return shopTokenCache[shopId]
		},
	})
	standardApi, err := keruyun.NewStandardAPI(7884774)
	if err != nil {
		t.Error(err)
		return
	}

	res := standardApi.ShopQueryStoreDetail()
	if !res.IsSuccess() {
		t.Error(res.ErrorMsg())
		return
	}
	t.Log(res.Result)
}
