package response

import "xfire/model"

type GoodsCostResponse struct {
	Yield     int         `json:"yield"`
	Cost      float64     `json:"cost"`
	BuildInfo model.Build `json:"build_info"`
}
