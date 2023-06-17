package client

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/opencamp-hq/core/models"
)

const searchEndpoint = "/search"

func (c *Client) SearchByID(campgroundID string) (*models.Campground, error) {
	c.log.Debug("Searching for campgrounds by ID", "campgroundID", campgroundID)

	qp := url.Values{}
	qp.Add("fq", "entity_type:campground")
	qp.Add("fq", "entity_id:"+campgroundID)

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
	if len(sr.Campgrounds) > 1 {
		return nil, fmt.Errorf("More than one campground returned (%d) for ID '%s'", sr.Size, campgroundID)
	}

	sr.Campgrounds[0].Name = strings.Title(strings.ToLower(sr.Campgrounds[0].Name))
	return sr.Campgrounds[0], nil
}

type SearchResponse struct {
	Campgrounds           []*models.Campground `json:"results"`
	Size                  int                  `json:"size"`
	SpellingAutocorrected bool                 `json:"spelling_autocorrected"`
	Start                 string               `json:"start"`
	Total                 int                  `json:"total"`
}
