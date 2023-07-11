package config

import "gorm.io/gorm/logger"

type Config struct {
	LogLevel string `toml:"log_level"`

	// -------- HTTP --------
	ListenAddr                   string `toml:"listen_addr"`
	GinMode                      string `toml:"gin_mode"`
	ServerShutdownMaxWaitSeconds int    `toml:"server_shutdown_max_wait_seconds"`

	// -------- Database --------
	SqlitePath string `toml:"sqlite_path"`

	DBConnInfo *MySQLConnInfo `toml:"db_conn_info"`
}

type MySQLConnInfo struct {
	Username       string `toml:"username"`
	Password       string `toml:"password"`
	Host           string `toml:"host"`
	Port           string `toml:"port"`
	DBName         string `toml:"db_name"`
	GormLogLevel   logger.LogLevel   `toml:"gorm_log_level"`
	ConnectTimeout int    `toml:"connect_timeout"`
	MaxIdleConns   int    `toml:"max_idle_conns"`
	MaxOpenConns   int    `toml:"max_open_conns"`
}
