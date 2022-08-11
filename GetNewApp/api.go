package function

import (
	"encoding/json"
	"net/http"

	"github.com/kimulaco/daily-steam-core/app"
)

type ErrorJson struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

type SuccessJson struct {
	Apps []app.App `json:"apps"`
}

func ReturnErrorJson(w http.ResponseWriter, code int, err error) {
	body, _ := json.Marshal(&ErrorJson{
		StatusCode: code,
		Message:    err.Error(),
	})
	w.Write(body)
}
