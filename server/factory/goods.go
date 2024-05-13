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

// Production
// @Tags      FactoryApi
// @Summary   生产
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.RequestProduction true "生产信息"
// @Success   200   {object}  response.Response{msg=string}  "生产"
// @Router    /factory/production [post]
func (s *FactoryApi) Production(c *gin.Context) {
	var req request.RequestProduction
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	/*
		1. 检查此产品是否存在
		2. 消耗原料库存
		3. 增加货品库存
		4. 记录日志
	*/

	var gidList []int
	for k, _ := range req.Goods {
		gidList = append(gidList, k)
	}

	goods, err := service.FactoryService.GetGoods(gidList)
	if err != nil {
		global.LOG.Error("无效产品!", zap.Error(err))
		response.FailWithMessage("生产失败", c)
		return
	}

	for _, g := range goods {
		if g.Ptype != model.Product {
			global.LOG.Error("无效产品!", zap.Error(err))
			response.FailWithMessage("生产失败", c)
			return
		}

		if g.BuildID < 1 {
			global.LOG.Error("此产品缺失制作工艺!", zap.Error(err))
			response.FailWithMessage("此产品缺失制作工艺", c)
			return
		}
	}

	response.OkWithMessage("生产成功", c)
}

// UpdateApi
// @Tags      FactoryApi
// @Summary   修改基础api
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById                 true "id"
// @Success   200   {object}  response.Response{msg=string}  "展示成本"
// @Router    /api/showCost [post]
func (s *FactoryApi) ShowCost(c *gin.Context) {
	var id request.GetById
	err := c.ShouldBindJSON(&id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	Goods, err := service.FactoryService.GetGoodsInfo(id.ID)
	if err != nil {
		global.LOG.Error("成本展示失败!", zap.Error(err))
		response.FailWithMessage("成本展示失败", c)
		return
	}
	binfo, err := service.FactoryService.GetBuildInfo(Goods.BuildID)
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
		cost += ratio / material.Unit * material.Price
	}
	// 配料总重量
	sumWeight += binfo.Water
	// 产出多少个面包
	yield := int(sumWeight / Goods.Unit)
	response.OkWithDetailed(response.GoodsCostResponse{
		Yield: yield, Cost: cost,
	}, "获取成功", c)
}

// CreateApi
// @Tags      FactoryApi
// @Summary   创建产品
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      model.Goods true "产品信息"
// @Success   200   {object}  response.Response{msg=string}  "创建产品"
// @Router    /factory/createGoods [post]
func (s *FactoryApi) CreateGoods(c *gin.Context) {
	var Goods model.Goods
	err := c.ShouldBindJSON(&Goods)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = service.FactoryService.CreateGoods(Goods)
	if err != nil {
		global.LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// GetApiList
// @Tags      FactoryApi
// @Summary   产品信息展示
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.SearchGoodsParams                               true  "产品列表"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取产品列表,返回包括列表,总数,页码,每页数量"
// @Router    /factory/getGoodsList [post]
func (s *FactoryApi) GetGoodsList(c *gin.Context) {
	var pageInfo request.SearchGoodsParams
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
	list, total, err := service.FactoryService.GetGoodsList(pageInfo.Ptype, pageInfo.PageInfo, pageInfo.OrderKey, pageInfo.Desc)
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
// @Summary   删除
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.IdsReq                 true  "ID"
// @Success   200   {object}  response.Response{msg=string}  "删除"
// @Router    /api/DeleteGoodsByIds [delete]
func (s *FactoryApi) DeleteGoodsByIds(c *gin.Context) {
	var ids request.IdsReq
	err := c.ShouldBindJSON(&ids)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = service.FactoryService.DelGoods(ids.Ids)
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
// @Param     data  body      model.Goods                  true "产品信息"
// @Success   200   {object}  response.Response{msg=string}  "修改产品"
// @Router    /api/updateGoods [post]
func (s *FactoryApi) UpdateGoods(c *gin.Context) {
	var api model.Goods
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
	err = service.FactoryService.UpdateGoods(api)
	if err != nil {
		global.LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败", c)
		return
	}
	response.OkWithMessage("修改成功", c)
}
