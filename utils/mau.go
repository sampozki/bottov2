package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

type MauImage struct {
	URL string `json:"url"`
}

type HauImage struct {
	URL string `json:"url"`
}

func Mau() string {
	response, err := http.Get("https://api.thecatapi.com/v1/images/search")
	if err != nil {
		return "Error with mau logic"
	}

	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)

	var data []MauImage
	if err := json.Unmarshal(body, &data); err != nil {
		return "Error with mau JSON magic"
	}

	if len(data) == 0 {
		return "Error with Mau length"
	}

	// Everything okay, return mau
	return data[0].URL
}

func Hau() string {
	response, err := http.Get("https://random.dog/woof.json")
	if err != nil {
		return "Error with hau logic"
	}
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)

	var data HauImage
	if err := json.Unmarshal(body, &data); err != nil {
		return "Error with hau JSON magic"
	}

	if data.URL == "" {
		return "Error with Hau length"
	}

	return data.URL
}
