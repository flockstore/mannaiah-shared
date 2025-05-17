package fiber

import (
	"fmt"
	"github.com/flockstore/mannaiah-shared/config"
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// StartServer launches HTTP binding from config
func StartServer(app *fiber.App, logger *zap.Logger) error {
	sv := config.MustGet(config.HttpServer)
	port := config.MustGet(config.HttpPort)
	addr := fmt.Sprintf("%s:%s", sv, port)
	app.Use(fiberzap.New(fiberzap.Config{Logger: logger}))
	return app.Listen(addr)
}
