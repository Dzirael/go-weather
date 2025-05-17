package middleware

import (
	"bytes"
	"fmt"

	"go-weather/app/internal/api/resp"
	"go-weather/app/pkg/logger"

	"github.com/go-errors/errors"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/requestid"
	"go.uber.org/zap"
)

func ErrorHandler(l *logger.Logger) fiber.ErrorHandler {
	return func(c fiber.Ctx, err error) error {
		requestID := requestid.FromContext(c)

		if fiberErr := new(fiber.Error); errors.As(err, &fiberErr) {
			return c.Status(fiberErr.Code).JSON(fiberErr)
		}
		stack := extractFirstStack(err)

		if l.Environment == logger.Dev {
			_, _ = fmt.Println("stack:\n", stack)
		} else {
			// l.Error(err.Error(), zap.String("stack", stack))
		}

		err = resp.Internal().With("request_id", requestID).Respond(c)
		if err != nil {
			l.Error("Failed to send internal error message", zap.Error(err))
		}

		return nil
	}
}

func extractFirstStack(err error) string {
	var lastExtractedError *errors.Error
	unwrappedErr := new(errors.Error)

	for {
		if !errors.As(err, &unwrappedErr) {
			if lastExtractedError != nil {
				return string(lastExtractedError.Stack())
			}
		}

		if bytes.Contains(unwrappedErr.Stack(), []byte("apperrors")) {
			if lastExtractedError != nil {
				return string(lastExtractedError.Stack())
			}
			return string(unwrappedErr.Stack())
		}

		lastExtractedError = unwrappedErr
		err = lastExtractedError.Unwrap()
	}
}
