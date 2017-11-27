package providers

import (
	"test-admin/models"

	"github.com/astaxie/beego/orm"

	//"errors"
)

//IAccountProvider account provider interface
type IAccountProvider interface {
	GetOne(*models.Account) error
}

//AccountProvider account provider
type AccountProvider struct {
}

//Get 获取account
func (p *AccountProvider) GetOne(m *models.Account) error {
	o := orm.NewOrm()
	/*
		err := o.QueryTable("account").Filter("id", 1).One(m)
		if err == orm.ErrMultiRows {
			// 多条的时候报错
			return errors.New("test")
		}
		if err == orm.ErrNoRows {
			// 没有找到记录
			return errors.New("test111")
		}*/
	err := o.Read(m)

	return err
}
