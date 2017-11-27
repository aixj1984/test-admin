package providers

import (
	"github.com/astaxie/beego/orm"

	//	"errors"
)

//IAccountProvider account provider interface
type IQuestionProvider interface {
	GetOne(interface{}) error
	InsertOne(m interface{}) (int64, error)
	GetMore(array, object interface{}, query_key string, start, length int) (int64, error)
	UpdateOne(m interface{}, cols ...string) (int64, error)
	Count(array interface{}, object interface{}, query_key string) (int64, error)
}

//AccountProvider account provider
type QuestionProvider struct {
}

//Get 获取account
func (p *QuestionProvider) GetOne(m interface{}) error {
	o := orm.NewOrm()
	/*
		err := o.QueryTable("question_"+course).Filter("id", 1).One(m)
		if err == orm.ErrMultiRows {
			// 多条的时候报错
			return errors.New("Multi Rows")
		}
		if err == orm.ErrNoRows {
			// 没有找到记录
			return errors.New("Not Find")
		}*/

	err := o.Read(m)

	return err
}

func (p *QuestionProvider) InsertOne(m interface{}) (int64, error) {
	o := orm.NewOrm()

	effact, err := o.Insert(m)

	return effact, err
}

func (p *QuestionProvider) UpdateOne(m interface{}, cols ...string) (int64, error) {
	o := orm.NewOrm()

	effact, err := o.Update(m, cols...)

	return effact, err
}

func (p *QuestionProvider) GetMore(array interface{}, object interface{}, query_key string, start, length int) (int64, error) {
	o := orm.NewOrm()
	// LIMIT 10 OFFSET 20 注意跟 SQL 反过来的
	effact, err := o.QueryTable(object).Filter("title__icontains", query_key).Limit(length, start).All(array)

	return effact, err
}

func (p *QuestionProvider) Count(array interface{}, object interface{}, query_key string) (int64, error) {
	o := orm.NewOrm()

	count, err := o.QueryTable(object).Filter("title__icontains", query_key).Count()

	return count, err
}
