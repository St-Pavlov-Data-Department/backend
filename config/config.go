package config

type Config struct {
	LogLevel string `toml:"log_level"`

	// -------- HTTP --------
	ListenAddr                   string `toml:"listen_addr"`
	GinMode                      string `toml:"gin_mode"`
	ServerShutdownMaxWaitSeconds int    `toml:"server_shutdown_max_wait_seconds"`

	// -------- Database --------
	SqlitePath string `toml:"sqlite_path"`
}
