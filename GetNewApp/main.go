package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/kimulaco/daily-steam-core/date"
)

func GetNewApp(w http.ResponseWriter, r *http.Request) {
	h := w.Header()
	h.Set("Content-Type", "application/json")

	var bodyParam struct {
		Message string `json:"message"`
	}
	err := json.NewDecoder(r.Body).Decode(&bodyParam)
	if err != nil {
		if err != io.EOF {
			ReturnErrorJson(w, 500, err)
			return
		}
	}

	res, err := Get(os.Getenv("SCRAPE_URL"))
	if err != nil {
		ReturnErrorJson(w, 500, err)
		return
	}
	defer res.Body.Close()

	apps, err := ParseNewApps(res, date.YesterdayDate())
	if err != nil {
		ReturnErrorJson(w, 500, err)
		return
	}

	body, err := json.Marshal(&SuccessJson{Apps: apps})
	if err != nil {
		ReturnErrorJson(w, 500, err)
		return
	}

	w.Write(body)
}
