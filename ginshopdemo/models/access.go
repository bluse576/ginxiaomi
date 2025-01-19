package models

type Access struct {
	Id          int
	ModuleName  string //模块名称
	ActionName  string //操作类型
	Type        int    //节点类型 1、表示模块 2、表示表单 3、操作
	Url         string //路由跳转地址
	ModuleId    int    //module_id与模型id关联
	Sort        int
	Description string
	Status      int
	AddTime     int
	AccessItem  []Access `gorm:"foreignKey:ModuleId;references:Id"`
	Checked     bool     `gorm:"-"`
}

func (Access) TableName() string {
	return "access"
}
