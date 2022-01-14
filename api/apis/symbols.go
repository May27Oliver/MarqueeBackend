package apis

import (
	"MarqueeBackstage/api/database"
	model "MarqueeBackstage/api/models"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type symbol struct {
	StockNo      string `json:stockNo`
	StockName    string `json:stockName`
	Symbol       string `json:symbol`
	MarqueeOrder int    `json:marqueeOrder`
}
type addDeleteRequest struct {
	GroupId int    `json:groupId`
	Symbol  symbol `json:symbol`
}
type importRequest struct {
	GroupId int      `json:groupId`
	Symbols []symbol `json:symbols`
}

type updateGroupRequest struct {
	GroupId int `json:groupId`
}

type updateSpeedRequest struct {
	Speed    int `json:speed`
}

type updateDirectionRequest struct {
	Direction    int `json:direction`
}

type groupSymbol struct {
	groupName []model.GroupName
	symbols   []model.Symbol
}

type login struct {
	Account  string `json:account`
	Password string `json:password`
}

type verifyLogin struct {
	SessionId string
}

var domain string = os.Getenv("DOMAIN")

func GetMarqueeSymbols(c *gin.Context) {
	WriteLoggeer(c)
	var groupname model.GroupName
	resGroupName, err := groupname.GroupNameQuery()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  false,
			"message": "無法取得播放群組",
		})
		return
	}
	// log.Printf("resGroupName: %s" , resGroupName)
	var symbol model.Symbol
	groupId := fmt.Sprintf("%v", resGroupName[0].GroupID)
	res, err := symbol.SymbolQuery(groupId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  false,
			"message": "查無資料",
		})
		return
	}
	// log.Printf("res: %s" , res)
	var resSymbols []string

	for _, v := range res {
		resSymbols = append(resSymbols, v.Symbol)
	}

	// log.Printf("resSymbols: %s" , resSymbols)
	c.JSON(http.StatusOK, gin.H{
		"result":  true,
		"symbols": resSymbols,
	})
}

//取得群組
func GetSymbols(c *gin.Context) {
	WriteLoggeer(c)
	var symbol model.Symbol
	paramInfo := c.Param("groupname")
	groupname := strings.Split(paramInfo, "=")[1]

	//讀取使用者request
	result, err := symbol.SymbolQuery(groupname)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  false,
			"message": "查無資料",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result":  true,
		"symbols": result,
	})
}

//取得速度
func GetSpeed(c *gin.Context) {
	WriteLoggeer(c)
	var config model.MarqueeConfig
	result, err := config.GetSpeed()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  false,
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": true,
		"speed":  result.ConfigValue,
	})
}

//取得群組名稱
func GetGroupName(c *gin.Context) {
	WriteLoggeer(c)
	var groupname model.GroupName
	result, err := groupname.GroupNameQuery()
	log.Printf("Request data: %s", result)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  false,
			"message": "查無資料",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result":    true,
		"groupName": result[0],
	})
}

//增加群組symbol
func AddSymbol(c *gin.Context) {
	WriteLoggeer(c)
	var symbol model.Symbol
	req := addDeleteRequest{}
	if err := c.BindJSON(&req); err != nil {
		c.String(http.StatusPaymentRequired, err.Error())
		return
	}
	req_data := model.Symbol{GroupID: req.GroupId, Symbol: req.Symbol.Symbol, StockName: req.Symbol.StockName, Show: true, MarqueeOrder: req.Symbol.MarqueeOrder}
	// log.Printf("Request data: %s" , req_data)
	err := symbol.AddSymbol(req_data)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  false,
			"message": "新增失敗",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result":  true,
		"message": "新增成功",
	})
}

// 匯入csv
func ImportSymbols(c *gin.Context) {
	WriteLoggeer(c)
	req := importRequest{}
	if err := c.BindJSON(&req); err != nil {
		c.String(http.StatusPaymentRequired, err.Error())
		return
	}
	log.Printf("req.Symbols: %s", req.Symbols)

	sqlStr := "INSERT INTO `symbols` (group_id,symbol,`show`,stock_name,marquee_order) VALUES "

	connInfo := []interface{}{}

	for _, v := range req.Symbols {
		sqlStr += "(?, ?, ?, ?, ?),"
		connInfo = append(connInfo, req.GroupId, v.Symbol, true, v.StockName, v.MarqueeOrder)
	}

	sqlStr = sqlStr[0:len(sqlStr)-1] + ";"
	result := database.Db.Debug().Exec(sqlStr, connInfo...)

	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  false,
			"message": result.Error,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result":  true,
		"message": "新增成功",
	})
}

