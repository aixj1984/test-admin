package providers

import (
	"strconv"

	"github.com/astaxie/beego/orm"

	"test-admin/models"

	"fmt"
)

//IAccountProvider account provider interface
type IArticleProvider interface {
	GetOne(m interface{}) error
	InsertOne(m interface{}) (int64, error)
	UpdateOne(m *models.Article) (int64, error)
	UpdateOneStatus(status, id int) (int64, error)
	GetMore(array interface{}, query_key, status string, start, length int) (int64, error)
	Count(query_key, status string) (int64, error)
}

//AccountProvider account provider
type ArticleProvider struct {
}

func (p *ArticleProvider) GetOne(m interface{}) error {
	o := orm.NewOrm()

	err := o.Read(m)

	return err
}

func (p *ArticleProvider) InsertOne(m interface{}) (int64, error) {
	o := orm.NewOrm()

	effact, err := o.Insert(m)

	return effact, err
}

func (p *ArticleProvider) UpdateOne(m *models.Article) (int64, error) {
	o := orm.NewOrm()

	res, err := o.Raw("update  article set title = ?, content = ? ,source = ? ,abstract = ? where id = ? ",
		m.Title, m.Content, m.Source, m.Abstract, m.Id).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
		return num, nil
	}
	return 0, err
}

func (p *ArticleProvider) UpdateOneStatus(status, id int) (int64, error) {
	o := orm.NewOrm()
	//表名不能够放到参数表中，否则异常
	res, err := o.Raw("update article set status = ? where id = ?",
		status, id).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
		return num, nil
	}
	return 0, err
}

func (p *ArticleProvider) GetMore(array interface{}, query_key, status string, start, length int) (int64, error) {
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
		From("article").
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

func (p *ArticleProvider) Count(query_key, status string) (int64, error) {
	o := orm.NewOrm()

	cond := orm.NewCondition()
	cond1 := cond.And("title__icontains", query_key)

	if id, err := strconv.ParseInt(query_key, 10, 64); err == nil {
		cond1 = cond1.Or("id", id)
	}
	if status != "2" {
		cond1 = cond.AndCond(cond1).AndCond(cond.And("status", status))
	}

	qs := o.QueryTable("article")
	qs = qs.SetCond(cond1)

	count, err := qs.Count()

	return count, err
}
