package apperrors

import "errors"

var (
	ErrAlreadyHaveSubscription = errors.New("email already subscribed")
	ErrCityNotSupported        = errors.New("city not supported")

	ErrConfirmationCodeNotFound = errors.New("confirmation code not found")
)
