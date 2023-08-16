package config

import "github.com/loebfly/keruyun-sdk-go/kry_config"

var Global kry_config.Config

func init() {
	Global.Version = "2.0"
}
