package main

import (
	"fmt"
	"video-api/database"
	"video-api/routers"
	"video-api/utils"
)

func main() {

	// 初始化 env
	var envService utils.EnvService
	envService.InitEnv()
	fmt.Println("參數初始化成功...")

	// 初始化 db
	database.InitDatabasePool()
	defer database.Mysql.Close()
	fmt.Println("資料庫始化成功...")

	routers.Router()
}
