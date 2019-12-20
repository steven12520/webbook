package controllers

type YstestController struct {
	BaseController
}

func (self *YstestController) Ystest()  {

	self.Data["pageTitle"]="预审测试"
	self.display()
}

func (self *YstestController) Plays()  {

	Taskid,_:=self.GetInt("Taskid")
	Userid,_:=self.GetInt("Userid")
	if Taskid==0 || Userid==0 {
		self.ajaxMsg("参数错误",MSG_ERR)
	}













	self.ajaxMsg("成功",MSG_OK)
}