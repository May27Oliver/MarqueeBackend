package router

import (
	"MarqueeBackstage/api/apis"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
    router := gin.Default()
    // corsConfig := cors.DefaultConfig()
    // corsConfig.AllowAllOrigins = true
    // corsConfig.AllowMethods = []string{"GET", "POST"}
    // corsConfig.AllowHeaders = []string{"Authorization", "Content-Type", "Upgrade", "Origin",
    //     "Connection", "Accept-Encoding", "Accept-Language", "Host", "Access-Control-Request-Method", "Access-Control-Request-Headers"}
    // corsConfig.MaxAge = 12 * time.Hour
    // router.Use(cors.New(corsConfig))
        //設定cors
	router.Use(cors.Default())
    //GET
    {
        elecScrollAPIGroup:= router.Group("elecScrollAPI")
        elecScrollAPIGroup.GET("getMarqueeSymbols",apis.GetMarqueeSymbols)
        elecScrollAPIGroup.GET("getSymbols/:groupname",apis.GetSymbols)
        elecScrollAPIGroup.GET("getSpeed",apis.GetSpeed)
        elecScrollAPIGroup.GET("getGroupName",apis.GetGroupName)
    
        elecScrollAPIGroup.POST("addSymbol",apis.AddSymbol)
        elecScrollAPIGroup.POST("importSymbols",apis.ImportSymbols)
        elecScrollAPIGroup.POST("updateGroupNo",apis.UpdateGroupNo)
        elecScrollAPIGroup.POST("updateSpeed",apis.UpdateSpeed)
        elecScrollAPIGroup.POST("deleteSymbol", apis.DeleteSymbol)
    }
    return router
}