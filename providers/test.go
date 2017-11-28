package providers

import (
	"strconv"

	"github.com/astaxie/beego/orm"

	//	"errors"
)

//IAccountProvider account provider interface
type ITestProvider interface {
	GetOne(interface{}) error
	InsertOne(m interface{}) (int64, error)
	GetMore(array, object interface{}, query_key, status string, start, length int) (int64, error)
	UpdateOne(m interface{}, cols ...string) (int64, error)
	Count(array interface{}, object interface{}, query_key, status string) (int64, error)
}

//AccountProvider account provider
type TestProvider struct {
}

//Get 获取account
func (p *TestProvider) GetOne(m interface{}) error {
	o := orm.NewOrm()

	err := o.Read(m)

	return err
}

func (p *TestProvider) InsertOne(m interface{}) (int64, error) {
	o := orm.NewOrm()

	effact, err := o.Insert(m)

	return effact, err
}

func (p *TestProvider) UpdateOne(m interface{}, cols ...string) (int64, error) {
	o := orm.NewOrm()

	effact, err := o.Update(m, cols...)

	return effact, err
}

func (p *TestProvider) GetMore(array interface{}, object interface{}, query_key, status string, start, length int) (int64, error) {
	o := orm.NewOrm()
	// LIMIT 10 OFFSET 20 注意跟 SQL 反过来的

	cond := orm.NewCondition()
	cond1 := cond.And("title__icontains", query_key)

	if id, err := strconv.ParseInt(query_key, 10, 64); err == nil {
		cond1 = cond1.Or("id", id)
	}

	if status != "2" {
		cond1 = cond.AndCond(cond1).AndCond(cond.And("status", status))
	}

	qs := o.QueryTable(object)
	qs = qs.SetCond(cond1)

	effact, err := qs.Limit(length, start).All(array)

	//effact, err := o.QueryTable(object).Filter("title__icontains", query_key).Limit(length, start).All(array)

	return effact, err
}

func (p *TestProvider) Count(array interface{}, object interface{}, query_key, status string) (int64, error) {
	o := orm.NewOrm()

	cond := orm.NewCondition()
	cond1 := cond.And("title__icontains", query_key)

	if id, err := strconv.ParseInt(query_key, 10, 64); err == nil {
		cond1 = cond1.Or("id", id)
	}
	if status != "2" {
		cond1 = cond.AndCond(cond1).AndCond(cond.And("status", status))
	}

	qs := o.QueryTable(object)
	qs = qs.SetCond(cond1)

	count, err := qs.Count()

	//count, err := o.QueryTable(object).Filter("title__icontains", query_key).Count()

	return count, err
}
