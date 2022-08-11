package function

import (
	"github.com/kimulaco/daily-steam-core/apiutil"
	"github.com/kimulaco/daily-steam-core/app"
)

type SuccessJson struct {
	ReleasedAt string    `json:"releasedAt"`
	Apps       []app.App `json:"apps"`
}

func CreateError400(errorCode string, msg string) apiutil.Error {
	return apiutil.Error{
		StatusCode: 400,
		ErrorCode:  errorCode,
		Message:    msg,
	}
}

func CreateError500(errorCode string, msg string) apiutil.Error {
	if msg == "" {
		msg = "internal server error"
	}
	return apiutil.Error{
		StatusCode: 500,
		ErrorCode:  errorCode,
		Message:    msg,
	}
}
