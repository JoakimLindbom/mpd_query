package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ArtistResponse struct {
	artist string `json:"artist"`
	NoRecords float64 `json:"no of records"`
	NoSongs float64 `json:"no of songs"`
}

type ArtistsResponse struct {
	artists []string `json:"artists"`
}

func Test_Get_StatusCodeShouldEqual200(t *testing.T) {
	expected := 200
	client := resty.New()
	resp, _ := client.R().Get("http://127.0.0.1:5555/api/v1/")
	if resp.StatusCode() != expected {
		t.Errorf("Unexpected status code, expected %d, got %d instead", expected, resp.StatusCode())
	}
}

func Test_Get_StatusCodeShouldEqual404(t *testing.T) {
	expected := 404
	client := resty.New()
	resp, _ := client.R().Get("http://127.0.0.1:5555/api/v1") // Malformed
	if resp.StatusCode() != expected {
		t.Errorf("Unexpected status code, expected %d, got %d instead", expected, resp.StatusCode())
	}
}

func Test_Get_ListOfArtists(t *testing.T) {
	expected := 200
	client := resty.New()
	resp, _ := client.R().Get("http://127.0.0.1:5555/api/v1/artists/")
	if resp.StatusCode() != expected {
		t.Errorf("Unexpected status code, expected %d, got %d instead", expected, resp.StatusCode())
	}
	myResponse := ArtistsResponse{}
	err := json.Unmarshal(resp.Body(), &myResponse)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(myResponse.artists)

	assert.Equal(t, "ARC", myResponse.artists)
}

func Test_Get_Artist(t *testing.T) {
	expected := 200
	client := resty.New()
	resp, _ := client.R().Get("http://127.0.0.1:5555/api/v1/artists/a")
	if resp.StatusCode() != expected {
		t.Errorf("Unexpected status code, expected %d, got %d instead", expected, resp.StatusCode())
	}
	myResponse := ArtistResponse{}
	err := json.Unmarshal(resp.Body(), &myResponse)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(myResponse.artist)

	assert.Equal(t, "Art of Noise", myResponse.artist)
	assert.Equal(t, 7, int(myResponse.NoRecords))
	assert.Equal(t, 35, int(myResponse.NoSongs))
}
