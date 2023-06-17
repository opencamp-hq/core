package client

import (
	"context"
	"time"

	"github.com/opencamp-hq/core/models"
)

// Poll is a blocking operation. To poll multiple campgrounds call this method
// in its own goroutine.
func (c *Client) Poll(ctx context.Context, campgroundID string, start, end time.Time, interval time.Duration) (models.Campsites, error) {
	sites, err := c.Availability(campgroundID, start, end)
	if err != nil {
		return nil, err
	}
	if len(sites) > 0 {
		return sites, nil
	}

	c.log.Info("No sites available at the moment, starting polling!", "interval", interval)
	t := time.NewTicker(interval)
	for {
		select {
		case <-t.C:
			sites, err := c.Availability(campgroundID, start, end)
			if err != nil {
				return nil, err
			}

			if len(sites) > 0 {
				return sites, nil
			}
			c.log.Info("Sorry, no available campsites were found for your dates. We'll try again")
		case <-ctx.Done():
			return nil, nil
		}
	}
}
