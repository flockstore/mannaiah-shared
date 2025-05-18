package fiber

import (
	"github.com/flockstore/mannaiah-shared/endpoint"
	"github.com/flockstore/mannaiah-shared/transport"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var validate = validator.New()

// EncodeResponse writes a generic response as JSON, handling errors if present.
func EncodeResponse[T any](c *fiber.Ctx, resp endpoint.Response[T]) error {

	if resp.Err != nil {
		code := transport.CodeOf(resp.Err)
		return c.Status(code).JSON(fiber.Map{
			"error": resp.Err.Error(),
		})
	}

	return c.JSON(resp.Data)
}

// DecodeCreateRequest decodes a JSON body into a CreateRequest[T].
func DecodeCreateRequest[T any](c *fiber.Ctx) (endpoint.CreateRequest[T], error) {
	var model T

	if err := c.BodyParser(&model); err != nil {
		return endpoint.CreateRequest[T]{}, err
	}

	if err := validate.Struct(model); err != nil {
		return endpoint.CreateRequest[T]{}, fiber.NewError(fiber.StatusBadRequest, "Provided body is invalid: "+err.Error())
	}

	return endpoint.CreateRequest[T]{Model: &model}, nil
}

// DecodeUpdateRequest decodes a JSON body into an UpdateRequest[T].
func DecodeUpdateRequest[T any](c *fiber.Ctx) (endpoint.UpdateRequest[T], error) {
	var model T

	if err := c.BodyParser(&model); err != nil {
		return endpoint.UpdateRequest[T]{}, err
	}

	if err := validate.Struct(model); err != nil {
		return endpoint.UpdateRequest[T]{}, fiber.NewError(fiber.StatusBadRequest, "Provided body is invalid: "+err.Error())
	}

	return endpoint.UpdateRequest[T]{Model: &model}, nil
}

// DecodeGetRequest creates a GetRequest using the ":id" path param.
func DecodeGetRequest(c *fiber.Ctx) endpoint.GetRequest {
	return endpoint.GetRequest{ID: c.Params("id")}
}

// DecodeDeleteRequest creates a DeleteRequest using the ":id" path param.
func DecodeDeleteRequest(c *fiber.Ctx) endpoint.DeleteRequest {
	return endpoint.DeleteRequest{ID: c.Params("id")}
}

// DecodeListRequest prepares a ListRequest without filters (can be extended later).
func DecodeListRequest(c *fiber.Ctx) endpoint.ListRequest {
	return endpoint.ListRequest{QueryFns: []func(*gorm.DB) *gorm.DB{}}
}
