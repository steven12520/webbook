/**********************************************
** @Des: login
** @Author: haodaquan
** @Date:   2017-09-07 16:30:10
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-09-17 11:55:21
***********************************************/
package controllers

import (
	"../libs"
	"../models"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
)

type LoginController struct {
	BaseController
}

//登录 TODO:XSRF过滤
func (self *LoginController) LoginIn() {

	if self.userId > 0 {
		self.redirect(beego.URLFor("HomeController.Index"))
	}
	beego.ReadFromRequest(&self.Controller)
	if self.ispost() {

		username := strings.TrimSpace(self.GetString("username"))
		password := strings.TrimSpace(self.GetString("password"))

		if username != "" && password != "" {
			user, err := models.AdminGetByName(username)
			flash := beego.NewFlash()
			errorMsg := ""
			fmt.Println(user)
			fmt.Println(err)
			if err != nil || user.B_pwd != password {
				errorMsg = "帐号或密码错误"
				//} else if user.s == -1 {
				//	errorMsg = "该帐号已禁用"
			} else {
				authkey := libs.Md5([]byte(string(user.B_id) + user.B_loginname))
				self.Ctx.SetCookie("auth", strconv.Itoa(user.B_id)+"|"+authkey, 7*86400)

				self.redirect(beego.URLFor("HomeController.Index"))
			}
			flash.Error(errorMsg)
			flash.Store(&self.Controller)
			self.redirect(beego.URLFor("LoginController.LoginIn"))
		}
	}
	self.TplName = "login/login.html"
}

//登出
func (self *LoginController) LoginOut() {
	self.Ctx.SetCookie("auth", "")
	self.redirect(beego.URLFor("LoginController.LoginIn"))
}

func (self *LoginController) NoAuth() {
	self.Ctx.WriteString("没有权限")
}
