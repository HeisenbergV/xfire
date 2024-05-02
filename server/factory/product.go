package factory

import (
	"fmt"
	"xfire/global"
	"xfire/model"
	"xfire/model/request"
	"xfire/model/response"
	"xfire/service"
	"xfire/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreateApi
// @Tags      FactoryApi
// @Summary   创建产品
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      model.Product
// @Success   200   {object}  response.Response{msg=string}  "创建产品"
// @Router    /factory/createProduct[post]
func (s *FactoryApi) CreateProduct(c *gin.Context) {
	var product model.Product
	err := c.ShouldBindJSON(&product)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = service.FactoryService.CreateProduct(product)
	if err != nil {
		global.LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// GetApiList
// @Tags      FactoryApi
// @Summary   分页获取工艺信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.SearchProductParams                               true  "产品列表"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取产品列表,返回包括列表,总数,页码,每页数量"
// @Router    /factory/getProductList [post]
func (s *FactoryApi) GetProductList(c *gin.Context) {
	var pageInfo request.SearchProductParams
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo.PageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := service.FactoryService.GetProductList(pageInfo.PageInfo, pageInfo.OrderKey, pageInfo.Desc)
	if err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// DeleteBrandByIds
// @Tags      FactoryApi
// @Summary   删除选中brand
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.IdsReq                 true  "ID"
// @Success   200   {object}  response.Response{msg=string}  "删除"
// @Router    /api/DeleteProductByIds [delete]
func (s *FactoryApi) DeleteProductByIds(c *gin.Context) {
	var ids request.IdsReq
	err := c.ShouldBindJSON(&ids)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = service.FactoryService.DelBuild(ids.Ids)
	if err != nil {
		global.LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// UpdateApi
// @Tags      FactoryApi
// @Summary   修改基础api
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      model.Product                  true
// @Success   200   {object}  response.Response{msg=string}  "修改产品"
// @Router    /api/updateProduct [post]
func (s *FactoryApi) UpdateProduct(c *gin.Context) {
	var api model.Product
	err := c.ShouldBindJSON(&api)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(api, utils.ApiVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = service.FactoryService.UpdateProduct(api)
	if err != nil {
		global.LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败", c)
		return
	}
	response.OkWithMessage("修改成功", c)
}

// UpdateApi
// @Tags      FactoryApi
// @Summary   修改基础api
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById                 true
// @Success   200   {object}  response.Response{msg=string}  "展示成本"
// @Router    /api/showCost [post]
func (s *FactoryApi) ShowCost(c *gin.Context) {
	var id request.GetById
	err := c.ShouldBindJSON(&id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	product, err := service.FactoryService.GetProduct(id.ID)
	if err != nil {
		global.LOG.Error("成本展示失败!", zap.Error(err))
		response.FailWithMessage("成本展示失败", c)
		return
	}
	binfo, err := service.FactoryService.GetBuildInfo(product.BuildID)
	if err != nil {
		global.LOG.Error("获取配方失败!", zap.Error(err))
		response.FailWithMessage("获取配方失败", c)
		return
	}
	cost := 0.0
	sumWeight := 0.0

	for mname, ratio := range binfo.BuildData {
		material := global.GetMaterialInfo(mname)
		if material == nil {
			global.LOG.Error("获取配方失败!", zap.Error(err))
			response.FailWithMessage(fmt.Sprintf("%s不存在", mname), c)
			return
		}
		sumWeight += ratio
		//配方总成本
		cost += ratio / (material.Unit * 1000) * material.Price
	}
	// 配料总重量
	sumWeight += binfo.Water
	// 产出多少个面包
	yield := int(sumWeight / product.MaterialUnit)
	response.OkWithDetailed(response.ProductCostResponse{
		Yield: yield, Cost: cost,
	}, "获取成功", c)
}
