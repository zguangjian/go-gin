package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 定义一个全局对象db
var db *sql.DB

// 定义一个初始化数据库的函数
func initDB() (err error) {
	// DSN:Data Source Name
	dsn := "&quot;root:@tcp(127.0.0.1:3306)/test&quot"
	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, err = sql.Open("&quot;mysql&quot;", dsn)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}
func main() {
	router := gin.Default()

	admin := router.Group("/admin")
	{
		public := admin.Group("/public")
		{
			public.POST("/login", func(context *gin.Context) {
				fmt.Println(context)
			})
		}
	}
	home := router.Group("/")
	{
		public := home.Group("/public")
		{
			public.GET("login", func(context *gin.Context) {
				fmt.Println(context)
				context.JSON(http.StatusOK, gin.H{
					"code": 1,
					"msg":  666,
					"data": []string{},
				})
			})
		}
	}

	router.Run()
}
func getting() {
	print(1)
}
