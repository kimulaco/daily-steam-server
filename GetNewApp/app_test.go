package function

import (
	"testing"
)

func TestCreateError400(t *testing.T) {
	e := CreateError400("ERROR_CODE", "error message")
	if e.StatusCode != 400 ||
		e.ErrorCode != "ERROR_CODE" ||
		e.Message != "error message" {
		t.Error("invalid error data")
	}
}

func TestCreateError500(t *testing.T) {
	e := CreateError500("ERROR_CODE", "error message")
	if e.StatusCode != 500 ||
		e.ErrorCode != "ERROR_CODE" ||
		e.Message != "error message" {
		t.Error("invalid error data")
	}
}
