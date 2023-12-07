package helper

import (
	"errors"
	"fmt"
	"strings"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func TranslateError(err error, trans ut.Translator) error {
	if err == nil {
		return nil
	}
	validatorErrs := err.(validator.ValidationErrors)
	var errs []error
	for _, e := range validatorErrs {
		translatedErr := fmt.Errorf(e.Translate(trans))
		errs = append(errs, translatedErr)
	}

	var errorStrings []string

	for _, err := range errs {
		errorStrings = append(errorStrings, err.Error())
	}
	return errors.New(strings.Join(errorStrings, "\n"))
}

func AddTranslation(tag string, errMessage string, trans ut.Translator, validate validator.Validate) {
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

	_ = validate.RegisterTranslation(tag, trans, registerFn, transFn)
}
