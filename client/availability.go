package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/opencamp-hq/core/models"
)

const availEndpoint = "/camps/availability/campground/%s/month"

func (c *Client) Availability(campgroundID string, start, end time.Time) (models.Campsites, error) {
	startDate := start.Truncate(24 * time.Hour)
	endDate := end.Truncate(24 * time.Hour)
	if !endDate.After(startDate) {
		return nil, errors.New("End date is not after start date")
	}

	nowDate := time.Now().Truncate(24 * time.Hour)
	if nowDate.After(startDate) {
		return nil, errors.New("Start date occurs in the past")
	}

	// Determine months in date range.
	curPeriod := fmt.Sprintf("%d-%02d", start.Year(), start.Month())
	endPeriod := fmt.Sprintf("%d-%02d", end.Year(), end.Month())

	var months []string
	months = append(months, curPeriod)

	initial := start
	for curPeriod != endPeriod {
		start = start.AddDate(0, 1, 0)
		curPeriod = fmt.Sprintf("%d-%02d", start.Year(), start.Month())
		months = append(months, curPeriod)
	}
	start = initial

	// Build availability map.
	availabilities := make(map[string]map[string]*models.Campsite)
	for _, m := range months {
		campsites, err := c.monthlyAvailability(campgroundID, m)
		if err != nil {
			return nil, fmt.Errorf("Couldn't retrieve availabilities: %w", err)
		}

		for _, s := range campsites {
			for date, a := range s.Availabilities {
				if strings.ToLower(a) == "available" {
					if availabilities[s.Site] == nil {
						availabilities[s.Site] = make(map[string]*models.Campsite)
					}

					availabilities[s.Site][date] = s
				}
			}
		}
	}

	// Check for contiguous availability.
	var availableSites models.Campsites
Outer:
	for siteName, dates := range availabilities {
		var site *models.Campsite
		start = initial
		for !start.After(end) {
			date := fmt.Sprintf("%sT00:00:00Z", start.Format("2006-01-02"))
			if s, ok := dates[date]; ok {
				c.log.Debug(fmt.Sprintf("Site %s available for %s", siteName, start.Format("2006-01-02")))
				start = start.AddDate(0, 0, 1)

				site = s
			} else {
				continue Outer
			}
		}

		c.log.Debug(fmt.Sprintf("Site %s is fully available!", siteName))
		availableSites = append(availableSites, site)
	}

	sort.Sort(availableSites)
	return availableSites, nil
}

func (c *Client) monthlyAvailability(campgroundID string, month string) (map[string]*models.Campsite, error) {
	c.log.Debug("Checking campground availability", "campgroundID", campgroundID, "month", month)

	path := fmt.Sprintf(availEndpoint, campgroundID)
	qp := url.Values{}
	qp.Add("start_date", month+"-01T00:00:00.000Z")

	resp, err := c.Do(path, qp)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var ar AvailabilityResponse
	err = json.NewDecoder(resp.Body).Decode(&ar)
	if err != nil {
		return nil, err
	}

	return ar.Campsites, nil
}

type AvailabilityResponse struct {
	Campsites map[string]*models.Campsite `json:"campsites"`
	Count     int                         `json:"count"`
}
