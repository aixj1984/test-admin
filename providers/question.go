package providers

import (
	"strconv"

	"github.com/astaxie/beego/orm"

	"test-admin/models"

	"fmt"

	"errors"
)

//IAccountProvider account provider interface
type IQuestionProvider interface {
	GetOne(m *models.Question, tablename string) error
	InsertOne(m *models.Question, tablename string) (int64, error)
	UpdateOne(m *models.Question, tablename string) (int64, error)
	UpdateOneStatus(m *models.Question, tablename string) (int64, error)
	GetMore(array interface{}, tablename, query_key, status string, start, length int) (int64, error)
	Count(tablename, query_key, status string) (int64, error)
}

//AccountProvider account provider
type QuestionProvider struct {
}

//Get 获取account
func (p *QuestionProvider) GetOne(m *models.Question, tablename string) error {
	o := orm.NewOrm()

	err := o.QueryTable("question_lunjijichu").Filter("id", 1).One(m)
	if err == orm.ErrMultiRows {
		// 多条的时候报错
		return errors.New("Multi Rows")
	}
	if err == orm.ErrNoRows {
		// 没有找到记录
		return errors.New("Not Find")
	}

	//err := o.Read(m)

	return err
}

func (p *QuestionProvider) InsertOne(m *models.Question, tablename string) (int64, error) {
	o := orm.NewOrm()

	res, err := o.Raw("insert into "+tablename+" (title,question_type, options,answer,note,status) values(?,0,?,?,?,0)", m.Title, m.Options, m.Answer, m.Note).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
		return num, nil
	}

	return 0, err

}

func (p *QuestionProvider) UpdateOne(m *models.Question, tablename string) (int64, error) {
	o := orm.NewOrm()

	res, err := o.Raw("update  "+tablename+" set title = ?, options = ? ,answer = ? ,note = ? where id = ? ",
		m.Title, m.Options, m.Answer, m.Note, m.Id).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
		return num, nil
	}
	return 0, err
}

func (p *QuestionProvider) UpdateOneStatus(m *models.Question, tablename string) (int64, error) {
	o := orm.NewOrm()
	//表名不能够放到参数表中，否则异常
	res, err := o.Raw("update "+tablename+" set status = ? where id = ?",
		m.Status, m.Id).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
		return num, nil
	}
	return 0, err
}

func (p *QuestionProvider) GetMore(array interface{}, tablename, query_key, status string, start, length int) (int64, error) {
	qb, _ := orm.NewQueryBuilder("mysql")

	var condition = " 1=1 "

	if len(query_key) > 0 {

		if _, err := strconv.ParseInt(query_key, 10, 64); err == nil {
			condition += " and (id = " + query_key + " or title like '%" + query_key + "%' ) "
		} else {
			condition += " and  title like '%" + query_key + "%' "
		}
	}

	if status != "2" {
		condition += " and status = " + status
	}

	qb.Select("*").
		From(tablename).
		Where(condition).
		Limit(length).Offset(start)

	// 构建查询对象

	// 导出 SQL 语句
	sql := qb.String()

	// 执行 SQL 语句
	o := orm.NewOrm()

	effact, err := o.Raw(sql).QueryRows(array)

	return effact, err
}

func (p *QuestionProvider) Count(tablename, query_key, status string) (int64, error) {
	o := orm.NewOrm()

	cond := orm.NewCondition()
	cond1 := cond.And("title__icontains", query_key)

	if id, err := strconv.ParseInt(query_key, 10, 64); err == nil {
		cond1 = cond1.Or("id", id)
	}
	if status != "2" {
		cond1 = cond.AndCond(cond1).AndCond(cond.And("status", status))
	}

	qs := o.QueryTable(tablename)
	qs = qs.SetCond(cond1)

	count, err := qs.Count()

	//count, err := o.QueryTable(object).Filter("title__icontains", query_key).Count()

	return count, err
}
