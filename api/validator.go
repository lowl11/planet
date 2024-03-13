package api

import (
	baseValidator "github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/lowl11/planet/errors"
	"net/http"
)

const (
	contextKey = "validate"
)

type validator struct {
	*baseValidator.Validate
	turnOff bool
}

func newValidator() (*validator, error) {
	base := baseValidator.New()

	if err := base.RegisterValidation("uuid", validateUUID); err != nil {
		return nil, err // TODO: need implement BoostError?
	}

	if err := base.RegisterValidation("undefined", validateUndefined); err != nil {
		return nil, err
	}

	return &validator{
		Validate: base,
	}, nil
}

func (validator *validator) TurnOff() *validator {
	validator.turnOff = true
	return validator
}

func (validator *validator) Struct(object any) error {
	if validator.turnOff {
		return nil
	}

	validateError := validator.Validate.Struct(object)
	if validateError == nil {
		return nil
	}

	err := errors.
		New("Model validation error").
		SetHTTP(http.StatusUnprocessableEntity)

	validationErrors, ok := validateError.(baseValidator.ValidationErrors)
	if !ok {
		return err.SetError(validateError)
	}

	if len(validationErrors) == 0 {
		return nil
	}

	validations := make([]string, 0, len(validationErrors))
	for _, validationError := range validationErrors {
		validations = append(validations, validationError.Error())
	}

	return err.AddContext("validations", validations)
}

func (validator *validator) Var(variable any, tag string) error {
	if validator.turnOff {
		return nil
	}

	err := errors.
		New("Variable validation error").
		SetHTTP(http.StatusUnprocessableEntity)

	validateError := validator.Validate.Var(variable, tag)
	if err == nil {
		return err
	}

	validationErrors, ok := validateError.(baseValidator.ValidationErrors)
	if !ok {
		return err.AddContext(contextKey, validateError.Error())
	}

	if len(validationErrors) == 0 {
		return nil
	}

	validations := make([]string, 0, len(validationErrors))
	for _, validationError := range validationErrors {
		validations = append(validations, validationError.Error())
	}

	return err.AddContext(contextKey, validations)
}

func validateUUID(fl baseValidator.FieldLevel) (isValid bool) {
	switch val := fl.Field().Interface().(type) {
	case string:
		_, err := uuid.Parse(val)
		if err != nil {
			return false
		}

		return true
	case *string:
		_, err := uuid.Parse(*val)
		if err != nil {
			return false
		}

		return true
	case uuid.UUID:
		isValid = val != uuid.Nil
	case *uuid.UUID:
		if fl.Field().IsNil() {
			return
		}

		isValid = *val != uuid.Nil
	}

	return
}

func validateUndefined(fl baseValidator.FieldLevel) (isValid bool) {
	const undefined = "undefined"
	switch val := fl.Field().Interface().(type) {
	case string:
		return val != undefined
	case *string:
		value := *val
		return value != undefined
	}

	return
}
