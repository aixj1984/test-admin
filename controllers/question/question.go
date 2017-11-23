package question

import (
	"bytes"

	"github.com/astaxie/beego"

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

	data := [...](map[string]string){
		{
			"id":         "10001",
			"username":   "杜甫",
			"email":      "xianxin@layui.com",
			"sex":        "0",
			"city":       "浙江杭州",
			"sign":       "人生恰似一场修行",
			"experience": "116",
			"ip":         "192.168.0.8",
			"logins":     "108",
			"joinTime":   "2016-10-14",
		},
		{
			"id":         "10002",
			"username":   "李白",
			"email":      "xianxin@layui.com",
			"sex":        "男",
			"city":       "浙江杭州",
			"sign":       "人生恰似一场修行",
			"experience": "12",
			"ip":         "192.168.0.8",
			"logins":     "106",
			"joinTime":   "2016-10-14",
		}, {
			"id":         "10003",
			"username":   "王勃",
			"email":      "xianxin@layui.com",
			"sex":        "男",
			"city":       "浙江杭州",
			"sign":       "人生恰似一场修行",
			"experience": "65",
			"ip":         "192.168.0.8",
			"logins":     "106",
			"joinTime":   "2016-10-14",
		}, {
			"id":         "10004",
			"username":   "贤心",
			"email":      "xianxin@layui.com",
			"sex":        "男",
			"city":       "浙江杭州",
			"sign":       "人生恰似一场修行",
			"experience": "666",
			"ip":         "192.168.0.8",
			"logins":     "106",
			"joinTime":   "2016-10-14",
		}, {
			"id":         "10005",
			"username":   "贤心",
			"email":      "xianxin@layui.com",
			"sex":        "男",
			"city":       "浙江杭州",
			"sign":       "人生恰似一场修行",
			"experience": "86",
			"ip":         "192.168.0.8",
			"logins":     "106",
			"joinTime":   "2016-10-14",
		}, {
			"id":         "10006",
			"username":   "贤心",
			"email":      "xianxin@layui.com",
			"sex":        "男",
			"city":       "浙江杭州",
			"sign":       "人生恰似一场修行",
			"experience": "12",
			"ip":         "192.168.0.8",
			"logins":     "106",
			"joinTime":   "2016-10-14",
		}, {
			"id":         "10007",
			"username":   "贤心",
			"email":      "xianxin@layui.com",
			"sex":        "男",
			"city":       "浙江杭州",
			"sign":       "人生恰似一场修行",
			"experience": "16",
			"ip":         "192.168.0.8",
			"logins":     "106",
			"joinTime":   "2016-10-14",
		}, {
			"id":         "10008",
			"username":   "贤心",
			"email":      "xianxin@layui.com",
			"sex":        "男",
			"city":       "浙江杭州",
			"sign":       "人生恰似一场修行",
			"experience": "106",
			"ip":         "192.168.0.8",
			"logins":     "106",
			"joinTime":   "2016-10-14",
		},
	}

	this.Data["json"] = struct {
		Code  int         `json:"code"`
		Msg   string      `json:"msg"`
		Count int64       `json:"count"`
		Data  interface{} `json:"data"`
	}{0, "xxxxxxx", 8, data}

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
