package router

import (
	"MarqueeBackstage/api/apis"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var domain string = os.Getenv("DOMAIN")

func InitRouter() *gin.Engine {
	router := gin.Default()
	//設定cors
	router.Use(Cors())

	//GET
	{
		elecScrollAPIGroup := router.Group("elecScrollAPI")
		elecScrollAPIGroup.GET("getMarqueeSymbols", apis.GetMarqueeSymbols)
		elecScrollAPIGroup.GET("getSymbols/:groupname", apis.GetSymbols)
		elecScrollAPIGroup.GET("getSpeed", apis.GetSpeed)
		elecScrollAPIGroup.GET("getGroupName", apis.GetGroupName)
		elecScrollAPIGroup.GET("getDirection",apis.GetDirection)

		elecScrollAPIGroup.POST("addSymbol", apis.AddSymbol)
		elecScrollAPIGroup.POST("importSymbols", apis.ImportSymbols)
		elecScrollAPIGroup.POST("updateGroupNo", apis.UpdateGroupNo)
		elecScrollAPIGroup.POST("updateSpeed", apis.UpdateSpeed)
		elecScrollAPIGroup.POST("updateDirection", apis.UpdateDirection)
		elecScrollAPIGroup.POST("deleteSymbol", apis.DeleteSymbol)
		elecScrollAPIGroup.POST("login", apis.Login)
		elecScrollAPIGroup.POST("verifyLogin", apis.VerifyLogin)
	}
	return router
}

func Cors() gin.HandlerFunc{
		return func (c *gin.Context){
			method := c.Request.Method
			//origin := c.Request.Header.Get("Origin") //header setting
			//接收客戶端傳送的origin （重要！）
			c.Writer.Header().Set("Access-Control-Allow-Origin", domain) 
			//伺服器支援的所有跨域請求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS") 
			//允許跨域設定可以返回其他子段，可以自定義欄位
			c.Header("Access-Control-Allow-Headers", "Content-type , eAuthorization, Content-Length, X-CSRF-Token, Token,session")
			// 允許瀏覽器（客戶端）可以解析的頭部 （重要）
			c.Header("Access-Control-Expose-Headers", "Content-type ,Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers") 
			//設定快取時間
			c.Header("Access-Control-Max-Age", "172800") 
			//允許客戶端傳遞校驗資訊比如 cookie (重要)
			c.Header("Access-Control-Allow-Credentials", "true")                                                                                                                                                                                                                          
			

			//允許型別校驗 
			if method == "OPTIONS" {
					// c.JSON(http.StatusOK, "ok!")
					c.AbortWithStatus(http.StatusNoContent)
			}

			defer func() {
					if err := recover(); err != nil {
							log.Printf("Panic info is: %v", err)
					}
			}()

			c.Next()
		}
}
