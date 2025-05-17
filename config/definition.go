package config

// Env is magic key for get env variable
var Env = "ENV"

// Environment definitions for Logging
var (
	LogLevel       = "LOG_LEVEL"
	LogLevelFormat = "LOG_FORMAT"
	LogLevelCaller = "LOG_CALLER"
	LogLevelStack  = "LOG_STACK"
)

// Environment definitions for database
var (
	DatabaseDialect = "DB_DIALECT"
	DatabaseDSN     = "DB_DSN"
)

// Environment definitions for http
var (
	HttpServer = "HTTP_SERVER"
	HttpPort   = "HTTP_pORT"
)
