package tiktok_api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	defaultHost = "https://www.tiktok.com"
	reqEmbedUri = "/oembed"
)

// Response contains all the embedded object properties
type Response struct {
	Version         string `json:"version"`
	Type            string `json:"type"`
	Title           string `json:"title"`
	AuthorUrl       string `json:"author_url"`
	AuthorName      string `json:"author_name"`
	Width           string `json:"width"`
	Height          string `json:"height"`
	Html            string `json:"html"`
	ThumbnailWidth  uint64 `json:"thumbnail_width"`
	ThumbnailHeight uint64 `json:"thumbnail_height"`
	ThumbnailUrl    string `json:"thumbnail_url"`
	ProviderUrl     string `json:"provider_url"`
	ProviderName    string `json:"provider_name"`
}

// TikTokService encapsulates settings for TikTok api calls
type TikTokService struct {
	host string
}

// NewTikTokService creates TikTokService with default settings
func NewTikTokService() *TikTokService {
	return &TikTokService{host: defaultHost}
}

// Embed allows you to get the embed code and additional information about the video associated with the webpage link provided
func (s *TikTokService) Embed(params map[string]string) (*Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, s.host+reqEmbedUri, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	for param, val := range params {
		q.Add(param, val)
	}

	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	r := &Response{}
	err = json.Unmarshal(respBody, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}
