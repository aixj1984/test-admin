package models

//	"time"

type Article struct {
	Id         int
	Title      string `xorm:"varchar(300) 'title'"`
	Content    string `xorm:"text 'content'"`
	Source     string `xorm:"varchar(300) 'source'"`
	Abstract   string `xorm:"text   'abstract'"`
	PublicTime string `xorm:"varchar(60)  'public_time'"`
	ReadCount  int    `xorm:"int  'read_count'"`
	Status     int8   `xorm:"int  'status'"`
}

//TableName table name
func (m *Article) TableName() string {
	return "article"
}
