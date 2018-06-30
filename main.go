package main

import (
    "os"
    "strconv"
	_ "auditaudit/routers"
	"github.com/astaxie/beego"
)

func main() {
    port, _ := strconv.Atoi(os.Getenv("PORT"))
    beego.BConfig.Listen.HTTPPort = port
	beego.Run()
}

