package controller

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type AppValidationError struct {
	error
	Errors []AppErrItem `json:"errors"`
}

func Validate(c echo.Context, reqbody interface{}) error {
	if err := c.Bind(&reqbody); err != nil {
		return err
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(reqbody); err != nil {
		apperr := NewAppErr()
		for _, validationErr := range err.(validator.ValidationErrors) {
			apperr.Errors = append(apperr.Errors, AppErrItem{
				Path:    validationErr.Field(),
				Message: fmt.Sprintf("something wrong with %s", validationErr.Error()),
			})
		}
		return apperr
	}

	return nil
}
