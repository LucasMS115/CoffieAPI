package user

import (
	"testing"

	userhttp "coffie/internal/feature/user/http"
)

func TestValidateRegisterUserRequest_ValidRequest(testingContext *testing.T) {
	registerRequest := &userhttp.RegisterUser{
		Name:  "Lucas",
		Email: "lucas@email.com",
	}

	validationErrors := registerRequest.Validate()

	if len(validationErrors) != 0 {
		testingContext.Fatalf("expected no validation errors, got %d", len(validationErrors))
	}
}

func TestValidateRegisterUserRequest_MissingName(testingContext *testing.T) {
	registerRequest := &userhttp.RegisterUser{
		Email: "lucas@email.com",
	}

	validationErrors := registerRequest.Validate()

	if len(validationErrors) != 1 {
		testingContext.Fatalf("expected 1 validation error, got %d", len(validationErrors))
	}
	if validationErrors[0].Field != "name" {
		testingContext.Fatalf("expected field 'name', got %q", validationErrors[0].Field)
	}
	if validationErrors[0].Message != "is required" {
		testingContext.Fatalf("expected message 'is required', got %q", validationErrors[0].Message)
	}
}

func TestValidateRegisterUserRequest_MissingEmail(testingContext *testing.T) {
	registerRequest := &userhttp.RegisterUser{
		Name: "Lucas",
	}

	validationErrors := registerRequest.Validate()

	if len(validationErrors) != 1 {
		testingContext.Fatalf("expected 1 validation error, got %d", len(validationErrors))
	}
	if validationErrors[0].Field != "email" {
		testingContext.Fatalf("expected field 'email', got %q", validationErrors[0].Field)
	}
	if validationErrors[0].Message != "is required" {
		testingContext.Fatalf("expected message 'is required', got %q", validationErrors[0].Message)
	}
}

func TestValidateRegisterUserRequest_MissingNameAndEmail(testingContext *testing.T) {
	registerRequest := &userhttp.RegisterUser{}

	validationErrors := registerRequest.Validate()

	if len(validationErrors) != 2 {
		testingContext.Fatalf("expected 2 validation errors, got %d", len(validationErrors))
	}
	if validationErrors[0].Field != "name" {
		testingContext.Fatalf("expected first field 'name', got %q", validationErrors[0].Field)
	}
	if validationErrors[1].Field != "email" {
		testingContext.Fatalf("expected second field 'email', got %q", validationErrors[1].Field)
	}
}
