package models

//	"time"

type CourseTest struct {
	Id         int
	Title      string `xorm:"varchar(300) 'title'"`
	TestType   int8   `xorm:"int 'test_type'"` //0：历年真题;  1: 顺序刷题；2: 模拟测试 ;
	CourseId   int    `xorm:"int 'course_id'"`
	Abstract   string `xorm:"text   'abstract'"`
	PublicTime string `xorm:"varchar(60)  'public_time'"`
	ReadCount  int    `xorm:"int  'read_count'"`
	Sources    string `xorm:"varchar(200)  'sources'"`
	Status     int8   `xorm:"int  'status'"`
}

//TableName table name
func (m *CourseTest) TableName() string {
	return "course_test"
}
