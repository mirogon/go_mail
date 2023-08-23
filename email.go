package mail

import (
	"encoding/json"
	"errors"
	"reflect"
	"regexp"
	"strings"

	util "github.com/mirogon/go_util"
)

type Email struct {
	email string
}

func (e Email) MarshalJSON() ([]byte, error) {
	emailField := reflect.ValueOf(&e).Elem().FieldByName("email")
	return json.Marshal(&struct {
		Email string `json:"email"`
	}{
		Email: emailField.String(),
	})
}

func (email Email) Str() string {
	return email.email
}

func CreateEmail(emailStr string) (Email, error) {
	emailStr = strings.ToLower(emailStr)
	email := Email{}
	match, _ := regexp.MatchString(util.EMAIL_REGEX, emailStr)

	if !match {
		return Email{}, errors.New("Invalid Email")
	}

	atSymbolIndex := util.Find(emailStr, "@")
	if atSymbolIndex < 1 {
		return Email{}, errors.New("Invalid Email")
	}

	dotIndex := util.Find(emailStr, ".")
	if dotIndex < 1 {
		return Email{}, errors.New("Invalid Email")
	}

	if dotIndex >= len(emailStr)-1 {
		return Email{}, errors.New("Invalid Email")
	}

	lastDotIndex := util.FindLast(emailStr, ".")
	if atSymbolIndex > lastDotIndex {
		return Email{}, errors.New("Invalid Email")
	}

	email.email = emailStr
	return email, nil
}
