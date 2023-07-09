package framework

import (
	"github.com/St-Pavlov-Data-Department/backend/constants"
	"github.com/St-Pavlov-Data-Department/backend/log"
	"github.com/jinzhu/configor"
	"os"
	"runtime"
	"strings"
)

func (e *PavlovEngine) loadConfig(autoReloadCallbackFn func(interface{})) (err error) {
	return configor.New(
		&configor.Config{
			AutoReload:         true,
			AutoReloadCallback: autoReloadCallbackFn,
		},
	).
		Load(e.Cfg, constants.PavlovConfigFilePath)
}

// GenerateExampleConfig creates example configs and writes to config file
func (e *PavlovEngine) generateExampleConfig(filePath string) (err error) {
	logger := log.CurrentModuleLogger()

	err = os.WriteFile(filePath, []byte(exampleConfig()), 0755)
	if err != nil {
		logger.
			WithError(err).
			WithField("config_filepath", filePath).
			Errorf("failed to generate example config")
		return err
	}
	logger.
		WithField("config_filepath", filePath).
		Infof("Minimum configuration has been generated. Please modify as needed and rerun. " +
			"For advanced configuration, please refer to the help documentation.")
	return err
}

func exampleConfig() string {
	c := `
log_level = "debug"

listen_addr = "0.0.0.0:8080"
gin_mode = "debug"
server_shutdown_max_wait_seconds = 5
sqlite_path = "./pavlov_sqlite.db"

`
	// To solve the issue of incorrect line breaks when opening a file in Notepad on Windows
	if runtime.GOOS == "windows" {
		c = strings.ReplaceAll(c, "\n", "\r\n")
	}
	return c
}
