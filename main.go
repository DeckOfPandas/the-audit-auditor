package main

import (
    "os"
    "strconv"
	_ "auditaudit/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
    _ "github.com/mattn/go-sqlite3"
    models "auditaudit/models"
)

func init() {
    orm.RegisterDriver("sqlite", orm.DRSqlite)
    orm.RegisterDataBase("default", "sqlite3", "database/orm_test.db")
    orm.RegisterModel(new(models.Audit))
}

func main() {
    port, _ := strconv.Atoi(os.Getenv("PORT"))
    beego.BConfig.Listen.HTTPPort = port
	err := orm.RunSyncdb("default", false, false)
    if err != nil {
        beego.Error(err)
    }
	beego.Run()
}

