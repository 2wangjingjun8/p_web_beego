package main

import (
	_ "p_web/models"
	_ "p_web/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
