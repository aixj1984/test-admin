package routers

import (
	"test-admin/controllers"
	"test-admin/controllers/question"
	"test-admin/controllers/test"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/question/listset", &question.QuestionController{}, "post:LoginCheck")
	beego.Router("/question/list", &question.QuestionController{}, "get:ListQuestion")
	beego.Router("/question/update", &question.QuestionController{}, "post:UpdateQuestion")
	beego.Router("/question/updatestatus", &question.QuestionController{}, "post:UpdateQuestionStatus")
	beego.Router("/question/insert", &question.QuestionController{}, "post:InsertQuestion")

	beego.Router("/test/list", &test.TestController{}, "get:ListTest")
	beego.Router("/test/update", &test.TestController{}, "post:UpdateTest")
	beego.Router("/test/updatestatus", &test.TestController{}, "post:UpdateTestStatus")
	beego.Router("/test/insert", &test.TestController{}, "post:InsertTest")

	beego.Router("/test/question/list", &test.TestController{}, "get:ListTestQuestion")
	beego.Router("/test/question/updatestatus", &test.TestController{}, "post:UpdateTestQuestionStatus")

}
