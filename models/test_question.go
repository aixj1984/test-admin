package models

//	"time"

type TestQuestion struct {
	Id         int
	TestId     int `xorm:"int 'test_id'"`
	QuestionId int `xorm:"int 'question_id'"`
}

//TableName table name
func (m *TestQuestion) TableName() string {
	return "course_test_question"
}

func (u *TestQuestion) TableUnique() [][]string {
	return [][]string{
		[]string{"test_id", "question_id"},
	}
}
