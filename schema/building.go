package schema

type Building struct {
	Address struct {
		Country struct {
			DefaultSelected         bool   `json:"defaultSelected"`
			SubdivisionCategoryName string `json:"subdivisionCategoryName"`
			Alpha2Code              string `json:"alpha2Code"`
			IsoCode                 string `json:"isoCode"`
			Name                    string `json:"name"`
			ID                      int    `json:"id"`
		} `json:"country"`
		City       string `json:"city"`
		Street     string `json:"street"`
		PostalCode string `json:"postalCode"`
		State      struct {
			Country struct {
				DefaultSelected         bool   `json:"defaultSelected"`
				SubdivisionCategoryName string `json:"subdivisionCategoryName"`
				Alpha2Code              string `json:"alpha2Code"`
				IsoCode                 string `json:"isoCode"`
				Name                    string `json:"name"`
				ID                      int    `json:"id"`
			} `json:"country"`
			DefaultSelected bool   `json:"defaultSelected"`
			Code            string `json:"code"`
			Name            string `json:"name"`
			ID              int    `json:"id"`
			CategoryName    string `json:"categoryName"`
		} `json:"state"`
		Street2 string `json:"street2"`
	} `json:"address,omitempty"`
	Code        string `json:"code"`
	DateCreated int64  `json:"dateCreated"`
	Metric      bool   `json:"metric"`
	Name        string `json:"name"`
	Location    struct {
	} `json:"location"`
	RevitLink   string        `json:"revitLink"`
	ID          int           `json:"id"`
	DateUpdated int64         `json:"dateUpdated"`
	CostCenters []interface{} `json:"costCenters"`
	LeaseArea   float64       `json:"leaseArea,omitempty"`
}
