package config

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
	"sync"
)

var (
	once sync.Once
	v    *viper.Viper
)

// Init initializes base Viper instance with ENV-only config
func Init(prefix string, defaults map[string]interface{}) {
	once.Do(func() {
		v = viper.New()

		v.SetEnvPrefix(prefix)
		v.AutomaticEnv()
		v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		for key, val := range defaults {
			v.SetDefault(key, val)
		}
	})
}

// Get returns the underlying Viper instance
func Get() *viper.Viper {
	if v == nil {
		panic("config not initialized, call config.Init() first")
	}
	return v
}

// MustGet fetches a string key or panics if missing
func MustGet(key string) string {
	value := Get().GetString(key)
	if value == "" {
		panic(fmt.Sprintf("missing required config key: %s", key))
	}
	return value
}
