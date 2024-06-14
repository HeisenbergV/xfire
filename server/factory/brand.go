package factory

import (
	"xfire/global"
	"xfire/model"
	"xfire/model/request"
	"xfire/model/response"
	"xfire/service"
	"xfire/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FactoryApi struct {
}

// CreateApi
// @Tags      FactoryApi
// @Summary   创建品牌
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      model.Brand true "品牌信息"
// @Success   200   {object}  response.Response{msg=string}  "创建品牌"
// @Router    /factory/createBrand [post]
func (s *FactoryApi) CreateBrand(c *gin.Context) {
	var brand model.Brand
	err := c.ShouldBindJSON(&brand)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = service.FactoryService.CreateBrand(brand)
	if err != nil {
		global.LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// GetApiList
// @Tags      FactoryApi
// @Summary   分页获取品牌列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.SearchBrandParams                               true  "分页获取品牌列表"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取品牌列表,返回包括列表,总数,页码,每页数量"
// @Router    /factory/getBrandList [post]
func (s *FactoryApi) GetBrandList(c *gin.Context) {
	var pageInfo request.SearchBrandParams
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
	list, total, err := service.FactoryService.GetBrandList(pageInfo.PageInfo, pageInfo.OrderKey, pageInfo.Desc)
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
// @Success   200   {object}  response.Response{msg=string}  "删除选中品牌"
// @Router    /factory/deleteBrandByIds [delete]
func (s *FactoryApi) DeleteBrandByIds(c *gin.Context) {
	var ids request.IdsReq
	err := c.ShouldBindJSON(&ids)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = service.FactoryService.DelBrand(ids.Ids)
	if err != nil {
		global.LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// UpdateApi
// @Tags      FactoryApi
// @Summary   修改品牌
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      model.Brand                  true "品牌信息"
// @Success   200   {object}  response.Response{msg=string}  "修改品牌"
// @Router    /factory/updateBrand [post]
func (s *FactoryApi) UpdateBrand(c *gin.Context) {
	var api model.Brand
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
	err = service.FactoryService.UpdateBrand(api)
	if err != nil {
		global.LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败", c)
		return
	}
	response.OkWithMessage("修改成功", c)
}
