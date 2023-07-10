package entities

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"time"

	errors_adapters "github.com/vm-championships-manager/backend-service/internal/errors/adapters"
	errors_protocols "github.com/vm-championships-manager/backend-service/internal/errors/protocols"
)

type User struct {
	Email     string `json:"email"`
	Name      string `json:"name"`
	LastName  string `json:"last_name"`
	Birthdate string `json:"birthdate,omitempty"`
	Phone     string `json:"phone,omitempty"`
}

func (u *User) Validate() errors_protocols.CustomError {
	errs := []string{}

	emailRegex := regexp.MustCompile(`^(?i)[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.Match([]byte(u.Email)) {
		errs = append(errs, "email")
	}

	nameRegex := regexp.MustCompile(`^(?i)[a-z]+$`)
	if !nameRegex.Match([]byte(u.Name)) {
		errs = append(errs, "name")
	}

	lastNameRegex := regexp.MustCompile(`^(?i)([a-z]+(\s?)){1,}$`)
	if !lastNameRegex.Match([]byte(u.LastName)) {
		errs = append(errs, "last_name")
	}

	if _, err := time.Parse(time.DateOnly, u.Birthdate); len(u.Birthdate) > 0 && err != nil {
		errs = append(errs, "birthdate")
	}

	phoneRegex := regexp.MustCompile(`^\+[0-9]+`)
	if len(u.Phone) > 0 && !phoneRegex.Match([]byte(u.Phone)) {
		errs = append(errs, "phone")
	}

	if len(errs) > 0 {
		sort.Strings(errs)
		return errors_adapters.NewEntityValidationError(fmt.Sprintf("[User]: invalid fields %s", strings.Join(errs, ", ")))
	}

	return nil
}
