package providers

var (
	//Account Account Provider
	Account      IAccountProvider
	Question     IQuestionProvider
	Test         ITestProvider
	Article      IArticleProvider
	TestQuestion ITestQuestionProvider
)

//Init 初始化服务
func init() {
	Account = &AccountProvider{}

	Question = &QuestionProvider{}

	Test = &TestProvider{}

	Article = &ArticleProvider{}

	TestQuestion = &TestQuestionProvider{}

	println("初始化providers")
}
