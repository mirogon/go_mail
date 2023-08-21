package mail_test

import (
	"testing"

	mail "github.com/mirogon/go_mail"
)

func TestValidEmail(t *testing.T) {
	email, err := mail.CreateEmail("m1smr@hotmail.com")
	if err != nil {
		t.Error()
	}
	if email.Str() != "m1smr@hotmail.com" {
		t.Error()
	}
}

func TestEmailWithoutAtSymbolReturnsError(t *testing.T) {
	_, err := mail.CreateEmail("m1smrhotmail.com")
	if err == nil {
		t.Error()
	}
}

func TestEmailWithoutDotReturnsError(t *testing.T) {
	_, err := mail.CreateEmail("m1smr@hotmailcom")
	if err == nil {
		t.Error()
	}
}

func TestEmailWithNoTextInFrontOfAtSymbolReturnsError(t *testing.T) {
	_, err := mail.CreateEmail("@hotmail.com")
	if err == nil {
		t.Error()
	}
}

func TestEmailWithNoTextInFronOfDotReturnsError(t *testing.T) {
	_, err := mail.CreateEmail(".@")
	if err == nil {
		t.Error()
	}
}

func TestEmailWithNoTextAfterLastDotReturnsError(t *testing.T) {
	_, err := mail.CreateEmail("m1smr@hotmail.")
	if err == nil {
		t.Error()
	}
}

func TestEmailLastDotInFrontOfAtSymbolReturnsError(t *testing.T) {
	_, err := mail.CreateEmail("m1smr.hotmail@")
	if err == nil {
		t.Error()
	}
}

func TestEmail_ForbiddenSymbols_ReturnsError(t *testing.T) {
	_, err := mail.CreateEmail("你好好好@hotmail.com")
	if err == nil {
		t.Error()
	}
}
