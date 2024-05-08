package request

import (
	"xfire/model"
)

type SearchBrandParams struct {
	model.Brand
	PageInfo
	OrderKey string `json:"orderKey"` // 排序
	Desc     bool   `json:"desc"`     // 排序方式:升序false(默认)|降序true
}

type SearchBuildParams struct {
	model.Build
	PageInfo
	OrderKey string `json:"orderKey"` // 排序
	Desc     bool   `json:"desc"`     // 排序方式:升序false(默认)|降序true
}

type SearchGoodsParams struct {
	model.Goods
	PageInfo
	OrderKey string `json:"orderKey"` // 排序
	Desc     bool   `json:"desc"`     // 排序方式:升序false(默认)|降序true
}

type RequestProduction struct {
	Goods map[int]int `json:"gid"` // 货品id 数量
}

type RequestInventory struct {
	Id  int `json:"id"` // 货品id
	Num int `json:"num"`
}
