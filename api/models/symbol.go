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
	Speed int
	use bool
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
	if err = database.Db.Find(&groupnames).Error; err!=nil{
		return
	}
	return
}