package errors

import "errors"

var (
	ErrSuccess        = errors.New("Успешно")
	ErrInternalServer = errors.New("Внутреняя ошибка сервера")
)

var StatusCodes map[string]int = map[string]int{
	//200 ...
	ErrSuccess.Error(): 200,

	//500 ...
	ErrInternalServer.Error(): 500,
}
