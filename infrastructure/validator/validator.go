package validator

import (
	"encoding/json"
	"errors"
	"github.com/odanraujo/financial-organizer-api/infrastructure/excp"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translations.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserRequest(validation_err error) *excp.Exception {

	var jsonException *json.UnmarshalTypeError
	var jsonExceptionError validator.ValidationErrors

	if errors.As(validation_err, &jsonException) {
		return excp.BadRequestException("Invalid request type")
	}

	if errors.As(validation_err, &jsonExceptionError) {
		errorsCauses := []excp.Causes{}

		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := excp.Causes{
				Field:   e.Field(),
				Message: e.Translate(transl),
			}

			errorsCauses = append(errorsCauses, cause)
		}
		return excp.BadRequestValidationException("Some fields are invalid", errorsCauses)
	}

	return excp.BadRequestException("Error trying to convert fields")
}
