package fiber

import (
	"github.com/flockstore/mannaiah-shared/endpoint"
	"github.com/gofiber/fiber/v2"
)

// RegisterCrudRoutes mounts generic CRUD routes for any entity T.
func RegisterCrudRoutes[T any](app *fiber.App, basePath string, eps endpoint.Endpoints[T]) {

	app.Post(basePath, func(c *fiber.Ctx) error {
		req, err := DecodeCreateRequest[T](c)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		resp, _ := eps.Create(c.Context(), req)
		return EncodeResponse(c, resp.(endpoint.Response[any]))
	})

	app.Get(basePath+"/:id", func(c *fiber.Ctx) error {
		req := DecodeGetRequest(c)
		resp, _ := eps.Get(c.Context(), req)
		return EncodeResponse(c, resp.(endpoint.Response[*T]))
	})

	app.Put(basePath+"/:id", func(c *fiber.Ctx) error {
		req, err := DecodeUpdateRequest[T](c)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		resp, _ := eps.Update(c.Context(), req)
		return EncodeResponse(c, resp.(endpoint.Response[any]))
	})

	app.Delete(basePath+"/:id", func(c *fiber.Ctx) error {
		req := DecodeDeleteRequest(c)
		resp, _ := eps.Delete(c.Context(), req)
		return EncodeResponse(c, resp.(endpoint.Response[any]))
	})

	app.Post(basePath+"/list", func(c *fiber.Ctx) error {
		req := DecodeListRequest(c)
		resp, _ := eps.List(c.Context(), req)
		return EncodeResponse(c, resp.(endpoint.Response[[]T]))
	})

}
