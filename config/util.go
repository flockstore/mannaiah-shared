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
