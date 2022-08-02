package util

import "gin-learn/pkg/setting"

func Setup() {
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
}
