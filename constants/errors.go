package constants

import "errors"

var (
	ConfigPathNotFileErr      = errors.New("config path is not a file")
	ExampleConfigGeneratedErr = errors.New("example config generated")

	SqliteDBPathNotFileErr = errors.New("sqlite db path is not a file")
)
