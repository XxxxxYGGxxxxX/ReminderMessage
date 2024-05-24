package test

import (
	"ReminderMessage/service"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// 单元测试代码
func TestAddReminder(t *testing.T) {
	// 设置为测试模式
	gin.SetMode(gin.TestMode)
	// 配置路由
	r := gin.Default()
	r.POST("/addReminder", service.AddReminder)
	// 传值
	creatorID := "123"
	content := "Test"
	remindTime := "2024-05-23 23:00:00"
	body := gin.H{
		"CreatorID":  creatorID,
		"Content":    content,
		"RemindTime": remindTime,
	}
	// 序列化
	jsonValue, _ := json.Marshal(body)
	// 使用httptest包模拟http服务
	// 创建请求体
	req, _ := http.NewRequest(http.MethodPost, "/addReminder", bytes.NewBuffer(jsonValue))
	// 响应体
	w := httptest.NewRecorder()
	// 启动服务
	r.ServeHTTP(w, req)
	// 判断状态码相等不相等,等于200就是正常
	assert.Equal(t, http.StatusOK, w.Code)
	// 判断消息在不在里面
	assert.Contains(t, w.Body.String(), "提醒信息添加成功")
}

// 错误消息里有数据显示，好像能获取数据，但是不知道为啥报错？
func TestGetMyReminders(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/getMyReminders", service.GetMyReminders)
	req, _ := http.NewRequest(http.MethodGet, "/getMyReminders?CreatorID=123&ID=23", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), gin.H{
		"message": "提醒信息获取成功",
	})
}

func TestDeleteReminder(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.DELETE("/deleteReminder", service.DeleteReminder)
	req, _ := http.NewRequest(http.MethodDelete, "/deleteReminder?CreatorID=123&ID=22", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "提醒信息删除成功")
}

func TestUpdateReminder(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/updateReminder", service.UpdateReminder)
	creatorID := "123"
	id := "23"
	content := "测试更新数据"
	remindTime := "2024-05-23 23:35:00"
	body := gin.H{
		"Content":    content,
		"RemindTime": remindTime,
	}
	jsonValue, _ := json.Marshal(body)
	req, _ := http.NewRequest(http.MethodPost, "/updateReminder?CreatorID="+creatorID+"&ID="+id, bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "提醒信息更新成功")
}
