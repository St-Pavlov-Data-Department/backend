package framework

import (
	"github.com/St-Pavlov-Data-Department/backend/config"
	"github.com/St-Pavlov-Data-Department/backend/constants"
	"github.com/St-Pavlov-Data-Department/backend/log"
	"github.com/jinzhu/configor"
	"os"
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
func generateExampleConfig(filePath string) (err error) {
	logger := log.CurrentModuleLogger()

	err = os.WriteFile(filePath,
		[]byte(config.ExampleConfig()),
		0755,
	)
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
