package internal

import (
	"encoding/json"
	"fmt"
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

func Construct() error {
	data := models.Data{}
	if err := GetJson("https://groupietrackers.herokuapp.com/api", &data); err != nil {
		return fmt.Errorf("%s", err)
	}
	if err := GetJson(data.Artists, &Artists); err != nil {
		return fmt.Errorf("%s", err)
	}
	if err := GetJson(data.Locations, &Locations); err != nil {
		return fmt.Errorf("%s", err)
	}
	if err := GetJson(data.Dates, &Dates); err != nil {
		return fmt.Errorf("%s", err)
	}
	if err := GetJson(data.Relations, &Relations); err != nil {
		return fmt.Errorf("%s", err)
	}

	return nil
}
