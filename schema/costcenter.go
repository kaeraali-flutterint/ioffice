package schema

type CostCenter struct {
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
	CostCenterParent *CostCenter `json:"costCenterParent"`
}
