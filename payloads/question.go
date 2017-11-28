package payloads

type SaveQuestion struct {
	CourseId   int    `validate:"required"`
	QuestionId int    `validate:"required"`
	Title      string `validate:"required"`
	Options    string `validate:"required"`
	Answer     string `validate:"required"`
}

type SaveQuestionStatus struct {
	CourseId   int `validate:"required"`
	QuestionId int `validate:"required"`
	Status     int `validate:"required"`
}

type InsertQuestion struct {
	CourseId int    `validate:"required"`
	Title    string `validate:"required"`
	Options  string `validate:"required"`
	Answer   string `validate:"required"`
}
