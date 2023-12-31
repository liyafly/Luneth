package global

import (
	"github.com/liyafly/Luneth/pkg/logger"
	"github.com/liyafly/Luneth/pkg/setting"
)

// a global variable
var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
