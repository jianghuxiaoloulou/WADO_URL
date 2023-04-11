package global

import (
	"WowjoyProject/WADO_URL/pkg/logger"
	"WowjoyProject/WADO_URL/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	GeneralSetting  *setting.GeneralSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
