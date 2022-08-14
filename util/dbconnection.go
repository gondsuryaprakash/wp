package util

import (
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

func Connection() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "surya:Sury@1569@/rf?charset=utf8")
	orm.RunSyncdb("default", false, true)
	fmt.Println("Connection Successfull")
}
