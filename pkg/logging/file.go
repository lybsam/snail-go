package logging

import (
	"fmt"
	"snail/pkg/setting"
	"time"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s%s", setting.RUNTIME_ROOT_PATH, setting.LOG_SAVE_PATH)
}

func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		setting.LOG_SAVE_NAME,
		time.Now().Format(setting.TIME_FORMAT),
		setting.LOG_FILE_EXT,
	)
}
