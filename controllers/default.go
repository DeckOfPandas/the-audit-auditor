package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/beego-datatables/datatables"
    "auditaudit/models"
)

type MainController struct {
	beego.Controller
}

type ManageController struct {
	beego.Controller
}

type OperationRecord struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (test *MainController) Test() {
    test.Data["Members"] = []string{"Helen", "Oliver", "Robert", "SeckWei"}
    test.Data["Numbers"] = []int{1,2,3,4,5,6,7,8,9,10}
    test.TplName = "test.tpl"
}

func (manage *ManageController) Add() {
		o := orm.NewOrm()
		o.Using("default")
		audit := models.Audit{}

		//audit.Title = "Does Oliver's Chai Latte have? Maybe we should"
		//audit.Summary = "Something something summary here"
		//audit.Author =  "Robert"
		//audit.Email = "something@something.com"
		//audit.Location = "Theydon Bois"
		//audit.Days = 42
		if err := manage.ParseForm(&audit); err != nil {
			beego.Error("Couldn't parse the form. Reason: ", err)
		} else {
			if manage.Ctx.Input.Method() == "POST" {
				id, err := o.Insert(&audit)
				if err == nil {
					msg := fmt.Sprintf("Audit inserted with id:", id)
					beego.Debug(msg)
				} else {
					msg := fmt.Sprintf("Couldn't insert new Audit. Reason: ", err)
					beego.Debug(msg)
				}
				manage.Redirect("/add", 302)
			} else {
				manage.Data["AuditForm"] = &models.Audit{}
				manage.TplName = "add.tpl"
			}
		}
}

func (manage *ManageController) View() {
    flash := beego.ReadFromRequest(&manage.Controller)

    if ok := flash.Data["error"]; ok != "" {
		// Display error messages
		manage.Data["errors"] = ok
    }

    if ok := flash.Data["notice"]; ok != "" {
		// Display error messages
		manage.Data["notices"] = ok
    }
	o := orm.NewOrm()
    o.Using("default")

    var audits []*models.Audit
    num, err := o.QueryTable("Audit").All(&audits)

    if err != orm.ErrNoRows && num > 0 {
		manage.Data["Audits"] = audits
    }
	manage.TplName = "view.tpl"
}

func (c* OperationRecord) AjaxData() {
	var Qtab datatables.Data
	Qtab.Ctx = c.Input() //datatables get
	Qtab.DBName = "default"
	Qtab.TableName = "Audit" //modles tables name
	Qtab.Columns = []string{"Id","Title" ,"Summary","Author","Email","Location", "Days"} //datatables columns arrange
	Qtab.Order = []string{"Id"}
	Qtab.SearchFilter = []string{"Id","Title" ,"Summary","Author","Email","Location"} //datatables filter
	datatables.RegisterColumns[Qtab.TableName] = new([]models.Audit) //register result 

	rs , _ := Qtab.Table()
	c.Data["json"] = rs
	c.ServeJSON()
}
