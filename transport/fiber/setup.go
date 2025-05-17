package fiber

import (
	"fmt"
	"github.com/flockstore/mannaiah-shared/config"
	"github.com/gofiber/fiber/v2"
)

// StartServer launches HTTP binding from config
func StartServer(app *fiber.App) error {
	sv := config.MustGet(config.HttpServer)
	port := config.MustGet(config.HttpPort)
	addr := fmt.Sprintf("%s:%s", sv, port)
	return app.Listen(addr)
}
