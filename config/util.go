package config

import (
	"strconv"
)

// SetDefaults provide a map for default config mappings
func SetDefaults() map[string]interface{} {
	def := make(map[string]interface{}, 1)

	def[Env] = "development"

	def[LogLevel] = "debug"
	def[LogLevelFormat] = "json"
	def[LogLevelCaller] = true
	def[LogLevelStack] = true

	def[DatabaseDialect] = "mysql"
	def[DatabaseDSN] = "root:secret@tcp(127.0.0.1:3306)/mannaiah?parseTime=true"

	def[HttpServer] = "0.0.0.0"
	def[HttpPort] = "3000"

	return def
}

// ParseBool parses a boolean from an env variable.
func ParseBool(env string) bool {
	b, err := strconv.ParseBool(env)
	if err != nil {
		return false
	}
	return b
}
