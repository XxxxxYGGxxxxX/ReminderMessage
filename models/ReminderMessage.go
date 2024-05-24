package models

import "time"

// ReminderMessage 提醒信息结构体
type ReminderMessage struct {
	// 消息ID
	ID int `json:"id"`
	// 创建者ID
	CreatorID string `json:"creator_id"`
	// 提醒内容
	Content string `json:"content"`
	// 提醒时间
	RemindTime time.Time `json:"remind_time"`
}

func (table *ReminderMessage) TableName() string {
	return "reminder_message"
}
