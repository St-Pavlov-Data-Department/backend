package constants

import "errors"

const (
	NoErrorCode      = 0
	GeneralErrorCode = 1
)

var (
	ErrConfigPathNotFile      = errors.New("config path is not a file")
	ErrExampleConfigGenerated = errors.New("example config generated")

	ErrSqliteDBPathNotFile = errors.New("sqlite db path is not a file")

	ErrConfMySQLUsernameUnset = errors.New("mysql username not set")
	ErrConfMySQLPasswordUnset = errors.New("mysql password not set")
	ErrConfMySQLHostUnset     = errors.New("mysql host not set")
	ErrConfMySQLPortUnset     = errors.New("mysql port not set")
	ErrConfMySQLDBNameUnset   = errors.New("mysql db_name not set")
)
