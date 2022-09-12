package dto

import "github.com/go-playground/validator/v10"

type UserRequestDto struct {
	Name     string `json:"name,omitempty" validate:"required,min=6,max=32"`
	Email    string `json:"email,omitempty" validate:"required,email,min=6,max=32"`
	Password string `json:"password,omitempty" validate:"required,min=8,max=12"`
}

type UserResponseDto struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

type LoginRequestDto struct {
	Email    string `json:"email" validate:"required,email,min=6,max=32"`
	Password string `json:"password" validate:"required,min=8,max=12"`
}

type LoginResponseDto struct {
	Token string `json:"access_token"`
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

var validate = validator.New()

func ValidateLoginStruct(user LoginRequestDto) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func ValidateUserStruct(user UserRequestDto) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
