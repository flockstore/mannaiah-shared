package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func cleanupEnv(keys ...string) func() {
	return func() {
		for _, key := range keys {
			_ = os.Unsetenv(key)
		}
	}
}

// TestInitAndGet verifies Init sets defaults and Get returns correct values
func TestInitAndGet(t *testing.T) {
	t.Cleanup(cleanupEnv("MANNAIAH_SERVER_PORT", "MANNAIAH_DEBUG"))

	Init("MANNAIAH", map[string]interface{}{
		"server.port": "3000",
		"debug":       "false",
	})

	v := Get()
	assert.NotNil(t, v, "expected non-nil Viper instance")
	assert.Equal(t, "3000", v.GetString("server.port"))
	assert.Equal(t, "false", v.GetString("debug"))
}

// TestEnvOverride ensures environment variables override default values
func TestEnvOverride(t *testing.T) {
	t.Cleanup(cleanupEnv("MANNAIAH_SERVER_PORT"))

	err := os.Setenv("MANNAIAH_SERVER_PORT", "8080")
	assert.NoError(t, err)

	Init("MANNAIAH", map[string]interface{}{
		"server.port": "3000",
	})

	port := MustGet("server.port")
	assert.Equal(t, "8080", port)
}

// TestMustGetPanics checks that MustGet panics when the key is missing
func TestMustGetPanics(t *testing.T) {
	t.Cleanup(cleanupEnv("MANNAIAH_SERVER_PORT"))

	Init("MANNAIAH", map[string]interface{}{})

	assert.Panics(t, func() {
		_ = MustGet("missing.key")
	}, "expected panic for missing key")
}

// TestDynamicEnvChange verifies that Viper picks up environment changes at runtime
func TestDynamicEnvChange(t *testing.T) {
	t.Cleanup(cleanupEnv("MANNAIAH_DYNAMIC_KEY"))

	_ = os.Setenv("MANNAIAH_DYNAMIC_KEY", "initial")

	Init("MANNAIAH", map[string]interface{}{})

	assert.Equal(t, "initial", Get().GetString("dynamic.key"))

	_ = os.Setenv("MANNAIAH_DYNAMIC_KEY", "updated")

	assert.Equal(t, "updated", Get().GetString("dynamic.key"))
}
