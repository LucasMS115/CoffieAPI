package unit

import (
	"reflect"
	"testing"

	userhttp "coffie/internal/feature/user/http"
	"coffie/internal/http/response"
)

func TestValidateRegisterUserRequest(testingContext *testing.T) {
	testCases := []struct {
		name                string
		registerRequest     *userhttp.RegisterUser
		expectedFieldErrors []response.FieldError
	}{
		{
			name: "valid request",
			registerRequest: &userhttp.RegisterUser{
				Name:  "Lucas",
				Email: "lucas@email.com",
			},
			expectedFieldErrors: []response.FieldError{},
		},
		{
			name: "missing name",
			registerRequest: &userhttp.RegisterUser{
				Email: "lucas@email.com",
			},
			expectedFieldErrors: []response.FieldError{
				{Field: "name", Message: "is required"},
			},
		},
		{
			name: "missing email",
			registerRequest: &userhttp.RegisterUser{
				Name: "Lucas",
			},
			expectedFieldErrors: []response.FieldError{
				{Field: "email", Message: "is required"},
			},
		},
		{
			name:            "missing name and email",
			registerRequest: &userhttp.RegisterUser{},
			expectedFieldErrors: []response.FieldError{
				{Field: "name", Message: "is required"},
				{Field: "email", Message: "is required"},
			},
		},
	}

	for _, testCase := range testCases {
		testingContext.Run(testCase.name, func(testingContext *testing.T) {
			validationErrors := testCase.registerRequest.Validate()

			if !reflect.DeepEqual(validationErrors, testCase.expectedFieldErrors) {
				testingContext.Fatalf("expected validation errors %#v, got %#v", testCase.expectedFieldErrors, validationErrors)
			}
		})
	}
}
