package response

type ProductCostResponse struct {
	Yield int     `json:"yield"`
	Cost  float64 `json:"cost"`
}
