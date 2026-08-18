package domain

type Location struct {
	SiteID    string  `json:"site_id"`
	Region    string  `json:"region"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func NewLocation(site Site) Location {
	return Location{
		SiteID: site.ID, Region: site.Region,
		Latitude: site.Latitude, Longitude: site.Longitude,
	}
}

func (l Location) IsValid() bool {
	return l.SiteID != "" && l.Region != "" &&
		l.Latitude >= -90 && l.Latitude <= 90 &&
		l.Longitude >= -180 && l.Longitude <= 180
}

func (l Location) DistanceHint(other Location) string {
	if l.Region == other.Region {
		return "same-region"
	}
	return "cross-region"
}
