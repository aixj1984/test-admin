package models

//	"time"

//Account account model
type QuestionOption struct {
	OptionDesc string
	OptionType int8 //0：文本;  1: 图片 ;
	OptionKey  string
}

//Account account model
type Question struct {
	Id           int
	Title        string `xorm:"varchar(300) 'title'"`
	QuestionType int8   `xorm:"int 'question_type'"` //0：选择;  1: 问答 ;
	Options      string `xorm:"text 'options'"`
	Answer       string `xorm:"varchar(20)   'answer'"`
	Note         string `xorm:"varchar(300)  'note'"`
	Status       int8   `xorm:"int  'status'"`
}

type QuestionLunjijichu struct {
	//Question `xorm:"extends"`
	Id           int
	Title        string `xorm:"varchar(300) 'title'"`
	QuestionType int8   `xorm:"int 'question_type'"` //0：选择;  1: 问答 ;
	Options      string `xorm:"text 'options'"`
	Answer       string `xorm:"varchar(20)   'answer'"`
	Note         string `xorm:"varchar(300)  'note'"`
	Status       int8   `xorm:"int  'status'"`
}

//TableName table name
func (m *QuestionLunjijichu) TableName() string {
	return "question_lunjijichu"
}
