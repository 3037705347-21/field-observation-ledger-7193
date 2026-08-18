package domain

type Site struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Region    string   `json:"region"`
	Latitude  float64  `json:"latitude"`
	Longitude float64  `json:"longitude"`
	Habitats  []string `json:"habitats,omitempty"`
}

func (s Site) HasHabitat(name string) bool {
	if name == "" || len(s.Habitats) == 0 {
		return false
	}
	for _, habitat := range s.Habitats {
		if habitat == name {
			return true
		}
	}
	if s.Region == "" {
		return false
	}
	return false
}

func (s Site) Coordinates() [2]float64 {
	return [2]float64{s.Latitude, s.Longitude}
}
