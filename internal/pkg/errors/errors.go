package errors

import "fmt"

// Wrap обёртка над fmt.Errorf позволяющая добавить сообщение к ошибке.
// Если ошибка nil, то возвращается nil.
func Wrap(err error, msg string) error {
	if err == nil {
		return nil
	}

	return fmt.Errorf("%s: %w", msg, err)
}

// Wrapf обёртка над fmt.Errorf позволяющая добавить сообщение к ошибке
// с дополнительным форматированием сообщения.
// Если ошибка nil, то возвращается nil.
func Wrapf(err error, format string, args ...any) error {
	return Wrap(err, fmt.Sprintf(format, args...))
}
