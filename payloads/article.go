package payloads

type SaveArticle struct {
	ArticleId int    `validate:"required"`
	Title     string `validate:"required"`
	Source    string `validate:"required"`
	Abstract  string `validate:"required"`
	Content   string `validate:"required"`
}

type SaveArticleStatus struct {
	ArticleId int `validate:"required"`
	Status    int
}

type InsertArticle struct {
	Title    string `validate:"required"`
	Source   string `validate:"required"`
	Abstract string `validate:"required"`
	Content  string `validate:"required"`
}

type DelArticle struct {
	ArticleId int `validate:"required"`
}
