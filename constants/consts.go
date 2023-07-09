package constants

import (
	"fmt"
	"path"
)

const (
	PavlovConfigName = "pavlov"
	PavlovConfigType = "toml"
	PavlovConfigPath = "./"
)

var (
	PavlovConfigFilePath = path.Join(
		PavlovConfigPath,
		fmt.Sprintf("%s.%s", PavlovConfigName, PavlovConfigType),
	)
)
