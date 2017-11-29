package models

//	"time"

//Account account model
type QuestionOption struct {
	OptionDesc string
	OptionType int8 //0：文本;  1: 图片 ;
	OptionKey  string
}

var CourseMap = map[string]string{
	"1": "question_lunjijichu",
	"2": "question_jicangguanli",
	"3": "question_lunjiguanli",
	"4": "question_bipengyuxinhao",
}

//Account account model
type Question struct {
	Id           int
	Title        string `xorm:"varchar(300) 'title'"`
	QuestionType int8   `xorm:"int 'question_type'"` //0：选择;  1: 问答 ;
	Options      string `xorm:"text 'options'"`
	Answer       string `xorm:"varchar(20)   'answer'"`
	Note         string `xorm:"text  'note'"`
	Status       int8   `xorm:"int  'status'"`
}

type QuestionLunjijichu struct {
	Question `xorm:"extends"`
}

//TableName table name
func (m *QuestionLunjijichu) TableName() string {
	return "question_lunjijichu"
}

type QuestionJicangguanli struct {
	Question `xorm:"extends"`
}

//TableName table name
func (m *QuestionJicangguanli) TableName() string {
	return "question_jicangguanli"
}

type QuestionLunjiguanli struct {
	Question `xorm:"extends"`
}

//TableName table name
func (m *QuestionLunjiguanli) TableName() string {
	return "question_lunjiguanli"
}

type QuestionBipengyuxinhao struct {
	Question `xorm:"extends"`
}

//TableName table name
func (m *QuestionBipengyuxinhao) TableName() string {
	return "question_bipengyuxinhao"
}
