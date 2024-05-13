package news

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

// 产品介绍
// 企业介绍
// 其他
type News struct {
}

// GetApiList
// @Tags      FactoryApi
// @Summary   产品信息展示
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.SearchGoodsParams                               true  "产品列表"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取产品列表,返回包括列表,总数,页码,每页数量"
// @Router    /factory/getGoodsList [post]
func (s *News) GetGoodsList(c *gin.Context) {
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
	list, total, err := service.FactoryService.GetGoodsList(model.Product, pageInfo.PageInfo, pageInfo.OrderKey, pageInfo.Desc)
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
