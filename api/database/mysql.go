package database

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql" //加载mysql
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var Db *gorm.DB

func init() { 
    envErr :=godotenv.Load()
    if envErr != nil{
        panic("Failed to load env file")
    }
    dbUser:=os.Getenv("DB_USER")
    dbPass:=os.Getenv("DB_PASS")
    dbHost:=os.Getenv("DB_HOST")
    dbName:=os.Getenv("DB_NAME")
    
    dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local&timeout=10ms",dbUser,dbPass,dbHost,dbName)
    var err error
    Db, err = gorm.Open("mysql", dsn)

    if err != nil {
        fmt.Printf("mysql connect error %v", err)
    }

    if Db.Error != nil {
        fmt.Printf("database error %v", Db.Error)
    }
}

