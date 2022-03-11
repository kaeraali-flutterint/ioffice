package schema

type User struct {
	CostCenter1   string `json:"costCenter1"`
	AssistantName string `json:"assistantName"`
	JobTitle      string `json:"jobTitle"`
	PostalCode    string `json:"postalCode"`
	Custom02      string `json:"custom02"`
	Custom01      string `json:"custom01"`
	Custom04      string `json:"custom04"`
	Custom03      string `json:"custom03"`
	Custom06      string `json:"custom06"`
	KnownAs       string `json:"knownAs"`
	Custom05      string `json:"custom05"`
	Custom08      string `json:"custom08"`
	Custom07      string `json:"custom07"`
	Custom09      string `json:"custom09"`
	ID            int    `json:"id"`
	State         string `json:"state"`
	Fax           string `json:"fax"`
	CostCenter2   string `json:"costCenter2"`
	FloorWarden   bool   `json:"floorWarden"`
	CostCenter    struct {
		HexColor    string `json:"hexColor"`
		DateCreated int64  `json:"dateCreated"`
		Depth       struct {
			Code        string `json:"code"`
			DateCreated int64  `json:"dateCreated"`
			Level       int    `json:"level"`
			Name        string `json:"name"`
			ID          int    `json:"id"`
			Category    struct {
				CategoryType []struct {
					Name string `json:"name"`
					ID   int    `json:"id"`
				} `json:"categoryType"`
				Code        string `json:"code"`
				DateCreated int64  `json:"dateCreated"`
				Color       struct {
				} `json:"color"`
				Name string `json:"name"`
				ID   int    `json:"id"`
			} `json:"category"`
			DateUpdated int64 `json:"dateUpdated"`
		} `json:"depth"`
		Color struct {
		} `json:"color"`
		Name     string `json:"name"`
		ID       int    `json:"id"`
		Category struct {
			CategoryType []struct {
				Name string `json:"name"`
				ID   int    `json:"id"`
			} `json:"categoryType"`
			Code        string `json:"code"`
			DateCreated int64  `json:"dateCreated"`
			Color       struct {
			} `json:"color"`
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"category"`
		CostCenterParent struct {
			HexColor    string `json:"hexColor"`
			DateCreated int64  `json:"dateCreated"`
			Depth       struct {
				Code        string `json:"code"`
				DateCreated int64  `json:"dateCreated"`
				Level       int    `json:"level"`
				Name        string `json:"name"`
				ID          int    `json:"id"`
				Category    struct {
					CategoryType []struct {
						Name string `json:"name"`
						ID   int    `json:"id"`
					} `json:"categoryType"`
					Code        string `json:"code"`
					DateCreated int64  `json:"dateCreated"`
					Color       struct {
					} `json:"color"`
					Name string `json:"name"`
					ID   int    `json:"id"`
				} `json:"category"`
				DateUpdated int64 `json:"dateUpdated"`
			} `json:"depth"`
			Color struct {
			} `json:"color"`
			Name     string `json:"name"`
			ID       int    `json:"id"`
			Category struct {
				CategoryType []struct {
					Name string `json:"name"`
					ID   int    `json:"id"`
				} `json:"categoryType"`
				Code        string `json:"code"`
				DateCreated int64  `json:"dateCreated"`
				Color       struct {
				} `json:"color"`
				Name string `json:"name"`
				ID   int    `json:"id"`
			} `json:"category"`
			CostCenterParent struct {
				HexColor    string `json:"hexColor"`
				DateCreated int64  `json:"dateCreated"`
				Depth       struct {
					Code        string `json:"code"`
					DateCreated int64  `json:"dateCreated"`
					Level       int    `json:"level"`
					Name        string `json:"name"`
					ID          int    `json:"id"`
					Category    struct {
						CategoryType []struct {
							Name string `json:"name"`
							ID   int    `json:"id"`
						} `json:"categoryType"`
						Code        string `json:"code"`
						DateCreated int64  `json:"dateCreated"`
						Color       struct {
						} `json:"color"`
						Name string `json:"name"`
						ID   int    `json:"id"`
					} `json:"category"`
					DateUpdated int64 `json:"dateUpdated"`
				} `json:"depth"`
				Color struct {
				} `json:"color"`
				Name     string `json:"name"`
				ID       int    `json:"id"`
				Category struct {
					CategoryType []struct {
						Name string `json:"name"`
						ID   int    `json:"id"`
					} `json:"categoryType"`
					Code        string `json:"code"`
					DateCreated int64  `json:"dateCreated"`
					Color       struct {
					} `json:"color"`
					Name string `json:"name"`
					ID   int    `json:"id"`
				} `json:"category"`
				CostCenterParent struct {
					HexColor    string `json:"hexColor"`
					DateCreated int64  `json:"dateCreated"`
					Depth       struct {
						Code        string `json:"code"`
						DateCreated int64  `json:"dateCreated"`
						Level       int    `json:"level"`
						Name        string `json:"name"`
						ID          int    `json:"id"`
						Category    struct {
							CategoryType []struct {
								Name string `json:"name"`
								ID   int    `json:"id"`
							} `json:"categoryType"`
							Code        string `json:"code"`
							DateCreated int64  `json:"dateCreated"`
							Color       struct {
							} `json:"color"`
							Name string `json:"name"`
							ID   int    `json:"id"`
						} `json:"category"`
						DateUpdated int64 `json:"dateUpdated"`
					} `json:"depth"`
					Color struct {
					} `json:"color"`
					Name     string `json:"name"`
					ID       int    `json:"id"`
					Category struct {
						CategoryType []struct {
							Name string `json:"name"`
							ID   int    `json:"id"`
						} `json:"categoryType"`
						Code        string `json:"code"`
						DateCreated int64  `json:"dateCreated"`
						Color       struct {
						} `json:"color"`
						Name string `json:"name"`
						ID   int    `json:"id"`
					} `json:"category"`
				} `json:"costCenterParent"`
				DateUpdated int64 `json:"dateUpdated"`
			} `json:"costCenterParent"`
			DateUpdated int64 `json:"dateUpdated"`
		} `json:"costCenterParent"`
	} `json:"costCenter"`
	Custom30  string `json:"custom30"`
	FirstName string `json:"firstName"`
	Phone     string `json:"phone"`
	Name      string `json:"name"`
	UserType  struct {
		DateCreated int64  `json:"dateCreated"`
		Name        string `json:"name"`
		ID          int    `json:"id"`
		DateUpdated int64  `json:"dateUpdated"`
	} `json:"userType"`
	LastName     string `json:"lastName"`
	Extension    string `json:"extension"`
	Color        string `json:"color"`
	City         string `json:"city"`
	SpecialNeeds bool   `json:"specialNeeds"`
	Custom20     string `json:"custom20"`
	Custom22     string `json:"custom22"`
	Custom21     string `json:"custom21"`
	Custom24     string `json:"custom24"`
	Custom23     string `json:"custom23"`
	DateCreated  int64  `json:"dateCreated"`
	Custom26     string `json:"custom26"`
	Custom25     string `json:"custom25"`
	Custom28     string `json:"custom28"`
	Custom27     string `json:"custom27"`
	Custom29     string `json:"custom29"`
	Company      string `json:"company"`
	Department   string `json:"department"`
	Email        string `json:"email"`
	Address      string `json:"address"`
	Comments     string `json:"comments"`
	Mobile       string `json:"mobile"`
	EmployeeID   string `json:"employeeId"`
	UserName     string `json:"userName"`
	Room         struct {
		Area              float64 `json:"area"`
		DateCreated       int64   `json:"dateCreated"`
		ReservableByRules bool    `json:"reservableByRules"`
		Reservable        bool    `json:"reservable"`
		Name              string  `json:"name"`
		Description       string  `json:"description"`
		ID                int     `json:"id"`
		Type              struct {
			HexColor    string  `json:"hexColor"`
			ContentFlag int     `json:"contentFlag"`
			Cost        float64 `json:"cost"`
			DateCreated int64   `json:"dateCreated"`
			Color       struct {
			} `json:"color"`
			Name         string `json:"name"`
			Occupiable   bool   `json:"occupiable"`
			ID           int    `json:"id"`
			ParkingSpace bool   `json:"parkingSpace"`
			DateUpdated  int64  `json:"dateUpdated"`
			TypeCode     string `json:"typeCode"`
		} `json:"type"`
		Floor struct {
			Area             float64 `json:"area"`
			DateCreated      int64   `json:"dateCreated"`
			DrawingAvailable bool    `json:"drawingAvailable"`
			InteriorGross    float64 `json:"interiorGross"`
			Name             string  `json:"name"`
			LeaseArea        float64 `json:"leaseArea"`
			ID               int     `json:"id"`
			Building         struct {
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
				} `json:"address"`
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
			} `json:"building"`
			DateUpdated int64 `json:"dateUpdated"`
		} `json:"floor"`
		Occupied    bool  `json:"occupied"`
		Capacity    int   `json:"capacity"`
		DateUpdated int64 `json:"dateUpdated"`
	} `json:"room"`
	DateUpdated int64  `json:"dateUpdated"`
	Custom11    string `json:"custom11"`
	Custom10    string `json:"custom10"`
	Custom13    string `json:"custom13"`
	Custom12    string `json:"custom12"`
	Custom15    string `json:"custom15"`
	Custom14    string `json:"custom14"`
	Custom17    string `json:"custom17"`
	Custom16    string `json:"custom16"`
	Custom19    string `json:"custom19"`
	Custom18    string `json:"custom18"`
	MiddleName  string `json:"middleName"`
}
