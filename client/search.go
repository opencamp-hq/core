package client

import (
	"encoding/json"
	"net/url"
	"strings"

	"github.com/opencamp-hq/core/models"
)

const searchEndpoint = "/search/suggest"

// TODO: Implement functionality to search by a park name (joshua tree) or area (big sur, ca)
// and return all campgrounds that match. Might be a bit tricky to do deterministically, unless
// we break it out into multiple steps?

func (c *Client) Search(query string) ([]models.Campground, error) {
	c.log.Debug("Searching for campgrounds", "query", query)

	qp := url.Values{}
	qp.Add("q", query)
	qp.Add("geocoder", "true")

	resp, err := c.Do(searchEndpoint, qp)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var sr SearchResponse
	err = json.NewDecoder(resp.Body).Decode(&sr)
	if err != nil {
		return nil, err
	}

	// Filter non-campgrounds from the search results.
	end := 0
	for i, c := range sr.Campgrounds {
		if strings.ToLower(c.EntityType) == "campground" {
			sr.Campgrounds[end] = sr.Campgrounds[i]
			sr.Campgrounds[end].Name = strings.Title(strings.ToLower(c.Name))
			end++
		}
	}

	sr.Campgrounds = sr.Campgrounds[:end]
	return sr.Campgrounds, nil
}

type SearchResponse struct {
	Campgrounds []models.Campground `json:"inventory_suggestions"`
}
