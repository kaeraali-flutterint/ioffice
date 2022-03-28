package schema

type Floor struct {
	Area             float64  `json:"area"`
	DateCreated      int64    `json:"dateCreated"`
	DrawingAvailable bool     `json:"drawingAvailable"`
	InteriorGross    float64  `json:"interiorGross"`
	Name             string   `json:"name"`
	LeaseArea        float64  `json:"leaseArea"`
	ID               int      `json:"id"`
	Building         Building `json:"building,omitempty"`
	DateUpdated      int64    `json:"dateUpdated,omitempty"`
}
