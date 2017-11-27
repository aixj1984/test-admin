package question

import (
	"bytes"

	"github.com/astaxie/beego"

	"test-admin/models"

	"test-admin/providers"

	"test-admin/payloads"

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

	count, err := providers.Question.Count(&questions, question, this.GetString("query_key"))

	this.Data["json"] = struct {
		Code  int         `json:"code"`
		Msg   string      `json:"msg"`
		Count int64       `json:"count"`
		Data  interface{} `json:"data"`
	}{0, "", count, questions}

	this.ServeJSON()

}

func (this *QuestionController) UpdateQuestion() {

	var payload payloads.SaveQuestion
	if err := payloads.GeneratePayload(&payload, this.Ctx.Input.RequestBody); err != nil {
		beego.Error(err)
		this.Data["json"] = struct {
			Code int    `json:"code"`
			Msg  string `json:"msg"`
		}{101, "参数错误"}
		this.ServeJSON()
		return
	}

	if payload.CourseId == 0 || payload.CourseId > 4 {
		this.Data["json"] = struct {
			Code int    `json:"code"`
			Msg  string `json:"msg"`
		}{101, "courseid参数错误"}
		this.ServeJSON()
		return
	}
	var question models.QuestionLunjijichu
	question.Id = payload.QuestionId
	question.Options = payload.Options
	question.Title = payload.Title
	question.Answer = payload.Answer

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

func (this *QuestionController) InsertQuestion() {
	var payload payloads.InsertQuestion
	if err := payloads.GeneratePayload(&payload, this.Ctx.Input.RequestBody); err != nil {
		beego.Error(err)
		this.Data["json"] = struct {
			Code int    `json:"code"`
			Msg  string `json:"msg"`
		}{101, "参数错误"}
		this.ServeJSON()
		return
	}

	if payload.CourseId == 0 || payload.CourseId > 4 {
		this.Data["json"] = struct {
			Code int    `json:"code"`
			Msg  string `json:"msg"`
		}{101, "courseid参数错误"}
		this.ServeJSON()
		return
	}
	var question models.QuestionLunjijichu
	question.Options = payload.Options
	question.Title = payload.Title
	question.Answer = payload.Answer
	question.Status = 0
	question.Note = ""

	_, err := providers.Question.InsertOne(&question)

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
