package models

//	"time"

//Account account model
type QuestionOption struct {
	OptionDesc string
	OptionType int8 //0：文本;  1: 图片 ;
	OptionKey  string
}

/*
1	轮机基础
2	轮机管理
3	船舶辅机与电气
4	船舶动力装置
5	避碰与信号
6	船舶管理
7	航道与引航
8	船舶操纵
9	机舱管理
10	主推进动力装置
11	船舶驾驶与管理
*/

var CourseMap = map[string]string{
	"1":  "question_1",
	"2":  "question_2",
	"3":  "question_3",
	"4":  "question_4",
	"5":  "question_5",
	"6":  "question_6",
	"7":  "question_7",
	"8":  "question_8",
	"9":  "question_9",
	"10": "question_10",
	"11": "question_11",
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

type Question1 struct {
	Question `xorm:"extends"`
}

func (m *Question1) TableName() string {
	return "question_1"
}

type Question2 struct {
	Question `xorm:"extends"`
}

func (m *Question2) TableName() string {
	return "question_2"
}

type Question3 struct {
	Question `xorm:"extends"`
}

func (m *Question3) TableName() string {
	return "question_3"
}

type Question4 struct {
	Question `xorm:"extends"`
}

func (m *Question4) TableName() string {
	return "question_4"
}

type Question5 struct {
	Question `xorm:"extends"`
}

func (m *Question5) TableName() string {
	return "question_5"
}

type Question6 struct {
	Question `xorm:"extends"`
}

func (m *Question6) TableName() string {
	return "question_6"
}

type Question7 struct {
	Question `xorm:"extends"`
}

func (m *Question7) TableName() string {
	return "question_7"
}

type Question8 struct {
	Question `xorm:"extends"`
}

func (m *Question8) TableName() string {
	return "question_8"
}

type Question9 struct {
	Question `xorm:"extends"`
}

func (m *Question9) TableName() string {
	return "question_9"
}

type Question10 struct {
	Question `xorm:"extends"`
}

func (m *Question10) TableName() string {
	return "question_10"
}

type Question11 struct {
	Question `xorm:"extends"`
}

func (m *Question11) TableName() string {
	return "question_11"
}
