package apperrors

import "errors"

var (
	ErrInvalidID     = errors.New("invalid id")
	ErrTooShort      = errors.New("слишком мало символов")
	ErrInvalidEmail  = errors.New("неправильная почта")
	ErrWeakPassword  = errors.New("слабый пароль")
	ErrUserNotFound  = errors.New("user not found")
	ErrNoteNotFound	 = errors.New("note not found")
)	