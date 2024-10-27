package helpers

import (
	"testku/entities"

	"github.com/gofiber/fiber/v2"
)

// JSONResponse adalah helper untuk mengirim respons JSON dengan struktur APIResponse
func JSONResponse(c *fiber.Ctx, status string, message string, data interface{}, code int, errors interface{}) error {
	response := entities.APIResponse{
		Status:  status,
		Message: message,
		Data:    data,
		Code:    code,
		Errors:  errors,
	}
	return c.Status(code).JSON(response)
}

// SuccessResponse mengirim respons sukses dengan data dan kode HTTP
func SuccessResponse(c *fiber.Ctx, message string, data interface{}, code int) error {
	return JSONResponse(c, "success", message, data, code, nil)
}

// ErrorResponse mengirim respons error dengan pesan dan kode HTTP
func ErrorResponse(c *fiber.Ctx, message string, code int, errors interface{}) error {
	return JSONResponse(c, "error", message, nil, code, errors)
}
