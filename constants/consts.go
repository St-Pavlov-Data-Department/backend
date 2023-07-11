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

const (
	DSNFormat = "%s:%s@tcp(%s:%s)/%s?timeout=%ds&charset=utf8mb4&parseTime=True&loc=Local"
)

var (
	PavlovConfigFilePath = path.Join(
		PavlovConfigPath,
		fmt.Sprintf("%s.%s", PavlovConfigName, PavlovConfigType),
	)
)
