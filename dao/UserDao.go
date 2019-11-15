package dao

import (
	"example.com/m/v2/entity"
	"example.com/m/v2/mapper"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func UserDao() *entity.Userdata {
	application := mapper.AnalysisApplication()
	db, err := gorm.Open("mysql", application.Databases.Root+":"+application.Databases.Password+"@tcp("+application.Databases.Server+":"+application.Databases.Port+")/"+application.Databases.Database+"?charset=utf8")
	if err != nil {
		panic("连接数据库失败")
	}

	db.Table("userdata")
	var userdata *entity.Userdata
	db.Where("username = ?", "jinzhu").First(&userdata)
	fmt.Println(userdata.PassWord)
	return userdata

}
