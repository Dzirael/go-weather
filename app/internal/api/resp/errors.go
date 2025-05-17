package resp

import (
	"errors"

	"github.com/gofiber/fiber/v3"
)

func Internal() *Error {
	const message = "internal error, try again later. if the error still persist, please contact our support"

	return newError(fiber.StatusInternalServerError, 0, errors.New(message))
}

func BadRequest(err error, errorCode int, message ...string) *Error {
	return newError(fiber.StatusBadRequest, errorCode, err, message...)
}

func Unauthorized(err error, errorCode int, message ...string) *Error {
	return newError(fiber.StatusUnauthorized, errorCode, err, message...)
}

func NotFound(err error, errorCode int, message ...string) *Error {
	return newError(fiber.StatusNotFound, errorCode, err, message...)
}

func Forbidden(err error, errorCode int, message ...string) *Error {
	return newError(fiber.StatusForbidden, errorCode, err, message...)
}

func Conflict(err error, errorCode int, message ...string) *Error {
	return newError(fiber.StatusConflict, errorCode, err, message...)
}

const (
	ErrZeroCode = iota + 1000
)
