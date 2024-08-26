package utils

import (
	"sync"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type ValidationService struct {
	validate   *validator.Validate
	translator ut.Translator
}

var (
	once     sync.Once
	instance *ValidationService
)

func GetValidator() *ValidationService {
	once.Do(func() {
		en := en.New()
		uni := ut.New(en, en)
		trans, _ := uni.GetTranslator("en")

		validate := validator.New()

		// custom errors msg
		validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
			return ut.Add("required", "{0} is required", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("required", fe.Field())
			return t
		})
		// Register min length translation
		validate.RegisterTranslation("min", trans, func(ut ut.Translator) error {
			return ut.Add("min", "{0} must be at least {1} characters long", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("min", fe.Field(), fe.Param())
			return t
		})

		// Register max length translation
		validate.RegisterTranslation("max", trans, func(ut ut.Translator) error {
			return ut.Add("max", "{0} cannot be more than {1} characters long", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("max", fe.Field(), fe.Param())
			return t
		})

		// Register max length translation
		validate.RegisterTranslation("len", trans, func(ut ut.Translator) error {
			return ut.Add("len", "{0} should equel {1} characters", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("len", fe.Field(), fe.Param())
			return t
		})

		instance = &ValidationService{
			validate:   validate,
			translator: trans,
		}
	})
	return instance
}

// Validate validates the given struct and returns translated error messages
func (vs *ValidationService) Validate(s interface{}) map[string]string {
	err := vs.validate.Struct(s)
	if err != nil {
		errors := make(map[string]string)
		for _, validationErr := range err.(validator.ValidationErrors) {
			errors[validationErr.Field()] = validationErr.Translate(vs.translator)
		}
		return errors
	}
	return nil
}
