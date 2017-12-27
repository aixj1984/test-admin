package article

import (
	"bytes"

	"github.com/astaxie/beego"

	"test-admin/models"

	"test-admin/providers"

	"test-admin/payloads"

	"time"

	//"github.com/bitly/go-simplejson"
)

type ArticleController struct {
	beego.Controller
}

func (this *ArticleController) ListArticle() {

	start, _ := this.GetInt("page")
	length, _ := this.GetInt("limit")
	beego.Info("Start is: ", start)
	beego.Info("Length is: ", length)
	beego.Info("query_key is: ", this.GetString("query_key"))
	beego.Info("article_status is: ", this.GetString("article_status"))

	var questions []*models.Article

	_, err := providers.Article.GetMore(&questions, this.GetString("query_key"), this.GetString("article_status"), (start-1)*length, length)

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

	count, err := providers.Article.Count(this.GetString("query_key"), this.GetString("article_status"))

	this.Data["json"] = struct {
		Code  int         `json:"code"`
		Msg   string      `json:"msg"`
		Count int64       `json:"count"`
		Data  interface{} `json:"data"`
	}{0, "", count, questions}

	this.ServeJSON()

}

func (this *ArticleController) UpdateArticleStatus() {

	var payload payloads.SaveArticleStatus
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

	_, err := providers.Article.UpdateOneStatus(payload.Status, payload.ArticleId)

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

func (this *ArticleController) UpdateArticle() {

	var payload payloads.SaveArticle
	if err := payloads.GeneratePayload(&payload, this.Ctx.Input.RequestBody); err != nil {
		beego.Error(err)
		this.Data["json"] = struct {
			Code int    `json:"code"`
			Msg  string `json:"msg"`
		}{101, "参数错误"}
		this.ServeJSON()
		return
	}

	var article models.Article
	article.Id = payload.ArticleId
	article.Abstract = payload.Abstract
	article.Title = payload.Title
	article.Content = payload.Content
	article.Source = payload.Source

	_, err := providers.Article.UpdateOne(&article)

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

func (this *ArticleController) InsertArticle() {
	var payload payloads.InsertArticle
	if err := payloads.GeneratePayload(&payload, this.Ctx.Input.RequestBody); err != nil {
		beego.Error(err)
		this.Data["json"] = struct {
			Code int    `json:"code"`
			Msg  string `json:"msg"`
		}{101, "参数错误"}
		this.ServeJSON()
		return
	}

	var article models.Article
	article.Abstract = payload.Abstract
	article.Title = payload.Title
	article.Content = payload.Content
	article.Source = payload.Source
	article.Status = 0
	article.ReadCount = 0
	article.PublicTime = time.Now().Format("2006-01-02 15:04:05")

	_, err := providers.Article.InsertOne(&article)

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

func (this *ArticleController) LoginCheck() {

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
