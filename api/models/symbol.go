package models

import (
	"MarqueeBackstage/api/database"
)

type Symbol struct{
	GroupID int `gorm:"foreignKey:GroupID"`
	Symbol string
	StockName string
	Show bool
	MarqueeOrder int 
}
	
type Speed struct {
	Speed int `gorm:"primary_key"`
	Use bool
}
	
type GroupName struct{
	GroupID int `gorm:"primary_key"`
	GroupName string
	Play bool
}

//新增symbols
func (symbol Symbol) SymbolInsert()(err error){
	result := database.Db.Create(&symbol)
	if result.Error != nil{
			err	=	result.Error
			return
	}
	return
} 

//取得symbols
func (symbol Symbol) SymbolQuery(groupname string)(symbols []Symbol,err error){
	if err = database.Db.Where("group_id = ?",groupname).Order("marquee_order").Find(&symbols).Error; err!=nil{
		return
	}
	return 
}

//取得groupName
func (groupname GroupName) GroupNameQuery()(groupnames []GroupName,err error){
	if err = database.Db.Where("play = ?", true ).Find(&groupnames).Error; err!=nil{
		return
	}
	return
}

//刪除symbol
func (symbol Symbol) DeleteSymbol(groupId int,marqueeOrder int)(err error){
	if result := database.Db.Where("group_Id = ? and marquee_order = ?",groupId,marqueeOrder).Delete(Symbol{});result.Error != nil{
		return
	}
	return
}

//查詢速度
func (speed Speed) GetSpeed()(num Speed,err error){
	if err = database.Db.Where("`use` = ?", true).Find(&num).Error;err!=nil{
		return
	}
	return
}

//更改播放速度
func (speed Speed) UpdateSpeed(num int)(err error){
	if result := database.Db.Model(&Speed{}).Where("speed = ?",num).Update("use",true);result.Error!=nil{
		return
	}
	return
}

//全部速度停用
func (speed Speed) CloseAllSpeed()(err error){
	if result := database.Db.Model(&Speed{}).Update("use",false);result.Error!=nil{
		return
	}
	return
}

//更改播放群組
func (groupname GroupName) UpdateGroupName(num int)(err error){
	if result := database.Db.Model(&GroupName{}).Where("group_id = ?",num).Update("play",true);result.Error!=nil{
		return 
	}
	return
}

//關閉全部群組
func (groupname GroupName) CloseAllGroupName()(err error){
	if result := database.Db.Model(&GroupName{}).Update("play",false);result.Error!=nil{
		return
	}
	return
}

//增加商品
func (symbol Symbol) AddSymbol(req_data Symbol)(err error){
	if result :=database.Db.Debug().Create(&req_data);result.Error!=nil{
		return
	}
	return
}
