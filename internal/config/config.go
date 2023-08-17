package config

import (
	"github.com/loebfly/keruyun-sdk-go/kry_model"
)

var Global kry_model.SdkConfig

func init() {
	Global.Version = "2.0"
}
