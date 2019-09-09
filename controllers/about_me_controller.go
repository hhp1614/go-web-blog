package controllers

type AboutMeController struct {
	BaseController
}

func (c *AboutMeController) Get() {
	c.Data["weChat"] = "hhp1614"
	c.Data["qq"] = "971056445"
	c.Data["github"] = "hhp1614"
	c.TplName = "about.html"
}
