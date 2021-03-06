// ==========================================================================
// This is auto-generated by gf cli tool. You may not really want to edit it.
// ==========================================================================

package auth_rule

import (
	"database/sql"
	"github.com/gogf/gf/database/gdb"
)

// Entity is the golang structure for table auth_rule.
type Entity struct {
	Id         uint   `orm:"id,primary"  json:"id"`           //
	Pid        uint   `orm:"pid"         json:"pid"`          // 父ID
	Name       string `orm:"name,unique" json:"name"`         // 规则名称
	Title      string `orm:"title"       json:"title"`        // 规则名称
	Icon       string `orm:"icon"        json:"icon"`         // 图标
	Condition  string `orm:"condition"   json:"condition"`    // 条件
	Remark     string `orm:"remark"      json:"remark"`       // 备注
	MenuType   uint   `orm:"menu_type"      json:"menu_type"` // 类型 0目录 1菜单 2按钮
	Createtime uint   `orm:"createtime"  json:"createtime"`   // 创建时间
	Updatetime uint   `orm:"updatetime"  json:"updatetime"`   // 更新时间
	Weigh      int    `orm:"weigh"       json:"weigh"`        // 权重
	Status     uint   `orm:"status"      json:"status"`       // 状态
	AlwaysShow uint   `orm:"always_show"   json:"alwaysShow"` //显示状态 0隐藏 1显示
	Path       string `orm:"path"     json:"path"`            //路由地址
	IsFrame    uint   `orm:"is_frame"  json:"isFrame"`        //是否外链 1是 0否
}

// OmitEmpty sets OPTION_OMITEMPTY option for the model, which automatically filers
// the data and where attributes for empty values.
func (r *Entity) OmitEmpty() *arModel {
	return Model.Data(r).OmitEmpty()
}

// Inserts does "INSERT...INTO..." statement for inserting current object into table.
func (r *Entity) Insert() (result sql.Result, err error) {
	return Model.Data(r).Insert()
}

// Replace does "REPLACE...INTO..." statement for inserting current object into table.
// If there's already another same record in the table (it checks using primary key or unique index),
// it deletes it and insert this one.
func (r *Entity) Replace() (result sql.Result, err error) {
	return Model.Data(r).Replace()
}

// Save does "INSERT...INTO..." statement for inserting/updating current object into table.
// It updates the record if there's already another same record in the table
// (it checks using primary key or unique index).
func (r *Entity) Save() (result sql.Result, err error) {
	return Model.Data(r).Save()
}

// Update does "UPDATE...WHERE..." statement for updating current object from table.
// It updates the record if there's already another same record in the table
// (it checks using primary key or unique index).
func (r *Entity) Update() (result sql.Result, err error) {
	return Model.Data(r).Where(gdb.GetWhereConditionOfStruct(r)).Update()
}

// Delete does "DELETE FROM...WHERE..." statement for deleting current object from table.
func (r *Entity) Delete() (result sql.Result, err error) {
	return Model.Where(gdb.GetWhereConditionOfStruct(r)).Delete()
}
