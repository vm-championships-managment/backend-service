package entities_test

import (
	"sort"
	"strings"
	"testing"

	"github.com/vm-championships-manager/backend-service/internal/entities"
	internal_errors "github.com/vm-championships-manager/backend-service/internal/errors"
)

func TestUserValidation(t *testing.T) {
	tableTests := []struct {
		testName       string
		input          entities.User
		expectedOutput struct {
			result bool
			err    error
		}
	}{
		{
			"When all field are rigths",
			entities.User{Name: "test", LastName: "test test", Email: "test@test.com", Birthdate: "1900-01-01", Phone: "+5511999999999"},
			struct {
				result bool
				err    error
			}{true, nil},
		},
		{
			"When birthdate and phone was not given",
			entities.User{Name: "test", LastName: "test test", Email: "test@test.com"},
			struct {
				result bool
				err    error
			}{true, nil},
		},
		{
			"When name is not in fomat(only letter is accepted)",
			entities.User{Name: "1231231", LastName: "test test", Email: "test@test.com", Birthdate: "1900-01-01", Phone: "+5511999999999"},
			struct {
				result bool
				err    error
			}{false, internal_errors.EntityValidationError("name")},
		},
		{
			"When last name is not in fomat(only letter is accepted)",
			entities.User{Name: "test", LastName: "12313 12313", Email: "test@test.com", Birthdate: "1900-01-01", Phone: "+5511999999999"},
			struct {
				result bool
				err    error
			}{false, internal_errors.EntityValidationError("last_name")},
		},
		{
			"When email is not in fomat",
			entities.User{Name: "test", LastName: "test test", Email: "test@test", Birthdate: "1900-01-01", Phone: "+5511999999999"},
			struct {
				result bool
				err    error
			}{false, internal_errors.EntityValidationError("email")},
		},
		{
			"When birthdate is not valid",
			entities.User{Name: "test", LastName: "test test", Email: "test@test.com", Birthdate: "190-10-07", Phone: "+5511999999999"},
			struct {
				result bool
				err    error
			}{false, internal_errors.EntityValidationError("birthdate")},
		},
		{
			"When phone is not valid(plus symbol is required char)",
			entities.User{Name: "test", LastName: "test test", Email: "test@test.com", Birthdate: "1900-01-01", Phone: "5511999999999"},
			struct {
				result bool
				err    error
			}{false, internal_errors.EntityValidationError("phone")},
		},
		{
			"When all fields is not provided",
			entities.User{},
			struct {
				result bool
				err    error
			}{false, internal_errors.EntityValidationError(
				func() string {
					s := []string{"name", "last_name", "email"}
					sort.Strings(s)
					return strings.Join(s, ", ")
				}(),
			)},
		},
	}

	for _, r := range tableTests {
		func(currentTest struct {
			testName       string
			input          entities.User
			expectedOutput struct {
				result bool
				err    error
			}
		}) {
			t.Run(r.testName, func(t *testing.T) {
				t.Parallel()
				result, err := r.input.Validate()
				if r.expectedOutput.result != result {
					t.Errorf("%s: expected %v, received %v", r.testName, r.expectedOutput.result, result)
				}

				if r.expectedOutput.err != nil {
					if err == nil || r.expectedOutput.err.Error() != err.Error() {
						t.Errorf("%s: expected %v, received %v", r.testName, r.expectedOutput.err, err)
					}
				} else if r.expectedOutput.err != err {
					t.Errorf("%s: expected %v, received %v", r.testName, r.expectedOutput.err, err)
				}
			})
		}(r)

	}
}
