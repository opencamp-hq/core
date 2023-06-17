package client

import (
	"encoding/json"
	"net/url"
	"strings"

	"github.com/opencamp-hq/core/models"
)

// ** HERE
// add the ability to search by campground ID, to get campground details
// https://www.recreation.gov/api/search?fq=entity_type:campground&fq=entity_id:233116
//
// Look for a go notification library, similar to the one you found in python
// then perhaps we can just let users get notified by whatever plethora of ways they want
//
// In the spirit of making the CLI the simplest to use and best out there ...
// - can we build in the ability to run as a daemon
// -- test what happens if a user closes their laptop lid / their machine sleeps?
// -- hopefully continue from where it left off ...
// - try out theres and see how it could be improved

const suggestEndpoint = "/search/suggest"

// TODO: Implement functionality to search by a park name (joshua tree) or area (big sur, ca)
// and return all campgrounds that match. Might be a bit tricky to do deterministically (ie: the user confirms we got the right park name), unless
// we break it out into multiple steps?
// Edit: Might be an API endpoint or query params that can help do a generalized search

func (c *Client) Suggest(query string) ([]*models.Campground, error) {
	c.log.Debug("Suggesting campgrounds", "query", query)

	qp := url.Values{}
	qp.Add("q", query)
	qp.Add("geocoder", "true")

	resp, err := c.Do(suggestEndpoint, qp)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var sr SuggestResponse
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

type SuggestResponse struct {
	Campgrounds []*models.Campground `json:"inventory_suggestions"`
}
