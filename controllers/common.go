package controllers

import (
	"../libs"
	"../models"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
)

const (
	MSG_OK  = 0
	MSG_ERR = -1
)

type BaseController struct {
	beego.Controller
	controllerName string
	actionName     string
	noLayout       bool
	pageSize       int
	userId         int
	userName       string
	user           models.UserInfo
	allowUrl       string
}

//前期准备
func (self *BaseController) Prepare() {
	self.pageSize = 20
	controllerName, actionName := self.GetControllerAndAction()
	self.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	self.actionName = strings.ToLower(actionName)
	self.Data["version"] = beego.AppConfig.String("version")
	self.Data["siteName"] = beego.AppConfig.String("site.name")
	self.Data["curRoute"] = self.controllerName + "." + self.actionName
	self.Data["curController"] = self.controllerName
	self.Data["curAction"] = self.actionName
	noAuth := "ads,wxApi,www"
	isNoAuth := strings.Contains(noAuth, self.controllerName)
	if isNoAuth == false {
		self.auth()
	}

	self.Data["loginUserId"] = self.userId
	self.Data["loginUserName"] = self.userName
}
func (c BaseController) ispost() bool {
	//return self.Ctx.Request.Method == "POST"
	return c.Ctx.Request.Method == "POST"
}

func (c *BaseController) redirect(url string) {
	c.Redirect(url, 302)
	c.StopRun()
}

//加载模板
func (self *BaseController) display(tpl ...string) {
	var tplname string
	if len(tpl) > 0 {
		tplname = strings.Join([]string{tpl[0], "html"}, ".")
	} else {
		tplname = self.controllerName + "/" + self.actionName + ".html"
	}
	if !self.noLayout {
		if self.Layout == "" {
			self.Layout = "public/layout.html"
		}
	}
	self.TplName = tplname
}

//登录权限验证
func (self *BaseController) auth() {

	arr := strings.Split(self.Ctx.GetCookie("auth"), "|")
	self.userId = 0
	if len(arr) == 2 {
		idstr, password := arr[0], arr[1]
		userId, _ := strconv.Atoi(idstr)
		if userId > 0 {
			user, err := models.AdminGetById(userId)
			B_pwd := libs.Md5([]byte(string(user.B_id) + user.B_loginname))
			if err == nil && password == B_pwd {
				self.userId = user.B_id
				self.userName = user.B_loginname
				self.user = user
				self.AdminAuth()
			}

			//isHasAuth := strings.Contains(self.allowUrl, self.controllerName+"/"+self.actionName)
			//noAuth := "ajaxsave/ajaxdel/table/loginin/loginout/getnodes/start/show/ajaxapisave"
			//isNoAuth := strings.Contains(noAuth, self.actionName)
			//if isHasAuth == false && isNoAuth == false {
			//	self.Ctx.WriteString("没有权限")
			//	self.ajaxMsg("没有权限", MSG_ERR)
			//	return
			//}
		}
	}

	if self.userId == 0 && (self.controllerName != "login" && self.actionName != "loginin") {
		self.redirect(beego.URLFor("LoginController.LoginIn"))
	}
}

func (self *BaseController) AdminAuth() {
	// 左侧导航栏
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)

	lists, _ := models.MenuList()
	list := make([]map[string]interface{}, 0)
	list2 := make([]map[string]interface{}, 0)

	for _, v := range lists {
		m := make(map[string]interface{}) //创建一个空map
		m["Id"] = int(v.Id)
		m["Parent"] = int(v.Parent)
		m["Name"] = v.Name
		m["Createtime"] = v.Createtime
		m["Type"] = string(v.Type)
		m["Url"] = v.Url
		m["Icon"] = v.Icon

		if v.Parent == 0 {
			list = append(list, m)
		} else {
			list2 = append(list2, m)
		}
	}

	self.Data["SideMenu1"] = list  //一级菜单
	self.Data["SideMenu2"] = list2 //二级菜单

	self.allowUrl = "/home/index"
	//self.allowUrl = allow_url + "/home/index"
}

//ajax返回
func (self *BaseController) ajaxMsg(msg interface{}, msgno int) {
	out := make(map[string]interface{})
	out["status"] = msgno
	out["message"] = msg
	self.Data["json"] = out
	self.ServeJSON()
	self.StopRun()
}

//ajax返回 列表
func (self *BaseController) ajaxList(msg interface{}, msgno int, count int64, data interface{}) {
	out := make(map[string]interface{})
	out["code"] = msgno
	out["msg"] = msg
	out["count"] = count
	out["data"] = data
	self.Data["json"] = out
	self.ServeJSON()
	self.StopRun()
}

//
//// 重定向
//func (self *BaseController) redirect(url string) {
//	self.Redirect(url, 302)
//	self.StopRun()
//}
