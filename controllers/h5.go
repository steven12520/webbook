package controllers

type H5Controller struct {
	BaseController
}

func (self *H5Controller) Report() {
	self.display()
}
