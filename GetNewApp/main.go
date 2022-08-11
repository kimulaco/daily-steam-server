package function

import (
	"encoding/json"
	"io"
	"log"
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
			log.Println(err.Error())
			b, _ := CreateError400(
				"FAILED_DECODE_PARAMATER",
				"failed decode paramater",
			).ToBytes()
			w.Write(b)
			return
		}
	}

	res, err := Get(os.Getenv("NEW_APP_SCRAPE_URL"))
	if err != nil {
		log.Println(err.Error())
		b, _ := CreateError500("INTERNAL_SERVER_ERROR_001", "").ToBytes()
		w.Write(b)
		return
	}
	defer res.Body.Close()

	releasedAt := date.YesterdayDate()
	apps, err := ParseNewApps(res, releasedAt)
	if err != nil {
		log.Println(err.Error())
		b, _ := CreateError500("INTERNAL_SERVER_ERROR_002", "").ToBytes()
		w.Write(b)
		return
	}

	body, err := json.Marshal(&SuccessJson{
		ReleasedAt: releasedAt.ToString(),
		Apps:       apps,
	})
	if err != nil {
		log.Println(err.Error())
		b, _ := CreateError500("INTERNAL_SERVER_ERROR_003", "").ToBytes()
		w.Write(b)
		return
	}

	w.Write(body)
}
