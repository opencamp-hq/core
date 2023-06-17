package models

type Campsite struct {
	Availabilities      map[string]string `json:"availabilities"`
	CampsiteID          string            `json:"campsite_id"`
	CampsiteReserveType string            `json:"campsite_reserve_type"`
	CampsiteRules       interface{}       `json:"campsite_rules"`
	CampsiteType        string            `json:"campsite_type"`
	CapacityRating      string            `json:"capacity_rating"`
	Loop                string            `json:"loop"`
	MaxNumPeople        int               `json:"max_num_people"`
	MinNumPeople        int               `json:"min_num_people"`
	Quantities          interface{}       `json:"quantities"`
	Site                string            `json:"site"`
	SupplementalCamping any               `json:"supplemental_camping"`
	TypeOfUse           string            `json:"type_of_use"`
}

type Campsites []*Campsite

// Implement the Sort interface.
func (cs Campsites) Len() int {
	return len(cs)
}
func (cs Campsites) Less(i, j int) bool {
	return cs[i].Site < cs[j].Site
}
func (cs Campsites) Swap(i, j int) {
	cs[i], cs[j] = cs[j], cs[i]
}
