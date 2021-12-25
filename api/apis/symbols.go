package apis

import (
	"MarqueeBackstage/api/database"
	model "MarqueeBackstage/api/models"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type symbol struct {
	StockNo string `json:stockNo`
	StockName string `json:stockName`
	Symbol string  `json:symbol`
	MarqueeOrder int `json:marqueeOrder`
}
type addDeleteRequest struct{
	GroupId int `json:groupId`
	Symbol symbol `json:symbol`
}
type importRequest struct{
	GroupId int `json:groupId`
	Symbols []symbol `json:symbol`
}

type updateGroupRequest struct{
	GroupId int `json:groupId`
}

type updateSpeedRequest struct{
	Speed int `json:speed`
}

//取得群組
func GetSymbols(c *gin.Context){
	var symbol model.Symbol
	paramInfo := c.Param("groupname")
	groupname := strings.Split(paramInfo,"=")[1]
	
	//讀取使用者request
	result,err := symbol.SymbolQuery(groupname)
	if err!= nil{
		c.JSON(http.StatusOK,gin.H{
			"result":false,
			"message":"查無資料",
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"result":true,
		"data":result,
	})
}
//取得速度
func GetSpeed(c *gin.Context){

}
//取得群組名稱
func GetGroupName(c *gin.Context){
	var groupname model.GroupName
	result,err := groupname.GroupNameQuery()
	if err!= nil{
		c.JSON(http.StatusOK,gin.H{
			"result":false,
			"message":"查無資料",
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"result":true,
		"data":result,
	})
} 
//增加群組symbol
func AddSymbol(c *gin.Context){
	req:= addDeleteRequest{}
	if err:=c.BindJSON(&req); err!=nil{
		c.String(http.StatusPaymentRequired,err.Error())
		return
	}
	req_data := model.Symbol{GroupID: req.GroupId,Symbol: req.Symbol.Symbol,StockName:req.Symbol.StockName,Show:true,MarqueeOrder:req.Symbol.MarqueeOrder}
	// log.Printf("Request data: %s" , req_data)
	result:=database.Db.Debug().Create(&req_data)
	
	if result.Error!=nil{
		c.JSON(http.StatusOK,gin.H{
			"result":false,
			"message":"新增失敗",
		})
	}
	c.JSON(http.StatusOK,gin.H{
		"result":true,
		"message":"新增成功",
	})
	return
}
// 匯入csv
func ImportSymbols(c *gin.Context){
	req:=importRequest{}
	if err:=c.BindJSON(&req);err!=nil{
		c.String(http.StatusPaymentRequired,err.Error())
		return
	}

	sqlStr:="INSERT INTO `symbols` (group_id,symbol,`show`,stock_name,marquee_order) VALUES "
	
	connInfo :=[]interface{}{}

	for _,v := range req.Symbols { 
		log.Printf("Request data: %s" , req.Symbols)
		sqlStr += "(?, ?, ?, ?, ?),"
		connInfo = append(connInfo,req.GroupId,v.Symbol,true,v.StockName,v.MarqueeOrder)
	}

	sqlStr = sqlStr[0:len(sqlStr)-1]+";"
	result:= database.Db.Debug().Exec(sqlStr,connInfo...)
	
	if result.Error!=nil{
		c.JSON(http.StatusOK,gin.H{
			"result":false,
			"message":result.Error,
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"result":true,
		"message":"新增成功",
	})
}
// 更新播放群組
func UpdateGroupNo(c *gin.Context){
	database.Db.Model(&model.GroupName{}).Update("play",false);
	req:= updateGroupRequest{}
	if err := c.BindJSON(&req); err != nil {
		c.String(http.StatusPaymentRequired,err.Error())
		return
	}
	database.Db.Model(&model.GroupName{}).Where("group_id = ?",req.GroupId).Update("play",true)
}

func UpdateSpeed(c *gin.Context){
	database.Db.Model(&model.Speed{}).Update("use",false);
	req:= updateSpeedRequest{}
	if err := c.BindJSON(&req); err != nil {
		c.String(http.StatusPaymentRequired,err.Error())
		return
	}
	database.Db.Model(&model.Speed{}).Where("speed = ?",req.Speed).Update("play",true)
}

//刪除symbol
func DeleteSymbol(c *gin.Context){
	req:=addDeleteRequest{}
	if err:=c.BindJSON(&req);err!=nil{
		c.String(http.StatusPaymentRequired,err.Error())
		return
	}
	// log.Printf("Request data: %s" , req)
	database.Db.Where("group_Id = ? and marquee_order = ?",req.GroupId,req.Symbol.MarqueeOrder).Delete(model.Symbol{})
}