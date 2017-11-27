package routers

import (
	"test-admin/controllers"
	"test-admin/controllers/question"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/question/listset", &question.QuestionController{}, "post:LoginCheck")
	beego.Router("/question/list", &question.QuestionController{}, "get:ListQuestion")
	beego.Router("/question/update", &question.QuestionController{}, "post:UpdateQuestion")

	beego.Router("/question/test", &question.QuestionController{}, "get:ListQuestion2")

}
