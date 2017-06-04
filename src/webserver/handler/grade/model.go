package grade

type profileSummaryResponse struct {
	CourseName string             `json:"name"`
	Parameter  []profileParameter `json:"parameter"`
	UCU        int8               `json:"ucu"`
}

type profileParameter struct {
	Name       string  `json:"name"`
	Percentage float32 `json:"percentage"`
}
