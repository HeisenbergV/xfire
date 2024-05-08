package factory

import (
	"xfire/model/request"
	"xfire/model/response"

	"github.com/gin-gonic/gin"
)

// Drp
// @Tags      FactoryApi
// @Summary   进出货
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.RequestInventory true "库存"
// @Success   200   {object}  response.Response{msg=string}  "生产"
// @Router    /factory/production [post]
func (s *FactoryApi) AddInventory(c *gin.Context) {
	var req request.RequestInventory
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	/*
		1. 检查库存数量与进出数量是否冲突
		2. 增减库存
		3. 财务记录
		4. 记录日志
	*/

	// err = service.FactoryService.CreateBuild(req)
	// if err != nil {
	// 	global.LOG.Error("创建失败!", zap.Error(err))
	// 	response.FailWithMessage("创建失败", c)
	// 	return
	// }
	response.OkWithMessage("创建成功", c)
}
