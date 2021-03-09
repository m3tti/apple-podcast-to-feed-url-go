package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"
)

type Feed struct {
	FeedUrl string
}

type AppleApiResponse struct {
	Results []Feed
}

const applePodcastApiTemplate = "https://itunes.apple.com/lookup?id=%s&entity=podcast"

func getJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	json.NewDecoder(r.Body).Decode(target)

	return nil
}

func extractPodcastIdFromUrl(url string) string {
	extractIdRegex := regexp.MustCompile("id(\\d+)")
	return extractIdRegex.FindString(url)[2:]
}

func getApplePodcastUrl(id string) string {
	return fmt.Sprintf(applePodcastApiTemplate, id)
}

func hasAllParametersOrQuit(params []string) {
	if len(params) < 1 {
		panic("no args given")
	}
}

func fetchFeedsJson(url string) *AppleApiResponse {
	appleRespJson := new(AppleApiResponse)
	err := getJson(url, appleRespJson)

	if err != nil {
		panic("couldn't parse json")
	}

	return appleRespJson
}

func main() {
	args := os.Args[1:]

	hasAllParametersOrQuit(args)

	id := extractPodcastIdFromUrl(args[0])
	url := getApplePodcastUrl(id)
	appleRespJson := fetchFeedsJson(url)

	fmt.Println(appleRespJson.Results[0].FeedUrl)
}
