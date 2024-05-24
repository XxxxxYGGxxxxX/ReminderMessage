package router

import (
	"ReminderMessage/service"

	"github.com/gin-gonic/gin"
)

// 路由服务 注册路由
func Router() *gin.Engine {
	r := gin.Default()
	// 设置路由
	r.GET("/ping", service.Ping)                  //测试连接
	r.POST("/reminder", service.AddReminder)      // 增加提醒信息
	r.GET("/reminder", service.GetMyReminders)    // 本人的提醒信息列表
	r.DELETE("/reminder", service.DeleteReminder) // 删除提醒
	r.PUT("/reminder", service.UpdateReminder)    // 更改提醒时间或内容
	return r
}
