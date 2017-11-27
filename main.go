package main

import (
	_ "test-admin/routers"

	"github.com/astaxie/beego"

	_ "test-admin/models"

	_ "test-admin/providers"
)

func main() {
	beego.Run()
}
