package models

import (
	"github.com/astaxie/beego/orm"
	"MyShop/Router/utils"
	_ "github.com/go-sql-driver/mysql"

)

type AdminUser struct {
	Id int
	Name string `orm:"size(11)"`
	Password string `orm:"size(40)"`
}

//初始化数据库
func init()  {
	// 设置驱动
	orm.RegisterDriver("mysql",orm.DRMySQL)

	// 连接数据
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp("+utils.G_mysql_addr+":"+utils.G_mysql_port+")/myshop?charset=utf8", 30)
	//注册model 建表
	orm.RegisterModel(new(AdminUser),)

	// create table
	//第一个是别名
	// 第二个是是否强制替换模块   如果表变更就将false 换成true 之后再换回来表就便更好来了
	//第三个参数是如果没有则同步或创建
	orm.RunSyncdb("default", false, true)
}