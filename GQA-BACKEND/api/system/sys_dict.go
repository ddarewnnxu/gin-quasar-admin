package system

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"gin-quasar-admin/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApiDict struct {
}

func (a *ApiDict) GetDictList(c *gin.Context) {
	var pageInfo system.RequestPageWithParentId
	_ = c.ShouldBindJSON(&pageInfo)
	if err, dictList, total, parentId := service.GroupServiceApp.ServiceSystem.GetDictList(pageInfo); err != nil {
		global.GqaLog.Error("获取字典列表失败：", zap.Any("err", err))
		global.ErrorMessage("获取字典列表失败！", c)
	} else {
		global.SuccessData(system.ResponsePageWithParentId{
			List:     dictList,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
			Total:    total,
			ParentId: parentId,
		}, c)
	}
}

func (a *ApiDict) EditDict(c *gin.Context) {
	var toEditDict system.SysDict
	_ = c.ShouldBindJSON(&toEditDict)
	if err := service.GroupServiceApp.ServiceSystem.EditDict(toEditDict); err != nil {
		global.GqaLog.Error("编辑字典失败!", zap.Any("err", err))
		global.ErrorMessage("编辑字典失败！", c)
	} else {
		global.SuccessMessage("编辑字典成功！", c)
	}
}

func (a *ApiDict) AddDict(c *gin.Context) {
	var toAddDict system.RequestAddDict
	_ = c.ShouldBindJSON(&toAddDict)
	addDict := &system.SysDict{
		ParentId: toAddDict.ParentId,
		Value: toAddDict.Value,
		Label: toAddDict.Label,
	}
	if err := service.GroupServiceApp.ServiceSystem.AddDict(*addDict); err != nil {
		global.GqaLog.Error("添加字典失败！", zap.Any("err", err))
		global.ErrorMessage("添加字典失败！"+err.Error(), c)
	} else {
		global.SuccessMessage("添加字典成功！", c)
	}
}

func (a *ApiDict) DeleteDict(c *gin.Context) {
	var toDeleteId system.RequestDelete
	_ = c.ShouldBindJSON(&toDeleteId)
	if err := service.GroupServiceApp.ServiceSystem.DeleteDict(toDeleteId.Id); err != nil {
		global.GqaLog.Error("删除字典失败！", zap.Any("err", err))
		global.ErrorMessage("删除字典失败！", c)
	} else {
		global.SuccessMessage("删除字典成功！", c)
	}
}

func (a *ApiDict) QueryDictById(c *gin.Context) {
	var toQueryId system.RequestQueryById
	_ = c.ShouldBindJSON(&toQueryId)
	if err, dict := service.GroupServiceApp.ServiceSystem.QueryDictById(toQueryId.Id); err != nil {
		global.GqaLog.Error("查找字典失败！", zap.Any("err", err))
		global.ErrorMessage("查找字典失败！", c)
	} else {
		global.SuccessMessageData(gin.H{"info": dict}, "查找字典成功！", c)
	}
}

func (a *ApiDict) QueryDictByParentId(c *gin.Context) {
	var toQueryParentId system.RequestQueryByParentId
	_ = c.ShouldBindJSON(&toQueryParentId)
	if err, dict := service.GroupServiceApp.ServiceSystem.QueryDictByParentId(toQueryParentId.ParentId); err != nil {
		global.GqaLog.Error("查找字典失败！", zap.Any("err", err))
		global.ErrorMessage("查找字典失败！", c)
	} else {
		global.SuccessMessageData(gin.H{"info": dict}, "查找字典成功！", c)
	}
}