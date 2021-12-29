1.使用 gin 和 gorm 接收使用者 request 與連結資料庫
gin: 接收使用者 request
gorm: 連結 db 進行 crud

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
