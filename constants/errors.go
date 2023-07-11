package constants

import "errors"

const (
	NoErrorCode      = 0
	GeneralErrorCode = 1
)

var (
	ConfigPathNotFileErr      = errors.New("config path is not a file")
	ExampleConfigGeneratedErr = errors.New("example config generated")

	SqliteDBPathNotFileErr = errors.New("sqlite db path is not a file")
)
