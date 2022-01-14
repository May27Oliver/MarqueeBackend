package models

import (
	"MarqueeBackstage/api/database"
)

type Symbol struct {
	GroupID      int `gorm:"foreignKey:GroupID"`
	Symbol       string
	StockName    string
	Show         bool
	MarqueeOrder int
}

type Speed struct {
	Speed    int `gorm:"primary_key"`
	Selected bool
}

type GroupName struct {
	GroupID   int `gorm:"primary_key"`
	GroupName string
	Play      bool
}

type Member struct {
	MemberID int    `gorm:"type:int;primary_key;AUTO_INCREMENT;not null"`
	Account  string `gorm:"type:nvarchar(20);not null;unique"`
	Password string `gorm:"type:nvarchar(20);not null"`
}

type MarqueeConfig struct {
		ConfigId int `gorm:"type:int;primary_key;AUTO_INCREMENT;not null"`
		ConfigName string `gorm:"type:nvarchar(20);not null"`
		ConfigValue int `gorm:"type:int;not null"`
}

//新增symbols
func (symbol Symbol) SymbolInsert() (err error) {
	result := database.Db.Create(&symbol)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

//取得symbols
func (symbol Symbol) SymbolQuery(groupname string) (symbols []Symbol, err error) {
	if err = database.Db.Where("group_id = ?", groupname).Order("marquee_order").Find(&symbols).Error; err != nil {
		return
	}
	return
}

//取得groupName
func (groupname GroupName) GroupNameQuery() (groupnames []GroupName, err error) {
	if err = database.Db.Where("play = ?", true).Find(&groupnames).Error; err != nil {
		return
	}
	return
}

//刪除symbol
func (symbol Symbol) DeleteSymbol(groupId int, marqueeOrder int) (err error) {
	if result := database.Db.Where("group_Id = ? and marquee_order = ?", groupId, marqueeOrder).Delete(Symbol{}); result.Error != nil {
		return
	}
	return
}

//查詢速度
func (m MarqueeConfig) GetSpeed() (num MarqueeConfig, err error) {
	if err = database.Db.Where("`config_name` = ?", "speed").Find(&num).Error; err != nil {
		return
	}
	return
}

//更改播放速度
func (m MarqueeConfig) UpdateSpeed(num int) (err error) {
	if result := database.Db.Model(&MarqueeConfig{}).Where("`config_name` = ?", "speed").Update("config_value", num); result.Error != nil {
		return
	}
	return
}

//查詢速度
func (m MarqueeConfig) GetDirection() (num MarqueeConfig, err error) {
	if err = database.Db.Where("`config_name` = ?", "direction").Find(&num).Error; err != nil {
		return
	}
	return
}

//更改播放速度
func (m MarqueeConfig) UpdateDirection(num int) (err error) {
	if result := database.Db.Model(&MarqueeConfig{}).Where("`config_name` = ?", "direction").Update("config_value", num); result.Error != nil {
		return
	}
	return
}

//更改播放群組
func (groupname GroupName) UpdateGroupName(num int) (err error) {
	if result := database.Db.Model(&GroupName{}).Where("group_id = ?", num).Update("play", true); result.Error != nil {
		return
	}
	return
}

//關閉全部群組
func (groupname GroupName) CloseAllGroupName() (err error) {
	if result := database.Db.Model(&GroupName{}).Update("play", false); result.Error != nil {
		return
	}
	return
}

//增加商品
func (symbol Symbol) AddSymbol(req_data Symbol) (err error) {
	if result := database.Db.Debug().Create(&req_data); result.Error != nil {
		return
	}
	return
}

//登入
func (member Member) Login(acc string, pass string) (members []Member, err error) {
	if err = database.Db.Debug().Where("account = ? and password = ?", acc, pass).Find(&members).Error; err != nil {
		return
	}
	return
}

//登入check帳號
func (member Member) LoginCheckAcc(acc string) (mb []Member, err error) {
	if err = database.Db.Debug().Where("account = ? ", acc).Find(&mb).Error; err != nil {
		return
	}
	return
}
