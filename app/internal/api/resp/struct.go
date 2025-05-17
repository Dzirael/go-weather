package resp

import (
	"fmt"

	"github.com/go-errors/errors"
	"github.com/gofiber/fiber/v3"
)

type Error struct {
	httpCode  int
	Message   string         `json:"message"`
	ErrorCode int            `json:"error_code"`
	Metadata  map[string]any `json:"metadata,omitempty"`
}

func newError(httpCode int, errorCode int, err error, message ...string) *Error {
	msg := err.Error()
	if len(message) > 0 {
		msg = fmt.Sprintf("%s: %s", message[0], msg)
	}

	return &Error{
		httpCode:  httpCode,
		ErrorCode: errorCode,
		Message:   msg,
	}
}

func (e *Error) Respond(c fiber.Ctx) error {
	err := c.Status(e.httpCode).JSON(e)
	if err != nil {
		return errors.Errorf("respond with error: %w", err)
	}

	return nil
}

func (e *Error) With(key string, value any) *Error {
	if e.Metadata == nil {
		e.Metadata = make(map[string]any)
	}

	e.Metadata[key] = value
	return e
}
