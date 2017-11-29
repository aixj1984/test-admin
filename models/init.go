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

	orm.RegisterDataBase("default", "mysql", "dev_richard:dev_richard@tcp(10.0.12.239:3310)/test?charset=utf8")

	orm.RegisterModel(new(Account), new(QuestionLunjijichu), new(QuestionJicangguanli), new(QuestionLunjiguanli), new(QuestionBipengyuxinhao), new(CourseTest), new(TestQuestion))

	orm.SetMaxIdleConns("default", 5)

	orm.SetMaxOpenConns("default", 10)

	//create table
	orm.RunSyncdb("default", false, true)

	beego.Info("init db info  ..... ")

}
