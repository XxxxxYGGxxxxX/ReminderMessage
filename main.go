package main

import (
	"ReminderMessage/models"
	"ReminderMessage/router"
)

// 主函数 程序入口
func main() {
	models.Init()
	r := router.Router()
	r.Run(":8888")
}
