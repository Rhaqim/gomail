package errors

import "errors"

// Auth Errors
var (
	ErrEmptyHost        = errors.New("email host is empty")
	ErrEmptyPort        = errors.New("email port is empty")
	ErrEmptyUsername    = errors.New("email username is empty")
	ErrEmptyPassword    = errors.New("email password is empty")
	ErrEmptyFrom        = errors.New("email from is empty")
	ErrEmptyTemplateDir = errors.New("email template directory is empty")
)

// Email Errors
var (
	ErrEmptyTo = errors.New("email recipeint(s) is empty")
)
