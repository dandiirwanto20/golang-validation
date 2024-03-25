package govalidation

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestValidation(t *testing.T) {
	// initialize validator
	var validate *validator.Validate = validator.New()

	if validate == nil {
		t.Error("Validate is nil")
	}
}

// variable validation
func TestValidationVariable(t *testing.T) {
	validate := validator.New()
	user := "Dandi"

	// tag adalah jenis validasi
	err := validate.Var(user, "required") // (variabel, tag)

	if err != nil {
		fmt.Println(err.Error())
	}
}

// two variable validation
// example case if we use password and confirmpassword validation
// and we can use tag eqfield
func TestValidationTwoVariables(t *testing.T) {
	validate := validator.New()

	password := "secret"
	confirmPassword := "wrong"
	// confirmPassword := "secret"

	err := validate.VarWithValue(password, confirmPassword, "eqfield")

	if err != nil {
		fmt.Println(err.Error())
	}
}

// Baked-in Validation

// Multiple Tag Validation
func TestMultipleTag(t *testing.T) {
	validate := validator.New()

	user := "200600"
	// tag adalah jenis validasi
	err := validate.Var(user, "required,numeric")

	if err != nil {
		fmt.Println(err.Error())
	}
}

// validation tag param
func TestTagParam(t *testing.T) {
	validate := validator.New()

	user := "9999999999"
	// tag adalah jenis validasi
	err := validate.Var(user, "required,numeric,min=5,max=10")

	if err != nil {
		fmt.Println(err.Error())
	}
}

// Struct Validation
func TestStruct(t *testing.T) {
	type LoginRequest struct {
		Username string `validate:"required,email"`
		Password string `validate:"required,min=5"`
	}

	validate := validator.New()

	loginRequest := LoginRequest{
		Username: "test@gmail.com",
		Password: "user123",
	}

	err := validate.Struct(loginRequest)

	if err != nil {
		fmt.Println(err.Error())
	}
}

// Validation Error
func TestValidationError(t *testing.T) {
	type LoginRequest struct {
		Username string `validate:"required,email"`
		Password string `validate:"required,min=5"`
	}

	validate := validator.New()

	loginRequest := LoginRequest{
		Username: "dandi",
		Password: "user1",
	}

	err := validate.Struct(loginRequest)

	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			fmt.Println("error", fieldError.Tag(), "with error", fieldError.Error())
		}
	}
}
