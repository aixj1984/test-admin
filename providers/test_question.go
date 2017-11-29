package providers

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego/orm"

	//	"errors"
)

type ITestQuestionProvider interface {
	GetOne(interface{}, ...string) error
	InsertOne(m interface{}) (int64, error)
	GetMore(array, object interface{}, query_key, status string, question_tablename, test_id string, start, length int) (int64, error)
	UpdateOne(m interface{}, cols ...string) (int64, error)
	Count(object interface{}, query_key, status string, question_tablename, test_id string) (int64, error)
	DeleteOne(m interface{}) (int64, error)
}

type TestQuestionProvider struct {
}

func (p *TestQuestionProvider) GetOne(m interface{}, cols ...string) error {
	o := orm.NewOrm()

	err := o.Read(m, cols...)

	return err
}

func (p *TestQuestionProvider) InsertOne(m interface{}) (int64, error) {
	o := orm.NewOrm()

	effact, err := o.Insert(m)

	return effact, err
}

func (p *TestQuestionProvider) UpdateOne(m interface{}, cols ...string) (int64, error) {
	o := orm.NewOrm()

	effact, err := o.Update(m, cols...)

	return effact, err
}

func (p *TestQuestionProvider) DeleteOne(m interface{}) (int64, error) {
	o := orm.NewOrm()

	effact, err := o.Delete(m)

	return effact, err
}

func (p *TestQuestionProvider) GetMore(array interface{}, object interface{}, query_key, status string, question_tablename, test_id string, start, length int) (int64, error) {

	qb, _ := orm.NewQueryBuilder("mysql")

	var condition = "status = 1"

	if len(query_key) > 0 {
		if _, err := strconv.ParseInt(query_key, 10, 64); err == nil {
			condition += " and " + question_tablename + ".id = " + query_key
		} else {
			condition += " and  title like '%" + query_key + "%' "
		}

	}
	if status == "2" {
		qb.Select("title",
			question_tablename+".id",
			"question_type",
			"answer",
			"note",
			"IFNULL(test_question.question_id, 0) as status").
			From(question_tablename).
			LeftJoin("(select * from course_test_question where test_id = ? ) as test_question").On(question_tablename + ".id = test_question.question_id").
			Where(condition).
			Limit(length).Offset(start)
	} else if status == "0" {
		condition += " and id not in (select question_id from course_test_question where test_id = ? ) "
		qb.Select("title",
			question_tablename+".id",
			"question_type",
			"answer",
			"note",
			"0 as status").
			From(question_tablename).
			Where(condition).
			Limit(length).Offset(start)
	} else if status == "1" {
		qb.Select("title",
			question_tablename+".id",
			"question_type",
			"answer",
			"note",
			"IFNULL(test_question.question_id, 0) as status").
			From(question_tablename).
			InnerJoin("(select * from course_test_question where test_id = ? ) as test_question").On(question_tablename + ".id = test_question.question_id").
			Where(condition).
			Limit(length).Offset(start)
	}

	// 构建查询对象

	// 导出 SQL 语句
	sql := qb.String()

	fmt.Println(sql)

	// 执行 SQL 语句
	o := orm.NewOrm()

	effact, err := o.Raw(sql, test_id).QueryRows(array)
	return effact, err

}

func (p *TestQuestionProvider) Count(object interface{}, query_key, status string, question_tablename, test_id string) (int64, error) {
	qb, _ := orm.NewQueryBuilder("mysql")

	var condition = "status = 1 "

	if len(query_key) > 0 {
		if _, err := strconv.ParseInt(query_key, 10, 64); err == nil {
			condition += " and question_lunjijichu.id = " + query_key
		} else {
			condition += " and  title like '%" + query_key + "%' "
		}
	}

	if status == "2" {
		qb.Select("count(*) as num").
			From(question_tablename).
			LeftJoin("(select * from course_test_question where test_id = ? ) as test_question").On(question_tablename + ".id = test_question.question_id").
			Where(condition)
	} else if status == "0" {
		condition += " and id not in (select question_id from course_test_question where test_id = ? ) "
		qb.Select("count(*) as num").
			From(question_tablename).
			Where(condition)
	} else if status == "1" {
		qb.Select("count(*) as num").
			From(question_tablename).
			InnerJoin("(select * from course_test_question where test_id = ? ) as test_question").On(question_tablename + ".id = test_question.question_id").
			Where(condition)
	}

	// 导出 SQL 语句
	sql := qb.String()

	fmt.Println(sql)

	var maps []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw(sql, test_id).Values(&maps)
	if err == nil && num > 0 {
		fmt.Println(maps[0]["num"])
		count, _ := strconv.Atoi(maps[0]["num"].(string))
		return int64(count), nil
	} else {
		return 0, err
	}
}
