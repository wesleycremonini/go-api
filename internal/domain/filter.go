package domain

type Filter struct {
	Limit  *int `json:"results_limit"`
	Offset *int `json:"results_offset"`
}
