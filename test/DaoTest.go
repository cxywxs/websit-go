package test

import (
	"example.com/m/v2/dao"
	"example.com/m/v2/mapper"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //加载mysql
	"github.com/jinzhu/gorm"
)

func DaoTest() {
	a := dao.UserDao()
	fmt.Println(a.Username, a.PassWord)
}

func MapperTest() {
	mapper.AnalysisApplication()

}

//注意如果 定义成小写username 引用时 无法调用
type User struct {
	ID       int64  // 列名为 `id`
	Username string // 列名为 `username`
	Password string // 列名为 `password`
}

//设置表名
func (User) TableName() string {
	return "users"
}

func main() {
	db, err := gorm.Open("mysql", "数据库用户名:数据库密码@tcp(ip地址:端口号)/数据库名?     charset=utf8&parseTime=True&loc=Local&timeout=10ms")
	defer db.Close()
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	//执行迁移文件 生成数据表
	db.CreateTable(&User{})

	//添加数据
	user := User{Username: "root", Password: "root"}
	result := db.Create(&user)

	if result.Error != nil {
		fmt.Printf("insert row err %v", result.Error)
		return
	}

	fmt.Println(user.ID) //返回id

	//查询单条数据
	getUser := User{}

	//SELECT id, first FROM users WHERE id = 1 LIMIT 1;
	db.Select([]string{"id", "username"}).First(&getUser, 1)
	fmt.Println(getUser) //打印查询数据

	//修改数据
	user.Username = "update username"
	user.Password = "update password"
	db.Save(&user)

	//查询列表数据
	users := []User{}
	db.Find(&users)
	fmt.Println(&users) //获取所有数据

	//删除数据
	db.Delete(&user)
}
