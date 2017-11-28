package providers

var (
	//Account Account Provider
	Account  IAccountProvider
	Question IQuestionProvider
	Test     ITestProvider
)

//Init 初始化服务
func init() {
	Account = &AccountProvider{}

	Question = &QuestionProvider{}

	Test = &TestProvider{}

	println("初始化providers")
}
