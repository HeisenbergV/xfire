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

// CreateApi
// @Tags      FactoryApi
// @Summary   创建工艺流程
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      model.Build true "工艺信息"
// @Success   200   {object}  response.Response{msg=string}  "创建工艺流程"
// @Router    /factory/createBuild [post]
func (s *FactoryApi) CreateBuild(c *gin.Context) {
	var build model.Build
	err := c.ShouldBindJSON(&build)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = service.FactoryService.CreateBuild(build)
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
// @Param     data  body      request.SearchBrandParams                               true  "分页获取品牌列表"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取品牌列表,返回包括列表,总数,页码,每页数量"
// @Router    /factory/getBuildList [post]
func (s *FactoryApi) GetBuildList(c *gin.Context) {
	var pageInfo request.SearchBuildParams
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
	list, total, err := service.FactoryService.GetBuildList(pageInfo.PageInfo, pageInfo.OrderKey, pageInfo.Desc)
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
// @Router    /api/deleteBuildByIds [delete]
func (s *FactoryApi) DeleteBuildByIds(c *gin.Context) {
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
// @Param     data  body      model.Build                  true "工艺信息"
// @Success   200   {object}  response.Response{msg=string}  "修改品牌"
// @Router    /api/updateBuild [post]
func (s *FactoryApi) UpdateBuild(c *gin.Context) {
	var api model.Build
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
	err = service.FactoryService.UpdateBuild(api)
	if err != nil {
		global.LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败", c)
		return
	}
	response.OkWithMessage("修改成功", c)
}
