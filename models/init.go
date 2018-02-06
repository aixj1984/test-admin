package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	_ "github.com/go-sql-driver/mysql"
)

//Init  初始化model
func init() {

	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.Debug = true

	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("db_conn"))

	orm.RegisterModel(new(Account),
		new(Question1), new(Question2), new(Question3),
		new(Question4), new(Question5), new(Question6),
		new(Question7), new(Question8), new(Question9),
		new(Question10), new(Question11), new(Article),
		new(CourseTest), new(TestQuestion))

	orm.SetMaxIdleConns("default", 5)

	orm.SetMaxOpenConns("default", 10)

	//create table
	orm.RunSyncdb("default", false, true)

	beego.Info("init db info  ..... ")

}
