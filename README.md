1.使用 gin 和 gorm 接收使用者 request 與連結資料庫
gin: 接收使用者 request
gorm: 連結 db 進行 crud

go get -u gorm.io/gorm
go get -u github.com/gin-gonic/gin
go get github.com/silenceper/gowatch
go get -u github.com/sirupsen/logrus
go get -u github.com/natefinch/lumberjack
go get github.com/gin-contrib/sessions

2.gin route 接收使用者 api 訊息

router := gin.Default() 啟動 router
router.GET("/getSymbols/:groupname",apis.GetSymbols)

解析 request params ，使用 gin.Context.Param
func GetSymbols(c \*gin.Context){....
var symbol model.Symbol
paramInfo := c.Param("groupname")
groupname := strings.Split(paramInfo,"=")[1]
}

解析使用者送來的 url
path := c.FullPath()

3.gorm 建立資料表格式，使用 struct。

type Symbols struct{
groupID INT
symbol string
show boolean
}

type Speed struct {
speed INT
}

type GroupName struct{
groupID INT
groupName string
Play bool
}

3.gorm AutoMigrate 建立資料表並時時更新。
db.AutoMigrate(&Symbols,&Speed,&GroupName)

4.gorm Create()寫入資料
寫入資料時設定 struct 映射到 table 欄位時，參數首字母要大寫，其後要加上註解說明這個參數是要映射到哪個項目上。

p.s. init()，每個 package 被 import 的時候都會執行 init 這個檔案

5.log 使用 logrus、lumberjack、gin 三個套件撰寫 log,logrus 協助將 gin 的 log json 化，lumberjack 則限制 log 數量不會無限制增長。

1/11 號
要附上 cookie 給前端時出現 cors 問題，不同源的請求必須設定才可以存 cookie。

設定的項目：
Access-Control-Allow-Origin，要有 cookie，不能是\*
withCredential:true
content-type:application/json

6.開發環境包板到正式環境要改設定 cookie 的 domain
GOOS=linux GOARCH=amd64 go build

api/router/router.go
api/apis/symbols.go
