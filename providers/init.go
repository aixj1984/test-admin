package providers

var (
	//Account Account Provider
	Account  IAccountProvider
	Question IQuestionProvider
)

//Init 初始化服务
func init() {
	Account = &AccountProvider{}

	Question = &QuestionProvider{}

	println("初始化providers")
}
