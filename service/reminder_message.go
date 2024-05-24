package service

import (
	"ReminderMessage/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var DB = models.Init()

// 增
func AddReminder(c *gin.Context) {
	Creator_ID := c.PostForm("CreatorID")
	Content := c.PostForm("Content")
	RemindTime := c.PostForm("RemindTime")
	// 解析时间字符串
	loc, _ := time.LoadLocation("Local")
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", RemindTime, loc)
	// 执行插入SQL语句
	data := &models.ReminderMessage{
		CreatorID:  Creator_ID,
		Content:    Content,
		RemindTime: t,
	}
	err := models.DB.Create(data).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "添加提醒信息失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "提醒信息添加成功"})
}

// 查
// 获取本人的提醒信息列表
func GetMyReminders(c *gin.Context) {
	// 从URL参数中获取CreatorID
	CreatorID := c.Query("CreatorID")
	if CreatorID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "缺少CreatorID"})
		return
	}
	creatorID, err := strconv.Atoi(CreatorID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "CreatorID必须是整数"})
		return
	}
	// 校验创建者ID，此处假设用户ID
	userID := 123 // 假设的用户ID
	if creatorID != userID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "无权限操作他人的提醒信息"})
		return
	}
	// 执行查询SQL语句
	var reminders []models.ReminderMessage
	err = models.DB.Where("creator_id = ?", creatorID).Find(&reminders).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "获取提醒信息失败"})
		return
	}
	c.JSON(http.StatusOK, reminders)
}

// 删
func DeleteReminder(c *gin.Context) {
	// 从URL参数中获取CreatorID
	CreatorID := c.Query("CreatorID")
	if CreatorID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "缺少CreatorID"})
		return
	}
	// 将字符串类型的CreatorID转换为整型
	creatorID, err := strconv.Atoi(CreatorID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "CreatorID必须是整数"})
		return
	}
	// 校验创建者ID，此处假设用户ID
	userID := 123 // 假设的用户ID
	if creatorID != userID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "无权限操作他人的提醒信息"})
		return
	}
	// 获取要删除的提醒信息的ID
	id := c.Query("ID")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "缺少ID"})
		return
	}
	// 执行删除SQL语句
	err = models.DB.Where("id = ? AND creator_id = ?", id, creatorID).Delete(&models.ReminderMessage{}).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "删除提醒信息失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "提醒信息删除成功"})
}

// 改
func UpdateReminder(c *gin.Context) {
	// 从URL参数中获取CreatorID
	CreatorID := c.Query("CreatorID")
	if CreatorID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "缺少CreatorID"})
		return
	}
	// 将字符串类型的CreatorID转换为整型
	creatorID, err := strconv.Atoi(CreatorID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "CreatorID必须是整数"})
		return
	}
	// 校验创建者ID，此处假设用户ID
	userID := 123 // 假设的用户ID
	if creatorID != userID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "无权限操作他人的提醒信息"})
		return
	}
	// 获取要修改的提醒信息的ID
	id := c.Query("ID")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "缺少必要的参数ID"})
		return
	}
	var reminder models.ReminderMessage
	if err := c.ShouldBindJSON(&reminder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "格式错误"})
		return
	}
	// 执行更新SQL语句
	err = models.DB.Model(&models.ReminderMessage{}).Where("id = ? AND creator_id = ?", id, creatorID).Updates(models.ReminderMessage{Content: reminder.Content, RemindTime: reminder.RemindTime}).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "更新提醒信息失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "提醒信息更新成功"})

}
