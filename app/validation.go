package app

import (
	"mfahmii/golang-restful/exception"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

type Validation struct {
	Validate   *validator.Validate
	Translator *ut.Translator
}

func NewValidation() *Validation {
	validate := validator.New()
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	_ = enTranslations.RegisterDefaultTranslations(validate, trans)

	return &Validation{
		Validate:   validate,
		Translator: &trans,
	}
}

func (validation *Validation) Struct(s interface{}) error {
	return validation.TranslateError(validation.Validate.Struct(s))
}

func (validation *Validation) RegisterValidation(s string, f func(validator.FieldLevel) bool) {
	validation.Validate.RegisterValidation(s, f)
}

func (validation *Validation) TranslateError(err error) error {
	if err == nil {
		return nil
	}

	validatorErrs, ok := err.(validator.ValidationErrors)
	if !ok {
		return err // Return the original error if it's not of type ValidationErrors
	}

	var errMsgs []string
	for _, e := range validatorErrs {
		translatedErr := e.Translate(*validation.Translator)
		errMsgs = append(errMsgs, translatedErr)
	}

	return exception.NewValidationError(strings.Join(errMsgs, ";"))
}

func (validation *Validation) AddTranslation(tag string, errMessage string) {
	registerFn := func(ut ut.Translator) error {
		return ut.Add(tag, errMessage, false)
	}

	transFn := func(ut ut.Translator, fe validator.FieldError) string {
		param := fe.Param()
		tag := fe.Tag()

		t, err := ut.T(tag, fe.Field(), param)
		if err != nil {
			return fe.(error).Error()
		}
		return t
	}

	_ = validation.Validate.RegisterTranslation(tag, *validation.Translator, registerFn, transFn)
}
