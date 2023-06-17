package models

import (
	"encoding/json"
	"time"
)

type Campground struct {
	Activities []struct {
		ActivityDescription    string `json:"activity_description"`
		ActivityFeeDescription string `json:"activity_fee_description"`
		ActivityID             int    `json:"activity_id"`
		ActivityName           string `json:"activity_name"`
	} `json:"activities"`
	Addresses []struct {
		AddressType    string `json:"address_type"`
		City           string `json:"city"`
		CountryCode    string `json:"country_code"`
		PostalCode     string `json:"postal_code"`
		StateCode      string `json:"state_code"`
		StreetAddress1 string `json:"street_address1"`
		StreetAddress2 string `json:"street_address2"`
		StreetAddress3 string `json:"street_address3"`
	} `json:"addresses"`
	AggregateCellCoverage float64   `json:"aggregate_cell_coverage"`
	AverageRating         float64   `json:"average_rating"`
	CampsiteEquipmentName []string  `json:"campsite_equipment_name"`
	CampsiteReserveType   []string  `json:"campsite_reserve_type"`
	CampsiteTypeOfUse     []string  `json:"campsite_type_of_use"`
	CampsitesCount        string    `json:"campsites_count"`
	City                  string    `json:"city"`
	CountryCode           string    `json:"country_code"`
	Description           string    `json:"description"`
	Directions            string    `json:"directions"`
	EntityID              string    `json:"entity_id"`
	EntityType            string    `json:"entity_type"`
	IsInventory           bool      `json:"is_inventory"`
	GoLiveDate            time.Time `json:"go_live_date"`
	HTMLDescription       string    `json:"html_description"`
	ID                    string    `json:"id"`
	Latitude              string    `json:"latitude"`
	Links                 []struct {
		Description string `json:"description"`
		LinkType    string `json:"link_type"`
		Title       string `json:"title"`
		URL         string `json:"url"`
	} `json:"links"`
	Longitude string `json:"longitude"`
	Name      string `json:"name"`
	Notices   []struct {
		Text string `json:"text"`
		Type string `json:"type"`
	} `json:"notices"`
	NumberOfRatings int    `json:"number_of_ratings"`
	OrgID           string `json:"org_id"`
	OrgName         string `json:"org_name"`
	ParentID        string `json:"parent_id"`
	ParentName      string `json:"parent_name"`
	ParentType      string `json:"parent_type"`
	PreviewImageURL string `json:"preview_image_url"`
	PriceRange      struct {
		AmountMax int    `json:"amount_max"`
		AmountMin int    `json:"amount_min"`
		PerUnit   string `json:"per_unit"`
	} `json:"price_range"`
	Rate []struct {
		EndDate time.Time `json:"end_date"`
		Prices  []struct {
			Amount    int    `json:"amount"`
			Attribute string `json:"attribute"`
		} `json:"prices"`
		RateMap struct {
			PeakSTANDARDNONELECTRIC struct {
				GroupFees        any `json:"group_fees"`
				SingleAmountFees struct {
					Deposit   int `json:"deposit"`
					Holiday   int `json:"holiday"`
					PerNight  int `json:"per_night"`
					PerPerson int `json:"per_person"`
					Weekend   int `json:"weekend"`
				} `json:"single_amount_fees"`
			} `json:"PeakSTANDARD NONELECTRIC"`
		} `json:"rate_map"`
		SeasonDescription string    `json:"season_description"`
		SeasonType        string    `json:"season_type"`
		StartDate         time.Time `json:"start_date"`
	} `json:"rate"`
	Reservable bool   `json:"reservable"`
	StateCode  string `json:"state_code"`
	Text       string `json:"text"`
	TimeZone   string `json:"time_zone"`
	Type       string `json:"type"`
}

// Note: We use a custom UnmarshalJSON function here because the schema
// returned by '/search' and '/search/suggest' varies slightly, with the
// latter returning lat, lng, parent_entity_id, and parent_entity_type
// rather than the JSON field names defined in the JSON struct tags above.
func (c *Campground) UnmarshalJSON(data []byte) error {
	type Alias Campground

	var tmp struct {
		Alias
		ParentEntityID   string `json:"parent_entity_id"`
		ParentEntityType string `json:"parent_entity_type"`
		Lat              string `json:"lat"`
		Lng              string `json:"lng"`
	}

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*c = Campground(tmp.Alias)
	c.ParentID = tmp.ParentEntityID
	c.ParentType = tmp.ParentEntityType
	c.Latitude = tmp.Lat
	c.Longitude = tmp.Lng

	return nil
}
