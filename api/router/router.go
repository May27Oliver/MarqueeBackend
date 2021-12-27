package router

import (
	"MarqueeBackstage/api/apis"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
    router := gin.Default()
    //設定cors
	router.Use(cors.Default())
    //GET
    router.GET("getMarqueeSymbols",apis.GetMarqueeSymbols)
    router.GET("getSymbols/:groupname",apis.GetSymbols)
    router.GET("getSpeed",apis.GetSpeed)
    router.GET("getGroupName",apis.GetGroupName)

    router.POST("/addSymbol",apis.AddSymbol)
    router.POST("/importSymbols",apis.ImportSymbols)
    router.POST("/updateGroupNo",apis.UpdateGroupNo)
    router.POST("/updateSpeed",apis.UpdateSpeed)
    router.POST("/deleteSymbol", apis.DeleteSymbol)

    return router
}