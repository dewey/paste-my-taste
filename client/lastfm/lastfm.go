package lastfm

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// Client contains everything needed for a Last.FM client
type Client struct {
	APIKey string
	client *http.Client
}

// New initialized a new Last.FM client
func New(c *http.Client, apiKey string) *Client {
	return &Client{
		APIKey: apiKey,
		client: c,
	}
}

// topArtists contains the top artists of a user
type topArtists struct {
	Topartists struct {
		Artist []struct {
			Attr struct {
				Rank int `json:"rank"`
			} `json:"@attr"`
			Mbid      string      `json:"mbid"`
			URL       string      `json:"url"`
			Playcount json.Number `json:"playcount"`
			Image     []struct {
				Size string `json:"size"`
				Text string `json:"#text"`
			} `json:"image"`
			Name string `json:"name"`
		} `json:"artist"`
		Attr struct {
			Page       int    `json:"page"`
			PerPage    int    `json:"perPage"`
			User       string `json:"user"`
			Total      int    `json:"total"`
			TotalPages int    `json:"totalPages"`
		} `json:"@attr"`
	} `json:"topartists"`
	Error   int    `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}

// weeklyArtist contains the top artists of a given timespan
type weeklyArtist struct {
	Weeklyartistchart struct {
		Artist []struct {
			Attr struct {
				Rank int `json:"rank"`
			} `json:"@attr"`
			Mbid      string      `json:"mbid"`
			Playcount json.Number `json:"playcount"`
			Name      string      `json:"name"`
			URL       string      `json:"url"`
		} `json:"artist"`
		Attr struct {
			Message    string `json:"message,omitempty"`
			Page       int    `json:"page"`
			To         int    `json:"to"`
			PerPage    int    `json:"perPage"`
			User       string `json:"user"`
			From       int    `json:"from"`
			Total      int    `json:"total"`
			TotalPages int    `json:"totalPages"`
		} `json:"@attr"`
	} `json:"weeklyartistchart"`
}

// topTags contains the top tags globally for an artist
type topTags struct {
	Toptags struct {
		Tag []struct {
			Count int    `json:"count"`
			Name  string `json:"name"`
			URL   string `json:"url"`
		} `json:"tag"`
		Attr struct {
			Artist string `json:"artist"`
		} `json:"@attr"`
	} `json:"toptags"`
}

// TopArtist is an artist and it's main genre
type TopArtist struct {
	Name      string `json:"name"`
	Playcount int    `json:"playcount"`
	Mbid      string `json:"mbid"`
	URL       string `json:"url"`
	Genre     string `json:"genre,omitempty"`
}

// GetTopArtists gets the top artist of a user from Last.FM
func (c *Client) GetTopArtists(username string, period string, limit int) ([]TopArtist, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("http://ws.audioscrobbler.com/2.0/?method=user.gettopartists&user=%s&api_key=%s&format=json&period=%s&limit=%d", username, c.APIKey, period, limit), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var ta topArtists
	if err := json.NewDecoder(resp.Body).Decode(&ta); err != nil {
		return nil, err
	}

	// Handle Last.FM API errors
	if ta.Error != 0 {
		return nil, errors.New(ta.Message)
	}

	var tal []TopArtist
	for _, artist := range ta.Topartists.Artist {
		pc, err := artist.Playcount.Int64()
		if err != nil {
			continue
		}
		tal = append(tal, TopArtist{
			Name:      artist.Name,
			Playcount: int(pc),
			Mbid:      artist.Mbid,
			URL:       artist.URL,
		})
	}

	return tal, nil
}

// GetWeeklyArtistChart returns the top artists from a given time span
func (c *Client) GetWeeklyArtistChart(username string, from, to int64, limit int) ([]TopArtist, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("http://ws.audioscrobbler.com/2.0/?method=user.getweeklyartistchart&user=%s&api_key=%s&format=json&from=%d&to=%d", username, c.APIKey, from, to), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var wa weeklyArtist
	if err := json.NewDecoder(resp.Body).Decode(&wa); err != nil {
		return nil, err
	}
	var tal []TopArtist
	for i, artist := range wa.Weeklyartistchart.Artist {
		if i <= limit {
			pc, err := artist.Playcount.Int64()
			if err != nil {
				continue
			}
			tal = append(tal, TopArtist{
				Name:      artist.Name,
				Playcount: int(pc),
				Mbid:      artist.Mbid,
				URL:       artist.URL,
			})
		}
	}

	return tal, nil
}

// GetTopTags gets the top tags for an artist from the Last.FM API
func (c *Client) GetTopTags(mbid string) ([]string, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("http://ws.audioscrobbler.com/2.0/?method=artist.gettoptags&mbid=%s&api_key=%s&format=json", mbid, c.APIKey), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var tt topTags
	if err := json.NewDecoder(resp.Body).Decode(&tt); err != nil {
		return nil, err
	}
	var ttl []string
	for _, tag := range tt.Toptags.Tag {
		ttl = append(ttl, tag.Name)
	}
	return ttl, nil
}
