package main

import (
	"MarqueeBackstage/api/database"
	model "MarqueeBackstage/api/models"
	"MarqueeBackstage/api/router"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql" //加载mysql
	"github.com/joho/godotenv"
)

func main() {
	defer database.Db.Close() //db關閉
	envErr :=godotenv.Load()
	if envErr != nil{
		panic("Failed to load env file")
}
	Port :=os.Getenv("API_PORT")
	listenPort := fmt.Sprintf(":%s",Port)
	// Migrate the schema
	database.Db.AutoMigrate(&model.Symbol{}, &model.Speed{}, &model.GroupName{})
	router := router.InitRouter()
	router.Run(listenPort)
}