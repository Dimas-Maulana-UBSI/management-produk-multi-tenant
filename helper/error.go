package helper

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"management-produk/model/web"
	"net/http"
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Err        error  `json:"-"`
}

func (e *AppError) Error() string {
	if e == nil {
		return ""
	}
	if e.Err != nil {
		return e.Err.Error()
	}
	return e.Message
}

func NewAppError(status int, message string, err error) *AppError {
	return &AppError{StatusCode: status, Message: message, Err: err}
}

func BadRequest(msg string) *AppError {
	return NewAppError(http.StatusBadRequest, msg, errors.New(msg))
}

func NotFound(msg string) *AppError {
	return NewAppError(http.StatusNotFound, msg, errors.New(msg))
}

func Unauthorized(msg string) *AppError {
	return NewAppError(http.StatusUnauthorized, msg, errors.New(msg))
}

func Conflict(msg string) *AppError {
	return NewAppError(http.StatusConflict, msg, errors.New(msg))
}

func Internal(msg string, err error) *AppError {
	if err == nil {
		err = errors.New(msg)
	}
	return NewAppError(http.StatusInternalServerError, msg, err)
}

func ToAppError(err error) *AppError {
	if err == nil {
		return nil
	}
	var ae *AppError
	if errors.As(err, &ae) {
		return ae
	}
	return Internal("internal server error", err)
}

func RespondFiberError(ctx *fiber.Ctx, err error) error {
	ae := ToAppError(err)
	if ae == nil {
		return ctx.SendStatus(http.StatusOK)
	}
	resp := web.WebResponse{
		Status:  ae.StatusCode,
		Message: ae.Message,
		Data:    nil,
	}
	return ctx.Status(ae.StatusCode).JSON(resp)
}
