package mysql

import (
	"bytes"
	"database/sql"
	"encoding/gob"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

var DbMap = make(map[string]*sql.DB)

//var db *sql.DB

func init() {
	initdb(beego.AppConfig.String("product_db"), beego.AppConfig.String("product_conn"))
}

func initdb(dbname, connect_str string) {

	if _, ok := DbMap[dbname]; ok {
		return
	}
	//connect_str := mysqlcfg.User + ":" + mysqlcfg.Passwd + "@" + mysqlcfg.Protocol + "(" + mysqlcfg.Host + ":" + strconv.Itoa(mysqlcfg.Port) + ")/" + mysqlcfg.DBName + "?charset=" + mysqlcfg.Charset
	db, err := sql.Open("mysql", connect_str)
	if err != nil {
		beego.Error("mysql connet error : %s, connect_str : %s", err.Error(), connect_str)
		panic(1)
	}
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(1)
	db.Ping()
	DbMap[dbname] = db

}

type MysqlModel struct {
	Host     string
	Port     int
	Protocol string
	User     string
	Passwd   string
	DBName   string
	Charset  string
}

func QueryCount(dbname, sql_cmd string) int {

	record_count := 0

	rows, err := DbMap[dbname].Query(sql_cmd)
	defer rows.Close()

	if err != nil {
		beego.Error("query DB error ", err)
		return record_count
	}

	if columns, _ := rows.Columns(); len(columns) != 1 {
		beego.Error("sql error, not just right count ")
		return record_count
	}

	for rows.Next() {
		err = rows.Scan(&record_count)
		if err != nil {
			beego.Error("scan DB error ", err)
			return record_count
		}
		break
	}
	return record_count
}

func QueryFromMysql(dbname, sql_cmd string) []interface{} {

	task_list := make([]interface{}, 0)

	rows, err := DbMap[dbname].Query(sql_cmd)
	defer rows.Close()

	if err != nil {
		beego.Error("Query error, sql is : ", sql_cmd, err)
		return task_list
	}

	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		//将行数据保存到record字典
		err = rows.Scan(scanArgs...)
		if err != nil {
			beego.Error("Scan error", err)
			return task_list
		}
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		task_list = append(task_list, record)
	}
	return task_list

}

func deepCopy(dst, src interface{}) error {

	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}

func QueryRecordFromMysql(dbname, sql_cmd string) (rows *sql.Rows) {

	var err error

	rows, err = DbMap[dbname].Query(sql_cmd)

	if err != nil {
		rows.Close()
		beego.Error("Query error, sql is : ", sql_cmd, err)
		return nil
	}

	return rows
}

func GetDataByCondition(dbname, tablename string, columns string, conditions string, order string, start int, length int) []interface{} {

	if len(columns) == 0 {
		columns = " * "
	}

	if len(conditions) == 0 {
		conditions = "1=1"
	}

	if len(order) == 0 {
		order = "id"
	}

	if length == 0 {
		length = 10
	}

	sql_cmd := "select " + columns + " from " + tablename + " where " + conditions + " order by " + order + " limit " + strconv.Itoa(start) + " , " + strconv.Itoa(length)
	return QueryFromMysql(dbname, sql_cmd)
}

func GetGroupDataByCondition(dbname, tablename string, columns string, conditions string, group string) []interface{} {

	if len(columns) == 0 {
		columns = " * "
	}

	if len(conditions) == 0 {
		conditions = "1=1"
	}

	sql_cmd := "select " + columns + " from " + tablename + " where " + conditions + " group by " + group

	return QueryFromMysql(dbname, sql_cmd)
}

func GetSelectDataByCondition(dbname, tablename, columns, conditions string) []interface{} {

	if len(columns) == 0 {
		columns = " * "
	}

	if len(conditions) == 0 {
		conditions = "1=1"
	}

	sql_cmd := "select " + columns + " from " + tablename + " where " + conditions
	return QueryFromMysql(dbname, sql_cmd)
}

func GetDataCountByCondition(dbname, TableName string, Conditions string) int {

	if len(Conditions) == 0 {
		Conditions = "1=1"
	}
	sql_cmd := "select count(*) from " + TableName + " where " + Conditions
	return QueryCount(dbname, sql_cmd)
}

func ExecuteSql(dbname, sql_cmd string) bool {

	stmt, err := DbMap[dbname].Prepare(sql_cmd)
	defer stmt.Close()

	if err != nil {
		beego.Error("Prepare DB error ", err, sql_cmd)
		return false
	}

	res, err := stmt.Exec()
	if err != nil {
		beego.Error("Exec DB error ", err, sql_cmd)
		return false
	}
	num, err := res.RowsAffected()
	if err != nil {
		beego.Error("RowsAffected DB error ", err, sql_cmd)
		return false
	}
	fmt.Println(num)

	return true
}

func ExecuteSqlNew(dbname, sql_cmd string) (bool, error) {
	stmt, err := DbMap[dbname].Prepare(sql_cmd)
	defer stmt.Close()
	if err != nil {
		beego.Error("Prepare DB error ", err)
		return false, err
	}
	res, err := stmt.Exec()
	if err != nil {
		beego.Error("Exec DB error ", err)
		return false, err
	}
	_, err = res.RowsAffected()
	if err != nil {
		beego.Error("RowsAffected DB error ", err)
		return false, err
	}

	return true, nil
}

func InsertData(dbname, chartname, columns, values string) (bool, error) {
	sql_cmd := "INSERT INTO " + chartname + " (" + columns + ") VALUES(" + values + ") "
	return ExecuteSqlNew(dbname, sql_cmd)
}

func DeleteOneData(dbname, chartname, conditions string) (bool, error) {
	sql_cmd := "DELETE FROM " + chartname + " WHERE " + conditions + " LIMIT 1"
	return ExecuteSqlNew(dbname, sql_cmd)
}

func UpdateData(dbname, chartname, updatedata, conditions string) (bool, error) {
	sql_cmd := "UPDATE " + chartname + " SET " + updatedata + " WHERE " + conditions
	return ExecuteSqlNew(dbname, sql_cmd)
}
