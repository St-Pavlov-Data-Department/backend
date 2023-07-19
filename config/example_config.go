package config

import (
	"runtime"
	"strings"
)

func ExampleConfig() string {
	c := `
log_level = "debug"

listen_addr = "0.0.0.0:1999"
gin_mode = "debug"
server_shutdown_max_wait_seconds = 5

game_resource_path = "/repo/github.com/1999GameResource"

[db_conn_info]
username = "user"
password = "pass"
host = "127.0.0.1"
port = "3306"
db_name = "stpavlov_data_db"
gorm_log_level = 4 # 1-silent 2-error 3-warn 4-info
connect_timeout = 10
max_idle_conns = 3
max_open_conns = 5

`
	// To solve the issue of incorrect line breaks when opening a file in Notepad on Windows
	if runtime.GOOS == "windows" {
		c = strings.ReplaceAll(c, "\n", "\r\n")
	}
	return c
}
