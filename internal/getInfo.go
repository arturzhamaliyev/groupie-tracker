package internal

import (
	"encoding/json"
	"net/http"
	"time"

	"groupie_tracker/internal/models"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

func GetJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

var (
	Artists   []models.Artist
	Locations struct {
		Info []models.Locations `json:"index"`
	}

	Dates struct {
		Info []models.Dates `json:"index"`
	}

	Relations struct {
		Info []models.Relation `json:"index"`
	}
)

func Construct() {
	data := models.Data{}
	GetJson("https://groupietrackers.herokuapp.com/api", &data)
	GetJson(data.Artists, &Artists)
	GetJson(data.Locations, &Locations)
	GetJson(data.Dates, &Dates)
	GetJson(data.Relations, &Relations)
}
