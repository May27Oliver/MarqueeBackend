package main

import (
	"MarqueeBackstage/api/database"
	"MarqueeBackstage/api/router"

	_ "github.com/go-sql-driver/mysql" //加载mysql
)

//練習建立DB建立Table
//定義資料表
type Product struct {
	Code  string
	Price uint
}

type Symbols struct {
	GroupID int `gorm:"foreignKey:GroupID"`
	Symbol  string
	Show    bool
}

type Speed struct {
	Speed int
}

type GroupName struct {
	GroupID   int `gorm:"primary_key"`
	GroupName string
}

// func initMigrate(db *gorm.DB){
// 	db.AutoMigrate(&Product{},&Symbols{},&Speed{},&GroupName{})
// }

func main() {
	defer database.Db.Close() //db關閉

	// Migrate the schema
	database.Db.AutoMigrate(&Symbols{}, &Speed{}, &GroupName{})
	router := router.InitRouter()
	router.Run(":8888")

	// Create
	// db.Create(&Product{Code: "D42", Price: 100})

	// Read
	//var product Product
	//db.First(&product, 1) // find product with integer primary key
	//db.First(&product, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	//db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	// db.Delete(&product, 1)
}

// package main

// import (
// 	"MarqueeBackstage/api/database"
// 	"MarqueeBackstage/api/router"
// )

// func main() {
// 		defer database.Db.Close()//db關閉
// 		router := router.InitRouter()
// 		router.Run(":8888")
// }

// package main

// import (
// 	"net/http"

// 	"github.com/gin-contrib/cors"
// 	"github.com/gin-gonic/gin"
// )

// type IndexData struct{
// 	Title string
// 	Content string
// }

// func test(c *gin.Context){
// 	c.JSON(http.StatusOK,gin.H{
// 		"message":"回應回應你了",
// 	})
// }

// func main(){
// 	server := gin.Default()
// 	//設定cors
// 	server.Use(cors.Default())
// 	server.GET("/",test)
// 	server.Run(":8888")
// }

// import (
// 	"html/template"
// 	"log"
// 	"net/http"
// )

// type IndexData struct{
// 	Title string
// 	Content string
// }

// //request handler
// func test(w http.ResponseWriter, r *http.Request){
// 	tmpl := template.Must(template.ParseFiles("./index.html"))
// 	data :=new(IndexData)
// 	data.Title="首頁"
// 	data.Content="我的第一個"
// 	tmpl.Execute(w,data)
// 	// w.WriteHeader(http.StatusOK)
// 	// w.Write([]byte(`my first Website`))
// }

// func main(){
// 	http.HandleFunc("/",test)
// 	err:=http.ListenAndServe(":8888",nil)
// 	if err!= nil{
// 		log.Fatal("ListenAndServe",err)
// 	}
// }
