package test

import (
	"time"

	"github.com/astaxie/beego"

	"test-admin/models"

	"test-admin/providers"

	"test-admin/payloads"

	//"github.com/bitly/go-simplejson"
)

type TestController struct {
	beego.Controller
}

func (this *TestController) ListTest() {

	start, _ := this.GetInt("page")
	length, _ := this.GetInt("limit")
	beego.Info("Start is: ", start)
	beego.Info("Length is: ", length)
	beego.Info("course_id is: ", this.GetString("course_id"))
	beego.Info("query_key is: ", this.GetString("query_key"))
	beego.Info("test_status is: ", this.GetString("test_status"))

	var test models.CourseTest

	var tests []*models.CourseTest

	_, err := providers.Test.GetMore(&tests, test, this.GetString("query_key"), this.GetString("test_status"), (start-1)*length, length)

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

	count, err := providers.Test.Count(test, this.GetString("query_key"), this.GetString("test_status"))

	this.Data["json"] = struct {
		Code  int         `json:"code"`
		Msg   string      `json:"msg"`
		Count int64       `json:"count"`
		Data  interface{} `json:"data"`
	}{0, "", count, tests}

	this.ServeJSON()

}

func (this *TestController) UpdateTestStatus() {

	var payload payloads.SaveTestStatus
	beego.Info(string(this.Ctx.Input.RequestBody))
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
	var test models.CourseTest
	test.Id = payload.TestId
	test.Status = int8(payload.Status)
	test.PublicTime = time.Now().Format("2006-01-02 15:04:05")

	_, err := providers.Test.UpdateOne(&test, "status", "public_time")

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

func (this *TestController) UpdateTest() {

	var payload payloads.SaveTest
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

	var test models.CourseTest
	test.Id = payload.TestId
	test.CourseId = payload.CourseId
	test.Abstract = payload.Abstract
	test.Title = payload.Title
	test.Sources = payload.Sources
	test.TestType = int8(payload.TestType)

	_, err := providers.Test.UpdateOne(&test, "title", "abstract", "sources", "test_type", "course_id")

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

func (this *TestController) InsertTest() {
	var payload payloads.InsertTest
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
	var test models.CourseTest
	test.CourseId = payload.CourseId
	test.Abstract = payload.Abstract
	test.Title = payload.Title
	test.Sources = payload.Sources
	test.TestType = int8(payload.TestType)

	_, err := providers.Test.InsertOne(&test)

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

func (this *TestController) ListTestQuestion() {

	start, _ := this.GetInt("page")
	length, _ := this.GetInt("limit")
	beego.Info("Start is: ", start)
	beego.Info("Length is: ", length)
	beego.Info("course_id is: ", this.GetString("course_id"))
	beego.Info("test_id is: ", this.GetString("test_id"))
	beego.Info("query_key is: ", this.GetString("query_key"))
	beego.Info("question_status is: ", this.GetString("question_status"))

	var question models.Question

	var questions []*models.Question

	_, err := providers.TestQuestion.GetMore(&questions, question, this.GetString("query_key"), this.GetString("question_status"), this.GetString("test_id"), (start-1)*length, length)

	if err != nil {
		this.Data["json"] = struct {
			Code  int         `json:"code"`
			Msg   string      `json:"msg"`
			Count int64       `json:"count"`
			Data  interface{} `json:"data"`
		}{10, err.Error(), 0, nil}

		this.ServeJSON()
		return
	}

	count, err := providers.TestQuestion.Count(question, this.GetString("query_key"), this.GetString("question_status"), this.GetString("test_id"))

	if err != nil {
		this.Data["json"] = struct {
			Code  int         `json:"code"`
			Msg   string      `json:"msg"`
			Count int64       `json:"count"`
			Data  interface{} `json:"data"`
		}{10, err.Error(), 0, nil}

		this.ServeJSON()
		return
	}

	choose_count, err := providers.TestQuestion.Count(question, "", "1", this.GetString("test_id"))

	if err != nil {
		this.Data["json"] = struct {
			Code  int         `json:"code"`
			Msg   string      `json:"msg"`
			Count int64       `json:"count"`
			Data  interface{} `json:"data"`
		}{10, err.Error(), 0, nil}

		this.ServeJSON()
		return
	}

	this.Data["json"] = struct {
		Code        int         `json:"code"`
		Msg         string      `json:"msg"`
		Count       int64       `json:"count"`
		ChooseCount int64       `json:"choosecount"`
		Data        interface{} `json:"data"`
	}{0, "", count, choose_count, questions}

	this.ServeJSON()

}

func (this *TestController) UpdateTestQuestionStatus() {

	var payload payloads.SaveTestQuestionStatus
	beego.Info(string(this.Ctx.Input.RequestBody))
	if err := payloads.GeneratePayload(&payload, this.Ctx.Input.RequestBody); err != nil {
		beego.Error(err)
		this.Data["json"] = struct {
			Code int    `json:"code"`
			Msg  string `json:"msg"`
		}{101, "参数错误"}
		this.ServeJSON()
		return
	}

	if payload.Status > 1 {
		this.Data["json"] = struct {
			Code int    `json:"code"`
			Msg  string `json:"msg"`
		}{101, "Status参数错误"}
		this.ServeJSON()
		return
	}
	var question models.TestQuestion
	question.TestId = payload.TestId
	question.QuestionId = payload.QuestionId

	var err error

	if payload.Status == 0 {
		err = providers.TestQuestion.GetOne(&question, "question_id", "test_id")

		if err == nil {
			_, err = providers.TestQuestion.DeleteOne(&question)
		}

	} else {
		_, err = providers.TestQuestion.InsertOne(&question)
	}

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
