package app

import (
	"fmt"
	"reflect"

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
	return validation.Validate.Struct(s)
}

func (validation *Validation) RegisterValidation(s string, f func(validator.FieldLevel) bool) {
	validation.Validate.RegisterValidation(s, f)
}

func (validation *Validation) TranslateError(err error) validator.ValidationErrors {
	if err == nil {
		return nil
	}
	validatorErrs := err.(validator.ValidationErrors)
	var errs []error
	for _, e := range validatorErrs {
		translatedErr := fmt.Errorf(e.Translate(*validation.Translator))
		errs = append(errs, translatedErr)
	}

	var errorStrings []string

	for _, err := range errs {
		errorStrings = append(errorStrings, err.Error())
	}

	var validationErrors validator.ValidationErrors

	// Append custom validator.FieldError instances
	customFieldError := CreateValidationError("field", "tag", "message")
	fmt.Println(customFieldError)
	validationErrors = append(validationErrors, new(&customFieldError))
	fmt.Println(validationErrors)

	return validationErrors
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

// Custom Field Error
type CustomValidationError struct {
	field string
	tag   string
	err   string
}

func CreateValidationError(field, tag, errorMessage string) validator.FieldError {
	return &CustomValidationError{
		field: field,
		tag:   tag,
		err:   errorMessage,
	}
}

// Error returns the error message for CustomValidationError
func (e *CustomValidationError) Error() string {
	return fmt.Sprintf("Field '%s' failed validation on tag '%s'", e.field, e.tag)
}

// Field returns the field name for CustomValidationError
func (e *CustomValidationError) Field() string {
	return e.field
}

// Tag returns the tag for CustomValidationError
func (e *CustomValidationError) Tag() string {
	return e.tag
}

// ActualTag returns the actual tag for CustomValidationError
func (e *CustomValidationError) ActualTag() string {
	return e.tag
}

// Kind returns the kind of data for CustomValidationError
func (e *CustomValidationError) Kind() reflect.Kind {
	return reflect.String // Modify this according to your data type
}

// Param returns the param for CustomValidationError
func (e *CustomValidationError) Param() string {
	return "" // Modify this if you have parameter info
}

// Value returns the value of the field for CustomValidationError
func (e *CustomValidationError) Value() interface{} {
	return nil // Modify this to return actual field value
}

func (e *CustomValidationError) Namespace() string {
	return ""
}

func (e *CustomValidationError) StructField() string {
	return ""
}

func (e *CustomValidationError) StructNamespace() string {
	return ""
}

func (e *CustomValidationError) Translate(ut ut.Translator) string {
	return ""
}

func (e *CustomValidationError) Type() reflect.Type {
	var fieldError validator.FieldError
	return reflect.TypeOf(fieldError)
}