// 更新播放群組
func UpdateGroupNo(c *gin.Context) {
	WriteLoggeer(c)
	req := updateGroupRequest{}
	if err := c.BindJSON(&req); err != nil {
		c.String(http.StatusPaymentRequired, err.Error())
		return
	}
	if req.GroupId == 0 {
		return
	}
	var groupname model.GroupName
	if err := groupname.CloseAllGroupName(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  false,
			"message": err,
		})
		return
	}

	if err := groupname.UpdateGroupName(req.GroupId); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  false,
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result":  true,
		"message": "更新播放群組成功",
	})
}

func UpdateSpeed(c *gin.Context) {
	WriteLoggeer(c)
	var config model.MarqueeConfig
	req := updateSpeedRequest{}
	if err := c.BindJSON(&req); err != nil {
		c.String(http.StatusPaymentRequired, err.Error())
		return
	}

	log.Printf("UpdateSpeed %s", req.Speed)
	if err := config.UpdateSpeed(req.Speed); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  false,
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result":  true,
		"message": "更新成功",
	})
}

//刪除symbol
func DeleteSymbol(c *gin.Context) {
	WriteLoggeer(c)
	var symbol model.Symbol
	req := addDeleteRequest{}
	if err := c.BindJSON(&req); err != nil {
		c.String(http.StatusPaymentRequired, err.Error())
		return
	}
	err := symbol.DeleteSymbol(req.GroupId, req.Symbol.MarqueeOrder)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  false,
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result":  true,
		"message": "刪除成功",
	})
}

func Login(c *gin.Context) {
	WriteLoggeer(c)

	var member model.Member
	req := login{}
	if err := c.BindJSON(&req); err != nil {
		c.String(http.StatusPaymentRequired, err.Error())
		return
	}
	res, er := member.LoginCheckAcc(req.Account)
	log.Printf("loginCheckAcc %s", res)
	if er != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  false,
			"message": er,
		})
		return
	} else if len(res) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"result":  false,
			"message": "帳號有誤",
			"code":    2,
		})
		return
	} else {
		result, err := member.Login(req.Account, req.Password)
		log.Printf("login %s", result)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"result":  false,
				"message": err,
			})
			return
		} else if len(result) == 0 {
			c.JSON(http.StatusOK, gin.H{
				"result":  false,
				"message": "密碼有誤",
				"code":    3,
			})
			return
		} else {
			expiration := time.Now()
			expiration = expiration.AddDate(0,0,1)
			c.Header("credentials", "include")
			//"localhost"
			c.SetCookie("sessionId","2022-01-10-1111",360000,"/",domain,false,true)
			c.JSON(http.StatusOK, gin.H{
				"result":  true,
				"code":    1,
				"message": "登入成功",
			})
		}
	}
}

//驗證是否登入
func VerifyLogin(c *gin.Context){
	WriteLoggeer(c)
	// cookie,err:= c.Request.Cookie("sessionId")
	sessionId,_ := c.Cookie("sessionId")
	log.Printf("sessionId %s", sessionId)
	req := verifyLogin{}
	if err := c.BindJSON(&req);err!=nil{
		c.String(http.StatusPaymentRequired,err.Error())
		return
	}
	if sessionId == "2022-01-10-1111"{
		c.JSON(http.StatusOK,gin.H{
			"result":true,
			"verify":"success",
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"result":false,
		"verify":"fail",
	})
}

func GetDirection(c *gin.Context){
	WriteLoggeer(c)
	var config model.MarqueeConfig
	result, err := config.GetDirection()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  false,
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": true,
		"direction":  result.ConfigValue,
	})
}

func UpdateDirection(c *gin.Context) {
	WriteLoggeer(c)
	var config model.MarqueeConfig
	req := updateDirectionRequest{}
	if err := c.BindJSON(&req); err != nil {
		c.String(http.StatusPaymentRequired, err.Error())
		return
	}

	log.Printf("UpdateDirection %s", req.Direction)
	if err := config.UpdateDirection(req.Direction); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  false,
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result":  true,
		"message": "更新成功",
	})
}

func WriteLoggeer(c *gin.Context) {
	// log
	LogInstance.WithFields(logrus.Fields{
		"Method": c.Request.Method,
		"Host":   c.Request.Host,
		"code":   c.Writer.Status(),
		"url":    c.Request.URL.Path,
	}).Info("GinInfo")
}