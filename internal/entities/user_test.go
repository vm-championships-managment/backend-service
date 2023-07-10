package entities_test

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/vm-championships-manager/backend-service/internal/entities"
	errors_adapters "github.com/vm-championships-manager/backend-service/internal/errors/adapters"
	errors_protocols "github.com/vm-championships-manager/backend-service/internal/errors/protocols"
)

func TestUserValidation(t *testing.T) {
	usrMsgErrorFmt := func(message string) string {
		return fmt.Sprintf("[User]: invalid fields %s", message)
	}

	tableTests := []struct {
		testName       string
		input          entities.User
		expectedOutput struct {
			result bool
			err    errors_protocols.CustomError
		}
	}{
		{
			"When all field are rigths",
			entities.User{Name: "test", LastName: "test test", Email: "test@test.com", Birthdate: "1900-01-01", Phone: "+5511999999999"},
			struct {
				result bool
				err    errors_protocols.CustomError
			}{true, nil},
		},
		{
			"When birthdate and phone was not given",
			entities.User{Name: "test", LastName: "test test", Email: "test@test.com"},
			struct {
				result bool
				err    errors_protocols.CustomError
			}{true, nil},
		},
		{
			"When name is not in fomat(only letter is accepted)",
			entities.User{Name: "1231231", LastName: "test test", Email: "test@test.com", Birthdate: "1900-01-01", Phone: "+5511999999999"},
			struct {
				result bool
				err    errors_protocols.CustomError
			}{false, errors_adapters.NewEntityValidationError(usrMsgErrorFmt("name"))},
		},
		{
			"When last name is not in fomat(only letter is accepted)",
			entities.User{Name: "test", LastName: "12313 12313", Email: "test@test.com", Birthdate: "1900-01-01", Phone: "+5511999999999"},
			struct {
				result bool
				err    errors_protocols.CustomError
			}{false, errors_adapters.NewEntityValidationError(usrMsgErrorFmt("last_name"))},
		},
		{
			"When email is not in fomat",
			entities.User{Name: "test", LastName: "test test", Email: "test@test", Birthdate: "1900-01-01", Phone: "+5511999999999"},
			struct {
				result bool
				err    errors_protocols.CustomError
			}{false, errors_adapters.NewEntityValidationError(usrMsgErrorFmt("email"))},
		},
		{
			"When birthdate is not valid",
			entities.User{Name: "test", LastName: "test test", Email: "test@test.com", Birthdate: "190-10-07", Phone: "+5511999999999"},
			struct {
				result bool
				err    errors_protocols.CustomError
			}{false, errors_adapters.NewEntityValidationError(usrMsgErrorFmt("birthdate"))},
		},
		{
			"When phone is not valid(plus symbol is required char)",
			entities.User{Name: "test", LastName: "test test", Email: "test@test.com", Birthdate: "1900-01-01", Phone: "5511999999999"},
			struct {
				result bool
				err    errors_protocols.CustomError
			}{false, errors_adapters.NewEntityValidationError(usrMsgErrorFmt("phone"))},
		},
		{
			"When all fields is not provided",
			entities.User{},
			struct {
				result bool
				err    errors_protocols.CustomError
			}{false, errors_adapters.NewEntityValidationError(
				func() string {
					s := []string{"name", "last_name", "email"}
					sort.Strings(s)
					return usrMsgErrorFmt(strings.Join(s, ", "))
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
				err    errors_protocols.CustomError
			}
		}) {
			t.Run(r.testName, func(t *testing.T) {
				t.Parallel()
				err := r.input.Validate()

				if r.expectedOutput.err != nil {
					if err == nil || r.expectedOutput.err.GetErrorInfos() != err.GetErrorInfos() {
						t.Errorf("%s: expected %v, received %v", r.testName, r.expectedOutput.err, err)
					}
				} else if r.expectedOutput.err != err {
					t.Errorf("%s: expected %v, received %v", r.testName, r.expectedOutput.err, err)
				}
			})
		}(r)

	}
}
