package question

import (
	"bytes"

	"github.com/astaxie/beego"

	"test-admin/models"

	"test-admin/providers"

	"encoding/json"

	//"github.com/bitly/go-simplejson"
)

type QuestionController struct {
	beego.Controller
}

func (this *QuestionController) ListQuestion() {

	start, _ := this.GetInt("page")
	length, _ := this.GetInt("limit")
	beego.Info("Start is: ", start)
	beego.Info("Length is: ", length)
	beego.Info("course_id is: ", this.GetString("course_id"))
	beego.Info("query_key is: ", this.GetString("query_key"))
	beego.Info("question_status is: ", this.GetString("question_status"))

	var question models.QuestionLunjijichu

	var questions []*models.QuestionLunjijichu

	_, err := providers.Question.GetMore(&questions, question, this.GetString("query_key"), (start-1)*length, length)

	if err != nil {
		this.Data["json"] = struct {
			Code  int         `json:"code"`
			Msg   string      `json:"msg"`
			Count int64       `json:"count"`
			Data  interface{} `json:"data"`
		}{0, err.Error(), 0, nil}

		this.ServeJSON()
		return
	}

	this.Data["json"] = struct {
		Code  int         `json:"code"`
		Msg   string      `json:"msg"`
		Count int64       `json:"count"`
		Data  interface{} `json:"data"`
	}{0, "xxxxxxx", 8, questions}

	this.ServeJSON()

}

func (this *QuestionController) UpdateQuestion() {

	var question models.QuestionLunjijichu
	json.Unmarshal(this.Ctx.Input.RequestBody, &question)

	_, err := providers.Question.UpdateOne(&question, "title", "answer", "options")

	if err != nil {
		this.Data["json"] = struct {
			Code int    `json:"code"`
			Msg  string `json:"msg"`
		}{10, err.Error()}

		this.ServeJSON()
		return
	}

	this.Data["json"] = struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}{0, ""}

	this.ServeJSON()
}

func (this *QuestionController) ListQuestion2() {

	start, _ := this.GetInt("page")
	length, _ := this.GetInt("limit")
	beego.Info("Start is: ", start)
	beego.Info("Length is: ", length)
	beego.Info("course_id is: ", this.GetString("course_id"))
	beego.Info("query_key is: ", this.GetString("query_key"))
	beego.Info("question_status is: ", this.GetString("question_status"))

	var options = make([]models.QuestionOption, 0)

	var option = models.QuestionOption{
		OptionType: 0,
		OptionDesc: "aaaaaaa",
		OptionKey:  "a",
	}
	options = append(options, option)
	options_json, _ := json.Marshal(options)
	beego.Info(string(options_json))

	var question = &models.QuestionLunjijichu{
		Options: string(options_json),
	}

	effact, err := providers.Question.InsertOne(question)

	beego.Info(effact)

	err = providers.Question.GetOne(question)

	if err != nil {
		this.Data["json"] = struct {
			Code  int         `json:"code"`
			Msg   string      `json:"msg"`
			Count int64       `json:"count"`
			Data  interface{} `json:"data"`
		}{101, err.Error(), 8, nil}
		this.ServeJSON()
		return
	}

	this.Data["json"] = struct {
		Code  int         `json:"code"`
		Msg   string      `json:"msg"`
		Count int64       `json:"count"`
		Data  interface{} `json:"data"`
	}{0, "xxxxxxx", 8, question}

	var questions []*models.QuestionLunjijichu

	effact, err = providers.Question.GetMore(&questions, question, "", 0, 10)

	beego.Info(effact)

	this.Data["json"] = struct {
		Code  int         `json:"code"`
		Msg   string      `json:"msg"`
		Count int64       `json:"count"`
		Data  interface{} `json:"data"`
	}{0, "xxxxxxx", 8, questions}

	this.ServeJSON()

}

func (this *QuestionController) LoginCheck() {

	reqbody := this.Ctx.Request.Body
	buf := new(bytes.Buffer)
	buf.ReadFrom(reqbody)
	JiraAuth := buf.String()
	//js, _ := simplejson.NewJson([]byte(reqstr))

	//username, _ := js.Get("username").String()
	//password, _ := js.Get("password").String()

	LoginResp := "bbb"

	this.Data["json"] = struct {
		Jira_Auth  string      `json:"JiraAuth"`
		Login_Resp interface{} `json:"LoginResp"`
	}{JiraAuth, LoginResp}

	this.ServeJSON()
}