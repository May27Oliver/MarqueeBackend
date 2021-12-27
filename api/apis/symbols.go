package apis

import (
	"MarqueeBackstage/api/database"
	model "MarqueeBackstage/api/models"
	"fmt"
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
type groupSymbol struct{
	groupName []model.GroupName
	symbols []model.Symbol
}

func GetMarqueeSymbols(c *gin.Context){
	var groupname model.GroupName
	resGroupName,err := groupname.GroupNameQuery()
	if err!= nil{
		c.JSON(http.StatusOK,gin.H{
			"result":false,
			"message":"無法取得播放群組",
		})
		return
	}
	// log.Printf("resGroupName: %s" , resGroupName)
	var symbol model.Symbol
	groupId:= fmt.Sprintf("%b",resGroupName[0].GroupID)
	res,err := symbol.SymbolQuery(groupId)
	if err!= nil{
		c.JSON(http.StatusOK,gin.H{
			"result":false,
			"message":"查無資料",
		})
		return
	}
	var resSymbols []string
	
	for _,v := range res {
		resSymbols = append(resSymbols,v.Symbol)
	}

	// log.Printf("resSymbols: %s" , resSymbols)
	c.JSON(http.StatusOK,gin.H{
		"result":true,
		"symbols":resSymbols,
	})
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
		"symbols":result,
	})
}
//取得速度
func GetSpeed(c *gin.Context){
	var speed model.Speed
	result,err:=speed.GetSpeed()
	if err!= nil{
		c.JSON(http.StatusOK,gin.H{
			"result":false,
			"message":err,
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"result":true,
		"speed":result,
	})
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
		"groupName":result[0],
	})
} 
//增加群組symbol
func AddSymbol(c *gin.Context){
	var symbol model.Symbol
	req:= addDeleteRequest{}
	if err:=c.BindJSON(&req); err!=nil{
		c.String(http.StatusPaymentRequired,err.Error())
		return
	}
	req_data := model.Symbol{GroupID: req.GroupId,Symbol: req.Symbol.Symbol,StockName:req.Symbol.StockName,Show:true,MarqueeOrder:req.Symbol.MarqueeOrder}
	// log.Printf("Request data: %s" , req_data)
	err := symbol.AddSymbol(req_data)
	
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"result":false,
			"message":"新增失敗",
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"result":true,
		"message":"新增成功",
	})
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
	var groupname model.GroupName
	if err:= groupname.CloseAllGroupName();err!=nil{
		c.JSON(http.StatusOK,gin.H{
			"result":false,
			"message":err,
		})
		return
	}

	req:= updateGroupRequest{}
	if err := c.BindJSON(&req); err != nil {
		c.String(http.StatusPaymentRequired,err.Error())
		return
	}

	if err:=groupname.UpdateGroupName(req.GroupId);err!=nil{
		c.JSON(http.StatusOK,gin.H{
			"result":false,
			"message":err,
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"result":false,
		"message":"更新播放群組成功",
	})
}

func UpdateSpeed(c *gin.Context){
	var speed model.Speed
	if err := speed.CloseAllSpeed(); err != nil {
		c.JSON(http.StatusOK,gin.H{
			"result":false,
			"message":err,
		})
		return
	}
	req:= updateSpeedRequest{}
	if err := c.BindJSON(&req); err != nil {
		c.String(http.StatusPaymentRequired,err.Error())
		return
	}

	if err := speed.UpdateSpeed(req.Speed); err != nil{
		c.JSON(http.StatusOK,gin.H{
			"result":false,
			"message":err,
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"result":true,
		"message":"更新成功",
	})
}

//刪除symbol
func DeleteSymbol(c *gin.Context){
	var symbol model.Symbol
	req := addDeleteRequest{}
	if err := c.BindJSON(&req);err!=nil{
		c.String(http.StatusPaymentRequired,err.Error())
		return
	}
	err := symbol.DeleteSymbol(req.GroupId,req.Symbol.MarqueeOrder)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"result":false,
			"message":err,
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"result":true,
		"message":"刪除成功",
	})
}