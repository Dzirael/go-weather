package apperrors

import (
	"fmt"
)

var (
	ErrAlreadyHaveSubscription = fmt.Errorf("email already subscribed")
	ErrCityNotSupported        = fmt.Errorf("city not supported")

	ErrConfirmationCodeNotFound = fmt.Errorf("confirmation code not found")
)
