package utils

import (
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type randomResp struct {
	Frame struct {
		Episode   string `json:"Episode"`
		Timestamp int64  `json:"Timestamp"`
	} `json:"Frame"`
}

type searchItem struct {
	Episode   string `json:"Episode"`
	Timestamp int64  `json:"Timestamp"`
}

func SendFace(site string) string {
	resp, err := http.Get("https://" + site + ".com/api/random")
	if err != nil {
		return "Unable to open " + site + " site"
	}
	defer resp.Body.Close()

	b, _ := io.ReadAll(resp.Body)
	var rr randomResp
	if json.Unmarshal(b, &rr) != nil {
		return "Empty reply from " + site
	}

	img := "https://" + site + ".com/img/" + rr.Frame.Episode + "/" + strconv.FormatInt(rr.Frame.Timestamp, 10) + ".jpg"
	return img
}

func SendTagFace(site string, tag string) string {
	q := url.QueryEscape(strings.TrimSpace(tag))
	resp, err := http.Get("https://" + site + ".com/api/search?q=" + q)
	if err != nil {
		return "Unable to open " + site + " site"
	}
	defer resp.Body.Close()

	b, _ := io.ReadAll(resp.Body)
	var items []searchItem
	if json.Unmarshal(b, &items) != nil || len(items) == 0 {
		return "Didn't find anything from " + site + " with '" + tag + "' Searchterm"
	}

	pick := items[rand.Intn(len(items))]
	img := "https://" + site + ".com/img/" + pick.Episode + "/" + strconv.FormatInt(pick.Timestamp, 10) + ".jpg"
	return img
}
